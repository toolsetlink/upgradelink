package menu

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/enum"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/ent/menu"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuLogic {
	return &UpdateMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMenuLogic) UpdateMenu(req *types.MenuPlainInfo) (resp *types.BaseMsgResp, err error) {
	// get parent level
	var menuLevel uint32
	if *req.ParentId != enum.DefaultParentId {
		m, err := l.svcCtx.DB.Menu.Query().Where(menu.IDEQ(*req.ParentId)).First(l.ctx)
		if err != nil {
			return nil, db_error.DefaultEntError(l.Logger, err, req)
		}

		menuLevel = m.MenuLevel + 1
	} else {
		menuLevel = 1
	}

	err = l.svcCtx.DB.Menu.UpdateOneID(*req.Id).
		SetNotNilMenuLevel(&menuLevel).
		SetNotNilMenuType(req.MenuType).
		SetNotNilParentID(req.ParentId).
		SetNotNilPath(req.Path).
		SetNotNilName(req.Name).
		SetNotNilRedirect(req.Redirect).
		SetNotNilComponent(req.Component).
		SetNotNilSort(req.Sort).
		SetNotNilDisabled(req.Disabled).
		SetNotNilServiceName(req.ServiceName).
		SetNotNilPermission(req.Permission).
		// meta
		SetNotNilTitle(req.Title).
		SetNotNilIcon(req.Icon).
		SetNotNilHideMenu(req.HideMenu).
		SetNotNilHideBreadcrumb(req.HideBreadcrumb).
		SetNotNilIgnoreKeepAlive(req.IgnoreKeepAlive).
		SetNotNilHideTab(req.HideTab).
		SetNotNilFrameSrc(req.FrameSrc).
		SetNotNilCarryParam(req.CarryParam).
		SetNotNilHideChildrenInMenu(req.HideChildrenInMenu).
		SetNotNilAffix(req.Affix).
		SetNotNilDynamicLevel(req.DynamicLevel).
		SetNotNilRealPath(req.RealPath).
		Exec(l.ctx)

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.UpdateSuccess)}, nil
}
