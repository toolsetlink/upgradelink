package upgrade_dev_group_relation

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"upgradelink-admin/server/api/internal/logic/upgrade_dev_group_relation"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"
)

func DeleteUpgradeDevGroupRelationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := upgrade_dev_group_relation.NewDeleteUpgradeDevGroupRelationLogic(r.Context(), svcCtx)
		resp, err := l.DeleteUpgradeDevGroupRelation(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
