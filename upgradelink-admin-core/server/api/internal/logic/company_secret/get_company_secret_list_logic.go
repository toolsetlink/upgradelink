package company_secret

import (
	"context"
	"encoding/json"
	"upgradelink-admin-core/server/api/internal/svc"
	"upgradelink-admin-core/server/api/internal/types"
	"upgradelink-admin-core/server/api/internal/utils"
	"upgradelink-admin-core/server/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/orm/ent/entctx/userctx"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCompanySecretListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCompanySecretListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCompanySecretListLogic {
	return &GetCompanySecretListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

type RuleDataInfo struct {
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

func (l *GetCompanySecretListLogic) GetCompanySecretList(req *types.CompanySecretListReq) (resp *types.CompanySecretListResp, err error) {

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

	isDelInt := uint32(0)
	data, err := l.svcCtx.CoreRpc.GetCompanySecretList(l.ctx, &core.CompanySecretListReq{
		Page:      req.Page,
		PageSize:  req.PageSize,
		CompanyId: userData.CompanyId,
		AccessKey: &req.AccessKey,
		SecretKey: &req.SecretKey,
		Enable:    req.Enable,
		IsDel:     &isDelInt,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.CompanySecretListResp{}
	resp.Data.Total = data.Total
	for _, v := range data.Data {
		// 整理有效期时间
		validityDatetimeStr := ""
		// 判断数据是否为默认值， 如果为默认值则代表为无期限数据
		if int(*v.ValidityDatetime) != 1000 && int(*v.ValidityDatetime) != 0 {
			validityDatetime, err := utils.IntUnixMilliToTime(int(*v.ValidityDatetime))
			if err != nil {
				return nil, err
			}
			validityDatetimeStr = validityDatetime.Format("2006-01-02 15:04:05")
		}

		// 整理 rule 数据
		var ruleData RuleDataInfo
		if v.RuleData != nil {
			if *v.RuleData != "" {
				err := json.Unmarshal([]byte(*v.RuleData), &ruleData)
				if err != nil {
					return nil, err
				}
			}
		}
		ruleDataUrl := []int{}
		if len(ruleData.Url) != 0 {
			ruleDataUrl = ruleData.Url
		}
		ruleDataFile := []int{}
		if len(ruleData.File) != 0 {
			ruleDataFile = ruleData.File
		}
		ruleDataConfiguration := []int{}
		if len(ruleData.Configuration) != 0 {
			ruleDataConfiguration = ruleData.Configuration
		}
		ruleDataTauri := []int{}
		if len(ruleData.Tauri) != 0 {
			ruleDataTauri = ruleData.Tauri
		}
		ruleDataElectron := []int{}
		if len(ruleData.Electron) != 0 {
			ruleDataElectron = ruleData.Electron
		}
		ruleDataApk := []int{}
		if len(ruleData.Apk) != 0 {
			ruleDataApk = ruleData.Apk
		}
		ruleDataWin := []int{}
		if len(ruleData.Win) != 0 {
			ruleDataWin = ruleData.Win
		}
		ruleDataMac := []int{}
		if len(ruleData.Mac) != 0 {
			ruleDataMac = ruleData.Mac
		}
		ruleDataLnx := []int{}
		if len(ruleData.Lnx) != 0 {
			ruleDataLnx = ruleData.Lnx
		}

		resp.Data.Data = append(resp.Data.Data, types.CompanySecretInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id:        v.Id,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			},
			AccessKey:             *v.AccessKey,
			SecretKey:             *v.SecretKey,
			ValidityDatetime:      validityDatetimeStr,
			RuleDataUrl:           ruleDataUrl,
			RuleDataFile:          ruleDataFile,
			RuleDataConfiguration: ruleDataConfiguration,
			RuleDataTauri:         ruleDataTauri,
			RuleDataElectron:      ruleDataElectron,
			RuleDataApk:           ruleDataApk,
			RuleDataWin:           ruleDataWin,
			RuleDataMac:           ruleDataMac,
			RuleDataLnx:           ruleDataLnx,
			Enable:                v.Enable,
			IsDel:                 v.IsDel,
			Description:           *v.Description,
		})
	}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	return resp, nil
}
