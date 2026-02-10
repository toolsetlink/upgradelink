package apk

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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

type ApkActionUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApkActionUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApkActionUploadLogic {
	return &ApkActionUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ApkActionUpload 处理 APK 上传
// 1. 验证应用信息
// 2. 下载文件
// 3. 上传文件到 S3
// 4. 创建 APK 版本
// 5. 创建 APK 升级策略
// 6. 返回上传结果
func (l *ApkActionUploadLogic) ApkActionUpload(req *types.ApkActionUploadReq) (resp *types.ApkActionUploadResp, err error) {
	// 1. 验证应用信息
	apkInfo, err := l.validateApkInfo(req)
	if err != nil {
		return nil, err
	}

	// 2. 下载文件
	filePath, fileGithubPath, err := l.downloadFile(req.Url, "downloads")
	if err != nil {
		l.Errorf("下载文件失败: %w", err)
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.downloadFailed"), l.svcCtx.Trans.Trans(l.ctx, "common.downloadFailedDocs"))
	}

	// 3. 上传文件到 S3
	fileUUID, err := l.uploadFileToOss(filePath, fileGithubPath)
	if err != nil {
		l.Errorf("上传文件失败: %w", err)
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.uploadFailed"), l.svcCtx.Trans.Trans(l.ctx, "common.uploadFailedDocs"))
	}

	// 4. 创建 APK 版本
	versionId, err := l.createApkVersion(apkInfo, fileUUID, req.Version)
	if err != nil {
		l.Errorf("创建 APK 版本失败: %w", err)
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.createFailed"), l.svcCtx.Trans.Trans(l.ctx, "common.createFailedDocs"))
	}

	// 5. 创建 APK 升级策略
	err = l.createApkUpgradeStrategy(apkInfo, versionId, req.Version, req.PromptUpgradeContent)
	if err != nil {
		l.Errorf("创建 APK 升级策略失败: %w", err)
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.createFailed"), l.svcCtx.Trans.Trans(l.ctx, "common.createFailedDocs"))
	}

	// 6. 返回上传结果
	return &types.ApkActionUploadResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 200,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, "common.success"),
		},
	}, nil
}

// validateApkInfo 验证 APK 信息
func (l *ApkActionUploadLogic) validateApkInfo(req *types.ApkActionUploadReq) (*model.UpgradeApk, error) {
	// 通过 key 获取 APK 信息
	apkInfo, err := l.svcCtx.ResourceCtx.GetApkInfoByKey(l.ctx, req.AppKey)
	if err != nil {
		return nil, http_handlers.NewLinkErr(l.ctx, http_handlers.ErrInternalServerError, l.svcCtx.Trans.Trans(l.ctx, "common.databaseError"), l.svcCtx.Trans.Trans(l.ctx, "common.databaseErrorDocs"))
	}

	return apkInfo, nil
}

// downloadFile 下载文件
func (l *ApkActionUploadLogic) downloadFile(url, outputDir string) (string, string, error) {
	// 从 URL 中提取文件名
	filename := filepath.Base(url)
	outputFile := filepath.Join(outputDir, filename)

	// 创建输出目录（如果不存在）
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		l.Errorf("创建目录失败: %w", err)
		return "", "", fmt.Errorf("创建目录失败: %w", err)
	}

	// 创建 HTTP 客户端
	client := &http.Client{
		Timeout: 10 * time.Minute,
	}

	// 发送 HTTP 请求
	httpReq, err := http.NewRequestWithContext(l.ctx, "GET", url, nil)
	if err != nil {
		l.Errorf("创建请求失败: %w", err)
		return "", "", fmt.Errorf("创建请求失败: %w", err)
	}

	resp, err := client.Do(httpReq)
	if err != nil {
		l.Errorf("下载请求失败: %w", err)
		return "", "", fmt.Errorf("下载请求失败: %w", err)
	}
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		l.Errorf("下载失败, 状态码: %d, 响应: %s", resp.StatusCode, string(body))
		return "", "", fmt.Errorf("下载失败, 状态码: %d", resp.StatusCode)
	}

	// 创建输出文件
	out, err := os.Create(outputFile)
	if err != nil {
		l.Errorf("创建文件失败: %w", err)
		return "", "", fmt.Errorf("创建文件失败: %w", err)
	}
	defer out.Close()

	// 复制文件内容
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		l.Errorf("复制文件内容失败: %w", err)
		return "", "", fmt.Errorf("复制文件内容失败: %w", err)
	}

	return outputFile, url, nil
}

