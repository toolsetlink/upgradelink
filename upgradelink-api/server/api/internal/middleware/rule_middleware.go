package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"upgradelink-api/server/api/internal/common"
	"upgradelink-api/server/api/internal/common/http_handlers"
	"upgradelink-api/server/api/internal/resource"
	"upgradelink-api/server/api/internal/resource/model"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type RuleMiddleware struct {
	serviceCtx *resource.Ctx
}

func NewRuleMiddleware(_service *resource.Ctx) *RuleMiddleware {
	return &RuleMiddleware{serviceCtx: _service}
}

type ReqPostDataInfo struct {
	UrlKey           string `json:"urlKey"`
	FileKey          string `json:"fileKey"`
	ConfigurationKey string `json:"configurationKey"`
	TauriKey         string `json:"tauriKey"`
	ElectronKey      string `json:"electronKey"`
	ApkKey           string `json:"apkKey"`
	WinKey           string `json:"winKey"`
	MacKey           string `json:"macKey"`
	LnxKey           string `json:"lnxKey"`
	AppKey           string `json:"appKey"`
}

type RuleDataInfo struct {
	Url           []int `json:"url,optional"`
	File          []int `json:"file,optional"`
	Configuration []int `json:"configuration,optional"`
	Tauri         []int `json:"tauri,optional"`
	Electron      []int `json:"electron,optional"`
	Apk           []int `json:"apk,optional"`
	Win           []int `json:"win,optional"`
	Mac           []int `json:"mac,optional"`
	Lnx           []int `json:"lnx,optional"`
}

// getAppID 根据应用类型和key获取应用ID
func (m *RuleMiddleware) getAppID(ctx context.Context, companyId int64, appType string, key string) (int64, error) {
	switch appType {
	case "url":
		info, err := m.serviceCtx.GetUrlInfoByCompanyIdAndKey(ctx, companyId, key)
		if err != nil {
			return 0, err
		}
		return info.Id, nil
	case "file":
		info, err := m.serviceCtx.GetFileInfoByCompanyIdAndKey(ctx, companyId, key)
		if err != nil {
			return 0, err
		}
		return info.Id, nil
	case "configuration":
		info, err := m.serviceCtx.GetConfigurationInfoByCompanyIdAndKey(ctx, companyId, key)
		if err != nil {
			return 0, err
		}
		return info.Id, nil
	case "tauri":
		info, err := m.serviceCtx.GetTauriInfoByCompanyIdAndKey(ctx, companyId, key)
		if err != nil {
			return 0, err
		}
		return info.Id, nil
	case "electron":
		info, err := m.serviceCtx.GetElectronInfoByCompanyIdAndKey(ctx, companyId, key)
		if err != nil {
			return 0, err
		}
		return info.Id, nil
	case "apk":
		info, err := m.serviceCtx.GetApkInfoByCompanyIdAndKey(ctx, companyId, key)
		if err != nil {
			return 0, err
		}
		return info.Id, nil
	case "win":
		info, err := m.serviceCtx.GetWinInfoByCompanyIdAndKey(ctx, companyId, key)
		if err != nil {
			return 0, err
		}
		return info.Id, nil
	case "mac":
		info, err := m.serviceCtx.GetMacInfoByCompanyIdAndKey(ctx, companyId, key)
		if err != nil {
			return 0, err
		}
		return info.Id, nil
	case "lnx":
		info, err := m.serviceCtx.GetLnxInfoByCompanyIdAndKey(ctx, companyId, key)
		if err != nil {
			return 0, err
		}
		return info.Id, nil
	default:
		return 0, errors.New("invalid app type")
	}
}

// verifyAppPermission 验证应用是否在配置的应用列表中
func (m *RuleMiddleware) verifyAppPermission(ruleData RuleDataInfo, appType string, appID int64) bool {

	switch appType {
	case "url":
		return common.IntInSlice(int(appID), ruleData.Url)
	case "file":
		return common.IntInSlice(int(appID), ruleData.File)
	case "configuration":
		return common.IntInSlice(int(appID), ruleData.Configuration)
	case "tauri":
		return common.IntInSlice(int(appID), ruleData.Tauri)
	case "electron":
		return common.IntInSlice(int(appID), ruleData.Electron)
	case "apk":
		return common.IntInSlice(int(appID), ruleData.Apk)
	case "win":
		return common.IntInSlice(int(appID), ruleData.Win)
	case "mac":
		return common.IntInSlice(int(appID), ruleData.Mac)
	case "lnx":
		return common.IntInSlice(int(appID), ruleData.Lnx)
	default:
		return false
	}
}

