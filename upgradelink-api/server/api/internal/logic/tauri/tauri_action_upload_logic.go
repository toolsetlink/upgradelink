package tauri

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"upgradelink-api/server/api/internal/common"
	"upgradelink-api/server/api/internal/common/http_handlers"
	"upgradelink-api/server/api/internal/common/utils/cloudstorage"
	"upgradelink-api/server/api/internal/common/utils/filex"
	"upgradelink-api/server/api/internal/resource"
	"upgradelink-api/server/api/internal/resource/model"
	"upgradelink-api/server/api/internal/svc"
	"upgradelink-api/server/api/internal/types"

	"upgradelink-api/server/api/internal/common/utils/uuidx"

	"github.com/duke-git/lancet/v2/datetime"
	"github.com/zeromicro/go-zero/core/logx"
)

type TauriActionUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTauriActionUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TauriActionUploadLogic {
	return &TauriActionUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 定义必要的结构体
type TauriLastJsonReq struct {
	Version   string           `json:"version" validate:"required"`
	Notes     string           `json:"notes"`
	PubDate   string           `json:"pub_date"`
	Platforms PlatformsRequest `json:"platforms"`
}

type PlatformInfoRequest struct {
	Signature string `json:"signature" form:"signature" validate:"required"`
	Url       string `json:"url" form:"url" validate:"required,url"`
}

type PlatformsRequest struct {
	DarwinAarch64  PlatformInfoRequest `json:"darwin-aarch64,optional"`
	DarwinX8664    PlatformInfoRequest `json:"darwin-x86_64,optional"`
	LinuxX8664     PlatformInfoRequest `json:"linux-x86_64,optional"`
	LinuxAarch64   PlatformInfoRequest `json:"linux-aarch64,optional"`
	WindowsX8664   PlatformInfoRequest `json:"windows-x86_64,optional"`
	WindowsAarch64 PlatformInfoRequest `json:"windows-aarch64,optional"`
	WindowsI686    PlatformInfoRequest `json:"windows-i686,optional"`
}

// TauriActionUpload 处理 Tauri 应用上传
// 1. 验证应用信息
// 2. 下载并解析版本信息
// 3. 构建下载任务
// 4. 删除旧版本数据
// 5. 并发处理平台任务
// 6. 返回上传结果
func (l *TauriActionUploadLogic) TauriActionUpload(req *types.TauriActionUploadReq) (resp *types.TauriActionUploadResp, err error) {
	// 1. 验证应用信息
	auriInfo, err := l.validateAppInfo(req)
	if err != nil {
		return nil, err
	}

	// 2. 下载并解析版本信息
	latestJsonReq, err := l.downloadAndParseVersionInfo(req.LatestJsonUrl)
	if err != nil {
		return nil, err
	}

	// 3. 构建下载任务
	downloadTasks := l.buildDownloadTasks(latestJsonReq.Platforms, latestJsonReq.Version, latestJsonReq.Notes)
	if len(downloadTasks) == 0 {
		var res types.TauriActionUploadResp
		res.Code = 200
		res.Msg = l.svcCtx.Trans.Trans(l.ctx, "common.success")
		return &res, nil
	}

	// 4. 删除旧版本数据
	err = l.deleteOldVersionData(auriInfo, latestJsonReq.Version)
	if err != nil {
		l.Errorf("删除旧版本数据失败: %w", err)
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.deleteFailed"), l.svcCtx.Trans.Trans(l.ctx, "common.internalErrorDocs"))
	}

	l.Infof("构建的下载任务: %+v", downloadTasks)

	// 5. 执行并发下载和上传
	outputDir := "downloads"
	err = l.processPlatformsConcurrently(downloadTasks, outputDir, auriInfo)
	if err != nil {
		l.Errorf("处理平台任务时出错: %v", err)
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.failed"), l.svcCtx.Trans.Trans(l.ctx, "common.internalErrorDocs"))
	}

	// 6. 返回上传结果
	var res types.TauriActionUploadResp
	res.Code = 200
	res.Msg = l.svcCtx.Trans.Trans(l.ctx, "common.success")
	return &res, nil
}

// validateAppInfo 验证应用信息
func (l *TauriActionUploadLogic) validateAppInfo(req *types.TauriActionUploadReq) (*model.UpgradeTauri, error) {
	// 通过 key 获取 tauri 信息
	tauriInfo, err := l.svcCtx.ResourceCtx.GetTauriInfoByKey(l.ctx, req.AppKey)
	if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.internalErrorDocs"))
	}

	return tauriInfo, nil
}

