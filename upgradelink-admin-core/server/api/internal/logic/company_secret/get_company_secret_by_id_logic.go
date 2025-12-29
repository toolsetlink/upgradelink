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

type GetCompanySecretByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCompanySecretByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCompanySecretByIdLogic {
	return &GetCompanySecretByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCompanySecretByIdLogic) GetCompanySecretById(req *types.IDReq) (resp *types.CompanySecretInfoResp, err error) {
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
	data, err := l.svcCtx.CoreRpc.GetCompanySecretById(l.ctx, &core.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	if data.CompanyId != userData.CompanyId {
		return nil, errorx.NewCodeInvalidArgumentError(i18n.TargetNotFound)
	}

	companySecretData, err := l.svcCtx.CoreRpc.GetCompanySecretById(l.ctx, &core.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	// 整理有效期时间
	validityDatetimeStr := ""
	// 判断数据是否为默认值， 如果为默认值则代表为无期限数据
	if int(*companySecretData.ValidityDatetime) != 1000 && int(*companySecretData.ValidityDatetime) != 0 {
		validityDatetime, err := utils.IntUnixMilliToTime(int(*companySecretData.ValidityDatetime))
		if err != nil {
			return nil, err
		}
		validityDatetimeStr = validityDatetime.Format("2006-01-02 15:04:05")
	}

	// 整理 rule 数据
	var ruleData RuleDataInfo
	if companySecretData.RuleData != nil {
		if *companySecretData.RuleData != "" {
			err := json.Unmarshal([]byte(*companySecretData.RuleData), &ruleData)
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

	return &types.CompanySecretInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.CompanySecretInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id:        companySecretData.Id,
				CreatedAt: companySecretData.CreatedAt,
				UpdatedAt: companySecretData.UpdatedAt,
			},
			AccessKey:             *companySecretData.AccessKey,
			SecretKey:             *companySecretData.SecretKey,
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
			Enable:                companySecretData.Enable,
			IsDel:                 companySecretData.IsDel,
			Description:           *companySecretData.Description,
		},
	}, nil
}
