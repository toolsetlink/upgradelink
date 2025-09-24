package upgrade_company_traffic_packet

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-upgrade/internal/logic/upgrade_company_traffic_packet"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
)

// swagger:route post /upgrade_company_traffic_packet/delete upgrade_company_traffic_packet DeleteUpgradeCompanyTrafficPacket
//
// Delete upgrade company traffic packet information | 删除UpgradeCompanyTrafficPacket信息
//
// Delete upgrade company traffic packet information | 删除UpgradeCompanyTrafficPacket信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDsReq
//
// Responses:
//  200: BaseMsgResp

func DeleteUpgradeCompanyTrafficPacketHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDsReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_company_traffic_packet.NewDeleteUpgradeCompanyTrafficPacketLogic(r.Context(), svcCtx)
		resp, err := l.DeleteUpgradeCompanyTrafficPacket(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
