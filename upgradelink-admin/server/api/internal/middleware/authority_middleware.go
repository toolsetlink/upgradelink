package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"upgradelink-admin/server/api/internal/common/enum"
	"upgradelink-admin/server/api/internal/common/http_error"
	"upgradelink-admin/server/api/internal/common/jwtctx/rolectx"
	"upgradelink-admin/server/api/internal/common/utils/jwt"

	"github.com/casbin/casbin/v3"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"

	"upgradelink-admin/server/api/internal/common/i18n"
)

type AuthorityMiddleware struct {
	Cbn   *casbin.Enforcer
	Rds   redis.UniversalClient
	Trans *i18n.Translator
}

func NewAuthorityMiddleware(cbn *casbin.Enforcer, rds redis.UniversalClient, trans *i18n.Translator) *AuthorityMiddleware {
	return &AuthorityMiddleware{
		Cbn:   cbn,
		Rds:   rds,
		Trans: trans,
	}
}

func (m *AuthorityMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		obj := r.URL.Path
		act := r.Method

		// get the role id
		roleIds, err := rolectx.GetRoleIDFromCtx(r.Context())
		if err != nil {
			httpx.Error(w, err)
			return
		}

		// check jwt blacklist
		jwtResult, err := m.Rds.Get(context.Background(), enum.RedisTokenPrefix+jwt.StripBearerPrefixFromToken(r.Header.Get("Authorization"))).Result()
		if err != nil && !errors.Is(err, redis.Nil) {
			logx.Errorw("redis error in jwt", logx.Field("detail", err.Error()))
			httpx.Error(w, http_error.NewApiError(http.StatusInternalServerError, err.Error()))
			return
		}
		if jwtResult == "1" {
			logx.Errorw("token in blacklist", logx.Field("detail", r.Header.Get("Authorization")))
			httpx.Error(w, http_error.NewApiErrorWithoutMsg(http.StatusUnauthorized))
			return
		}

		// 设置语言
		lang := r.Header.Get("Accept-Language")
		// 创建包含语言信息的新上下文
		newCtx := context.WithValue(r.Context(), "lang", lang)
		// 使用新上下文创建新请求
		newReq := r.WithContext(newCtx)

		// 效验权限
		result := batchCheck(m.Cbn, roleIds, act, obj)
		if result {
			logx.Infow("HTTP/HTTPS Request", logx.Field("UUID", newCtx.Value("userId")),
				logx.Field("path", obj), logx.Field("method", act))
			next(w, newReq)
			return
		} else {
			logx.Errorw("the role is not permitted to access the API", logx.Field("roleId", roleIds),
				logx.Field("path", obj), logx.Field("method", act))
			httpx.Error(w, http_error.NewApiError(http.StatusForbidden, m.Trans.Trans(
				newCtx,
				i18n.PermissionDeny)))
			return
		}
	}
}

func batchCheck(cbn *casbin.Enforcer, roleIds []string, act, obj string) bool {
	var checkReq [][]any
	for _, v := range roleIds {
		checkReq = append(checkReq, []any{v, obj, act})
	}
	fmt.Println(checkReq)

	result, err := cbn.BatchEnforce(checkReq)
	fmt.Println(result)
	fmt.Println(err)

	if err != nil {
		logx.Errorw("Casbin enforce error", logx.Field("detail", err.Error()))
		return false
	}

	for _, v := range result {
		if v {
			return true
		}
	}

	return false
}
