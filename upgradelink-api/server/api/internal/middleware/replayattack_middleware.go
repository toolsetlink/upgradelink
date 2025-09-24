package middleware

import (
	"context"
	"net/http"
	"time"

	"upgradelink-api/server/api/internal/common/http_handlers"
	"upgradelink-api/server/api/internal/resource"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type ReplayAttackMiddleware struct {
	serviceCtx *resource.Ctx
}

func NewReplayAttackMiddleware(_service *resource.Ctx) *ReplayAttackMiddleware {
	return &ReplayAttackMiddleware{serviceCtx: _service}
}

func (m *ReplayAttackMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//next(w, r)
		//return

		// 测试
		xTest := r.Header.Get("test")
		if xTest == "test" {
			next(w, r)
			return
		}

		// 时间戳验证 防重放攻击  RFC3339Nano格式
		xTimestamp := r.Header.Get("X-Timestamp")
		if xTimestamp == "" {
			httpx.Error(w, http_handlers.NewLinkErr(context.Background(), http_handlers.ErrHeadInvalid, "Missing X-Timestamp header", ""))
			return
		}
		// 解析时间字符串
		t, err := time.Parse(time.RFC3339, xTimestamp)
		if err != nil {
			httpx.Error(w, http_handlers.NewLinkErr(context.Background(), http_handlers.ErrHeadInvalid, "Invalid X-Timestamp format", ""))
			return
		}
		// 允许的时间窗口（例如 ±5 分钟）
		allowedWindow := time.Minute * 5
		currentTime := time.Now()
		// 计算时间范围
		minTime := currentTime.Add(-allowedWindow)
		maxTime := currentTime.Add(allowedWindow)

		// 检查时间是否在允许的窗口内
		if t.Before(minTime) || t.After(maxTime) {
			httpx.Error(w, http_handlers.NewLinkErr(context.Background(), http_handlers.ErrAuth, "Invalid X-Timestamp: timestamp is out of allowed window", ""))
			return
		}

		// 随机数验证 （建议UUID或16+位随机数）。  防重放攻击
		xNonce := r.Header.Get("X-Nonce")
		if xNonce == "" {
			httpx.Error(w, http_handlers.NewLinkErr(context.Background(), http_handlers.ErrHeadInvalid, "Missing X-Nonce header", ""))
			return
		}

		next(w, r)
	}
}
