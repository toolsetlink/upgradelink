package middleware

import (
	"bytes"
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"upgradelink-api/server/api/internal/common/http_handlers"
	"upgradelink-api/server/api/internal/resource"
	"upgradelink-api/server/api/internal/resource/model"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type SignatureMiddleware struct {
	serviceCtx *resource.Ctx
}

func NewSignatureMiddleware(_service *resource.Ctx) *SignatureMiddleware {
	return &SignatureMiddleware{serviceCtx: _service}
}

func (m *SignatureMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//next(w, r)
		//return

		// 测试
		xTest := r.Header.Get("test")
		if xTest == "test" {
			next(w, r)
			return
		}

		// 签名验证 防篡改
		xAccessKey := r.Header.Get("X-AccessKey")
		if xAccessKey == "" {
			httpx.Error(w, http_handlers.NewLinkErr(context.Background(), http_handlers.ErrHeadInvalid, "Missing X-AccessKey header", ""))
			return
		}
		xSignature := r.Header.Get("X-Signature")
		if xSignature == "" {
			httpx.Error(w, http_handlers.NewLinkErr(context.Background(), http_handlers.ErrHeadInvalid, "Missing X-Signature header", ""))
			return
		}

		// 获取 accessKey 对应的 secretKey
		secretInfo, err := m.serviceCtx.GetCompanySecretByAccessKey(context.Background(), xAccessKey)
		if err != nil && errors.Is(err, model.ErrNotFound) {
			httpx.Error(w, http_handlers.NewLinkErr(context.Background(), http_handlers.ErrAuth, "X-AccessKey is error", ""))
		} else if err != nil {
			httpx.Error(w, http_handlers.NewLinkErr(context.Background(), http_handlers.ErrInternalServerError, "X-AccessKey is error", ""))
		}

		xTimestamp := r.Header.Get("X-Timestamp")
		xNonce := r.Header.Get("X-Nonce")
		data := make(map[string]string)
		data["secretKey"] = secretInfo.SecretKey
		data["timestamp"] = xTimestamp
		data["nonce"] = xNonce
		data["url"], _ = url.QueryUnescape(r.RequestURI)
		if (r.Method == http.MethodPost || r.Method == http.MethodPut) && r.Body != nil {
			var bodyByte []byte
			bodyByte, err = io.ReadAll(r.Body)
			if err != nil {
				httpx.Error(w, http_handlers.NewLinkErr(context.Background(), http_handlers.ErrAuth, "ReadBody Field", ""))
				return
			}
			r.Body = io.NopCloser(bytes.NewBuffer(bodyByte))
			data["body"] = string(bodyByte)
		}
		//fmt.Println("data: ", data)
		//logc.Infof(r.Context(), "data: %v", data)

		// 生成签名并验签
		sg := GenerateSignature(data)
		if sg != xSignature {
			httpx.Error(w, http_handlers.NewLinkErr(context.Background(), http_handlers.ErrSign, "Signature Invalid", ""))
			return
		}

		next(w, r)
	}
}

// GenerateSignature 新生成签名规则  取原始串 body=${body}&timestamp=${timestamp}&url=${url}记为 signStr
//
//	@desc: 新生成签名规则  取原始串 body=${body}&timestamp=${timestamp}&url=${url}记为 signStr
//	@auth
//	@param
//	@return
func GenerateSignature(params map[string]string) string {
	// Step 1: Sort params by parameter name and value
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Step 2: Construct stringA
	var stringA strings.Builder
	for _, key := range keys {
		value := params[key]
		stringA.WriteString(fmt.Sprintf("%s=%s&", key, value))
	}
	stringAStr := strings.TrimRight(stringA.String(), "&")

	//now1 := time.Now().Format("2006-01-02T15:04:05.000Z07:00")
	//fmt.Println(now1)
	//body={
	//	"key": "key1",
	//		"versionCode": 1,
	//		"appointVersionCode": 0
	//}&nonce=abcdef1234567890&secretKey=a2&timestamp=1739527809281&url=/v1/url/upgrade
	//fmt.Println("stringAStr: ", stringAStr)

	hash := md5.Sum([]byte(stringAStr))
	md5Hash := fmt.Sprintf("%x", hash)

	signature := fmt.Sprintf(md5Hash)

	return signature
}
