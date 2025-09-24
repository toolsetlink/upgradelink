package download

import (
	"context"

	"upgradelink-api/server/apils/internal/svc"
	"upgradelink-api/server/apils/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUrlDownloadInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUrlDownloadInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUrlDownloadInfoLogic {
	return &GetUrlDownloadInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUrlDownloadInfoLogic) GetUrlDownloadInfo(req *types.GetUrlDownloadInfoReq) (resp *string, err error) {
	// todo: add your logic here and delete this line

	return
}
