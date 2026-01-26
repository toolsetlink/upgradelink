package token

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/common/utils/uuidx"

	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTokenLogic {
	return &CreateTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTokenLogic) CreateToken(req *types.TokenInfo) (resp *types.BaseMsgResp, err error) {
	_, err = l.svcCtx.DB.Token.Create().
		SetNotNilStatus(pointy.GetStatusPointer(req.Status)).
		SetNotNilUUID(uuidx.ParseUUIDStringToPointer(req.Uuid)).
		SetNotNilToken(req.Token).
		SetNotNilSource(req.Source).
		SetNotNilUsername(req.Username).
		SetNotNilExpiredAt(pointy.GetTimeMilliPointer(req.ExpiredAt)).
		Save(l.ctx)
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.CreateSuccess)}, nil
}
