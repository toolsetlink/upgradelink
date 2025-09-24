package publicuser

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/suyuan32/simple-admin-common/config"
	"github.com/suyuan32/simple-admin-common/enum/errorcode"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"

	"upgradelink-admin-core/server/rpc/types/core"

	"upgradelink-admin-core/server/api/internal/svc"
	"upgradelink-admin-core/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterBySmsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterBySmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterBySmsLogic {
	return &RegisterBySmsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *RegisterBySmsLogic) RegisterBySms(req *types.RegisterBySmsReq) (resp *types.BaseMsgResp, err error) {
	if l.svcCtx.Config.ProjectConf.RegisterVerify != "sms" && l.svcCtx.Config.ProjectConf.ResetVerify != "sms_or_email" {
		return nil, errorx.NewCodeAbortedError("login.registerTypeForbidden")
	}

	captchaData, err := l.svcCtx.Redis.Get(l.ctx, config.RedisCaptchaPrefix+req.PhoneNumber).Result()
	if err != nil {
		logx.Errorw("failed to get captcha data in redis for sms validation", logx.Field("detail", err),
			logx.Field("data", req))
		return nil, errorx.NewCodeInvalidArgumentError(i18n.Failed)
	}
	fmt.Println(captchaData)
	fmt.Println(req.Captcha)
	if captchaData == req.Captcha {
		// 先创建一个公司，公司先随机生成一个id，然后再创建用户
		company, err := l.svcCtx.CoreRpc.CreateCompany(l.ctx, &core.CompanyInfo{
			Name: &req.Username,
		})
		if err != nil {
			return nil, err
		}

		// 创建公司密钥
		accessKey, secretKey, _ := generateAPIKeys()
		_, err = l.svcCtx.CoreRpc.CreateCompanySecret(l.ctx, &core.CompanySecretInfo{
			CompanyId: &company.Id,
			AccessKey: &accessKey,
			SecretKey: &secretKey,
		})
		fmt.Println("CreateCompanySecret err: ", err)
		if err != nil {
			return nil, err
		}

		_, err = l.svcCtx.CoreRpc.CreateUser(l.ctx,
			&core.UserInfo{
				Username:     &req.Username,
				Password:     &req.Password,
				Mobile:       &req.PhoneNumber,
				Nickname:     &req.Username,
				Status:       pointy.GetPointer(uint32(1)),
				HomePath:     pointy.GetPointer("/dashboard"),
				RoleIds:      []uint64{l.svcCtx.Config.ProjectConf.DefaultRoleId},
				DepartmentId: pointy.GetPointer(l.svcCtx.Config.ProjectConf.DefaultDepartmentId),
				CompanyId:    &company.Id,
				PositionIds:  []uint64{l.svcCtx.Config.ProjectConf.DefaultPositionId},
			})
		if err != nil {
			return nil, err
		}

		err = l.svcCtx.Redis.Del(l.ctx, config.RedisCaptchaPrefix+req.PhoneNumber).Err()
		if err != nil {
			logx.Errorw("failed to delete captcha in redis", logx.Field("detail", err))
		}

		resp = &types.BaseMsgResp{
			Msg: l.svcCtx.Trans.Trans(l.ctx, "login.signupSuccessTitle"),
		}
		return resp, nil
	} else {
		return nil, errorx.NewCodeError(errorcode.InvalidArgument,
			l.svcCtx.Trans.Trans(l.ctx, "login.wrongCaptcha"))
	}
}

// generateAPIKeys 生成安全随机的AccessKey和SecretKey
func generateAPIKeys() (accessKey string, secretKey string, err error) {
	// 生成Access Key (16字节 -> 24字符Base64)
	accessBytes := make([]byte, 16)
	if _, err := rand.Read(accessBytes); err != nil {
		return "", "", err
	}
	accessKey = base64.RawURLEncoding.EncodeToString(accessBytes)

	// 生成Secret Key (32字节 -> 43字符Base64)
	secretBytes := make([]byte, 32)
	if _, err := rand.Read(secretBytes); err != nil {
		return "", "", err
	}
	secretKey = base64.RawURLEncoding.EncodeToString(secretBytes)

	return accessKey, secretKey, nil
}
