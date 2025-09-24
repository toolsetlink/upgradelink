package upgrade_company_income

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_company_income"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_company_income/create upgrade_company_income CreateUpgradeCompanyIncome
//
// Create upgrade company income information | 创建UpgradeCompanyIncome信息
//
// Create upgrade company income information | 创建UpgradeCompanyIncome信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeCompanyIncomeInfo
//
// Responses:
//  200: BaseMsgResp

func CreateUpgradeCompanyIncomeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeCompanyIncomeInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_company_income.NewCreateUpgradeCompanyIncomeLogic(r.Context(), svcCtx)
		resp, err := l.CreateUpgradeCompanyIncome(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
