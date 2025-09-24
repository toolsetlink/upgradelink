package url

import (
	"context"

	"upgradelink-api/server/apils/internal/svc"
	"upgradelink-api/server/apils/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUrlUpgradeInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUrlUpgradeInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUrlUpgradeInfoLogic {
	return &GetUrlUpgradeInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUrlUpgradeInfoLogic) GetUrlUpgradeInfo(req *types.GetUrlUpgradeInfoReq) (resp *types.GetUrlUpgradeInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
