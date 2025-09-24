package configuration

import (
	"context"

	"upgradelink-api/server/apils/internal/svc"
	"upgradelink-api/server/apils/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfigurationUpgradeInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetConfigurationUpgradeInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigurationUpgradeInfoLogic {
	return &GetConfigurationUpgradeInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetConfigurationUpgradeInfoLogic) GetConfigurationUpgradeInfo(req *types.GetConfigurationUpgradeInfoReq) (resp *types.GetConfigurationUpgradeInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
