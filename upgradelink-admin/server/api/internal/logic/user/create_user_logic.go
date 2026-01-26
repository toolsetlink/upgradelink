package user

import (
	"context"
	"net/http"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/http_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/utils/encrypt"
	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/ent/user"

	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.UserInfo) (resp *types.BaseMsgResp, err error) {
	if req.Mobile != nil {
		checkMobile, err := l.svcCtx.DB.User.Query().Where(user.MobileEQ(*req.Mobile)).Exist(l.ctx)
		if err != nil {
			return nil, db_error.DefaultEntError(l.Logger, err, req)
		}

		if checkMobile {
			return nil, http_error.NewApiError(http.StatusBadRequest, "login.mobileExist")
		}
	}

	if req.Email != nil {
		checkEmail, err := l.svcCtx.DB.User.Query().Where(user.EmailEQ(*req.Email)).Exist(l.ctx)
		if err != nil {
			return nil, db_error.DefaultEntError(l.Logger, err, req)
		}

		if checkEmail {
			return nil, http_error.NewApiError(http.StatusBadRequest, "login.signupUserExist")
		}
	}

	_, err = l.svcCtx.DB.User.Create().
		SetNotNilUsername(req.Username).
		SetNotNilPassword(pointy.GetPointer(encrypt.BcryptEncrypt(*req.Password))).
		SetNotNilNickname(req.Nickname).
		SetNotNilEmail(req.Email).
		SetNotNilMobile(req.Mobile).
		SetNotNilAvatar(req.Avatar).
		AddRoleIDs(req.RoleIds...).
		SetNotNilHomePath(req.HomePath).
		SetNotNilDescription(req.Description).
		Save(l.ctx)
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.CreateSuccess)}, nil
}