// fetchUrlContent 获取URL内容并返回响应体
// 参数：
// - ctx: 上下文
// - url: 请求URL
// 返回：
// - []byte: 响应体内容
// - error: 错误信息
func (l *TauriActionUploadLogic) fetchUrlContent(ctx context.Context, url string) ([]byte, error) {
	// 创建自定义HTTP客户端，设置超时时间为10分钟
	client := &http.Client{
		Timeout: 10 * time.Minute,
	}

	// 发送HTTP请求
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		l.Errorf("创建请求失败: %w", err)
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		l.Errorf("请求失败: %w", err)
		return nil, fmt.Errorf("请求失败: %w", err)
	}
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		l.Errorf("请求失败, 状态码: %d, 响应: %s", resp.StatusCode, string(body))
		return nil, fmt.Errorf("请求失败, 状态码: %d, 响应: %s", resp.StatusCode, string(body))
	}

	// 读取响应体
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		l.Errorf("读取响应体失败: %w", err)
		return nil, fmt.Errorf("读取响应体失败: %w", err)
	}

	return content, nil
}

// downloadAndParseVersionInfo 下载并解析版本信息
func (l *TauriActionUploadLogic) downloadAndParseVersionInfo(latestJsonUrl string) (*TauriLastJsonReq, error) {
	// 获取URL内容
	content, err := l.fetchUrlContent(l.ctx, latestJsonUrl)
	if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.downloadFailed"), l.svcCtx.Trans.Trans(l.ctx, "common.internalErrorDocs"))
	}

	// 解析JSON内容
	latestJsonReq := &TauriLastJsonReq{}
	err = json.Unmarshal(content, latestJsonReq)
	if err != nil {
		l.Errorf("解析JSON失败: %w", err)
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.validationError"), l.svcCtx.Trans.Trans(l.ctx, "common.internalErrorDocs"))
	}

	return latestJsonReq, nil
}

// deleteOldVersionData 删除旧版本数据
func (l *TauriActionUploadLogic) deleteOldVersionData(tauriInfo *model.UpgradeTauri, version string) error {
	// 转换版本号为数字格式
	versionCode, err := common.SemVerToInt64(version)
	if err != nil {
		l.Errorf("版本号转换失败: %w", err)
		return http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "upgrade.versionFormatError"), l.svcCtx.Trans.Trans(l.ctx, "common.internalErrorDocs"))
	}

	// 通过版本号删除对应的升级任务与版本
	return l.DeleteTauriVersionAndUpgradeStrategyByVersionCode(context.Background(), tauriInfo, versionCode)
}

// buildDownloadTasks 动态构建下载任务，过滤掉无效平台
// 参数：
// - platforms: 平台信息
// - version: 版本号
// - notes: 版本说明
// 返回：
// - []struct: 下载任务列表
func (l *TauriActionUploadLogic) buildDownloadTasks(
	platforms PlatformsRequest, version, notes string,
) []struct {
	name    string
	plat    PlatformInfoRequest
	target  string
	arch    string
	version string
	notes   string
} {
	var tasks []struct {
		name    string
		plat    PlatformInfoRequest
		target  string
		arch    string
		version string
		notes   string
	}

	// 定义平台映射关系，便于统一处理
	platformMappings := []struct {
		name   string
		plat   *PlatformInfoRequest
		target string
		arch   string
	}{
		{"darwin-x86_64", &platforms.DarwinX8664, "darwin", "x86_64"},
		{"darwin-aarch64", &platforms.DarwinAarch64, "darwin", "aarch64"},
		{"linux-x86_64", &platforms.LinuxX8664, "linux", "x86_64"},
		{"linux-aarch64", &platforms.LinuxAarch64, "linux", "aarch64"},
		{"windows-x86_64", &platforms.WindowsX8664, "windows", "x86_64"},
		{"windows-aarch64", &platforms.WindowsAarch64, "windows", "aarch64"},
		{"windows-i686", &platforms.WindowsI686, "windows", "i686"},
	}

	// 遍历所有平台，过滤出有效平台
	for _, mapping := range platformMappings {
		// 检查平台是否有效（URL不为空）
		if mapping.plat == nil || mapping.plat.Url == "" {
			l.Infof("跳过平台 %s: 未提供有效URL", mapping.name)
			continue
		}

		// 检查签名是否存在（如果需要）
		if mapping.plat.Signature == "" {
			l.Infof("平台 %s: 未提供签名信息", mapping.name)
		}

		tasks = append(tasks, struct {
			name    string
			plat    PlatformInfoRequest
			target  string
			arch    string
			version string
			notes   string
		}{
			name:    mapping.name,
			plat:    *mapping.plat,
			target:  mapping.target,
			arch:    mapping.arch,
			version: version,
			notes:   notes,
		})
	}

	return tasks
}

