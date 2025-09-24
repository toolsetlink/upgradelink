package download

import (
	"context"

	"upgradelink-api/server/apils/internal/svc"
	"upgradelink-api/server/apils/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFileDownloadInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFileDownloadInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileDownloadInfoLogic {
	return &GetFileDownloadInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFileDownloadInfoLogic) GetFileDownloadInfo(req *types.GetFileDownloadInfoReq) (resp *string, err error) {
	// todo: add your logic here and delete this line

	return
}
