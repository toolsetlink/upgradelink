package apk_patch

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
	"upgradelink-admin/server/api/internal/common/utils/cloudstorage"
	"upgradelink-admin/server/api/internal/common/utils/filex"
	"upgradelink-admin/server/api/internal/common/utils/uuidx"
	"upgradelink-admin/server/api/internal/resource"
	"upgradelink-admin/server/api/internal/resource/model"
	"upgradelink-admin/server/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logc"
)

// ApKPatch 处理APK差分包生成任务
// 主要功能包括：查询待处理任务、下载APK文件、生成差分包、上传到云存储
// 参数：
//   - ctx: 上下文对象，用于控制超时和取消
//   - svcCtx: 服务上下文，包含数据库连接等资源
//
// 返回值：无
func ApKPatch(ctx context.Context, svcCtx *svc.ServiceContext) {
	logc.Info(ctx, "开始执行APK差分包处理任务", "time", time.Now().Format("2006-01-02 15:04:05"))

	// 查询待处理的patch文件列表
	patchList, err := resource.GetApkPatchListByStatus(ctx, svcCtx, PatchStatusNotProcessed)
	if err != nil {
		logc.Error(ctx, "查询待处理的patch文件失败",
			"error", err.Error(),
			"status", PatchStatusNotProcessed)
		return
	}

	if len(patchList) == 0 {
		logc.Info(ctx, "没有待处理的APK差分包任务")
		return
	}

	logc.Info(ctx, "发现待处理的APK差分包任务",
		"count", len(patchList),
		"patchIds", getPatchIds(patchList))

	// 遍历处理每个patch文件
	successCount := 0
	for _, patch := range patchList {
		if err := processSinglePatch(ctx, svcCtx, patch); err != nil {
			logc.Error(ctx, "处理单个patch任务失败",
				"patchId", patch.Id,
				"apkId", patch.ApkId,
				"error", err.Error())
		} else {
			successCount++
		}
	}

	logc.Info(ctx, "APK差分包处理任务执行完成",
		"total", len(patchList),
		"success", successCount,
		"failed", len(patchList)-successCount)
}

// getPatchIds 获取patch ID列表用于日志记录
// 参数：
//   - patchList: APK差分包列表
//
// 返回值：包含所有patch ID的切片
func getPatchIds(patchList []model.UpgradeApkPatch) []int64 {
	ids := make([]int64, len(patchList))
	for i, patch := range patchList {
		ids[i] = patch.Id
	}
	return ids
}

// processSinglePatch 处理单个APK差分包任务
// 主要步骤：
//  1. 更新状态为处理中
//  2. 下载旧版本APK文件
//  3. 下载新版本APK文件
//  4. 创建差分包存储目录
//  5. 生成差分包
//  6. 验证差分包大小
//  7. 上传差分包到云存储
//
// 参数：
//   - ctx: 上下文对象
//   - svcCtx: 服务上下文
//   - patch: 单个APK差分包任务
//
// 返回值：处理过程中发生的错误
func processSinglePatch(ctx context.Context, svcCtx *svc.ServiceContext, patch model.UpgradeApkPatch) error {
	logc.Info(ctx, "开始处理单个patch任务",
		"patchId", patch.Id,
		"apkId", patch.ApkId,
		"lowVersionId", patch.LowApkVersionId,
		"highVersionId", patch.HighApkVersionId)

	// 修改patch状态为处理中
	if err := resource.UpdateApkPatchStatus(ctx, svcCtx, patch.Id, PatchStatusProcessing, patch.CloudFileId, patch.Description); err != nil {
		logc.Error(ctx, "修改patch状态为处理中失败",
			"patchId", patch.Id,
			"error", err.Error())
		return err
	}

	// 下载旧包文件
	oldApkFile, err := DownloadApkFile(ctx, svcCtx, patch.LowApkVersionId)
	if err != nil {
		logc.Error(ctx, "下载旧包文件失败",
			"patchId", patch.Id,
			"versionId", patch.LowApkVersionId,
			"error", err.Error())
		return updatePatchStatusWithError(ctx, svcCtx, patch, PatchStatusError, err)
	}
	defer cleanupTempFile(oldApkFile)

	// 下载新包文件
	newApkFile, err := DownloadApkFile(ctx, svcCtx, patch.HighApkVersionId)
	if err != nil {
		logc.Error(ctx, "下载新包文件失败",
			"patchId", patch.Id,
			"versionId", patch.HighApkVersionId,
			"error", err.Error())
		return updatePatchStatusWithError(ctx, svcCtx, patch, PatchStatusError, err)
	}
	defer cleanupTempFile(newApkFile)

	logc.Info(ctx, "APK文件下载完成",
		"patchId", patch.Id,
		"oldApkFile", oldApkFile,
		"newApkFile", newApkFile)

	// 创建差分包存储目录
	patchDir := getPatchDirectory(patch.ApkId, patch.Id)
	if err := os.MkdirAll(patchDir, 0755); err != nil {
		logc.Error(ctx, "创建差分包目录失败",
			"patchId", patch.Id,
			"directory", patchDir,
			"error", err.Error())
		return updatePatchStatusWithError(ctx, svcCtx, patch, PatchStatusError, err)
	}

	// 生成差分包文件路径
	outPatchFilePath := getPatchFilePath(patchDir, patch.LowApkVersionId, patch.HighApkVersionId)

	// 生成差分包
	if err := HDiffPatchHDiffZ(ctx, oldApkFile, newApkFile, outPatchFilePath); err != nil {
		logc.Error(ctx, "生成差分包文件失败",
			"patchId", patch.Id,
			"error", err.Error())
		return updatePatchStatusWithError(ctx, svcCtx, patch, PatchStatusError, err)
	}

	logc.Info(ctx, "差分包生成成功",
		"patchId", patch.Id,
		"patchFile", outPatchFilePath)

	// 检查差分包大小是否合理
	if err := validatePatchSize(ctx, svcCtx, patch, outPatchFilePath, newApkFile); err != nil {
		return err
	}

	// 上传差分包到云存储
	if err := uploadPatchToCloud(ctx, svcCtx, patch, outPatchFilePath); err != nil {
		return err
	}

	logc.Info(ctx, "patch任务处理完成", "patchId", patch.Id)
	return nil
}

