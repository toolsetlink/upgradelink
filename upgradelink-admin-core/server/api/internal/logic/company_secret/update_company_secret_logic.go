package company_secret

import (
	"context"
	"encoding/json"
	"upgradelink-admin-core/server/api/internal/utils"
	"upgradelink-admin-core/server/rpc/types/core"

	"upgradelink-admin-core/server/api/internal/svc"
	"upgradelink-admin-core/server/api/internal/types"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/orm/ent/entctx/userctx"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCompanySecretLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCompanySecretLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCompanySecretLogic {
	return &UpdateCompanySecretLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type UpdateRuleDataInfo struct {
	Url           []int `json:"url,optional"`
	File          []int `json:"file,optional"`
	Configuration []int `json:"configuration,optional"`
	Tauri         []int `json:"tauri,optional"`
	Electron      []int `json:"electron,optional"`
	Apk           []int `json:"apk,optional"`
	Win           []int `json:"win,optional"`
	Mac           []int `json:"mac,optional"`
	Lnx           []int `json:"lnx,optional"`
}

func (l *UpdateCompanySecretLogic) UpdateCompanySecret(req *types.CompanySecretInfo) (resp *types.BaseMsgResp, err error) {

	// 获取公司 id
	userId, err := userctx.GetUserIDFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}
	userData, err := l.svcCtx.CoreRpc.GetUserById(l.ctx, &core.UUIDReq{
		Id: userId,
	})
	if err != nil {
		return nil, err
	}

	// 效验公司数据
	data, err := l.svcCtx.CoreRpc.GetCompanySecretById(l.ctx, &core.IDReq{Id: *req.Id})
	if err != nil {
		return nil, err
	}

	if *data.CompanyId != *userData.CompanyId {
		return nil, errorx.NewCodeInvalidArgumentError(i18n.TargetNotFound)
	}

	// 整理有效期时间
	// 判断数据传输是否为空，如果为空，则代表无期限
	validityDatetimeInt := int64(1000)
	if req.ValidityDatetime != "" {
		validityDatetimeTime, _ := utils.StringToTime(req.ValidityDatetime)
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

	updateData, err := l.svcCtx.CoreRpc.UpdateCompanySecret(l.ctx,
		&core.CompanySecretInfo{
			Id:               req.Id,
			Enable:           req.Enable,
			ValidityDatetime: &validityDatetimeInt,
			RuleData:         &ruleDataStr,
			Description:      &req.Description,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, updateData.Msg)}, nil
}