// processPlatformsConcurrently 并发处理多个平台的下载和上传任务
// 参数：
// - tasks: 下载任务列表
// - outputDir: 输出目录
// - tauriInfo: tauri 应用信息
// 返回：
// - error: 错误信息
func (l *TauriActionUploadLogic) processPlatformsConcurrently(
	tasks []struct {
		name    string
		plat    PlatformInfoRequest
		target  string
		arch    string
		version string
		notes   string
	},
	outputDir string,
	tauriInfo *model.UpgradeTauri,
) error {
	if len(tasks) == 0 {
		return nil
	}

	ctx, cancel := context.WithCancel(l.ctx)
	defer cancel()

	var wg sync.WaitGroup
	errCh := make(chan error, len(tasks)) // 缓冲通道，容纳所有可能的错误

	for _, task := range tasks {
		wg.Add(1)
		go func(task struct {
			name    string
			plat    PlatformInfoRequest
			target  string
			arch    string
			version string
			notes   string
		}) {
			defer wg.Done()

			// 下载升级文件
			filePath, fileGithubPath, err := l.DownloadFile(ctx, task.plat, outputDir)
			if err != nil {
				errCh <- fmt.Errorf("平台 %s 下载失败: %w", task.name, err)
				return
			}
			l.Infof("平台 %s 下载成功, 保存至: %s, github: %s", task.name, filePath, fileGithubPath)

			// 下载安装包文件， 排查只需要 替换 mac 类型，才有独立的 dmg类型安装包，其他平台不需要
			installFilePath := filePath
			installFileGithubPath := ""
			if task.target == "darwin" {
				installFilePath, installFileGithubPath, err = l.DownloadInstallFile(ctx, task.plat, task.version, task.arch, outputDir)
				if err != nil {
					errCh <- fmt.Errorf("平台 %s 下载 install 失败: %w", task.name, err)
					return
				}
			}

			// 上传文件到 S3 并处理后续逻辑
			err = l.processFileUpload(ctx, filePath, fileGithubPath, installFilePath, installFileGithubPath, tauriInfo, task.target, task.arch, task.version, task.plat.Signature, task.notes)
			if err != nil {
				errCh <- fmt.Errorf("平台 %s 上传处理失败: %w", task.name, err)
				return
			}
			l.Infof("平台 %s 处理完成", task.name)
		}(task)
	}

	// 等待所有任务完成或上下文取消
	go func() {
		wg.Wait()
		close(errCh)
	}()

	// 收集所有错误（如果有）
	var allErrors []string
	for err := range errCh {
		allErrors = append(allErrors, err.Error())
	}

	// 如果有错误，返回合并的错误信息
	if len(allErrors) > 0 {
		l.Errorf("处理过程中发生错误: %s", strings.Join(allErrors, "; "))
	}

	return nil
}

