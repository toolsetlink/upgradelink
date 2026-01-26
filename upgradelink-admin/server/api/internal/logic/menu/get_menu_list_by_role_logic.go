package menu

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/rolectx"
	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/ent"
	"upgradelink-admin/server/api/internal/ent/menu"
	"upgradelink-admin/server/api/internal/ent/role"

	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuListByRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuListByRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuListByRoleLogic {
	return &GetMenuListByRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuListByRoleLogic) GetMenuListByRole() (resp *types.MenuListResp, err error) {

	roleIds, err := rolectx.GetRoleIDFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}

	roles, err := l.svcCtx.DB.Role.Query().Where(role.CodeIn(roleIds...)).WithMenus(func(query *ent.MenuQuery) {
		query.Order(ent.Asc(menu.FieldSort))
		query.Where(menu.Disabled(false))
	}).All(l.ctx)

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, "")
	}

	resp = &types.MenuListResp{}
	resp.Data.Total = uint64(len(roles))

	existMap := map[uint64]struct{}{}

	for _, r := range roles {
		for _, m := range r.Edges.Menus {
			createdAt := m.CreatedAt.Unix()
			updatedAt := m.UpdatedAt.Unix()

			if _, ok := existMap[m.ID]; !ok {
				resp.Data.Data = append(resp.Data.Data, types.MenuInfo{
					BaseIDInfo: types.BaseIDInfo{
						Id:        &m.ID,
						CreatedAt: &createdAt,
						UpdatedAt: &updatedAt,
					},
					MenuType:  &m.MenuType,
					Level:     &m.MenuLevel,
					Path:      &m.Path,
					Name:      &m.Name,
					Redirect:  &m.Redirect,
					Component: &m.Component,
					Sort:      &m.Sort,
					ParentId:  &m.ParentID,
					Meta: types.Meta{
						Title:              pointy.GetPointer(l.svcCtx.Trans.Trans(l.ctx, m.Title)),
						Icon:               &m.Icon,
						HideMenu:           &m.HideMenu,
						HideBreadcrumb:     &m.HideBreadcrumb,
						IgnoreKeepAlive:    &m.IgnoreKeepAlive,
						HideTab:            &m.HideTab,
						FrameSrc:           &m.FrameSrc,
						CarryParam:         &m.CarryParam,
						HideChildrenInMenu: &m.HideChildrenInMenu,
						Affix:              &m.Affix,
						DynamicLevel:       &m.DynamicLevel,
						RealPath:           &m.RealPath,
					},
					Disabled:    &m.Disabled,
					ServiceName: &m.ServiceName,
					Permission:  &m.Permission,
				})
			}

		}

	}

	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	return resp, nil
}
