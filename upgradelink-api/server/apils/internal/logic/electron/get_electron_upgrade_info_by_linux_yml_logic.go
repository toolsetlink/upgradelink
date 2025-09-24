package electron

import (
	"context"

	"upgradelink-api/server/apils/internal/svc"
	"upgradelink-api/server/apils/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetElectronUpgradeInfoByLinuxYmlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetElectronUpgradeInfoByLinuxYmlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetElectronUpgradeInfoByLinuxYmlLogic {
	return &GetElectronUpgradeInfoByLinuxYmlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetElectronUpgradeInfoByLinuxYmlLogic) GetElectronUpgradeInfoByLinuxYml(req *types.GetElectronUpgradeInfoReq) (resp *types.GetElectronUpgradeInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
