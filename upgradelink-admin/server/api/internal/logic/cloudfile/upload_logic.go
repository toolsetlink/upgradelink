package cloudfile

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
	"time"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/http_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/utils/cloudstorage"
	"upgradelink-admin/server/api/internal/common/utils/filex"
	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/common/utils/uuidx"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/duke-git/lancet/v2/datetime"
	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUploadLogic(ctx context.Context, r *http.Request, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *UploadLogic) Upload() (resp *types.CloudFileInfoResp, err error) {

	fmt.Println("in Upload")

	file, handler, err := l.r.FormFile("file")
	if err != nil {
		logx.Error("the value of file cannot be found")
		return nil, http_error.NewApiBadRequestError("file.parseFormFailed")
	}
	defer file.Close()

	// 计算文件 MD5
	md5Hash := md5.New()
	if _, err := io.Copy(md5Hash, file); err != nil {
		logx.Errorw("计算文件 MD5 失败", logx.Field("error", err), logx.Field("fileName", handler.Filename))
		return nil, http_error.NewApiInternalError("file.calculateMD5Failed")
	}
	md5Sum := hex.EncodeToString(md5Hash.Sum(nil)) // 转为 32 位十六进制字符串
	//fmt.Println("md5Sum:", md5Sum)

	// 重置文件指针到开头（否则后续上传会读取空内容）
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		logx.Errorw("重置文件指针失败", logx.Field("error", err))
		return nil, http_error.NewApiInternalError("file.resetPointerFailed")
	}

	// judge if the suffix is legal
	// 校验后缀是否合法
	dotIndex := strings.LastIndex(handler.Filename, ".")
	// if there is no suffix, reject it
	// 拒绝无后缀文件
	if dotIndex == -1 {
		logx.Errorw("reject the file which does not have suffix")
		return nil, http_error.NewApiBadRequestError("file.wrongTypeError")
	}

	fileName, fileSuffix := handler.Filename[:dotIndex], handler.Filename[dotIndex+1:]
	fileUUID := uuidx.NewUUID()
	storeFileName := fileUUID.String() + "." + fileSuffix
	userId := l.ctx.Value("userId").(string)

	// judge if the file size is over max size
	fileType := strings.Split(handler.Header.Get("Content-Type"), "/")[0]
	if fileType != "image" && fileType != "video" && fileType != "audio" {
		fileType = "other"
	}

	relativeSrc := fmt.Sprintf("%s/%s/%s/%s",
		l.svcCtx.Config.UploadConf.Folder,
		datetime.FormatTimeToStr(time.Now(), "yyyy-mm-dd"),
		fileType,
		storeFileName)

	url, err := l.UploadToProvider(file, relativeSrc)
	if err != nil {
		return nil, err
	}

	if l.svcCtx.Config.UploadConf.CdnUrl != "" {
		url = fmt.Sprintf("%s%s", l.svcCtx.Config.UploadConf.CdnUrl, relativeSrc)
	}

	// store to database
	query := l.svcCtx.DB.FmsCloudFile.Create().
		SetID(fileUUID.String()).
		SetName(fileName).
		SetFileType(filex.ConvertFileTypeToUint8(fileType)).
		SetURL(url).
		SetSize(uint64(handler.Size)).
		SetMd5(md5Sum).
		SetUserID(userId)

	data, err := query.Save(l.ctx)
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, nil)
	}

	return &types.CloudFileInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
			Data: "",
		},
		Data: types.CloudFileInfo{
			BaseUUIDInfo: types.BaseUUIDInfo{
				Id:        pointy.GetPointer(data.ID),
				CreatedAt: pointy.GetPointer(data.CreatedAt.UnixMilli()),
			},
			State:       pointy.GetPointer(data.State),
			Name:        pointy.GetPointer(data.Name),
			Url:         pointy.GetPointer(data.URL),
			RelativeSrc: pointy.GetPointer(relativeSrc),
			Size:        pointy.GetPointer(data.Size),
			FileType:    pointy.GetPointer(data.FileType),
			UserId:      pointy.GetPointer(data.UserID),
		},
	}, nil
}

func (l *UploadLogic) UploadToProvider(file multipart.File, fileName string) (url string, err error) {
	// 创建 S3 服务实例
	s3Config := &cloudstorage.S3Config{
		Bucket:    l.svcCtx.Config.UploadConf.Bucket,
		SecretID:  l.svcCtx.Config.UploadConf.SecretID,
		SecretKey: l.svcCtx.Config.UploadConf.SecretKey,
		Endpoint:  l.svcCtx.Config.UploadConf.Endpoint,
		Region:    l.svcCtx.Config.UploadConf.Region,
	}

	s3Service := cloudstorage.NewS3Service(s3Config)

	// 上传文件
	uploadedURL, err := s3Service.UploadFile(l.ctx, file, fileName)
	if err != nil {
		logx.Errorw("failed to upload object", logx.Field("detail", err))
		return url, http_error.NewApiInternalError(err.Error())
	}

	return uploadedURL, nil
}
