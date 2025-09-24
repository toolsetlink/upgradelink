package electron

import (
	"context"

	"upgradelink-api/server/api/internal/svc"
	"upgradelink-api/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetElectronUpgradeInfoByMacYmlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetElectronUpgradeInfoByMacYmlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetElectronUpgradeInfoByMacYmlLogic {
	return &GetElectronUpgradeInfoByMacYmlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetElectronUpgradeInfoByMacYmlLogic) GetElectronUpgradeInfoByMacYml(req *types.GetElectronUpgradeInfoReq) (resp *types.GetElectronUpgradeInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
