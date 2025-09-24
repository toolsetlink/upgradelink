package download

import (
	"context"

	"upgradelink-api/server/apils/internal/svc"
	"upgradelink-api/server/apils/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetElectronDownloadInfoExeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetElectronDownloadInfoExeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetElectronDownloadInfoExeLogic {
	return &GetElectronDownloadInfoExeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetElectronDownloadInfoExeLogic) GetElectronDownloadInfoExe(req *types.GetElectronDownloadInfoExeReq) (resp *string, err error) {
	// todo: add your logic here and delete this line

	return
}
