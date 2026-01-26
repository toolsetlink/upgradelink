package menu

import (
	"context"
	"fmt"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/enum"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/ent"
	"upgradelink-admin/server/api/internal/ent/menu"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMenuLogic {
	return &CreateMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateMenuLogic) CreateMenu(req *types.MenuPlainInfo) (resp *types.BaseMsgResp, err error) {
	if *req.MenuType == 0 {
		req.Component = pointy.GetPointer("LAYOUT")
		req.Redirect = pointy.GetPointer("")
		req.FrameSrc = pointy.GetPointer("")
	}

	// if exists , return success
	if req.Name != nil && req.Component != nil && req.Path != nil {
		check, err := l.svcCtx.DB.Menu.Query().Where(menu.Name(*req.Name), menu.Component(*req.Component), menu.Path(*req.Path)).Only(l.ctx)
		if err != nil && !ent.IsNotFound(err) {
			return nil, db_error.DefaultEntError(l.Logger, err, req)
		}

		if check != nil {
			return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.CreateSuccess)}, nil
		}
	}

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

	result, err := l.svcCtx.DB.Menu.Create().
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
		Save(l.ctx)
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	fmt.Println(result.ID)

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.CreateSuccess)}, nil

}
