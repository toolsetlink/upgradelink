package api

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/ent"
	"upgradelink-admin/server/api/internal/ent/api"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateApiLogic {
	return &CreateApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateApiLogic) CreateApi(req *types.ApiInfo) (resp *types.BaseMsgResp, err error) {
	// if exist , return success
	if req.Path != nil && req.Method != nil {
		check, err := l.svcCtx.DB.API.Query().Where(api.Path(*req.Path), api.Method(*req.Method)).Only(l.ctx)
		if err != nil && !ent.IsNotFound(err) {
			return nil, db_error.DefaultEntError(l.Logger, err, req)
		}

		if check != nil {
			return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.CreateSuccess)}, nil
		}
	}

	_, err = l.svcCtx.DB.API.Create().
		SetNotNilPath(req.Path).
		SetNotNilDescription(req.Description).
		SetNotNilAPIGroup(req.Group).
		SetNotNilMethod(req.Method).
		SetNotNilIsRequired(req.IsRequired).
		SetNotNilServiceName(req.ServiceName).
		Save(l.ctx)

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.CreateSuccess)}, nil
}