// updatePatchStatusWithError 更新patch状态并记录错误
// 参数：
//   - ctx: 上下文对象
//   - svcCtx: 服务上下文
//   - patch: APK差分包任务
//   - status: 要更新的状态
//   - err: 发生的错误
//
// 返回值：原始错误
func updatePatchStatusWithError(ctx context.Context, svcCtx *svc.ServiceContext, patch model.UpgradeApkPatch, status int64, err error) error {
	updateErr := resource.UpdateApkPatchStatus(ctx, svcCtx, patch.Id, status, patch.CloudFileId, err.Error())
	if updateErr != nil {
		logc.Error(ctx, "更新patch状态失败", "patchId", patch.Id, "status", status, "error", updateErr)
	}
	return err
}

// getPatchDirectory 获取差分包存储目录
// 参数：
//   - apkId: APK ID
//   - patchId: 差分包 ID
//
// 返回值：完整的目录路径
func getPatchDirectory(apkId, patchId int64) string {
	return fmt.Sprintf("%s%s%d/%s%d/",
		STORAGE_PATH,
		DOWNLOAD_APK_PATH,
		apkId,
		PATCH_PATH,
		patchId)
}

// getPatchFilePath 获取差分包文件路径
// 参数：
//   - patchDir: 差分包目录
//   - lowVersionId: 低版本ID
//   - highVersionId: 高版本ID
//
// 返回值：完整的差分包文件路径
func getPatchFilePath(patchDir string, lowVersionId, highVersionId int64) string {
	return fmt.Sprintf("%s%d_%d%s", patchDir, lowVersionId, highVersionId, PATCH_FILE_SUFFIX)
}

// validatePatchSize 验证差分包大小是否合理
// 检查差分包是否大于新版本APK文件，如果大于则标记为无效
// 参数：
//   - ctx: 上下文对象
//   - svcCtx: 服务上下文
//   - patch: APK差分包任务
//   - patchFilePath: 差分包文件路径
//   - newApkFilePath: 新版本APK文件路径
//
// 返回值：验证过程中发生的错误
func validatePatchSize(ctx context.Context, svcCtx *svc.ServiceContext, patch model.UpgradeApkPatch, patchFilePath, newApkFilePath string) error {
	patchFileSize, err := GetFileSize(patchFilePath)
	if err != nil {
		logc.Error(ctx, "获取差分包文件大小失败",
			"patchId", patch.Id,
			"filePath", patchFilePath,
			"error", err.Error())
		return updatePatchStatusWithError(ctx, svcCtx, patch, PatchStatusError, err)
	}

	newApkFileSize, err := GetFileSize(newApkFilePath)
	if err != nil {
		logc.Error(ctx, "获取新包文件大小失败",
			"patchId", patch.Id,
			"filePath", newApkFilePath,
			"error", err.Error())
		return updatePatchStatusWithError(ctx, svcCtx, patch, PatchStatusError, err)
	}

	logc.Info(ctx, "文件大小对比",
		"patchId", patch.Id,
		"patchSize", patchFileSize,
		"newApkSize", newApkFileSize,
		"ratio", fmt.Sprintf("%.2f%%", float64(patchFileSize)/float64(newApkFileSize)*100))

	// 如果差分包大于新包，标记为差分包过大
	if patchFileSize > newApkFileSize {
		logc.Error(ctx, "差分包大于新版本全量包",
			"patchId", patch.Id,
			"patchSize", patchFileSize,
			"newApkSize", newApkFileSize,
			"excess", patchFileSize-newApkFileSize)
		return updatePatchStatusWithError(ctx, svcCtx, patch, PatchStatusPatchGreaterThanFull,
			fmt.Errorf("差分包大小(%d)大于新包大小(%d)", patchFileSize, newApkFileSize))
	}

	return nil
}

