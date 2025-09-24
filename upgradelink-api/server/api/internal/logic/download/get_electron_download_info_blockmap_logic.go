package download

import (
	"context"

	"upgradelink-api/server/api/internal/svc"
	"upgradelink-api/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetElectronDownloadInfoBlockmapLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetElectronDownloadInfoBlockmapLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetElectronDownloadInfoBlockmapLogic {
	return &GetElectronDownloadInfoBlockmapLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetElectronDownloadInfoBlockmapLogic) GetElectronDownloadInfoBlockmap(req *types.GetElectronDownloadInfoReq) (resp *string, err error) {
	// todo: add your logic here and delete this line

	return
}
