package upgrade_company_traffic_packet

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/server/internal/logic/upgrade_company_traffic_packet"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
)

// swagger:route post /upgrade_company_traffic_packet/create upgrade_company_traffic_packet CreateUpgradeCompanyTrafficPacket
//
// Create upgrade company traffic packet information | 创建UpgradeCompanyTrafficPacket信息
//
// Create upgrade company traffic packet information | 创建UpgradeCompanyTrafficPacket信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: UpgradeCompanyTrafficPacketCreateInfo
//
// Responses:
//  200: BaseMsgResp

func CreateUpgradeCompanyTrafficPacketHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpgradeCompanyTrafficPacketCreateInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_company_traffic_packet.NewCreateUpgradeCompanyTrafficPacketLogic(r.Context(), svcCtx)
		resp, err := l.CreateUpgradeCompanyTrafficPacket(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
