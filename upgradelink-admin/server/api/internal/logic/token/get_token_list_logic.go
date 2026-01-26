package token

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/common/utils/uuidx"
	"upgradelink-admin/server/api/internal/ent"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/token"
	"upgradelink-admin/server/api/internal/ent/user"

	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTokenListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTokenListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenListLogic {
	return &GetTokenListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTokenListLogic) GetTokenList(req *types.TokenListReq) (resp *types.TokenListResp, err error) {
	var tokens *ent.TokenPageList

	if req.Username == nil && req.Uuid == nil && req.Nickname == nil && req.Email == nil {
		tokens, err = l.svcCtx.DB.Token.Query().Page(l.ctx, req.Page, req.PageSize)

		if err != nil {
			return nil, db_error.DefaultEntError(l.Logger, err, req)
		}
	} else {
		var predicates []predicate.User

		if req.Uuid != nil {
			predicates = append(predicates, user.IDEQ(uuidx.ParseUUIDString(*req.Uuid)))
		}

		if req.Username != nil {
			predicates = append(predicates, user.Username(*req.Username))
		}

		if req.Email != nil {
			predicates = append(predicates, user.EmailEQ(*req.Email))
		}

		if req.Nickname != nil {
			predicates = append(predicates, user.NicknameEQ(*req.Nickname))
		}

		u, err := l.svcCtx.DB.User.Query().Where(predicates...).First(l.ctx)
		if err != nil {
			return nil, db_error.DefaultEntError(l.Logger, err, req)
		}

		tokens, err = l.svcCtx.DB.Token.Query().Where(token.UUIDEQ(u.ID)).Page(l.ctx, req.Page, req.PageSize)

		if err != nil {
			return nil, db_error.DefaultEntError(l.Logger, err, req)
		}
	}

	resp = &types.TokenListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = tokens.PageDetails.Total

	for _, v := range tokens.List {
		resp.Data.Data = append(resp.Data.Data,
			types.TokenInfo{
				BaseUUIDInfo: types.BaseUUIDInfo{
					Id:        pointy.GetPointer(v.ID.String()),
					CreatedAt: pointy.GetPointer(v.CreatedAt.UnixMilli()),
					UpdatedAt: pointy.GetPointer(v.UpdatedAt.UnixMilli()),
				},
				Status:    pointy.GetPointer(uint32(v.Status)),
				Uuid:      pointy.GetPointer(v.UUID.String()),
				Token:     &v.Token,
				Source:    &v.Source,
				Username:  &v.Username,
				ExpiredAt: pointy.GetPointer(v.ExpiredAt.UnixMilli()),
			})
	}
	return resp, nil
}
