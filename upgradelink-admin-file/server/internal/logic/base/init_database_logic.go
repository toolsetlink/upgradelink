package base

import (
	"context"

	"entgo.io/ent/dialect/sql/schema"

	"github.com/suyuan32/simple-admin-common/msg/logmsg"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/suyuan32/simple-admin-common/enum/errorcode"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"

	"upgradelink-admin-file/server/internal/svc"
	"upgradelink-admin-file/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InitDatabaseLogic) InitDatabase() (resp *types.BaseMsgResp, err error) {
	if err := l.svcCtx.DB.Schema.Create(l.ctx, schema.WithForeignKeys(false)); err != nil {
		logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
		return nil, errorx.NewCodeError(errorcode.Internal, err.Error())
	}

	if l.svcCtx.Config.CoreRpc.Enabled {
		err = l.insertApiData()
		if err != nil {
			if status.Code(err) == codes.InvalidArgument {
				return nil, errorx.NewCodeInvalidArgumentError(
					"init.alreadyInit")
			}
			return nil, err
		}

		err = l.insertMenuData()
		if err != nil {
			return nil, err
		}
	}

	err = l.svcCtx.Casbin.LoadPolicy()
	if err != nil {
		logx.Errorw("failed to load Casbin Policy", logx.Field("detail", err))
		return nil, errorx.NewCodeInternalError(i18n.DatabaseError)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.Success)}, nil
}
