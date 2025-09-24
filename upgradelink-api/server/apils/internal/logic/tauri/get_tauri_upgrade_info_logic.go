package tauri

import (
	"context"

	"upgradelink-api/server/apils/internal/svc"
	"upgradelink-api/server/apils/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTauriUpgradeInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTauriUpgradeInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTauriUpgradeInfoLogic {
	return &GetTauriUpgradeInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTauriUpgradeInfoLogic) GetTauriUpgradeInfo(req *types.GetTauriUpgradeInfoReq) (resp *types.GetTauriUpgradeInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
