package file

import (
	"context"

	"upgradelink-api/server/apils/internal/svc"
	"upgradelink-api/server/apils/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFileUpgradeInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFileUpgradeInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileUpgradeInfoLogic {
	return &GetFileUpgradeInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFileUpgradeInfoLogic) GetFileUpgradeInfo(req *types.GetFileUpgradeInfoReq) (resp *types.GetFileUpgradeInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
