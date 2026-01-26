package token

import (
	"context"
	"net/http"
	"time"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/enum"
	"upgradelink-admin/server/api/internal/common/http_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/msg/log_msg"
	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/common/utils/uuidx"

	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTokenLogic {
	return &UpdateTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTokenLogic) UpdateToken(req *types.TokenInfo) (resp *types.BaseMsgResp, err error) {
	token, err := l.svcCtx.DB.Token.UpdateOneID(uuidx.ParseUUIDString(*req.Id)).
		SetNotNilStatus(pointy.GetStatusPointer(req.Status)).
		SetNotNilSource(req.Source).
		Save(l.ctx)
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	if uint8(*req.Status) == enum.StatusBanned {
		expiredTime := token.ExpiredAt.Sub(time.Now())
		if expiredTime > 0 {
			err = l.svcCtx.Redis.Set(l.ctx, enum.RedisTokenPrefix+token.Token, "1", expiredTime).Err()
			if err != nil {
				logx.Errorw(log_msg.RedisError, logx.Field("detail", err.Error()))
				return nil, http_error.NewApiError(http.StatusInternalServerError, i18n.RedisError)
			}
		}
	} else if uint8(*req.Status) == enum.StatusNormal {
		err := l.svcCtx.Redis.Del(l.ctx, enum.RedisTokenPrefix+token.Token).Err()
		if err != nil {
			logx.Errorw(log_msg.RedisError, logx.Field("detail", err.Error()))
			return nil, http_error.NewApiError(http.StatusInternalServerError, i18n.RedisError)
		}
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.UpdateSuccess)}, nil
}
