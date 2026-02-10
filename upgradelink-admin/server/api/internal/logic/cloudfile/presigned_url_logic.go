package cloudfile

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"upgradelink-admin/server/api/internal/common/http_error"
	"upgradelink-admin/server/api/internal/common/utils/cloudstorage"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/duke-git/lancet/v2/datetime"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type PresignedUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewPresignedUrlLogic(ctx context.Context, r *http.Request, svcCtx *svc.ServiceContext) *PresignedUrlLogic {
	return &PresignedUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *PresignedUrlLogic) GeneratePresignedUrl() (resp *types.PresignedUrlResp, err error) {
	// 解析请求参数
	var req types.PresignedUrlReq
	if err := httpx.ParseJsonBody(l.r, &req); err != nil {
		logx.Errorw("解析请求参数失败", logx.Field("error", err))
		return nil, http_error.NewApiBadRequestError("request.parseFailed")
	}

	// 校验文件名
	if req.Filename == "" {
		logx.Errorw("文件名不能为空")
		return nil, http_error.NewApiBadRequestError("file.nameRequired")
	}

	// 校验后缀是否合法
	dotIndex := strings.LastIndex(req.Filename, ".")
	if dotIndex == -1 {
		logx.Errorw("拒绝无后缀文件")
		return nil, http_error.NewApiBadRequestError("file.wrongTypeError")
	}

	// 生成文件信息
	fileName, fileSuffix := req.Filename[:dotIndex], req.Filename[dotIndex+1:]
	// 生成带时间戳的文件名
	timestamp := time.Now().Format("20060102150405")
	storeFileName := fmt.Sprintf("%s_%s.%s", timestamp, fileName, fileSuffix)

	// 确定文件类型
	fileType := "other"
	if req.ContentType != "" {
		contentTypeParts := strings.Split(req.ContentType, "/")
		if len(contentTypeParts) > 0 {
			subType := contentTypeParts[0]
			if subType == "image" || subType == "video" || subType == "audio" {
				fileType = subType
			}
		}
	}

	// 生成存储路径
	relativeSrc := fmt.Sprintf("%s/%s/%s/%s",
		l.svcCtx.Config.UploadConf.Folder,
		datetime.FormatTimeToStr(time.Now(), "yyyy-mm-dd"),
		fileType,
		storeFileName)

	// 生成预签名URL
	presignedURL, err := l.GeneratePresignedUrlToProvider(relativeSrc)
	if err != nil {
		return nil, err
	}

	// 构建文件访问URL
	fileURL := fmt.Sprintf("https://%s.%s%s", l.svcCtx.Config.UploadConf.Bucket, l.svcCtx.Config.UploadConf.Endpoint, relativeSrc)

	return &types.PresignedUrlResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  "success",
			Data: "",
		},
		Data: types.PresignedUrlInfo{
			PresignedUrl: presignedURL,
			FileUrl:      fileURL,
			FileName:     fileName,
			FileSuffix:   fileSuffix,
			FileType:     fileType,
		},
	}, nil
}

func (l *PresignedUrlLogic) GeneratePresignedUrlToProvider(key string) (string, error) {
	// 创建 S3 服务实例
	s3Config := &cloudstorage.S3Config{
		Bucket:    l.svcCtx.Config.UploadConf.Bucket,
		SecretID:  l.svcCtx.Config.UploadConf.SecretID,
		SecretKey: l.svcCtx.Config.UploadConf.SecretKey,
		Endpoint:  l.svcCtx.Config.UploadConf.Endpoint,
		Region:    l.svcCtx.Config.UploadConf.Region,
	}

	s3Service := cloudstorage.NewS3Service(s3Config)

	// 生成预签名URL，设置1小时过期
	expiry := 1 * time.Hour
	presignedURL, err := s3Service.GeneratePresignedURL(l.ctx, key, expiry)
	if err != nil {
		logx.Errorw("生成预签名URL失败", logx.Field("detail", err))
		return "", http_error.NewApiInternalError("file.generatePresignedUrlFailed")
	}

	return presignedURL, nil
}
