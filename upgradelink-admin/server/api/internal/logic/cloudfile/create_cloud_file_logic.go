package cloudfile

import (
	"context"
	"fmt"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/utils/filex"
	"upgradelink-admin/server/api/internal/common/utils/uuidx"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCloudFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCloudFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCloudFileLogic {
	return &CreateCloudFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCloudFileLogic) CreateCloudFile(req *types.CloudFileInfo) (resp *types.CloudFileInfoResp, err error) {

	fileUUID := uuidx.NewUUID()

	// store to database
	query := l.svcCtx.DB.FmsCloudFile.Create().
		SetID(fileUUID.String()).
		SetName(*req.Name).
		SetFileType(filex.ConvertFileTypeToUint8(*req.FileType)).
		SetURL(*req.Url).
		SetSize(*req.Size).
		SetMd5(*req.Md5).
		SetUserID(l.ctx.Value("userId").(string))

	data, err := query.Save(l.ctx)
	if err != nil {
		fmt.Println("err", err)
		logx.Errorw("创建云文件记录失败", logx.Field("error", err), logx.Field("fileName", req.Name))
		return nil, db_error.DefaultEntError(l.Logger, err, nil)
	}

	return &types.CloudFileInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  "success",
			Data: "",
		},
		Data: types.CloudFileInfo{
			BaseUUIDInfo: types.BaseUUIDInfo{
				Id: &data.ID,
			},
		},
	}, nil
}
