package middleware

import (
	"context"
	"net"
	"net/http"
	"strings"

	"upgradelink-api/server/api/internal/resource"

	"github.com/zeromicro/go-zero/core/logc"
)

type CdnRateLimitMiddleware struct {
	serviceCtx *resource.Ctx
}

func NewCdnRateLimitMiddleware(_service *resource.Ctx) *CdnRateLimitMiddleware {
	return &CdnRateLimitMiddleware{serviceCtx: _service}
}

func (m *CdnRateLimitMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		next(w, r)
		return
		
		// 测试
		xTest := r.Header.Get("test")
		if xTest == "test" {
			next(w, r)
			return
		}

		// 检查请求ip请求是否超出频率限制
		ip := getClientIP(r)

		logc.Info(context.Background(), "请求的ip为:", ip)

		// 检查秒级限制 (1秒最多2次)
		err := m.serviceCtx.GetRateLimitSecondByIp(context.Background(), ip)
		if err != nil {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}

		// 检查分钟级限制 (1分钟最多8次)
		err = m.serviceCtx.GetRateLimitMinuteByIp(context.Background(), ip)
		if err != nil {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}
		// 检查小时级限制 (1小时最多20次)
		err = m.serviceCtx.GetRateLimitHourByIp(context.Background(), ip)
		if err != nil {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}

		next(w, r)
	}
}

// getClientIP 获取客户端真实IP
func getClientIP(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip != "" {
		parts := strings.Split(ip, ",")
		return strings.TrimSpace(parts[0])
	}

	ip = r.Header.Get("X-Real-IP")
	if ip != "" {
		return ip
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}
