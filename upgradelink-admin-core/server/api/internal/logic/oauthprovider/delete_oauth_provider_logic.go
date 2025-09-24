package oauthprovider

import (
	"context"

	"upgradelink-admin-core/server/api/internal/svc"
	"upgradelink-admin-core/server/api/internal/types"
	"upgradelink-admin-core/server/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteOauthProviderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteOauthProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOauthProviderLogic {
	return &DeleteOauthProviderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteOauthProviderLogic) DeleteOauthProvider(req *types.IDsReq) (resp *types.BaseMsgResp, err error) {
	result, err := l.svcCtx.CoreRpc.DeleteOauthProvider(l.ctx, &core.IDsReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, result.Msg)}, nil
}
