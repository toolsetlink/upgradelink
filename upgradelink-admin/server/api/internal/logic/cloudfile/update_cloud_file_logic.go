package cloudfile

import (
	"context"
	"net/http"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCloudFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUpdateCloudFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCloudFileLogic {
	return &UpdateCloudFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCloudFileLogic) UpdateCloudFile(req *types.CloudFileInfo) (resp *types.BaseMsgResp, err error) {
	return
}
