package publicuser

import (
	"context"
	"strconv"
	"strings"
	"time"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/enum"
	"upgradelink-admin/server/api/internal/common/http_error"
	"upgradelink-admin/server/api/internal/common/utils/encrypt"
	"upgradelink-admin/server/api/internal/common/utils/jwt"
	"upgradelink-admin/server/api/internal/ent"
	"upgradelink-admin/server/api/internal/ent/user"

	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// 效验验证码
	if ok := l.svcCtx.Captcha.Verify(enum.RedisCaptchaPrefix+req.CaptchaId, req.Captcha, true); !ok {
		return nil, http_error.NewApiBadRequestError(l.svcCtx.Trans.Trans(l.ctx, "login.wrongCaptcha"))
	}

	userData, err := l.svcCtx.DB.User.Query().Where(user.UsernameEQ(req.Username)).WithRoles().First(l.ctx)
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	if userData.Status != enum.StatusNormal {
		return nil, http_error.NewApiBadRequestError("login.userBanned")
	}

	if !encrypt.BcryptCheck(req.Password, userData.Password) {
		return nil, http_error.NewApiBadRequestError("login.wrongUsernameOrPassword")
	}

	token, err := jwt.NewJwtToken(
		l.svcCtx.Config.Auth.AccessSecret, time.Now().Unix(),
		l.svcCtx.Config.Auth.AccessExpire,
		jwt.WithOption("userId", userData.ID),
		jwt.WithOption("roleId", strings.Join(GetRoleCodes(userData.Edges.Roles), ",")),
		jwt.WithOption("companyId", strconv.FormatUint(userData.CompanyID, 10)),
	)
	if err != nil {
		return nil, err
	}

	// add token into database
	expiredAt := time.Now().Add(time.Second * time.Duration(l.svcCtx.Config.Auth.AccessExpire))
	_, err = l.svcCtx.DB.Token.Create().SetStatus(enum.StatusNormal).
		SetUUID(userData.ID).
		SetToken(token).
		SetSource("core_user").
		SetUsername(userData.Username).
		SetExpiredAt(expiredAt).
		Save(l.ctx)
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	err = l.svcCtx.Redis.Del(l.ctx, enum.RedisCaptchaPrefix+req.CaptchaId).Err()
	if err != nil {
		logx.Errorw("failed to delete captcha in redis", logx.Field("detail", err))
	}

	resp = &types.LoginResp{
		BaseDataInfo: types.BaseDataInfo{Msg: l.svcCtx.Trans.Trans(l.ctx, "login.loginSuccessTitle")},
		Data: types.LoginInfo{
			UserId: userData.ID.String(),
			Token:  token,
			Expire: uint64(expiredAt.Unix()),
		},
	}

	return resp, nil
}

func GetRoleCodes(data []*ent.Role) []string {
	var codes []string
	for _, v := range data {
		codes = append(codes, v.Code)
	}
	return codes
}
