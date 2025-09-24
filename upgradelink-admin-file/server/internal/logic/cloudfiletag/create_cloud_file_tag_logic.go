package cloudfiletag

import (
	"context"

	"upgradelink-admin-file/server/internal/svc"
	"upgradelink-admin-file/server/internal/types"
	"upgradelink-admin-file/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCloudFileTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCloudFileTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCloudFileTagLogic {
	return &CreateCloudFileTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCloudFileTagLogic) CreateCloudFileTag(req *types.CloudFileTagInfo) (*types.BaseMsgResp, error) {
	_, err := l.svcCtx.DB.CloudFileTag.Create().
		SetNotNilStatus(req.Status).
		SetNotNilName(req.Name).
		SetNotNilRemark(req.Remark).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.CreateSuccess)}, nil
}
