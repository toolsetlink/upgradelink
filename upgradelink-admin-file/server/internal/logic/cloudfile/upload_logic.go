package cloudfile

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/duke-git/lancet/v2/datetime"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/errorx"

	"upgradelink-admin-file/server/internal/svc"
	"upgradelink-admin-file/server/internal/types"
	"upgradelink-admin-file/server/internal/utils/dberrorhandler"
	"upgradelink-admin-file/server/internal/utils/filex"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUploadLogic(r *http.Request, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *UploadLogic) Upload() (resp *types.CloudFileInfoResp, err error) {
	err = l.r.ParseMultipartForm(l.svcCtx.Config.UploadConf.MaxVideoSize)
	if err != nil {
		logx.Error("fail to parse the multipart form")
		return nil, errorx.NewCodeInvalidArgumentError(
			"file.parseFormFailed")
	}

	file, handler, err := l.r.FormFile("file")
	if err != nil {
		logx.Error("the value of file cannot be found")
		return nil, errorx.NewCodeInvalidArgumentError("file.parseFormFailed")
	}
	defer file.Close()

	// judge if the suffix is legal
	// 校验后缀是否合法
	dotIndex := strings.LastIndex(handler.Filename, ".")
	// if there is no suffix, reject it
	// 拒绝无后缀文件
	if dotIndex == -1 {
		logx.Errorw("reject the file which does not have suffix")
		return nil, errorx.NewCodeInvalidArgumentError("file.wrongTypeError")
	}

	fileName, fileSuffix := handler.Filename[:dotIndex], handler.Filename[dotIndex+1:]
	fileUUID := uuidx.NewUUID()
	storeFileName := fileUUID.String() + "." + fileSuffix
	userId := l.ctx.Value("userId").(string)

	// judge if the file size is over max size
	// 判断文件大小是否超过设定值
	fileType := strings.Split(handler.Header.Get("Content-Type"), "/")[0]
	if fileType != "image" && fileType != "video" && fileType != "audio" {
		fileType = "other"
	}
	err = filex.CheckOverSize(l.ctx, l.svcCtx, fileType, handler.Size)
	if err != nil {
		logx.Errorw("the file is over size", logx.Field("type", fileType),
			logx.Field("userId", userId), logx.Field("size", handler.Size),
			logx.Field("fileName", handler.Filename))
		return nil, err
	}

	var provider string
	if l.r.MultipartForm.Value["provider"] != nil && l.r.MultipartForm.Value["provider"][0] != "" {
		provider = l.r.MultipartForm.Value["provider"][0]
	} else {
		provider = l.svcCtx.CloudStorage.DefaultProvider
	}

	var fileTagId uint64
	if l.r.MultipartForm.Value["tagId"] != nil && l.r.MultipartForm.Value["tagId"][0] != "" {
		tagId, err := strconv.Atoi(l.r.MultipartForm.Value["tagId"][0])
		if err != nil {
			return nil, errorx.NewCodeInvalidArgumentError("wrong tag ID")
		}

		fileTagId = uint64(tagId)
	}

	relativeSrc := fmt.Sprintf("%s/%s/%s/%s",
		l.svcCtx.CloudStorage.ProviderData[provider].Folder,
		datetime.FormatTimeToStr(time.Now(), "yyyy-mm-dd"),
		fileType,
		storeFileName)

	url, err := l.UploadToProvider(file, relativeSrc, provider)
	if err != nil {
		return nil, err
	}

	if l.svcCtx.CloudStorage.ProviderData[provider].UseCdn {
		url = fmt.Sprintf("%s%s", l.svcCtx.CloudStorage.ProviderData[provider].CdnUrl, relativeSrc)
	}

	// store to database
	query := l.svcCtx.DB.CloudFile.Create().
		SetID(fileUUID).
		SetName(fileName).
		SetFileType(filex.ConvertFileTypeToUint8(fileType)).
		SetStorageProvidersID(l.svcCtx.CloudStorage.ProviderData[provider].Id).
		SetURL(url).
		SetSize(uint64(handler.Size)).
		SetUserID(userId)

	if fileTagId != 0 {
		query = query.AddTagIDs(fileTagId)
	}

	data, err := query.Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, nil)
	}

	return &types.CloudFileInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  i18n.Success,
			Data: "",
		},
		Data: types.CloudFileInfo{
			BaseUUIDInfo: types.BaseUUIDInfo{
				Id:        pointy.GetPointer(data.ID.String()),
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

func (l *UploadLogic) UploadToProvider(file multipart.File, fileName, provider string) (url string, err error) {
	if client, ok := l.svcCtx.CloudStorage.CloudStorage[provider]; ok {
		_, err := client.PutObjectWithContext(l.ctx, &s3.PutObjectInput{
			Bucket: aws.String(l.svcCtx.CloudStorage.ProviderData[provider].Bucket),
			Key:    aws.String(fileName),
			Body:   file,
		})
		if err != nil {
			logx.Errorw("failed to upload object", logx.Field("detail", err))
			var aerr awserr.Error
			if errors.As(err, &aerr) && aerr.Code() == request.CanceledErrorCode {
				return url, errorx.NewCodeInternalError("upload canceled due to timeout")
			} else {
				return url, errorx.NewCodeInternalError("failed to upload object")
			}
		}

		return fmt.Sprintf("https://%s.%s%s",
			l.svcCtx.CloudStorage.ProviderData[provider].Bucket,
			l.svcCtx.CloudStorage.ProviderData[provider].Endpoint, fileName), nil
	}

	return url, nil
}
