package upgrade_company_income

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_company_income"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_company_income/list upgrade_company_income GetUpgradeCompanyIncomeList
//
// Get upgrade company income list | 获取UpgradeCompanyIncome信息列表
//
// Get upgrade company income list | 获取UpgradeCompanyIncome信息列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeCompanyIncomeListReq
//
// Responses:
//  200: UpgradeCompanyIncomeListResp

func GetUpgradeCompanyIncomeListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeCompanyIncomeListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_company_income.NewGetUpgradeCompanyIncomeListLogic(r.Context(), svcCtx)
		resp, err := l.GetUpgradeCompanyIncomeList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
