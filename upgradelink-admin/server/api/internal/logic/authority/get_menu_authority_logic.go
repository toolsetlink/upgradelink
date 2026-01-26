package authority

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/ent/role"

	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuAuthorityLogic {
	return &GetMenuAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuAuthorityLogic) GetMenuAuthority(req *types.IDReq) (resp *types.MenuAuthorityInfoResp, err error) {
	menus, err := l.svcCtx.DB.Role.Query().Where(role.ID(req.Id)).QueryMenus().IDs(l.ctx)
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	resp = &types.MenuAuthorityInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.MenuAuthorityInfoReq{
			RoleId:  req.Id,
			MenuIds: menus,
		},
	}

	return resp, nil
}
