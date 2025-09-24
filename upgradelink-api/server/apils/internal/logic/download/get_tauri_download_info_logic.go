package download

import (
	"context"

	"upgradelink-api/server/apils/internal/svc"
	"upgradelink-api/server/apils/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTauriDownloadInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTauriDownloadInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTauriDownloadInfoLogic {
	return &GetTauriDownloadInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTauriDownloadInfoLogic) GetTauriDownloadInfo(req *types.GetTauriDownloadInfoReq) (resp *string, err error) {
	// todo: add your logic here and delete this line

	return
}
