package tasklog

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-core/server/api/internal/logic/tasklog"
	"upgradelink-admin-core/server/api/internal/svc"
	"upgradelink-admin-core/server/api/internal/types"
)

// swagger:route post /task_log/update tasklog UpdateTaskLog
//
// Update task log information | 更新任务日志
//
// Update task log information | 更新任务日志
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: TaskLogInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateTaskLogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TaskLogInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tasklog.NewUpdateTaskLogLogic(r.Context(), svcCtx)
		resp, err := l.UpdateTaskLog(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