// processFileUpload 处理文件的上传到 S3 及后续数据库操作
// 参数：
// - ctx: 上下文
// - filePath: 升级文件路径
// - fileGithubPath: 升级文件 GitHub 路径
// - installFilePath: 安装文件路径
// - installFileGithubPath: 安装文件 GitHub 路径
// - tauriInfo: tauri 应用信息
// - target: 目标平台
// - arch: 架构
// - version: 版本号
// - signature: 签名
// - notes: 版本说明
// 返回：
// - error: 错误信息
func (l *TauriActionUploadLogic) processFileUpload(ctx context.Context, filePath, fileGithubPath, installFilePath, installFileGithubPath string, tauriInfo *model.UpgradeTauri,
	target, arch, version, signature, notes string) error {

	fileUUID := ""
	installFileUUID := ""

	// 上传安装包到 S3
	if installFilePath != "" {
		var err error
		installFileUUID, err = l.UploadFileToOss(ctx, installFilePath, installFileGithubPath)
		if err != nil {
			return fmt.Errorf("上传安装包失败: %w", err)
		}
	}

	// 上传升级包到 S3
	if filePath != "" {
		var err error
		fileUUID, err = l.UploadFileToOss(ctx, filePath, fileGithubPath)
		if err != nil {
			return fmt.Errorf("上传升级包失败: %w", err)
		}
	}

	// 转换版本号为数字格式
	versionCode, err := common.SemVerToInt64(version)
	if err != nil {
		l.Errorf("版本号转换失败: %w", err)
		return fmt.Errorf("版本号转换失败: %w", err)
	}

	// 创建版本信息
	versionId, err := l.createTauriVersion(ctx, tauriInfo, fileUUID, installFileUUID, version, versionCode, target, arch, signature, notes)
	if err != nil {
		return err
	}

	// 创建升级任务
	return l.createTauriUpgradeStrategy(ctx, tauriInfo, versionId, version, target, arch, notes)
}

// getFileSuffix 处理文件后缀，支持多段扩展名
func (l *TauriActionUploadLogic) getFileSuffix(fileName string) (string, string, error) {
	// 校验文件后缀
	dotIndex := strings.LastIndex(fileName, ".")
	if dotIndex == -1 {
		return "", "", fmt.Errorf("拒绝无后缀的文件: %s", fileName)
	}

	baseName, suffix := fileName[:dotIndex], fileName[dotIndex+1:]

	// 特殊处理常见的多段扩展名
	commonMultiPartExtensions := map[string]bool{
		"tar.gz":    true,
		"tar.bz2":   true,
		"tar.xz":    true,
		"tar.lz":    true,
		"tar.lzma":  true,
		"jpg.jpg":   true, // 处理一些可能的错误命名
		"jpeg.jpeg": true,
		"png.png":   true,
	}

	// 检查是否可能是多段扩展名
	if len(baseName) > 0 {
		prevDotIndex := strings.LastIndex(baseName, ".")
		if prevDotIndex > 0 {
			potentialMultiPartExt := baseName[prevDotIndex+1:] + "." + suffix
			if commonMultiPartExtensions[potentialMultiPartExt] {
				// 确认为多段扩展名，调整文件名和扩展名
				baseName = baseName[:prevDotIndex]
				suffix = potentialMultiPartExt
			}
		}
	}

	return baseName, suffix, nil
}

// getFileType 根据文件扩展名判断文件类型
func (l *TauriActionUploadLogic) getFileType(fileName string) string {
	parts := strings.Split(fileName, ".")
	if len(parts) > 1 {
		ext := strings.ToLower(parts[len(parts)-1])
		if strings.Contains("jpg jpeg png gif bmp webp", ext) {
			return "image"
		} else if strings.Contains("mp4 mp3 wav avi mov", ext) {
			return "video"
		} else if strings.Contains("mp3 wav aac flac", ext) {
			return "audio"
		}
	}
	return "other"
}

