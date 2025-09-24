package middleware

import (
	"net/http"

	"upgradelink-api/server/api/internal/resource"
)

type RateLimitMiddleware struct {
	serviceCtx *resource.Ctx
}

func NewRateLimitMiddleware(_service *resource.Ctx) *RateLimitMiddleware {
	return &RateLimitMiddleware{serviceCtx: _service}
}

func (m *RateLimitMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		next(w, r)
	}
}