// uploadFileToOss 上传文件到 S3
func (l *ApkActionUploadLogic) uploadFileToOss(filePath, fileGithubPath string) (string, error) {
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
	md5Sum := hex.EncodeToString(md5Hash.Sum(nil))

	// 重置文件指针到开头
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

	// 上传文件到 S3
	s3Config := &cloudstorage.S3Config{
		Bucket:    l.svcCtx.Config.UploadConf.Bucket,
		SecretID:  l.svcCtx.Config.UploadConf.SecretID,
		SecretKey: l.svcCtx.Config.UploadConf.SecretKey,
		Endpoint:  l.svcCtx.Config.UploadConf.Endpoint,
		Region:    l.svcCtx.Config.UploadConf.Region,
	}

	s3Service := cloudstorage.NewS3Service(s3Config)
	uploadedURL, err := s3Service.UploadFile(l.ctx, file, relativeSrc)
	if err != nil {
		l.Errorf("S3 上传失败: %w", err)
		return "", fmt.Errorf("S3 上传失败: %w", err)
	}

	// 数据录入 cloud_file 表中
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
		// 数据录入 github url 表中
		_, err = l.svcCtx.ResourceCtx.AddFileGithub(l.ctx, resource.AddFileGithubReq{
			FileId: fileUUID.String(),
			Url:    fileGithubPath,
		})
		if err != nil {
			return "", fmt.Errorf("录入 github 文件记录失败: %w", err)
		}
	}

	return fileUUID.String(), nil
}

// getFileSuffix 处理文件后缀
func (l *ApkActionUploadLogic) getFileSuffix(fileName string) (string, string, error) {
	// 校验文件后缀
	dotIndex := filepath.Ext(fileName)
	if dotIndex == "" {
		return "", "", fmt.Errorf("拒绝无后缀的文件: %s", fileName)
	}

	baseName := strings.TrimSuffix(fileName, dotIndex)
	suffix := strings.TrimPrefix(dotIndex, ".")

	return baseName, suffix, nil
}

// getFileType 根据文件扩展名判断文件类型
func (l *ApkActionUploadLogic) getFileType(fileName string) string {
	ext := filepath.Ext(fileName)
	if ext == "" {
		return "other"
	}

	ext = strings.ToLower(strings.TrimPrefix(ext, "."))
	if strings.Contains("jpg jpeg png gif bmp webp", ext) {
		return "image"
	} else if strings.Contains("mp4 mp3 wav avi mov", ext) {
		return "video"
	} else if strings.Contains("mp3 wav aac flac", ext) {
		return "audio"
	} else if ext == "apk" {
		return "apk"
	} else {
		return "other"
	}
}

// createApkVersion 创建 APK 版本
func (l *ApkActionUploadLogic) createApkVersion(apkInfo *model.UpgradeApk, fileUUID, version string) (int64, error) {
	// 转换版本号为数字格式
	versionCode, err := common.SemVerToInt64(version)
	if err != nil {
		l.Errorf("版本号转换失败: %w", err)
		return 0, fmt.Errorf("版本号转换失败: %w", err)
	}

	// 创建 APK 版本
	versionId, err := l.svcCtx.ResourceCtx.AddApkVersion(l.ctx, resource.AddApkVersionReq{
		CompanyId:   apkInfo.CompanyId,
		ApkId:       apkInfo.Id,
		CloudFileId: fileUUID,
		VersionName: version,
		VersionCode: versionCode,
		Description: "",
	})
	if err != nil {
		l.Errorf("创建 APK 版本失败: %w", err)
		return 0, fmt.Errorf("创建 APK 版本失败: %w", err)
	}

	return versionId, nil
}

// createApkUpgradeStrategy 创建 APK 升级策略
func (l *ApkActionUploadLogic) createApkUpgradeStrategy(apkInfo *model.UpgradeApk, versionId int64, version, promptUpgradeContent string) error {
	// 使用时区统一时间点
	now := time.Now().Local()
	end := now.AddDate(1, 0, 0)

	// 创建 APK 升级策略
	_, err := l.svcCtx.ResourceCtx.AddApkStrategy(l.ctx, resource.AddApkStrategyReq{
		CompanyId:            apkInfo.CompanyId,
		Enable:               1,
		Name:                 fmt.Sprintf("%s-%s", apkInfo.Name, version),
		Description:          "自动生成的升级策略",
		ApkId:                apkInfo.Id,
		ApkVersionId:         versionId,
		BeginDatetime:        now,
		EndDatetime:          end,
		UpgradeType:          1,
		PromptUpgradeContent: promptUpgradeContent,
	})
	if err != nil {
		l.Errorf("创建 APK 升级策略失败: %w", err)
		return fmt.Errorf("创建 APK 升级策略失败: %w", err)
	}

	return nil
}
