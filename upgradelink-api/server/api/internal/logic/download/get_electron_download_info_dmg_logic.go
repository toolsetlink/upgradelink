package download

import (
	"context"
	"upgradelink-api/server/api/internal/svc"
	"upgradelink-api/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetElectronDownloadInfoDmgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetElectronDownloadInfoDmgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetElectronDownloadInfoDmgLogic {
	return &GetElectronDownloadInfoDmgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetElectronDownloadInfoDmgLogic) GetElectronDownloadInfoDmg(req *types.GetElectronDownloadInfoReq) (resp *string, err error) {
	// todo: add your logic here and delete this line

	return
}
