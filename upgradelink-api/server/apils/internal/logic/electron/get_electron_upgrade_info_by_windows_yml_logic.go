package electron

import (
	"context"

	"upgradelink-api/server/apils/internal/svc"
	"upgradelink-api/server/apils/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetElectronUpgradeInfoByWindowsYmlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetElectronUpgradeInfoByWindowsYmlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetElectronUpgradeInfoByWindowsYmlLogic {
	return &GetElectronUpgradeInfoByWindowsYmlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetElectronUpgradeInfoByWindowsYmlLogic) GetElectronUpgradeInfoByWindowsYml(req *types.GetElectronUpgradeInfoReq) (resp *types.GetElectronUpgradeInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
