package company_secret

import (
	"context"
	"encoding/json"
	"upgradelink-admin/server/api/internal/common"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

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

		reqRuleDataUrl := []int{}
		if len(req.RuleDataUrl) != 0 {
			reqRuleDataUrl = req.RuleDataUrl
		}
		reqRuleDataFile := []int{}
		if len(req.RuleDataFile) != 0 {
			reqRuleDataFile = req.RuleDataFile
		}
		reqRuleDataConfiguration := []int{}
		if len(req.RuleDataConfiguration) != 0 {
			reqRuleDataConfiguration = req.RuleDataConfiguration
		}
		reqRuleDataTauri := []int{}
		if len(req.RuleDataTauri) != 0 {
			reqRuleDataTauri = req.RuleDataTauri
		}
		reqRuleDataElectron := []int{}
		if len(req.RuleDataElectron) != 0 {
			reqRuleDataElectron = req.RuleDataElectron
		}
		reqRuleDataApk := []int{}
		if len(req.RuleDataApk) != 0 {
			reqRuleDataApk = req.RuleDataApk
		}
		reqRuleDataWin := []int{}
		if len(req.RuleDataWin) != 0 {
			reqRuleDataWin = req.RuleDataWin
		}
		reqRuleDataMac := []int{}
		if len(req.RuleDataMac) != 0 {
			reqRuleDataMac = req.RuleDataMac
		}
		reqRuleDataLnx := []int{}
		if len(req.RuleDataLnx) != 0 {
			reqRuleDataLnx = req.RuleDataLnx
		}
		// struct 转 json 字符串
		var updateRuleDataInfo UpdateRuleDataInfo
		updateRuleDataInfo = UpdateRuleDataInfo{
			Url:           reqRuleDataUrl,
			File:          reqRuleDataFile,
			Configuration: reqRuleDataConfiguration,
			Tauri:         reqRuleDataTauri,
			Electron:      reqRuleDataElectron,
			Apk:           reqRuleDataApk,
			Win:           reqRuleDataWin,
			Mac:           reqRuleDataMac,
			Lnx:           reqRuleDataLnx,
		}

		// updateRuleDataInfo 转为 json 字符串
		ruleDataJsonStr, err := json.Marshal(updateRuleDataInfo)
		if err != nil {
			return nil, err
		}
		ruleDataStr = string(ruleDataJsonStr)
	}

	err = l.svcCtx.DB.CompanySecret.UpdateOneID(*req.Id).
		SetNotNilValidityDatetime(pointy.GetTimeMilliPointer(&validityDatetimeInt)).
		SetNotNilRuleData(&ruleDataStr).
		SetNotNilEnable(req.Enable).
		SetNotNilDescription(&req.Description).
		SetNotNilIsDel(req.IsDel).
		Exec(l.ctx)

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.UpdateSuccess)}, nil
}