// uploadPatchToCloud 上传差分包到云存储
// 参数：
//   - ctx: 上下文对象
//   - svcCtx: 服务上下文
//   - patch: APK差分包任务
//   - patchFilePath: 差分包文件路径
//
// 返回值：上传过程中发生的错误
func uploadPatchToCloud(ctx context.Context, svcCtx *svc.ServiceContext, patch model.UpgradeApkPatch, patchFilePath string) error {
	// 修改状态为正在上传
	if err := resource.UpdateApkPatchStatus(ctx, svcCtx, patch.Id, PatchStatusUploading, patch.CloudFileId, patch.Description); err != nil {
		logc.Error(ctx, "修改patch状态为上传中失败",
			"patchId", patch.Id,
			"error", err.Error())
		return err
	}

	logc.Info(ctx, "开始上传差分包到云存储",
		"patchId", patch.Id,
		"filePath", patchFilePath)

	// 1. 打开文件
	file, err := os.Open(patchFilePath)
	if err != nil {
		logc.Error(ctx, "打开差分包文件失败",
			"patchId", patch.Id,
			"filePath", patchFilePath,
			"error", err.Error())
		return updatePatchStatusWithError(ctx, svcCtx, patch, PatchStatusUploadFailed, err)
	}
	defer file.Close()

	// 2. 获取文件信息
	fileInfo, err := file.Stat()
	if err != nil {
		logc.Error(ctx, "获取文件信息失败",
			"patchId", patch.Id,
			"filePath", patchFilePath,
			"error", err.Error())
		return updatePatchStatusWithError(ctx, svcCtx, patch, PatchStatusUploadFailed, err)
	}

	// 3. 计算文件 MD5
	md5Hash := md5.New()
	if _, err := io.Copy(md5Hash, file); err != nil {
		logc.Error(ctx, "计算文件 MD5 失败",
			"patchId", patch.Id,
			"filePath", patchFilePath,
			"error", err.Error())
		return updatePatchStatusWithError(ctx, svcCtx, patch, PatchStatusUploadFailed, err)
	}
	md5Sum := hex.EncodeToString(md5Hash.Sum(nil))

	// 4. 重置文件指针到开头
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		logc.Error(ctx, "重置文件指针失败",
			"patchId", patch.Id,
			"filePath", patchFilePath,
			"error", err.Error())
		return updatePatchStatusWithError(ctx, svcCtx, patch, PatchStatusUploadFailed, err)
	}

	// 5. 构建 S3 配置
	s3Config := &cloudstorage.S3Config{
		Bucket:    svcCtx.Config.UploadConf.Bucket,
		SecretID:  svcCtx.Config.UploadConf.SecretID,
		SecretKey: svcCtx.Config.UploadConf.SecretKey,
		Endpoint:  svcCtx.Config.UploadConf.Endpoint,
		Region:    svcCtx.Config.UploadConf.Region,
	}

	// 6. 创建 S3 服务实例并上传文件
	s3Service := cloudstorage.NewS3Service(s3Config)

	// 构建 S3 文件路径
	fileName := fmt.Sprintf("%s/%s/%d_%d%s",
		svcCtx.Config.UploadConf.Folder,
		time.Now().Format("2006-01-02"),
		patch.LowApkVersionId,
		patch.HighApkVersionId,
		PATCH_FILE_SUFFIX)

	// 上传文件
	uploadedURL, err := s3Service.UploadFile(ctx, file, fileName)
	if err != nil {
		logc.Error(ctx, "上传差分包到 S3 失败",
			"patchId", patch.Id,
			"filePath", patchFilePath,
			"s3Key", fileName,
			"error", err.Error())
		return updatePatchStatusWithError(ctx, svcCtx, patch, PatchStatusUploadFailed, err)
	}

	// 7. 如果配置了 CDN URL，使用 CDN URL
	finalURL := uploadedURL
	if svcCtx.Config.UploadConf.CdnUrl != "" {
		finalURL = fmt.Sprintf("%s%s", svcCtx.Config.UploadConf.CdnUrl, fileName)
	}

	// 8. 生成云文件 ID
	cloudFileID := uuidx.NewUUID().String()

	// 9. 构建云文件请求结构体
	now := time.Now()
	addReq := resource.AddFmsCloudFilesReq{
		Id:                        cloudFileID,
		CreatedAt:                 now,
		UpdatedAt:                 now,
		State:                     1, // 正常状态
		Name:                      fmt.Sprintf("%d_%d", patch.LowApkVersionId, patch.HighApkVersionId),
		Url:                       finalURL,
		Size:                      uint64(fileInfo.Size()),
		Md5:                       md5Sum,
		FileType:                  uint64(filex.ConvertFileTypeToUint8("other")), // 差分包类型为其他
		UserId:                    "system",                                      // 系统上传
		CloudFileStorageProviders: 1,                                             // S3 存储提供商
	}

	// 10. 插入到云文件表
	if _, err := resource.AddFmsCloudFiles(ctx, svcCtx, addReq); err != nil {
		logc.Error(ctx, "插入云文件记录失败",
			"patchId", patch.Id,
			"cloudFileID", cloudFileID,
			"error", err.Error())
		return updatePatchStatusWithError(ctx, svcCtx, patch, PatchStatusUploadFailed, err)
	}

	// 11. 修改patch状态为处理完成
	if err := resource.UpdateApkPatchStatus(ctx, svcCtx, patch.Id, PatchStatusCompleted, cloudFileID, patch.Description); err != nil {
		logc.Error(ctx, "修改patch状态为完成失败",
			"patchId", patch.Id,
			"cloudFileId", cloudFileID,
			"error", err.Error())
		return err
	}

	logc.Info(ctx, "差分包上传成功",
		"patchId", patch.Id,
		"cloudFileId", cloudFileID,
		"s3Key", fileName,
		"fileUrl", finalURL)
	return nil
}

