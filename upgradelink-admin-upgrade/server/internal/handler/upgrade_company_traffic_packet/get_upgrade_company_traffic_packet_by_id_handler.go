package upgrade_company_traffic_packet

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_company_traffic_packet"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_company_traffic_packet upgrade_company_traffic_packet GetUpgradeCompanyTrafficPacketById
//
// Get upgrade company traffic packet by ID | 通过ID获取UpgradeCompanyTrafficPacket信息
//
// Get upgrade company traffic packet by ID | 通过ID获取UpgradeCompanyTrafficPacket信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
// Responses:
//  200: UpgradeCompanyTrafficPacketInfoResp

func GetUpgradeCompanyTrafficPacketByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_company_traffic_packet.NewGetUpgradeCompanyTrafficPacketByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetUpgradeCompanyTrafficPacketById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
