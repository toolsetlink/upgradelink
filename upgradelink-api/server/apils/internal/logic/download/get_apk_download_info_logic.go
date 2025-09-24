package download

import (
	"context"

	"upgradelink-api/server/apils/internal/svc"
	"upgradelink-api/server/apils/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApkDownloadInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApkDownloadInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApkDownloadInfoLogic {
	return &GetApkDownloadInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApkDownloadInfoLogic) GetApkDownloadInfo(req *types.GetApkDownloadInfoReq) (resp string, err error) {
	// todo: add your logic here and delete this line

	return
}
