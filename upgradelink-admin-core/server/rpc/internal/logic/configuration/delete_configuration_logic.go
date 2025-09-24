package configuration

import (
	"context"

	"upgradelink-admin-core/server/rpc/ent/configuration"
	"upgradelink-admin-core/server/rpc/internal/svc"
	"upgradelink-admin-core/server/rpc/internal/utils/dberrorhandler"
	"upgradelink-admin-core/server/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteConfigurationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteConfigurationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteConfigurationLogic {
	return &DeleteConfigurationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteConfigurationLogic) DeleteConfiguration(in *core.IDsReq) (*core.BaseResp, error) {
	_, err := l.svcCtx.DB.Configuration.Delete().Where(configuration.IDIn(in.Ids...)).Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