// cleanupTempFile 清理临时文件
// 参数：
//   - filePath: 要清理的文件路径
func cleanupTempFile(filePath string) {
	if filePath != "" {
		if err := os.Remove(filePath); err != nil && !os.IsNotExist(err) {
			// 记录警告但不中断流程
			logc.Info(context.Background(), "清理临时文件失败", "filePath", filePath, "error", err)
		}
	}
}

// HDiffPatchHDiffZ 使用hdiffz工具生成APK差分包
// 参数说明：
//   - ctx: 上下文对象
//   - oldFilePath: 旧版本APK文件路径
//   - newFilePath: 新版本APK文件路径
//   - outPatchFilePath: 输出的差分包文件路径
//
// 返回值：命令执行过程中发生的错误
func HDiffPatchHDiffZ(ctx context.Context, oldFilePath, newFilePath, outPatchFilePath string) error {
	// 验证输入文件存在
	if err := validateInputFiles(oldFilePath, newFilePath); err != nil {
		return err
	}

	// 构建hdiffz命令
	diffOptions := "-f"
	command := fmt.Sprintf("/usr/local/bin/hdiffz %s %s %s %s",
		diffOptions, oldFilePath, newFilePath, outPatchFilePath)

	logc.Info(ctx, "执行hdiffz命令",
		"command", command,
		"oldFile", oldFilePath,
		"newFile", newFilePath,
		"outputFile", outPatchFilePath)

	// 转义命令中的特殊字符
	command = escapeCommandCharacters(command)

	// 执行命令
	cmd := exec.Command("/bin/sh", "-c", command)

	// 设置超时时间（5分钟）
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	cmd = exec.CommandContext(ctxWithTimeout, "/bin/sh", "-c", command)

	output, err := cmd.CombinedOutput()

	// 检查是否超时
	if errors.Is(ctxWithTimeout.Err(), context.DeadlineExceeded) {
		logc.Error(ctx, "hdiffz命令执行超时",
			"timeout", "5分钟",
			"command", command)
		return fmt.Errorf("hdiffz命令执行超时（5分钟）")
	}

	if err != nil {
		logc.Error(ctx, "hdiffz命令执行失败",
			"error", err.Error(),
			"output", string(output),
			"exitCode", getExitCode(cmd))
		return fmt.Errorf("hdiffz执行失败: %w, 输出: %s", err, string(output))
	}

	// 验证输出文件是否生成
	if err := validateOutputFile(outPatchFilePath); err != nil {
		logc.Error(ctx, "差分包文件验证失败",
			"outputFile", outPatchFilePath,
			"error", err.Error())
		return fmt.Errorf("差分包文件验证失败: %w", err)
	}

	logc.Info(ctx, "hdiffz命令执行成功",
		"output", string(output),
		"outputFile", outPatchFilePath,
		"fileSize", getFileSizeForLog(outPatchFilePath))
	return nil
}

// validateInputFiles 验证输入文件是否存在且可读
// 参数：
//   - oldFilePath: 旧版本文件路径
//   - newFilePath: 新版本文件路径
//
// 返回值：验证过程中发生的错误
func validateInputFiles(oldFilePath, newFilePath string) error {
	// 检查旧文件
	if err := checkFileExistsAndReadable(oldFilePath); err != nil {
		return fmt.Errorf("旧版本文件验证失败: %w", err)
	}

	// 检查新文件
	if err := checkFileExistsAndReadable(newFilePath); err != nil {
		return fmt.Errorf("新版本文件验证失败: %w", err)
	}

	return nil
}

// checkFileExistsAndReadable 检查文件是否存在且可读
// 参数：
//   - filePath: 文件路径
//
// 返回值：验证过程中发生的错误
func checkFileExistsAndReadable(filePath string) error {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("文件不存在: %s", filePath)
		}
		return fmt.Errorf("无法访问文件: %w", err)
	}

	// 检查是否为文件
	if fileInfo.IsDir() {
		return fmt.Errorf("路径是目录而不是文件: %s", filePath)
	}

	// 检查文件大小
	if fileInfo.Size() == 0 {
		return fmt.Errorf("文件大小为0: %s", filePath)
	}

	// 检查文件权限
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("文件不可读: %w", err)
	}
	file.Close()

	return nil
}

