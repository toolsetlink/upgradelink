package role

import (
	"context"

	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"upgradelink-admin-core/server/rpc/internal/svc"
	"upgradelink-admin-core/server/rpc/internal/utils/dberrorhandler"
	"upgradelink-admin-core/server/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-common/i18n"
)

type CreateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoleLogic {
	return &CreateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateRoleLogic) CreateRole(in *core.RoleInfo) (*core.BaseIDResp, error) {
	result, err := l.svcCtx.DB.Role.Create().
		SetNotNilStatus(pointy.GetStatusPointer(in.Status)).
		SetNotNilName(in.Name).
		SetNotNilCode(in.Code).
		SetNotNilDefaultRouter(in.DefaultRouter).
		SetNotNilRemark(in.Remark).
		SetNotNilSort(in.Sort).
		Save(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil
}