// UploadFileToOss 上传文件到S3并返回相关信息
func (l *TauriActionUploadLogic) UploadFileToOss(ctx context.Context, filePath, fileGithubPath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("无法打开文件: %w", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return "", fmt.Errorf("无法获取文件信息: %w", err)
	}

	// 处理文件后缀
	fileName, fileSuffix, err := l.getFileSuffix(fileInfo.Name())
	if err != nil {
		return "", err
	}

	// 计算文件 MD5
	md5Hash := md5.New()
	if _, err := io.Copy(md5Hash, file); err != nil {
		l.Errorf("计算文件 MD5 失败: %w", err)
		return "", fmt.Errorf("计算文件 MD5 失败: %w", err)
	}
	md5Sum := hex.EncodeToString(md5Hash.Sum(nil)) // 转为 32 位十六进制字符串

	// 重置文件指针到开头（否则后续上传会读取空内容）
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		l.Errorf("重置文件指针失败: %w", err)
		return "", fmt.Errorf("重置文件指针失败: %w", err)
	}

	fileUUID := uuidx.NewUUID()
	storeFileName := fileUUID.String() + "." + fileSuffix
	userId := "upload"

	// 判断文件类型
	fileType := l.getFileType(fileInfo.Name())

	relativeSrc := fmt.Sprintf("%s/%s/%s/%s",
		l.svcCtx.Config.UploadConf.Folder,
		datetime.FormatTimeToStr(time.Now(), "yyyy-mm-dd"),
		fileType,
		storeFileName)

	// 上传文件到S3
	s3Config := &cloudstorage.S3Config{
		Bucket:    l.svcCtx.Config.UploadConf.Bucket,
		SecretID:  l.svcCtx.Config.UploadConf.SecretID,
		SecretKey: l.svcCtx.Config.UploadConf.SecretKey,
		Endpoint:  l.svcCtx.Config.UploadConf.Endpoint,
		Region:    l.svcCtx.Config.UploadConf.Region,
	}

	s3Service := cloudstorage.NewS3Service(s3Config)
	uploadedURL, err := s3Service.UploadFile(ctx, file, relativeSrc)
	if err != nil {
		l.Errorf("S3上传失败: %w", err)
		return "", fmt.Errorf("S3上传失败: %w", err)
	}

	// 数据录入cloud_file表中
	_, err = l.svcCtx.ResourceCtx.AddFmsCloudFiles(l.ctx, resource.AddFmsCloudFilesReq{
		Id:        fileUUID.String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		State:     int64(1),
		Name:      fileName,
		FileType:  uint64(filex.ConvertFileTypeToUint8(fileType)),
		Url:       uploadedURL,
		Size:      uint64(fileInfo.Size()),
		UserId:    userId,
		Md5:       md5Sum,
	})
	if err != nil {
		return "", fmt.Errorf("录入云文件记录失败: %w", err)
	}

	if fileGithubPath != "" {
		// 数据录入github url表中
		_, err = l.svcCtx.ResourceCtx.AddFileGithub(l.ctx, resource.AddFileGithubReq{
			FileId: fileUUID.String(),
			Url:    fileGithubPath,
		})
		if err != nil {
			return "", fmt.Errorf("录入github文件记录失败: %w", err)
		}
	}

	return fileUUID.String(), nil
}

// DeleteTauriVersionAndUpgradeStrategyByVersionCode 删除版本号与对应的升级任务
func (l *TauriActionUploadLogic) DeleteTauriVersionAndUpgradeStrategyByVersionCode(ctx context.Context, tauriInfo *model.UpgradeTauri, versionCode int64) error {

	// 查询出旧版本
	oldVersionList, err := l.svcCtx.ResourceCtx.GetTauriVersionListByTauriIdAndVersionCode(l.ctx, tauriInfo.Id, versionCode)
	if err != nil {
		l.Errorf("查询Tauri旧版本失败: %w", err)
		return fmt.Errorf("查询旧版本失败: %w", err)
	}

	// 循环删除旧版本对应的升级策略和版本
	for _, tauriVersion := range oldVersionList {
		// 删除旧版本对应的升级策略
		_, err = l.svcCtx.ResourceCtx.DelTauriStrategyByVersionId(l.ctx, tauriVersion.Id)
		if err != nil {
			l.Errorf("删除Tauri旧版本对应的升级策略失败: %w", err)
			return fmt.Errorf("删除旧版本对应的升级策略失败: %w", err)
		}

		// 删除旧版本
		_, err = l.svcCtx.ResourceCtx.DelTauriVersionById(l.ctx, tauriVersion.Id)
		if err != nil {
			l.Errorf("删除Tauri旧版本失败: %w", err)
			return fmt.Errorf("删除旧版本失败: %w", err)
		}
	}

	return nil
}

// createTauriVersion 创建设备版本信息
func (l *TauriActionUploadLogic) createTauriVersion(ctx context.Context, tauriInfo *model.UpgradeTauri,
	fileUUID, installFileUUID string, version string, versionCode int64, target string, arch string, signature string, notes string) (int64, error) {
	versionId, err := l.svcCtx.ResourceCtx.AddTauriVersion(l.ctx, resource.AddTauriVersionReq{
		CompanyId:          tauriInfo.CompanyId,
		TauriId:            tauriInfo.Id,
		CloudFileId:        fileUUID,
		InstallCloudFileId: installFileUUID,
		VersionName:        version,
		VersionCode:        versionCode,
		Target:             target,
		Arch:               arch,
		Signature:          signature,
		Description:        notes,
	})
	if err != nil {
		l.Errorf("创建设备版本失败: %w", err)
		return 0, fmt.Errorf("创建设备版本失败: %w", err)
	}
	return versionId, nil
}

