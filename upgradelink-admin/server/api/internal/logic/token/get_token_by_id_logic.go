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

type GetTokenByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTokenByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenByIdLogic {
	return &GetTokenByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTokenByIdLogic) GetTokenById(req *types.UUIDReq) (resp *types.TokenInfoResp, err error) {
	result, err := l.svcCtx.DB.Token.Get(l.ctx, uuidx.ParseUUIDString(req.Id))
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.TokenInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.TokenInfo{
			BaseUUIDInfo: types.BaseUUIDInfo{
				Id:        pointy.GetPointer(result.ID.String()),
				CreatedAt: pointy.GetPointer(result.CreatedAt.UnixMilli()),
				UpdatedAt: pointy.GetPointer(result.UpdatedAt.UnixMilli()),
			},
			Status:    pointy.GetPointer(uint32(result.Status)),
			Uuid:      pointy.GetPointer(result.UUID.String()),
			Token:     &result.Token,
			Source:    &result.Source,
			Username:  &result.Username,
			ExpiredAt: pointy.GetPointer(result.ExpiredAt.UnixMilli()),
		},
	}, nil

}
