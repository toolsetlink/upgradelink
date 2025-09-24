package tasklog

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin-core/server/api/internal/logic/tasklog"
	"upgradelink-admin-core/server/api/internal/svc"
	"upgradelink-admin-core/server/api/internal/types"
)

// swagger:route post /task_log/list tasklog GetTaskLogList
//
// Get task log list | 获取任务日志列表
//
// Get task log list | 获取任务日志列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: TaskLogListReq
//
// Responses:
//  200: TaskLogListResp

func GetTaskLogListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TaskLogListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tasklog.NewGetTaskLogListLogic(r.Context(), svcCtx)
		resp, err := l.GetTaskLogList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
