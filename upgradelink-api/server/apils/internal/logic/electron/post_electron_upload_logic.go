package electron

import (
	"context"

	"upgradelink-api/server/apils/internal/svc"
	"upgradelink-api/server/apils/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostElectronUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostElectronUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostElectronUploadLogic {
	return &PostElectronUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostElectronUploadLogic) PostElectronUpload(req *types.PostElectronUploadReq) (resp *types.PostElectronUploadResp, err error) {
	// todo: add your logic here and delete this line

	return
}
