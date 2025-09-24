package user

import (
	"context"

	"upgradelink-admin-core/server/rpc/ent/user"
	"upgradelink-admin-core/server/rpc/internal/svc"
	"upgradelink-admin-core/server/rpc/internal/utils/dberrorhandler"
	"upgradelink-admin-core/server/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/encrypt"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// User management
func (l *CreateUserLogic) CreateUser(in *core.UserInfo) (*core.BaseUUIDResp, error) {
	if in.Mobile != nil {
		checkMobile, err := l.svcCtx.DB.User.Query().Where(user.MobileEQ(*in.Mobile)).Exist(l.ctx)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
		}

		if checkMobile {
			return nil, errorx.NewInvalidArgumentError("login.mobileExist")
		}
	}

	if in.Email != nil {
		checkEmail, err := l.svcCtx.DB.User.Query().Where(user.EmailEQ(*in.Email)).Exist(l.ctx)
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
		}

		if checkEmail {
			return nil, errorx.NewInvalidArgumentError("login.signupUserExist")
		}
	}

	result, err := l.svcCtx.DB.User.Create().
		SetNotNilUsername(in.Username).
		SetNotNilPassword(pointy.GetPointer(encrypt.BcryptEncrypt(*in.Password))).
		SetNotNilNickname(in.Nickname).
		SetNotNilEmail(in.Email).
		SetNotNilMobile(in.Mobile).
		SetNotNilAvatar(in.Avatar).
		AddRoleIDs(in.RoleIds...).
		SetNotNilHomePath(in.HomePath).
		SetNotNilDescription(in.Description).
		SetNotNilDepartmentID(in.DepartmentId).
		SetNotNilCompanyID(in.CompanyId).
		AddPositionIDs(in.PositionIds...).
		Save(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseUUIDResp{Id: result.ID.String(), Msg: i18n.CreateSuccess}, nil
}
