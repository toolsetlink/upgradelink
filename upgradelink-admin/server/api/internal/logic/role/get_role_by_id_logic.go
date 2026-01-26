package role

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoleByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleByIdLogic {
	return &GetRoleByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoleByIdLogic) GetRoleById(req *types.IDReq) (resp *types.RoleInfoResp, err error) {
	result, err := l.svcCtx.DB.Role.Get(l.ctx, req.Id)
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.RoleInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.RoleInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id:        pointy.GetPointer(result.ID),
				CreatedAt: pointy.GetPointer(result.CreatedAt.UnixMilli()),
				UpdatedAt: pointy.GetPointer(result.UpdatedAt.UnixMilli()),
			},
			Status:        pointy.GetPointer(uint32(result.Status)),
			Name:          &result.Name,
			Code:          &result.Code,
			DefaultRouter: &result.DefaultRouter,
			Remark:        &result.Remark,
			Sort:          pointy.GetPointer(result.Sort),
		},
	}, nil
}
