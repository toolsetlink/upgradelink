package cloudfiletag

import (
	"context"

	"upgradelink-admin-file/server/ent/cloudfiletag"
	"upgradelink-admin-file/server/internal/svc"
	"upgradelink-admin-file/server/internal/types"
	"upgradelink-admin-file/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCloudFileTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCloudFileTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCloudFileTagLogic {
	return &DeleteCloudFileTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCloudFileTagLogic) DeleteCloudFileTag(req *types.IDsReq) (*types.BaseMsgResp, error) {
	_, err := l.svcCtx.DB.CloudFileTag.Delete().Where(cloudfiletag.IDIn(req.Ids...)).Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.DeleteSuccess)}, nil
}
