package upgrade_company_traffic_packet

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_company_traffic_packet"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_company_traffic_packet/update upgrade_company_traffic_packet UpdateUpgradeCompanyTrafficPacket
//
// Update upgrade company traffic packet information | 更新UpgradeCompanyTrafficPacket信息
//
// Update upgrade company traffic packet information | 更新UpgradeCompanyTrafficPacket信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeCompanyTrafficPacketInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateUpgradeCompanyTrafficPacketHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeCompanyTrafficPacketInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_company_traffic_packet.NewUpdateUpgradeCompanyTrafficPacketLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUpgradeCompanyTrafficPacket(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