// validateOutputFile 验证输出文件是否成功生成
// 参数：
//   - filePath: 输出文件路径
//
// 返回值：验证过程中发生的错误
func validateOutputFile(filePath string) error {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("输出文件未生成: %s", filePath)
		}
		return fmt.Errorf("无法访问输出文件: %w", err)
	}

	// 检查是否为文件
	if fileInfo.IsDir() {
		return fmt.Errorf("输出路径是目录而不是文件: %s", filePath)
	}

	// 检查文件大小
	if fileInfo.Size() == 0 {
		return fmt.Errorf("输出文件大小为0: %s", filePath)
	}

	return nil
}

// getExitCode 获取命令的退出码
// 参数：
//   - cmd: 执行的命令
//
// 返回值：退出码，如果无法获取则返回-1
func getExitCode(cmd *exec.Cmd) int {
	if cmd.ProcessState == nil {
		return -1
	}
	return cmd.ProcessState.ExitCode()
}

// escapeCommandCharacters 转义命令中的特殊字符
// 参数：
//   - command: 原始命令字符串
//
// 返回值：转义后的命令字符串
func escapeCommandCharacters(command string) string {
	// 转义括号等特殊字符
	replacements := map[string]string{
		"（": "\\（",
		"）": "\\）",
		"(": "\\(",
		")": "\\)",
	}

	for old, new := range replacements {
		command = strings.ReplaceAll(command, old, new)
	}

	return command
}

// DownloadApkFile 通过apk_version_id下载APK文件到本地
// 参数：
//   - ctx: 上下文对象
//   - svcCtx: 服务上下文
//   - apkVersionId: APK版本ID
//
// 返回值：本地文件路径和可能的错误
func DownloadApkFile(ctx context.Context, svcCtx *svc.ServiceContext, apkVersionId int64) (string, error) {
	logc.Info(ctx, "开始下载APK文件", "apkVersionId", apkVersionId)

	// 验证输入参数
	if apkVersionId <= 0 {
		logc.Error(ctx, "APK版本ID无效", "apkVersionId", apkVersionId)
		return "", fmt.Errorf("APK版本ID无效: %d", apkVersionId)
	}

	// 查询APK版本信息
	apkVersionInfo, err := resource.GetApkVersionInfoById(ctx, svcCtx, apkVersionId)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			logc.Error(ctx, "APK版本信息不存在", "apkVersionId", apkVersionId)
			return "", fmt.Errorf("APK版本信息不存在, apkVersionId: %d", apkVersionId)
		}
		logc.Error(ctx, "查询APK版本信息失败", "apkVersionId", apkVersionId, "error", err.Error())
		return "", fmt.Errorf("查询APK版本信息失败: %w", err)
	}

	// 验证APK版本信息完整性
	if err := validateApkVersionInfo(apkVersionInfo); err != nil {
		logc.Error(ctx, "APK版本信息不完整",
			"apkVersionId", apkVersionId,
			"error", err.Error())
		return "", fmt.Errorf("APK版本信息不完整: %w", err)
	}

	// 查询云文件信息
	cloudFileInfo, err := resource.GetCloudFileInfoById(ctx, svcCtx, apkVersionInfo.CloudFileId)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			logc.Error(ctx, "云文件信息不存在",
				"apkVersionId", apkVersionId,
				"cloudFileId", apkVersionInfo.CloudFileId)
			return "", fmt.Errorf("云文件信息不存在, cloudFileId: %s", apkVersionInfo.CloudFileId)
		}
		logc.Error(ctx, "查询云文件信息失败",
			"apkVersionId", apkVersionId,
			"cloudFileId", apkVersionInfo.CloudFileId,
			"error", err.Error())
		return "", fmt.Errorf("查询云文件信息失败: %w", err)
	}

	// 验证云文件信息
	if err := validateCloudFileInfo(cloudFileInfo); err != nil {
		logc.Error(ctx, "云文件信息无效",
			"apkVersionId", apkVersionId,
			"cloudFileId", cloudFileInfo.Id,
			"error", err.Error())
		return "", fmt.Errorf("云文件信息无效: %w", err)
	}

	logc.Info(ctx, "获取到云文件信息",
		"apkVersionId", apkVersionId,
		"cloudFileId", cloudFileInfo.Id,
		"fileUrl", cloudFileInfo.Url,
		"fileSize", cloudFileInfo.Size)

	// 创建下载目录
	downloadPath := getApkDownloadPath(apkVersionInfo.ApkId, apkVersionId)
	if err := MkApkDir(downloadPath); err != nil {
		logc.Error(ctx, "创建下载目录失败",
			"apkVersionId", apkVersionId,
			"directory", downloadPath,
			"error", err.Error())
		return "", fmt.Errorf("创建下载目录失败: %w", err)
	}

	// 构建本地文件路径
	localFilePath := getLocalApkFilePath(downloadPath, cloudFileInfo.Url)

	// 检查文件是否已存在且有效
	if fileExistsAndValid(localFilePath) {
		logc.Info(ctx, "APK文件已存在且有效，跳过下载",
			"apkVersionId", apkVersionId,
			"localPath", localFilePath,
			"fileSize", getFileSizeForLog(localFilePath))
		return localFilePath, nil
	}

	// 下载文件
	if err := FileDownloadWithRetry(ctx, cloudFileInfo.Url, localFilePath); err != nil {
		logc.Error(ctx, "下载APK文件失败",
			"apkVersionId", apkVersionId,
			"url", cloudFileInfo.Url,
			"localPath", localFilePath,
			"error", err.Error())
		return "", fmt.Errorf("下载APK文件失败: %w", err)
	}

	// 验证下载的文件
	if err := validateDownloadedFile(localFilePath, cloudFileInfo.Size); err != nil {
		logc.Error(ctx, "下载文件验证失败",
			"apkVersionId", apkVersionId,
			"localPath", localFilePath,
			"error", err.Error())
		// 清理无效文件
		if removeErr := os.Remove(localFilePath); removeErr != nil {
			logc.Info(ctx, "清理无效文件失败",
				"apkVersionId", apkVersionId,
				"localPath", localFilePath,
				"error", removeErr.Error())
		}
		return "", fmt.Errorf("下载文件验证失败: %w", err)
	}

	logc.Info(ctx, "APK文件下载成功",
		"apkVersionId", apkVersionId,
		"localPath", localFilePath,
		"fileSize", getFileSizeForLog(localFilePath))
	return localFilePath, nil
}