// createTauriUpgradeStrategy 创建tauri升级策略
func (l *TauriActionUploadLogic) createTauriUpgradeStrategy(ctx context.Context, tauriInfo *model.UpgradeTauri,
	versionId int64, version, target, arch, notes string) error {
	// 使用时区统一时间点
	now := time.Now().Local()   // 使用服务器本地时区
	end := now.AddDate(1, 0, 0) // 默认有效期1年
	_, err := l.svcCtx.ResourceCtx.AddTauriStrategy(l.ctx, resource.AddTauriStrategyReq{
		CompanyId:            tauriInfo.CompanyId,
		Enable:               1,
		Name:                 fmt.Sprintf("%s-%s-%s-%s", tauriInfo.Name, version, target, arch),
		Description:          "自动生成的升级策略",
		TauriId:              tauriInfo.Id,
		TauriVersionId:       versionId,
		BeginDatetime:        now,
		EndDatetime:          end,
		UpgradeType:          1,
		PromptUpgradeContent: notes,
	})
	if err != nil {
		return fmt.Errorf("创建升级策略失败: %w", err)
	}
	return nil
}

// downloadFile 通用下载文件方法
func (l *TauriActionUploadLogic) downloadFile(ctx context.Context, url string, outputDir string) (string, string, error) {
	// 从URL中提取文件名
	filename := filepath.Base(url)
	outputFile := filepath.Join(outputDir, filename)

	// 创建输出目录（如果不存在）
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		l.Errorf("创建目录失败: %w", err)
		return "", "", fmt.Errorf("创建目录失败: %w", err)
	}

	// 创建自定义HTTP客户端，设置超时时间为10分钟
	client := &http.Client{
		Timeout: 10 * time.Minute,
	}

	// 发送HTTP请求
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		l.Errorf("创建请求失败: %w", err)
		return "", "", fmt.Errorf("创建请求失败: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		l.Errorf("下载请求失败: %w", err)
		return "", "", fmt.Errorf("下载请求失败: %w", err)
	}
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		l.Errorf("下载失败, 状态码: %d, 响应: %s", resp.StatusCode, string(body))
		return "", "", fmt.Errorf("下载失败, 状态码: %d, 响应: %s", resp.StatusCode, string(body))
	}

	// 创建输出文件
	out, err := os.Create(outputFile)
	if err != nil {
		l.Errorf("创建文件失败: %w", err)
		return "", "", fmt.Errorf("创建文件失败: %w", err)
	}
	defer func() {
		if err := out.Close(); err != nil {
			l.Errorf("关闭文件失败: %v", err)
		}
	}()

	// 复制文件内容
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		l.Errorf("复制文件内容失败: %w", err)
		return "", "", fmt.Errorf("复制文件内容失败: %w", err)
	}

	return outputFile, url, nil
}

// DownloadFile 下载单个升级文件
// 参数：
// - ctx: 上下文
// - task: 平台信息请求
// - outputDir: 输出目录
// 返回：
// - string: 文件路径
// - string: 文件 GitHub 路径
// - error: 错误信息
func (l *TauriActionUploadLogic) DownloadFile(ctx context.Context, task PlatformInfoRequest, outputDir string) (string, string, error) {
	return l.downloadFile(ctx, task.Url, outputDir)
}

// DownloadInstallFile 下载单个安装文件 目前只有mac类型需要下载安装包
// 参数：
// - ctx: 上下文
// - task: 平台信息请求
// - version: 版本号
// - arch: 架构
// - outputDir: 输出目录
// 返回：
// - string: 文件路径
// - string: 文件 GitHub 路径
// - error: 错误信息
func (l *TauriActionUploadLogic) DownloadInstallFile(ctx context.Context, task PlatformInfoRequest, version, arch, outputDir string) (string, string, error) {
	// 判断架构 拼接出 dmg 文件下载地址
	if arch == "aarch64" {
		// 提取前缀（去掉后缀 "aarch64.app.tar.gz"）
		prefix := strings.TrimSuffix(task.Url, "aarch64.app.tar.gz")
		task.Url = prefix + version + "_aarch64.dmg"
	} else if arch == "x86_64" {
		// 提取前缀（去掉后缀 "x64.app.tar.gz"）
		prefix := strings.TrimSuffix(task.Url, "x64.app.tar.gz")
		task.Url = prefix + version + "_x64.dmg"
	}

	return l.downloadFile(ctx, task.Url, outputDir)
}