// parseRequest 解析请求并返回应用类型和key，自动处理GET和POST请求
func (m *RuleMiddleware) parseRequest(r *http.Request) (string, string, error) {
	switch r.Method {
	case http.MethodGet:
		// 解析GET请求参数
		if urlKey := r.URL.Query().Get("urlKey"); urlKey != "" {
			return "url", urlKey, nil
		}
		if fileKey := r.URL.Query().Get("fileKey"); fileKey != "" {
			return "file", fileKey, nil
		}
		if configurationKey := r.URL.Query().Get("configurationKey"); configurationKey != "" {
			return "configuration", configurationKey, nil
		}
		if tauriKey := r.URL.Query().Get("tauriKey"); tauriKey != "" {
			return "tauri", tauriKey, nil
		}
		if electronKey := r.URL.Query().Get("electronKey"); electronKey != "" {
			return "electron", electronKey, nil
		}
		if apkKey := r.URL.Query().Get("apkKey"); apkKey != "" {
			return "apk", apkKey, nil
		}
		if winKey := r.URL.Query().Get("winKey"); winKey != "" {
			return "win", winKey, nil
		}
		if macKey := r.URL.Query().Get("macKey"); macKey != "" {
			return "mac", macKey, nil
		}
		if lnxKey := r.URL.Query().Get("lnxKey"); lnxKey != "" {
			return "lnx", lnxKey, nil
		}
		if appKey := r.URL.Query().Get("appKey"); appKey != "" {
			appType, err := m.serviceCtx.GetAppTypeByAppKey(context.Background(), appKey)
			if err != nil {
				return "", "", err
			}
			return appType, appKey, nil
		}
	case http.MethodPost:
		// 解析POST请求体
		bodyByte, err := io.ReadAll(r.Body)
		if err != nil {
			return "", "", err
		}
		r.Body = io.NopCloser(bytes.NewBuffer(bodyByte))

		var reqPostDataInfo ReqPostDataInfo
		err = json.Unmarshal(bodyByte, &reqPostDataInfo)
		if err != nil {
			return "", "", err
		}

		if reqPostDataInfo.UrlKey != "" {
			return "url", reqPostDataInfo.UrlKey, nil
		}
		if reqPostDataInfo.FileKey != "" {
			return "file", reqPostDataInfo.FileKey, nil
		}
		if reqPostDataInfo.ConfigurationKey != "" {
			return "configuration", reqPostDataInfo.ConfigurationKey, nil
		}
		if reqPostDataInfo.TauriKey != "" {
			return "tauri", reqPostDataInfo.TauriKey, nil
		}
		if reqPostDataInfo.ElectronKey != "" {
			return "electron", reqPostDataInfo.ElectronKey, nil
		}
		if reqPostDataInfo.ApkKey != "" {
			return "apk", reqPostDataInfo.ApkKey, nil
		}
		if reqPostDataInfo.WinKey != "" {
			return "win", reqPostDataInfo.WinKey, nil
		}
		if reqPostDataInfo.MacKey != "" {
			return "mac", reqPostDataInfo.MacKey, nil
		}
		if reqPostDataInfo.LnxKey != "" {
			return "lnx", reqPostDataInfo.LnxKey, nil
		}
		if reqPostDataInfo.AppKey != "" {
			appType, err := m.serviceCtx.GetAppTypeByAppKey(context.Background(), reqPostDataInfo.AppKey)
			if err != nil {
				return "", "", err
			}
			return appType, reqPostDataInfo.AppKey, nil
		}
	}

	return "", "", errors.New("no app key found in request")
}

// handleRequest 处理请求并验证应用权限，自动支持GET和POST请求
func (m *RuleMiddleware) handleRequest(w http.ResponseWriter, r *http.Request, companyId int64, ruleData RuleDataInfo) bool {
	ctx := r.Context()
	appType, appKey, err := m.parseRequest(r)
	if err != nil {
		httpx.Error(w, http_handlers.NewLinkErr(ctx, http_handlers.ErrAuth, "X-AccessKey check appKey is error", ""))
		return false
	}

	appID, err := m.getAppID(ctx, companyId, appType, appKey)
	if err != nil && errors.Is(err, model.ErrNotFound) {
		httpx.Error(w, http_handlers.NewLinkErr(ctx, http_handlers.ErrNotFound, common.ErrApp2Msg, common.ErrApp2Docs))
		return false
	} else if err != nil {
		httpx.Error(w, http_handlers.NewLinkErr(ctx, http_handlers.ErrInternalServerError, common.Err1Msg, common.Err1Docs))
		return false
	}

	if !m.verifyAppPermission(ruleData, appType, appID) {
		httpx.Error(w, http_handlers.NewLinkErr(ctx, http_handlers.ErrAuth, "X-AccessKey check appKey is error", ""))
		return false
	}

	return true
}

func (m *RuleMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// 签名验证 防篡改
		xAccessKey := r.Header.Get("X-AccessKey")
		if xAccessKey == "" {
			httpx.Error(w, http_handlers.NewLinkErr(r.Context(), http_handlers.ErrHeadInvalid, "Missing X-AccessKey header", ""))
			return
		}

		// accessKey 配置信息
		secretInfo, err := m.serviceCtx.GetCompanySecretByAccessKey(r.Context(), xAccessKey)
		if err != nil && errors.Is(err, model.ErrNotFound) {
			httpx.Error(w, http_handlers.NewLinkErr(r.Context(), http_handlers.ErrAuth, "X-AccessKey is invalid", ""))
			return
		} else if err != nil {
			httpx.Error(w, http_handlers.NewLinkErr(r.Context(), http_handlers.ErrInternalServerError, common.Err1Msg, common.Err1Docs))
			return
		}

		// 判断 accessKey 配置的应用权限
		if secretInfo.RuleData != "" {
			// 解析 RuleData 数据结构
			var ruleData RuleDataInfo
			err = json.Unmarshal([]byte(secretInfo.RuleData), &ruleData)
			if err != nil {
				httpx.Error(w, http_handlers.NewLinkErr(r.Context(), http_handlers.ErrInternalServerError, common.Err1Msg, common.Err1Docs))
				return
			}

			// 处理GET和POST请求的应用权限验证
			if r.Method == http.MethodGet || r.Method == http.MethodPost {
				if !m.handleRequest(w, r, secretInfo.CompanyId, ruleData) {
					httpx.Error(w, http_handlers.NewLinkErr(r.Context(), http_handlers.ErrAuth, "X-AccessKey check appKey is error", ""))
					return
				}
			}

		}

		next(w, r)
	}
}