// validateApkVersionInfo 验证APK版本信息的完整性
// 参数：
//   - info: APK版本信息
//
// 返回值：验证过程中发生的错误
func validateApkVersionInfo(info *model.UpgradeApkVersion) error {
	if info == nil {
		return errors.New("APK版本信息为空")
	}

	if info.CloudFileId == "" {
		return errors.New("cloud_file_id为空")
	}

	if info.ApkId <= 0 {
		return errors.New("apk_id无效")
	}

	return nil
}

// validateCloudFileInfo 验证云文件信息的有效性
// 参数：
//   - info: 云文件信息
//
// 返回值：验证过程中发生的错误
func validateCloudFileInfo(info *model.FmsCloudFiles) error {
	if info == nil {
		return errors.New("云文件信息为空")
	}

	if info.Url == "" {
		return errors.New("文件URL为空")
	}

	if !strings.HasPrefix(info.Url, "http://") && !strings.HasPrefix(info.Url, "https://") {
		return fmt.Errorf("文件URL格式无效: %s", info.Url)
	}

	if info.Size <= 0 {
		return errors.New("文件大小无效")
	}

	return nil
}

// fileExistsAndValid 检查文件是否存在且有效
// 参数：
//   - filePath: 文件路径
//
// 返回值：文件是否有效
func fileExistsAndValid(filePath string) bool {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return false
	}

	// 检查是否为文件
	if fileInfo.IsDir() {
		return false
	}

	// 检查文件大小
	if fileInfo.Size() == 0 {
		return false
	}

	// 检查文件权限
	file, err := os.Open(filePath)
	if err != nil {
		return false
	}
	file.Close()

	return true
}

// FileDownloadWithRetry 带重试机制的文件下载
// 参数：
//   - ctx: 上下文对象
//   - url: 文件下载URL
//   - path: 本地保存路径
//
// 返回值：下载过程中发生的错误
func FileDownloadWithRetry(ctx context.Context, url, path string) error {
	const maxRetries = 3
	const retryDelay = 2 * time.Second

	var lastErr error

	for attempt := 1; attempt <= maxRetries; attempt++ {
		logc.Info(ctx, "开始下载文件（重试机制）",
			"attempt", attempt,
			"maxRetries", maxRetries,
			"url", url,
			"path", path)

		if err := FileDownload(url, path); err != nil {
			lastErr = err
			logc.Error(ctx, "文件下载失败，准备重试",
				"attempt", attempt,
				"error", err.Error())

			// 如果不是最后一次尝试，等待后重试
			if attempt < maxRetries {
				time.Sleep(retryDelay)
				continue
			}
		} else {
			logc.Info(ctx, "文件下载成功", "attempt", attempt, "url", url)
			return nil
		}
	}

	return fmt.Errorf("文件下载失败，重试%d次后仍然失败: %w", maxRetries, lastErr)
}

