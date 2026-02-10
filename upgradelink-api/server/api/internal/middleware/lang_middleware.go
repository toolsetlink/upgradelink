package middleware

import (
	"context"
	"net/http"
	"upgradelink-api/server/api/internal/resource"
)

type LangMiddleware struct {
	serviceCtx *resource.Ctx
}

func NewLangMiddleware(_service *resource.Ctx) *LangMiddleware {
	return &LangMiddleware{serviceCtx: _service}
}

func (m *LangMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 设置语言
		lang := r.Header.Get("Accept-Language")
		// 创建包含语言信息的新上下文
		newCtx := context.WithValue(r.Context(), "lang", lang)
		// 使用新上下文创建新请求
		newReq := r.WithContext(newCtx)

		next(w, newReq)
	}
}
