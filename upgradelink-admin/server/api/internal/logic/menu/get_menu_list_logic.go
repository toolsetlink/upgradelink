package menu

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/ent"
	"upgradelink-admin/server/api/internal/ent/menu"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuListLogic {
	return &GetMenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMenuListLogic) GetMenuList() (resp *types.MenuPlainInfoListResp, err error) {
	menus, err := l.svcCtx.DB.Menu.Query().Page(l.ctx, 1, 100, func(pager *ent.MenuPager) {
		pager.Order = ent.Asc(menu.FieldSort)
	})
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, "")
	}

	resp = &types.MenuPlainInfoListResp{}
	resp.Data.Total = menus.PageDetails.Total
	for _, v := range menus.List {
		resp.Data.Data = append(resp.Data.Data, types.MenuPlainInfo{
			Id:                 &v.ID,
			CreatedAt:          pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:          pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Trans:              pointy.GetPointer(l.svcCtx.Trans.Trans(l.ctx, v.Title)),
			MenuType:           &v.MenuType,
			Level:              &v.MenuLevel,
			Path:               &v.Path,
			Name:               &v.Name,
			Redirect:           &v.Redirect,
			Component:          &v.Component,
			Sort:               &v.Sort,
			ParentId:           &v.ParentID,
			Title:              &v.Title,
			Icon:               &v.Icon,
			HideMenu:           &v.HideMenu,
			HideBreadcrumb:     &v.HideBreadcrumb,
			IgnoreKeepAlive:    &v.IgnoreKeepAlive,
			HideTab:            &v.HideTab,
			FrameSrc:           &v.FrameSrc,
			CarryParam:         &v.CarryParam,
			HideChildrenInMenu: &v.HideChildrenInMenu,
			Affix:              &v.Affix,
			DynamicLevel:       &v.DynamicLevel,
			RealPath:           &v.RealPath,
			Disabled:           &v.Disabled,
			ServiceName:        &v.ServiceName,
			Permission:         &v.Permission,
		})
	}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	return resp, nil
}