// validateDownloadedFile 验证下载的文件
// 参数：
//   - filePath: 文件路径
//   - expectedSize: 期望的文件大小
//
// 返回值：验证过程中发生的错误
func validateDownloadedFile(filePath string, expectedSize uint64) error {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return fmt.Errorf("无法访问下载的文件: %w", err)
	}

	// 检查是否为文件
	if fileInfo.IsDir() {
		return fmt.Errorf("下载路径是目录而不是文件: %s", filePath)
	}

	// 检查文件大小
	if fileInfo.Size() == 0 {
		return fmt.Errorf("下载文件大小为0: %s", filePath)
	}

	// 如果提供了期望大小，验证文件大小
	if expectedSize > 0 && fileInfo.Size() != int64(expectedSize) {
		return fmt.Errorf("文件大小不匹配，期望: %d, 实际: %d", expectedSize, fileInfo.Size())
	}

	return nil
}

// getFileSizeForLog 获取文件大小用于日志记录
// 参数：
//   - filePath: 文件路径
//
// 返回值：格式化后的文件大小字符串
func getFileSizeForLog(filePath string) string {
	if filePath == "" {
		return "unknown"
	}

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return "unknown"
	}

	return formatFileSize(fileInfo.Size())
}

// formatFileSize 格式化文件大小
// 参数：
//   - size: 文件大小（字节）
//
// 返回值：格式化后的文件大小字符串
func formatFileSize(size int64) string {
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
	)

	switch {
	case size >= GB:
		return fmt.Sprintf("%.2f GB", float64(size)/float64(GB))
	case size >= MB:
		return fmt.Sprintf("%.2f MB", float64(size)/float64(MB))
	case size >= KB:
		return fmt.Sprintf("%.2f KB", float64(size)/float64(KB))
	default:
		return fmt.Sprintf("%d B", size)
	}
}

// getApkDownloadPath 获取APK下载目录路径
// 参数：
//   - apkId: APK ID
//   - versionId: 版本ID
//
// 返回值：完整的下载目录路径
func getApkDownloadPath(apkId, versionId int64) string {
	return fmt.Sprintf("%s%s%d/%d/",
		STORAGE_PATH,
		DOWNLOAD_APK_PATH,
		apkId,
		versionId)
}

// getLocalApkFilePath 获取本地APK文件路径
// 参数：
//   - downloadPath: 下载目录路径
//   - fileUrl: 文件URL
//
// 返回值：完整的本地文件路径
func getLocalApkFilePath(downloadPath, fileUrl string) string {
	fileName := GetFileNameFromUrl(fileUrl)
	return downloadPath + fileName
}

// FileDownload 下载文件到本地路径
// 参数：
//   - url: 文件下载URL
//   - path: 本地保存路径
//
// 返回值：下载过程中发生的错误
func FileDownload(url, path string) error {
	logc.Info(context.Background(), "开始下载文件", "url", url, "path", path)

	// 验证URL格式
	if !isValidURL(url) {
		return fmt.Errorf("URL格式无效: %s", url)
	}

	// 创建临时文件，避免下载过程中文件损坏
	tempPath := path + ".tmp"

	// 发起HTTP请求
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("HTTP请求失败: %w", err)
	}
	defer resp.Body.Close()

	// 检查HTTP状态码
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("下载失败, HTTP状态码: %d", resp.StatusCode)
	}

	// 检查Content-Type
	contentType := resp.Header.Get("Content-Type")
	if !isValidContentType(contentType) {
		logc.Error(context.Background(), "文件Content-Type可能不匹配",
			"url", url,
			"contentType", contentType)
	}

	// 创建临时文件
	out, err := os.Create(tempPath)
	if err != nil {
		return fmt.Errorf("创建临时文件失败: %w", err)
	}

	// 复制文件内容
	written, err := io.Copy(out, resp.Body)
	if err != nil {
		out.Close()
		os.Remove(tempPath) // 清理临时文件
		return fmt.Errorf("复制文件内容失败: %w", err)
	}

	// 关闭文件
	if err := out.Close(); err != nil {
		os.Remove(tempPath) // 清理临时文件
		return fmt.Errorf("关闭临时文件失败: %w", err)
	}

	// 验证下载的文件大小
	if written == 0 {
		os.Remove(tempPath) // 清理临时文件
		return fmt.Errorf("下载文件大小为0")
	}

	// 重命名临时文件为最终文件
	if err := os.Rename(tempPath, path); err != nil {
		os.Remove(tempPath) // 清理临时文件
		return fmt.Errorf("重命名文件失败: %w", err)
	}

	logc.Info(context.Background(), "文件下载成功",
		"url", url,
		"path", path,
		"size", written)
	return nil
}

// isValidURL 验证URL格式是否有效
// 参数：
//   - url: 要验证的URL
//
// 返回值：URL是否有效
func isValidURL(url string) bool {
	if url == "" {
		return false
	}

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return false
	}

	// 简单的URL格式检查
	if len(url) < 10 { // http://a.b 是最短的有效URL
		return false
	}

	return true
}

