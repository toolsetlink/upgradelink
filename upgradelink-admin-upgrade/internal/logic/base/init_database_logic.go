package base

import (
	"context"

	"upgradelink-admin-upgrade/ent"
	"upgradelink-admin-upgrade/internal/utils/dberrorhandler"

	"entgo.io/ent/dialect/sql/schema"

	"github.com/pkg/errors"
	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/suyuan32/simple-admin-common/msg/logmsg"
	"github.com/suyuan32/simple-admin-common/orm/ent/entctx/userctx"
	"github.com/zeromicro/go-zero/core/errorx"

	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *InitDatabaseLogic) InitDatabase() (resp *types.BaseMsgResp, err error) {
	if err := l.svcCtx.DB.Schema.Create(l.ctx, schema.WithForeignKeys(false)); err != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		return nil, errorx.NewInternalError(err.Error())
	}

	return &types.BaseMsgResp{Msg: errormsg.Success}, nil
}

// GetCompanyId  获取公司id
func GetCompanyId(ctx context.Context, svcCtx *svc.ServiceContext, logger logx.Logger) (int, error) {
	// 获取公司 id
	userId, err := userctx.GetUserIDFromCtx(ctx)
	if err != nil {
		return 0, dberrorhandler.DefaultEntError(logger, err, "GetUserIDFromCtx error")
	}
	userData, err := svcCtx.DB.SysUser.Get(ctx, userId)
	if err != nil {
		return 0, dberrorhandler.DefaultEntError(logger, err, "GetCompanyId error")
	}
	
	if userData.CompanyID == 0 {
		return 0, dberrorhandler.DefaultEntError(logger, err, "GetCompanyId error")
	}

	companyID := int(userData.CompanyID)

	return companyID, nil
}

// WithTx 事务处理
func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = errors.Wrapf(err, "rolling back transaction: %v", rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return errors.Wrapf(err, "committing transaction: %v", err)
	}
	return nil
}
