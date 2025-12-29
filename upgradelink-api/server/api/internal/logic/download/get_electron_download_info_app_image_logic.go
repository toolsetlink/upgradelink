package download

import (
	"context"
	"upgradelink-api/server/api/internal/svc"
	"upgradelink-api/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetElectronDownloadInfoAppImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetElectronDownloadInfoAppImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetElectronDownloadInfoAppImageLogic {
	return &GetElectronDownloadInfoAppImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetElectronDownloadInfoAppImageLogic) GetElectronDownloadInfoAppImage(req *types.GetElectronDownloadInfoReq) (resp *string, err error) {
	// todo: add your logic here and delete this line

	return
}