// isValidContentType 验证Content-Type是否有效
// 参数：
//   - contentType: Content-Type字符串
//
// 返回值：Content-Type是否有效
func isValidContentType(contentType string) bool {
	if contentType == "" {
		return false
	}

	// 允许的Content-Type列表
	allowedTypes := []string{
		"application/vnd.android.package-archive", // APK文件
		"application/octet-stream",                // 二进制文件
		"application/zip",                         // ZIP文件
		"application/x-zip-compressed",            // 压缩文件
	}

	for _, allowedType := range allowedTypes {
		if strings.Contains(contentType, allowedType) {
			return true
		}
	}

	return false
}

// MkApkDir 创建APK存储目录
// 参数：
//   - path: 目录路径
//
// 返回值：创建过程中发生的错误
func MkApkDir(path string) error {
	if path == "" {
		return fmt.Errorf("目录路径为空")
	}

	// 检查路径是否已存在
	if _, err := os.Stat(path); err == nil {
		// 目录已存在，检查权限
		if err := checkDirectoryPermissions(path); err != nil {
			return fmt.Errorf("目录权限检查失败: %w", err)
		}
		return nil
	}

	// 创建目录
	if err := os.MkdirAll(path, 0755); err != nil {
		return fmt.Errorf("创建目录失败: %w", err)
	}

	// 验证目录是否创建成功
	if _, err := os.Stat(path); err != nil {
		return fmt.Errorf("目录创建验证失败: %w", err)
	}

	logc.Info(context.Background(), "目录创建成功", "path", path)
	return nil
}

// checkDirectoryPermissions 检查目录权限
// 参数：
//   - path: 目录路径
//
// 返回值：权限检查过程中发生的错误
func checkDirectoryPermissions(path string) error {
	// 检查目录是否可读
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("目录不可读: %w", err)
	}
	file.Close()

	// 检查目录是否可写
	testFile := path + "/.permission_test"
	if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
		return fmt.Errorf("目录不可写: %w", err)
	}
	os.Remove(testFile)

	return nil
}

// GetFileSize 获取文件大小（字节）
// 参数：
//   - filePath: 文件路径
//
// 返回值：文件大小和可能的错误
func GetFileSize(filePath string) (int64, error) {
	if filePath == "" {
		return 0, fmt.Errorf("文件路径为空")
	}

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return 0, fmt.Errorf("文件不存在: %s", filePath)
		}
		return 0, fmt.Errorf("获取文件信息失败: %w", err)
	}

	// 检查是否为文件
	if fileInfo.IsDir() {
		return 0, fmt.Errorf("路径是目录而不是文件: %s", filePath)
	}

	return fileInfo.Size(), nil
}

// GetFileNameFromUrl 从URL中提取文件名
// 参数：
//   - url: 文件URL
//
// 返回值：提取的文件名，如果无法提取则返回MD5值
func GetFileNameFromUrl(url string) string {
	if url == "" {
		return "unknown_file"
	}

	// 分割URL路径
	pathSegments := strings.Split(url, "/")
	if len(pathSegments) == 0 {
		return md5V(url)
	}

	// 获取最后一个路径段
	lastSegment := pathSegments[len(pathSegments)-1]

	// 去除查询参数
	fileName := strings.Split(lastSegment, "?")[0]

	// 检查文件名是否有效
	if isValidFileName(fileName) {
		return fileName
	}

	// 如果最后一个路径段为空，尝试使用倒数第二个
	if len(pathSegments) > 1 {
		secondLastSegment := pathSegments[len(pathSegments)-2]
		fileName = strings.Split(secondLastSegment, "?")[0]
		if isValidFileName(fileName) {
			return fileName
		}
	}

	// 如果无法提取有效文件名，使用URL的MD5作为文件名
	return md5V(url)
}

// isValidFileName 验证文件名是否有效
// 参数：
//   - fileName: 文件名
//
// 返回值：文件名是否有效
func isValidFileName(fileName string) bool {
	if fileName == "" {
		return false
	}

	// 检查文件名长度
	if len(fileName) > 255 {
		return false
	}

	// 检查是否包含非法字符
	illegalChars := []string{"/", "\\", ":", "*", "?", "\"", "<", ">", "|"}
	for _, char := range illegalChars {
		if strings.Contains(fileName, char) {
			return false
		}
	}

	// 检查是否为系统保留名称
	reservedNames := []string{"CON", "PRN", "AUX", "NUL",
		"COM1", "COM2", "COM3", "COM4", "COM5", "COM6", "COM7", "COM8", "COM9",
		"LPT1", "LPT2", "LPT3", "LPT4", "LPT5", "LPT6", "LPT7", "LPT8", "LPT9"}

	baseName := strings.ToUpper(strings.TrimSuffix(fileName, ".apk"))
	for _, reserved := range reservedNames {
		if baseName == reserved {
			return false
		}
	}

	return true
}

// md5V 计算字符串的MD5值
// 参数：
//   - str: 输入字符串
//
// 返回值：MD5哈希值的十六进制字符串
func md5V(str string) string {
	if str == "" {
		return "empty_string_md5"
	}

	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
