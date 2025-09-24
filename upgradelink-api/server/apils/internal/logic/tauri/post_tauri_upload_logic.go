package tauri

import (
	"context"

	"upgradelink-api/server/apils/internal/svc"
	"upgradelink-api/server/apils/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostTauriUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostTauriUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostTauriUploadLogic {
	return &PostTauriUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostTauriUploadLogic) PostTauriUpload(req *types.PostTauriUploadReq) (resp *types.PostTauriUploadResp, err error) {
	// todo: add your logic here and delete this line

	return
}
