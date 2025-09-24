package download

import (
	"context"

	"upgradelink-api/server/apils/internal/svc"
	"upgradelink-api/server/apils/internal/types"

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

func (l *GetElectronDownloadInfoDmgLogic) GetElectronDownloadInfoDmg(req *types.GetElectronDownloadInfoDmgReq) (resp *string, err error) {
	// todo: add your logic here and delete this line

	return
}
