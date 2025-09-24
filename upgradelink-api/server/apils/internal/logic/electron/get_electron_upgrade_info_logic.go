package electron

import (
	"context"

	"upgradelink-api/server/apils/internal/svc"
	"upgradelink-api/server/apils/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetElectronUpgradeInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetElectronUpgradeInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetElectronUpgradeInfoLogic {
	return &GetElectronUpgradeInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetElectronUpgradeInfoLogic) GetElectronUpgradeInfo(req *types.GetElectronUpgradeInfoReq) (resp *types.GetElectronUpgradeInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
