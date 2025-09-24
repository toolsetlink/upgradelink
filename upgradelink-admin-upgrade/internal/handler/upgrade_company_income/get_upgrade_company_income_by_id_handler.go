package upgrade_company_income

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/internal/logic/upgrade_company_income"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
)

// swagger:route post /upgrade_company_income upgrade_company_income GetUpgradeCompanyIncomeById
//
// Get upgrade company income by ID | 通过ID获取UpgradeCompanyIncome信息
//
// Get upgrade company income by ID | 通过ID获取UpgradeCompanyIncome信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
// Responses:
//  200: UpgradeCompanyIncomeInfoResp

func GetUpgradeCompanyIncomeByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_company_income.NewGetUpgradeCompanyIncomeByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetUpgradeCompanyIncomeById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
