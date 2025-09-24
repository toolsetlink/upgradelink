package base

import (
	"context"

	"entgo.io/ent/dialect/sql/schema"
	"github.com/suyuan32/simple-admin-common/enum/errorcode"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/msg/logmsg"
	"github.com/zeromicro/go-zero/core/errorx"

	"upgradelink-admin-message/server/ent"
	"upgradelink-admin-message/server/internal/enum/emailauthtype"
	"upgradelink-admin-message/server/internal/utils/dberrorhandler"

	"upgradelink-admin-message/server/internal/svc"
	"upgradelink-admin-message/server/types/mcms"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InitDatabaseLogic) InitDatabase(in *mcms.Empty) (*mcms.BaseResp, error) {
	if err := l.svcCtx.DB.Schema.Create(l.ctx, schema.WithForeignKeys(false)); err != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		return nil, errorx.NewCodeError(errorcode.Internal, err.Error())
	}

	err := l.InsertEmailProviderData()
	if err != nil {
		return nil, err
	}

	err = l.InsertSmsProviderData()
	if err != nil {
		return nil, err
	}

	return &mcms.BaseResp{Msg: i18n.Success}, nil
}

func (l *InitDatabaseLogic) InsertEmailProviderData() error {
	var emailProviders []*ent.EmailProviderCreate

	// tencent
	emailProviders = append(emailProviders, l.svcCtx.DB.EmailProvider.Create().
		SetName("tencent").
		SetAuthType(emailauthtype.Plain).
		SetEmailAddr("input your email address").
		SetPassword("input your password").
		SetPort(465).
		SetHostName("smtp.qq.com").
		SetTLS(true).
		SetIsDefault(true))

	err := l.svcCtx.DB.EmailProvider.CreateBulk(emailProviders...).Exec(l.ctx)
	if err != nil {
		return dberrorhandler.DefaultEntError(logx.WithContext(nil), err, nil)
	}

	return nil
}

func (l *InitDatabaseLogic) InsertSmsProviderData() error {
	var smsProviders []*ent.SmsProviderCreate

	// tencent
	smsProviders = append(smsProviders, l.svcCtx.DB.SmsProvider.Create().
		SetName("tencent").
		SetSecretID("input your secret ID").
		SetSecretKey("input your secret key").
		SetRegion("ap-nanjing").
		SetIsDefault(true))

	err := l.svcCtx.DB.SmsProvider.CreateBulk(smsProviders...).Exec(l.ctx)
	if err != nil {
		return dberrorhandler.DefaultEntError(logx.WithContext(nil), err, nil)
	}

	return nil
}
