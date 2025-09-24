package upgrade_company_income

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_company_income"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_company_income/update upgrade_company_income UpdateUpgradeCompanyIncome
//
// Update upgrade company income information | 更新UpgradeCompanyIncome信息
//
// Update upgrade company income information | 更新UpgradeCompanyIncome信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeCompanyIncomeInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateUpgradeCompanyIncomeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeCompanyIncomeInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_company_income.NewUpdateUpgradeCompanyIncomeLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUpgradeCompanyIncome(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
