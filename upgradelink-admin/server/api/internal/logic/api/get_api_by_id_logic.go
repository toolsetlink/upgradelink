package api

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApiByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApiByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApiByIdLogic {
	return &GetApiByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApiByIdLogic) GetApiById(req *types.IDReq) (resp *types.ApiInfoResp, err error) {
	result, err := l.svcCtx.DB.API.Get(l.ctx, req.Id)
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.ApiInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.ApiInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id:        &result.ID,
				CreatedAt: pointy.GetPointer(result.CreatedAt.UnixMilli()),
				UpdatedAt: pointy.GetPointer(result.UpdatedAt.UnixMilli()),
			},
			Trans:       l.svcCtx.Trans.Trans(l.ctx, result.Description),
			Path:        &result.Path,
			Description: &result.Description,
			Group:       &result.APIGroup,
			Method:      &result.Method,
			IsRequired:  &result.IsRequired,
			ServiceName: &result.ServiceName,
		},
	}, nil
}
