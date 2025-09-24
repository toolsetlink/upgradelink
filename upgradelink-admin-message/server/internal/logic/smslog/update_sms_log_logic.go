package smslog

import (
	"context"

	"upgradelink-admin-message/server/internal/svc"
	"upgradelink-admin-message/server/internal/utils/dberrorhandler"
	"upgradelink-admin-message/server/types/mcms"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSmsLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSmsLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSmsLogLogic {
	return &UpdateSmsLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSmsLogLogic) UpdateSmsLog(in *mcms.SmsLogInfo) (*mcms.BaseResp, error) {
	query := l.svcCtx.DB.SmsLog.UpdateOneID(uuidx.ParseUUIDString(*in.Id)).
		SetNotNilPhoneNumber(in.PhoneNumber).
		SetNotNilContent(in.Content).
		SetNotNilProvider(in.Provider)

	if in.SendStatus != nil {
		query.SetNotNilSendStatus(pointy.GetPointer(uint8(*in.SendStatus)))
	}

	err := query.Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &mcms.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
