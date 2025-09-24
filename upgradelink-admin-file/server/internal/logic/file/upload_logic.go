package file

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/duke-git/lancet/v2/datetime"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/suyuan32/simple-admin-common/enum/errorcode"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"

	"upgradelink-admin-file/server/internal/svc"
	"upgradelink-admin-file/server/internal/types"
	"upgradelink-admin-file/server/internal/utils/dberrorhandler"
	"upgradelink-admin-file/server/internal/utils/filex"
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

func (l *UploadLogic) Upload() (resp *types.UploadResp, err error) {
	err = l.r.ParseMultipartForm(l.svcCtx.Config.UploadConf.MaxVideoSize)
	if err != nil {
		logx.Error("fail to parse the multipart form")
		return nil, errorx.NewCodeInvalidArgumentError("file.parseFormFailed")
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
	timeString := datetime.FormatTimeToStr(time.Now(), "yyyy-mm-dd")
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

	// generate path
	// 生成路径

	//judge if the file directory exists, if not create it. Both private and public need
	//to be created in order to move files when status changed
	//判断文件夹是否已创建, 同时创建好私人和公开文件夹防止文件状态改变时无法移动

	publicStoreDir := path.Join(l.svcCtx.Config.UploadConf.PublicStorePath,
		l.svcCtx.Config.Name, fileType, timeString)
	privateStoreDir := path.Join(l.svcCtx.Config.UploadConf.PrivateStorePath,
		l.svcCtx.Config.Name, fileType, timeString)

	if !fileutil.IsExist(publicStoreDir) {
		err = fileutil.CreateDir(publicStoreDir + "/")
		if err != nil {
			logx.Errorw("failed to create directory for storing public files", logx.Field("path", publicStoreDir))
			return nil, errorx.NewCodeError(errorcode.Internal,
				l.svcCtx.Trans.Trans(l.ctx, i18n.Failed))
		}
	}

	if !fileutil.IsExist(privateStoreDir) {
		err = fileutil.CreateDir(privateStoreDir + "/")
		if err != nil {
			logx.Errorw("failed to create directory for storing private files", logx.Field("path", privateStoreDir))
			return nil, errorx.NewCodeError(errorcode.Internal,
				l.svcCtx.Trans.Trans(l.ctx, i18n.Failed))
		}
	}

	// default is public
	// 默认是公开的
	targetFile, err := os.Create(path.Join(publicStoreDir, storeFileName))
	_, err = io.Copy(targetFile, file)
	if err != nil {
		logx.Errorw("fail to create file", logx.Field("path", path.Join(publicStoreDir, storeFileName)))
		return nil, errorx.NewCodeError(errorcode.Internal,
			l.svcCtx.Trans.Trans(l.ctx, i18n.Failed))
	}

	// store in database
	// 提交数据库
	relativePath := fmt.Sprintf("/%s/%s/%s/%s", l.svcCtx.Config.Name,
		fileType, timeString, storeFileName)

	var md5 string
	if l.r.MultipartForm.Value["md5"] != nil {
		md5 = l.r.MultipartForm.Value["md5"][0]
	} else {
		md5 = ""
	}

	var fileTagId uint64
	if l.r.MultipartForm.Value["tagId"] != nil && l.r.MultipartForm.Value["tagId"][0] != "" {
		tagId, err := strconv.Atoi(l.r.MultipartForm.Value["tagId"][0])
		if err != nil {
			return nil, errorx.NewCodeInvalidArgumentError("wrong tag ID")
		}

		fileTagId = uint64(tagId)
	}

	query := l.svcCtx.DB.File.Create().
		SetID(fileUUID).
		SetNotNilName(&fileName).
		SetNotNilFileType(pointy.GetPointer(filex.ConvertFileTypeToUint8(fileType))).
		SetNotNilPath(&relativePath).
		SetNotNilUserID(&userId).
		SetNotNilMd5(pointy.GetPointer(md5)).
		SetNotNilStatus(pointy.GetPointer(uint8(1))).
		SetNotNilSize(pointy.GetPointer(uint64(handler.Size)))

	if fileTagId != 0 {
		query = query.AddTagIDs(fileTagId)
	}

	err = query.Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, "upload failed")
	}

	return &types.UploadResp{
		BaseDataInfo: types.BaseDataInfo{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.Success)},
		Data: types.UploadInfo{Name: handler.Filename, Url: l.svcCtx.Config.UploadConf.ServerURL +
			relativePath},
	}, nil
}
