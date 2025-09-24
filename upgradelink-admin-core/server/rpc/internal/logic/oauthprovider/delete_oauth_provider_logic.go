package oauthprovider

import (
	"context"

	"upgradelink-admin-core/server/rpc/ent/oauthprovider"

	"upgradelink-admin-core/server/rpc/internal/svc"
	"upgradelink-admin-core/server/rpc/internal/utils/dberrorhandler"
	"upgradelink-admin-core/server/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-common/i18n"
)

type DeleteOauthProviderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteOauthProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOauthProviderLogic {
	return &DeleteOauthProviderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteOauthProviderLogic) DeleteOauthProvider(in *core.IDsReq) (*core.BaseResp, error) {
	_, err := l.svcCtx.DB.OauthProvider.Delete().Where(oauthprovider.IDIn(in.Ids...)).Exec(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
