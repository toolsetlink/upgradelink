package upgrade_company_traffic_packet

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/internal/logic/upgrade_company_traffic_packet"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
)

// swagger:route post /upgrade_company_traffic_packet/list upgrade_company_traffic_packet GetUpgradeCompanyTrafficPacketList
//
// Get upgrade company traffic packet list | 获取UpgradeCompanyTrafficPacket信息列表
//
// Get upgrade company traffic packet list | 获取UpgradeCompanyTrafficPacket信息列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeCompanyTrafficPacketListReq
//
// Responses:
//  200: UpgradeCompanyTrafficPacketListResp

func GetUpgradeCompanyTrafficPacketListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeCompanyTrafficPacketListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_company_traffic_packet.NewGetUpgradeCompanyTrafficPacketListLogic(r.Context(), svcCtx)
		resp, err := l.GetUpgradeCompanyTrafficPacketList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
