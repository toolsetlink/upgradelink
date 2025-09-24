package filetag

import (
	"context"

	"upgradelink-admin-file/internal/svc"
	"upgradelink-admin-file/internal/types"
	"upgradelink-admin-file/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetFileTagByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFileTagByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileTagByIdLogic {
	return &GetFileTagByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFileTagByIdLogic) GetFileTagById(req *types.IDReq) (*types.FileTagInfoResp, error) {
	data, err := l.svcCtx.DB.FileTag.Get(l.ctx, req.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.FileTagInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.FileTagInfo{
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
