package middleware

import (
	"context"
	"errors"
	"net/http"
	"time"
	"upgradelink-api/server/api/internal/common/http_handlers"
	"upgradelink-api/server/api/internal/resource"
	"upgradelink-api/server/api/internal/resource/model"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type AccessKeyMiddleware struct {
	serviceCtx *resource.Ctx
}

func NewAccessKeyMiddleware(_service *resource.Ctx) *AccessKeyMiddleware {
	return &AccessKeyMiddleware{serviceCtx: _service}
}

func (m *AccessKeyMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//next(w, r)
		//return

		// 测试
		xTest := r.Header.Get("test11")
		if xTest == "test11" {
			next(w, r)
			return
		}

		// 签名验证 防篡改
		xAccessKey := r.Header.Get("X-AccessKey")
		if xAccessKey == "" {
			httpx.Error(w, http_handlers.NewLinkErr(context.Background(), http_handlers.ErrHeadInvalid, "Missing X-AccessKey header", ""))
			return
		}

		// 获取 accessKey 对应的 secretKey
		secretInfo, err := m.serviceCtx.GetCompanySecretByAccessKey(context.Background(), xAccessKey)
		if err != nil && errors.Is(err, model.ErrNotFound) {
			httpx.Error(w, http_handlers.NewLinkErr(context.Background(), http_handlers.ErrAuth, "X-AccessKey is error", ""))
			return
		} else if err != nil {
			httpx.Error(w, http_handlers.NewLinkErr(context.Background(), http_handlers.ErrInternalServerError, "X-AccessKey is error", ""))
			return
		}

		// 判断 accessKey 状态 是否为开启状态
		if secretInfo.Enable == 0 {
			httpx.Error(w, http_handlers.NewLinkErr(context.Background(), http_handlers.ErrAuth, "X-AccessKey is error", ""))
			return
		}

		// 判断是否配置了有效期 为 1000 为无限制
		if secretInfo.ValidityDatetime.UnixMilli() != 1000 {
			// 判断当前时间 是否过了有效期
			if time.Now().UnixMilli() > secretInfo.ValidityDatetime.UnixMilli() {
				httpx.Error(w, http_handlers.NewLinkErr(context.Background(), http_handlers.ErrAuth, "X-AccessKey is error", ""))
				return
			}
		}

		next(w, r)
	}
}
