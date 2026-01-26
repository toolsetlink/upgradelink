package company_secret

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"upgradelink-admin/server/api/internal/common"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/userctx"
	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/common/utils/uuidx"
	"upgradelink-admin/server/api/internal/ent/user"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCompanySecretLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCompanySecretLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCompanySecretLogic {
	return &CreateCompanySecretLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCompanySecretLogic) CreateCompanySecret(req *types.CompanySecretInfo) (resp *types.BaseMsgResp, err error) {
	// 获取公司 id
	userId, err := userctx.GetUserIDFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}
	userData, err := l.svcCtx.DB.User.Query().Where(user.IDEQ(uuidx.ParseUUIDString(userId))).WithRoles().First(l.ctx)
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, userId)
	}

	// 生产密钥
	accessKey, secretKey, _ := generateAPIKeys()

	// 整理有效期时间
	// 判断数据传输是否为空，如果为空，则代表无期限
	validityDatetimeInt := int64(1000)
	if req.ValidityDatetime != "" {
		validityDatetimeTime, _ := common.StringToTime(req.ValidityDatetime)
		validityDatetimeInt = validityDatetimeTime.UnixMilli()
	}

	// 整理权限数据 判断数据是否都为空，如果都为空的话，字段录入为空字符
	ruleDataStr := ""
	if len(req.RuleDataUrl) != 0 || len(req.RuleDataFile) != 0 || len(req.RuleDataConfiguration) != 0 || len(req.RuleDataTauri) != 0 || len(req.RuleDataElectron) != 0 || len(req.RuleDataApk) != 0 || len(req.RuleDataWin) != 0 || len(req.RuleDataMac) != 0 || len(req.RuleDataLnx) != 0 {
		// struct 转 json 字符串
		var updateRuleDataInfo UpdateRuleDataInfo
		updateRuleDataInfo = UpdateRuleDataInfo{
			Url:           req.RuleDataUrl,
			File:          req.RuleDataFile,
			Configuration: req.RuleDataConfiguration,
			Tauri:         req.RuleDataTauri,
			Electron:      req.RuleDataElectron,
			Apk:           req.RuleDataApk,
			Win:           req.RuleDataWin,
			Mac:           req.RuleDataMac,
			Lnx:           req.RuleDataLnx,
		}

		// updateRuleDataInfo 转为 json 字符串
		ruleDataJsonStr, err := json.Marshal(updateRuleDataInfo)
		if err != nil {
			return nil, err
		}
		ruleDataStr = string(ruleDataJsonStr)
	}

	_, err = l.svcCtx.DB.CompanySecret.Create().
		SetNotNilCompanyID(&userData.CompanyID).
		SetNotNilAccessKey(&accessKey).
		SetNotNilSecretKey(&secretKey).
		SetNotNilValidityDatetime(pointy.GetTimeMilliPointer(&validityDatetimeInt)).
		SetNotNilRuleData(&ruleDataStr).
		SetNotNilEnable(req.Enable).
		SetNotNilDescription(&req.Description).
		Save(l.ctx)

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.CreateSuccess)}, nil
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
