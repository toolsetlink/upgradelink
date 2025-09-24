package smsprovider

import (
	"context"

	"upgradelink-admin-message/server/ent/smsprovider"
	"upgradelink-admin-message/server/internal/svc"
	"upgradelink-admin-message/server/internal/utils/dberrorhandler"
	"upgradelink-admin-message/server/types/mcms"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSmsProviderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteSmsProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSmsProviderLogic {
	return &DeleteSmsProviderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteSmsProviderLogic) DeleteSmsProvider(in *mcms.IDsReq) (*mcms.BaseResp, error) {
	_, err := l.svcCtx.DB.SmsProvider.Delete().Where(smsprovider.IDIn(in.Ids...)).Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &mcms.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
