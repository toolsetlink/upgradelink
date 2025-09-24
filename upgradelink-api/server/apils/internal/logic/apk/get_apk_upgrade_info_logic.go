package apk

import (
	"context"

	"upgradelink-api/server/apils/internal/svc"
	"upgradelink-api/server/apils/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApkUpgradeInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApkUpgradeInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApkUpgradeInfoLogic {
	return &GetApkUpgradeInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApkUpgradeInfoLogic) GetApkUpgradeInfo(req *types.GetApkUpgradeInfoReq) (resp *types.GetApkUpgradeInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
