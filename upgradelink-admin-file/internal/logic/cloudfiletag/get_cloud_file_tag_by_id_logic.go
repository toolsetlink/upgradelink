package cloudfiletag

import (
	"context"

	"upgradelink-admin-file/internal/svc"
	"upgradelink-admin-file/internal/types"
	"upgradelink-admin-file/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCloudFileTagByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCloudFileTagByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCloudFileTagByIdLogic {
	return &GetCloudFileTagByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCloudFileTagByIdLogic) GetCloudFileTagById(req *types.IDReq) (*types.CloudFileTagInfoResp, error) {
	data, err := l.svcCtx.DB.CloudFileTag.Get(l.ctx, req.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.CloudFileTagInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.CloudFileTagInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id:        &data.ID,
				CreatedAt: pointy.GetPointer(data.CreatedAt.UnixMilli()),
				UpdatedAt: pointy.GetPointer(data.UpdatedAt.UnixMilli()),
			},
			Status: &data.Status,
			Name:   &data.Name,
			Remark: &data.Remark,
		},
	}, nil
}
