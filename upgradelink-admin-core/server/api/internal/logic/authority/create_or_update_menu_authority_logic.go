package authority

import (
	"context"

	"upgradelink-admin-core/server/api/internal/svc"
	"upgradelink-admin-core/server/api/internal/types"
	"upgradelink-admin-core/server/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateMenuAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrUpdateMenuAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateMenuAuthorityLogic {
	return &CreateOrUpdateMenuAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrUpdateMenuAuthorityLogic) CreateOrUpdateMenuAuthority(req *types.MenuAuthorityInfoReq) (resp *types.BaseMsgResp, err error) {
	authority, err := l.svcCtx.CoreRpc.CreateOrUpdateMenuAuthority(l.ctx, &core.RoleMenuAuthorityReq{
		RoleId:  req.RoleId,
		MenuIds: req.MenuIds,
	})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, authority.Msg)}, nil
}
