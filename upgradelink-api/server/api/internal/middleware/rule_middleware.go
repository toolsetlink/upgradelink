package middleware

import (
	"net/http"
	"upgradelink-api/server/api/internal/resource"
)

type RuleMiddleware struct {
	serviceCtx *resource.Ctx
}

func NewRuleMiddleware(_service *resource.Ctx) *RuleMiddleware {
	return &RuleMiddleware{serviceCtx: _service}
}

func (m *RuleMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
	}
}
