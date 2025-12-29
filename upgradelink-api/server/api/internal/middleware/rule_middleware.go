package middleware

import "net/http"

type RuleMiddleware struct {
}

func NewRuleMiddleware() *RuleMiddleware {
	return &RuleMiddleware{}
}

func (m *RuleMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need
		next(w, r)
	}
}
