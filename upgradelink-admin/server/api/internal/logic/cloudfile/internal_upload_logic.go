package cloudfile

import (
	"context"

	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InternalUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInternalUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InternalUploadLogic {
	return &InternalUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InternalUploadLogic) InternalUpload() (resp *types.CloudFileInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
