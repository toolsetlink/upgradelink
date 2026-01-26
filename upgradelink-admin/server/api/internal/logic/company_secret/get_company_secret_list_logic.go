package company_secret

import (
	"context"
	"encoding/json"
	"upgradelink-admin/server/api/internal/common"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/userctx"
	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/common/utils/uuidx"
	"upgradelink-admin/server/api/internal/ent/companysecret"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/user"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

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
		svcCtx: svcCtx,
	}
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
	userData, err := l.svcCtx.DB.User.Query().Where(user.IDEQ(uuidx.ParseUUIDString(userId))).WithRoles().First(l.ctx)
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, userId)
	}
	if err != nil {
		return nil, err
	}

	var predicates []predicate.CompanySecret
	predicates = append(predicates, companysecret.CompanyIDEQ(userData.CompanyID))

	if req.AccessKey != "" {
		predicates = append(predicates, companysecret.AccessKeyContains(req.AccessKey))
	}
	if req.SecretKey != "" {
		predicates = append(predicates, companysecret.SecretKeyContains(req.SecretKey))
	}
	if req.Enable != nil {
		predicates = append(predicates, companysecret.EnableEQ(*req.Enable))
	}
	isDelInt := uint32(0)
	predicates = append(predicates, companysecret.IsDelEQ(isDelInt))

	result, err := l.svcCtx.DB.CompanySecret.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	resp = &types.CompanySecretListResp{}
	resp.Data.Total = result.PageDetails.Total
	for _, v := range result.List {
		// 整理有效期时间
		validityDatetimeStr := ""
		// 判断数据是否为默认值， 如果为默认值则代表为无期限数据
		if v.ValidityDatetime.UnixMilli() != 1000 && v.ValidityDatetime.UnixMilli() != 0 {
			validityDatetime, err := common.IntUnixMilliToTime(int(v.ValidityDatetime.UnixMilli()))
			if err != nil {
				return nil, err
			}
			validityDatetimeStr = validityDatetime.Format("2006-01-02 15:04:05")
		}

		// 整理 rule 数据
		var ruleData RuleDataInfo
		if v.RuleData != "" {
			if v.RuleData != "" {
				err := json.Unmarshal([]byte(v.RuleData), &ruleData)
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
				Id:        &v.ID,
				CreatedAt: pointy.GetPointer(v.CreatedAt.UnixMilli()),
				UpdatedAt: pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			},
			AccessKey:             v.AccessKey,
			SecretKey:             v.SecretKey,
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
			Enable:                &v.Enable,
			IsDel:                 &v.IsDel,
			Description:           v.Description,
		})
	}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	return resp, nil
}
