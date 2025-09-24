package download

import (
	"context"

	"upgradelink-api/server/apils/internal/svc"
	"upgradelink-api/server/apils/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetElectronDownloadInfoZipLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetElectronDownloadInfoZipLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetElectronDownloadInfoZipLogic {
	return &GetElectronDownloadInfoZipLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetElectronDownloadInfoZipLogic) GetElectronDownloadInfoZip(req *types.GetElectronDownloadInfoZipReq) (resp *string, err error) {
	// todo: add your logic here and delete this line

	return
}
