package download

import (
	"context"

	"upgradelink-api/server/apils/internal/svc"
	"upgradelink-api/server/apils/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetElectronDownloadInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetElectronDownloadInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetElectronDownloadInfoLogic {
	return &GetElectronDownloadInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetElectronDownloadInfoLogic) GetElectronDownloadInfo(req *types.GetElectronDownloadInfoReq) (resp *string, err error) {
	// todo: add your logic here and delete this line

	return
}
