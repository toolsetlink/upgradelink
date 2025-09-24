package smslog

import (
	"context"

	"upgradelink-admin-message/server/internal/svc"
	"upgradelink-admin-message/server/internal/utils/dberrorhandler"
	"upgradelink-admin-message/server/types/mcms"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSmsLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateSmsLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSmsLogLogic {
	return &CreateSmsLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateSmsLogLogic) CreateSmsLog(in *mcms.SmsLogInfo) (*mcms.BaseUUIDResp, error) {
	query := l.svcCtx.DB.SmsLog.Create().
		SetNotNilPhoneNumber(in.PhoneNumber).
		SetNotNilContent(in.Content).
		SetNotNilProvider(in.Provider)

	if in.SendStatus != nil {
		query.SetNotNilSendStatus(pointy.GetPointer(uint8(*in.SendStatus)))
	}

	result, err := query.Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &mcms.BaseUUIDResp{Id: result.ID.String(), Msg: i18n.CreateSuccess}, nil
}
