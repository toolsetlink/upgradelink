package role

import (
	"context"

	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"upgradelink-admin-core/server/rpc/internal/svc"
	"upgradelink-admin-core/server/rpc/internal/utils/dberrorhandler"
	"upgradelink-admin-core/server/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoleByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleByIdLogic {
	return &GetRoleByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRoleByIdLogic) GetRoleById(in *core.IDReq) (*core.RoleInfo, error) {
	result, err := l.svcCtx.DB.Role.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.RoleInfo{
		Id:            &result.ID,
		CreatedAt:     pointy.GetPointer(result.CreatedAt.UnixMilli()),
		UpdatedAt:     pointy.GetPointer(result.UpdatedAt.UnixMilli()),
		Status:        pointy.GetPointer(uint32(result.Status)),
		Name:          &result.Name,
		Code:          &result.Code,
		DefaultRouter: &result.DefaultRouter,
		Remark:        &result.Remark,
		Sort:          &result.Sort,
	}, nil
}
