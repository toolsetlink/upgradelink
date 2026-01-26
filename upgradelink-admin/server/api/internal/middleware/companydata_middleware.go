package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"upgradelink-admin/server/api/internal/common/http_error"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"
	"upgradelink-admin/server/api/internal/ent"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// CompanyDataMiddleware 数据所属权校验中间件
type CompanyDataMiddleware struct {
	DB *ent.Client
}

// NewCompanyDataMiddleware 创建数据所属权校验中间件
func NewCompanyDataMiddleware(db *ent.Client) *CompanyDataMiddleware {
	return &CompanyDataMiddleware{
		DB: db,
	}
}

// Handle 处理数据所属权校验
func (m *CompanyDataMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 获取请求路径和方法
		path := r.URL.Path
		method := r.Method

		// 只有特定的请求方法需要校验数据所属权
		if method != http.MethodPost {
			next(w, r)
			return
		}

		// 获取API信息
		apiInfo, ok := getApiInfo(path)
		if !ok {
			// 不是需要校验的API，直接放行
			next(w, r)
			return
		}

		// 获取companyId
		companyId := companyctx.GetCompanyIDFromCtx(r.Context())

		// 读取请求体内容
		var buf bytes.Buffer
		_, err := io.ReadAll(io.TeeReader(r.Body, &buf))
		if err != nil {
			httpx.Error(w, http_error.NewApiInternalError("failed to read request body"))
			return
		}

		// 重置请求体，以便后续handler可以再次读取
		r.Body = io.NopCloser(bytes.NewReader(buf.Bytes()))

		// 解析请求体，获取ID
		var reqMap map[string]interface{}
		if err := json.NewDecoder(bytes.NewReader(buf.Bytes())).Decode(&reqMap); err != nil {
			httpx.Error(w, http_error.NewApiInternalError("failed to parse request body"))
			return
		}

		// 获取ID值
		idValue, ok := reqMap["id"]
		if !ok {
			// 没有ID参数，直接放行
			next(w, r)
			return
		}

		// 校验数据所属权
		if err := m.validateDataOwnership(r.Context(), apiInfo.table, idValue, companyId); err != nil {
			httpx.Error(w, err)
			return
		}

		// 校验通过，继续处理请求
		next(w, r)
	}
}

// apiInfo API信息
type apiInfo struct {
	path  string
	table string
}

// getApiInfo 根据请求路径获取API信息
func getApiInfo(path string) (apiInfo, bool) {
	// 去除可能的前缀
	path = strings.TrimPrefix(path, "/")

	// API路径到数据库表的映射
	apiMap := map[string]string{
		"upgrade_apk":                            "upgrade_apk",
		"upgrade_apk_upgrade_strategy":           "upgrade_apk_upgrade_strategy",
		"upgrade_apk_version":                    "upgrade_apk_version",
		"upgrade_electron":                       "upgrade_electron",
		"upgrade_electron_upgrade_strategy":      "upgrade_electron_upgrade_strategy",
		"upgrade_electron_version":               "upgrade_electron_version",
		"upgrade_tauri":                          "upgrade_tauri",
		"upgrade_tauri_upgrade_strategy":         "upgrade_tauri_upgrade_strategy",
		"upgrade_tauri_version":                  "upgrade_tauri_version",
		"upgrade_win":                            "upgrade_win",
		"upgrade_win_upgrade_strategy":           "upgrade_win_upgrade_strategy",
		"upgrade_win_version":                    "upgrade_win_version",
		"upgrade_mac":                            "upgrade_mac",
		"upgrade_mac_upgrade_strategy":           "upgrade_mac_upgrade_strategy",
		"upgrade_mac_version":                    "upgrade_mac_version",
		"upgrade_lnx":                            "upgrade_lnx",
		"upgrade_lnx_upgrade_strategy":           "upgrade_lnx_upgrade_strategy",
		"upgrade_lnx_version":                    "upgrade_lnx_version",
		"upgrade_file":                           "upgrade_file",
		"upgrade_file_upgrade_strategy":          "upgrade_file_upgrade_strategy",
		"upgrade_file_version":                   "upgrade_file_version",
		"upgrade_url":                            "upgrade_url",
		"upgrade_url_upgrade_strategy":           "upgrade_url_upgrade_strategy",
		"upgrade_url_version":                    "upgrade_url_version",
		"upgrade_configuration":                  "upgrade_configuration",
		"upgrade_configuration_upgrade_strategy": "upgrade_configuration_upgrade_strategy",
		"upgrade_configuration_version":          "upgrade_configuration_version",
		"upgrade_dev":                            "upgrade_dev",
		"upgrade_dev_group":                      "upgrade_dev_group",
		"upgrade_dev_model":                      "upgrade_dev_model",
	}

	// 查找匹配的API路径
	for apiPath, table := range apiMap {
		if strings.HasPrefix(path, apiPath) {
			return apiInfo{
				path:  apiPath,
				table: table,
			}, true
		}
	}

	return apiInfo{}, false
}

// validateDataOwnership 校验数据所属权
func (m *CompanyDataMiddleware) validateDataOwnership(ctx context.Context, table string, idValue interface{}, companyId int) error {
	// 将ID转换为int
	id, ok := idValue.(float64)
	if !ok {
		return http_error.NewApiInternalError("invalid id type")
	}

	// 根据表名查询数据并校验companyId
	switch table {
	case "upgrade_apk":
		data, err := m.DB.UpgradeApk.Get(ctx, int(id))
		if err != nil {
			return http_error.NewApiInternalError("failed to get data")
		}
		if data.CompanyID != companyId {
			return http_error.NewApiForbiddenError("no permission to access this data")
		}
	case "upgrade_apk_upgrade_strategy":
		data, err := m.DB.UpgradeApkUpgradeStrategy.Get(ctx, int(id))
		if err != nil {
			return http_error.NewApiInternalError("failed to get data")
		}
		if data.CompanyID != companyId {
			return http_error.NewApiForbiddenError("no permission to access this data")
		}
	case "upgrade_apk_version":
		data, err := m.DB.UpgradeApkVersion.Get(ctx, int(id))
		if err != nil {
			return http_error.NewApiInternalError("failed to get data")
		}
		if data.CompanyID != companyId {
			return http_error.NewApiForbiddenError("no permission to access this data")
		}
	case "upgrade_electron":
		data, err := m.DB.UpgradeElectron.Get(ctx, int(id))
		if err != nil {
			return http_error.NewApiInternalError("failed to get data")
		}
		if data.CompanyID != companyId {
			return http_error.NewApiForbiddenError("no permission to access this data")
		}
	case "upgrade_electron_upgrade_strategy":
		data, err := m.DB.UpgradeElectronUpgradeStrategy.Get(ctx, int(id))
		if err != nil {
			return http_error.NewApiInternalError("failed to get data")
		}
		if data.CompanyID != companyId {
			return http_error.NewApiForbiddenError("no permission to access this data")
		}
	case "upgrade_electron_version":
		data, err := m.DB.UpgradeElectronVersion.Get(ctx, int(id))
		if err != nil {
			return http_error.NewApiInternalError("failed to get data")
		}
		if data.CompanyID != companyId {
			return http_error.NewApiForbiddenError("no permission to access this data")
		}
	case "upgrade_tauri":
		data, err := m.DB.UpgradeTauri.Get(ctx, int(id))
		if err != nil {
			return http_error.NewApiInternalError("failed to get data")
		}
		if data.CompanyID != companyId {
			return http_error.NewApiForbiddenError("no permission to access this data")
		}
	case "upgrade_tauri_upgrade_strategy":
		data, err := m.DB.UpgradeTauriUpgradeStrategy.Get(ctx, int(id))
		if err != nil {
			return http_error.NewApiInternalError("failed to get data")
		}
		if data.CompanyID != companyId {
			return http_error.NewApiForbiddenError("no permission to access this data")
		}
	case "upgrade_tauri_version":
		data, err := m.DB.UpgradeTauriVersion.Get(ctx, int(id))
		if err != nil {
			return http_error.NewApiInternalError("failed to get data")
		}
		if data.CompanyID != companyId {
			return http_error.NewApiForbiddenError("no permission to access this data")
		}
	case "upgrade_win":
		data, err := m.DB.UpgradeWin.Get(ctx, int(id))
		if err != nil {
			return http_error.NewApiInternalError("failed to get data")
		}
		if data.CompanyID != companyId {
			return http_error.NewApiForbiddenError("no permission to access this data")
		}
	case "upgrade_win_upgrade_strategy":
		data, err := m.DB.UpgradeWinUpgradeStrategy.Get(ctx, int(id))
		if err != nil {
			return http_error.NewApiInternalError("failed to get data")
		}
		if data.CompanyID != companyId {
			return http_error.NewApiForbiddenError("no permission to access this data")
		}
	case "upgrade_win_version":
		data, err := m.DB.UpgradeWinVersion.Get(ctx, int(id))
		if err != nil {
			return http_error.NewApiInternalError("failed to get data")
		}
		if data.CompanyID != companyId {
			return http_error.NewApiForbiddenError("no permission to access this data")
		}
	case "upgrade_mac":
		data, err := m.DB.UpgradeMac.Get(ctx, int(id))
		if err != nil {
			return http_error.NewApiInternalError("failed to get data")
		}
		if data.CompanyID != companyId {
			return http_error.NewApiForbiddenError("no permission to access this data")
		}
	case "upgrade_mac_upgrade_strategy":
		data, err := m.DB.UpgradeMacUpgradeStrategy.Get(ctx, int(id))
		if err != nil {
			return http_error.NewApiInternalError("failed to get data")
		}
		if data.CompanyID != companyId {
			return http_error.NewApiForbiddenError("no permission to access this data")
		}
	case "upgrade_mac_version":
		data, err := m.DB.UpgradeMacVersion.Get(ctx, int(id))
		if err != nil {
			return http_error.NewApiInternalError("failed to get data")
		}
		if data.CompanyID != companyId {
			return http_error.NewApiForbiddenError("no permission to access this data")
		}
	case "upgrade_lnx":
		data, err := m.DB.UpgradeLnx.Get(ctx, int(id))
		if err != nil {
			return http_error.NewApiInternalError("failed to get data")
		}
		if data.CompanyID != companyId {
			return http_error.NewApiForbiddenError("no permission to access this data")
		}
	case "upgrade_lnx_upgrade_strategy":
		data, err := m.DB.UpgradeLnxUpgradeStrategy.Get(ctx, int(id))
		if err != nil {
			return http_error.NewApiInternalError("failed to get data")
		}
		if data.CompanyID != companyId {
			return http_error.NewApiForbiddenError("no permission to access this data")
		}
	case "upgrade_lnx_version":
		data, err := m.DB.UpgradeLnxVersion.Get(ctx, int(id))
		if err != nil {
			return http_error.NewApiInternalError("failed to get data")
		}
		if data.CompanyID != companyId {
			return http_error.NewApiForbiddenError("no permission to access this data")
		}
	case "upgrade_file":
		data, err := m.DB.UpgradeFile.Get(ctx, int(id))
		if err != nil {
			return http_error.NewApiInternalError("failed to get data")
		}
		if data.CompanyID != companyId {
			return http_error.NewApiForbiddenError("no permission to access this data")
		}
	case "upgrade_file_upgrade_strategy":
		data, err := m.DB.UpgradeFileUpgradeStrategy.Get(ctx, int(id))
		if err != nil {
			return http_error.NewApiInternalError("failed to get data")
		}
		if data.CompanyID != companyId {
			return http_error.NewApiForbiddenError("no permission to access this data")
		}
	case "upgrade_file_version":
		data, err := m.DB.UpgradeFileVersion.Get(ctx, int(id))
		if err != nil {
			return http_error.NewApiInternalError("failed to get data")
		}
		if data.CompanyID != companyId {
			return http_error.NewApiForbiddenError("no permission to access this data")
		}
	case "upgrade_url":
		data, err := m.DB.UpgradeUrl.Get(ctx, int(id))
		if err != nil {
			return http_error.NewApiInternalError("failed to get data")
		}
		if data.CompanyID != companyId {
			return http_error.NewApiForbiddenError("no permission to access this data")
		}
	case "upgrade_url_upgrade_strategy":
		data, err := m.DB.UpgradeUrlUpgradeStrategy.Get(ctx, int(id))
		if err != nil {
			return http_error.NewApiInternalError("failed to get data")
		}
		if data.CompanyID != companyId {
			return http_error.NewApiForbiddenError("no permission to access this data")
		}
	case "upgrade_url_version":
		data, err := m.DB.UpgradeUrlVersion.Get(ctx, int(id))
		if err != nil {
			return http_error.NewApiInternalError("failed to get data")
		}
		if data.CompanyID != companyId {
			return http_error.NewApiForbiddenError("no permission to access this data")
		}
	case "upgrade_configuration":
		data, err := m.DB.UpgradeConfiguration.Get(ctx, int(id))
		if err != nil {
			return http_error.NewApiInternalError("failed to get data")
		}
		if data.CompanyID != companyId {
			return http_error.NewApiForbiddenError("no permission to access this data")
		}
	case "upgrade_configuration_upgrade_strategy":
		data, err := m.DB.UpgradeConfigurationUpgradeStrategy.Get(ctx, int(id))
		if err != nil {
			return http_error.NewApiInternalError("failed to get data")
		}
		if data.CompanyID != companyId {
			return http_error.NewApiForbiddenError("no permission to access this data")
		}
	case "upgrade_configuration_version":
		data, err := m.DB.UpgradeConfigurationVersion.Get(ctx, int(id))
		if err != nil {
			return http_error.NewApiInternalError("failed to get data")
		}
		if data.CompanyID != companyId {
			return http_error.NewApiForbiddenError("no permission to access this data")
		}
	case "upgrade_dev":
		data, err := m.DB.UpgradeDev.Get(ctx, int(id))
		if err != nil {
			return http_error.NewApiInternalError("failed to get data")
		}
		if data.CompanyID != companyId {
			return http_error.NewApiForbiddenError("no permission to access this data")
		}
	case "upgrade_dev_group":
		data, err := m.DB.UpgradeDevGroup.Get(ctx, int(id))
		if err != nil {
			return http_error.NewApiInternalError("failed to get data")
		}
		if data.CompanyID != companyId {
			return http_error.NewApiForbiddenError("no permission to access this data")
		}
	case "upgrade_dev_model":
		data, err := m.DB.UpgradeDevModel.Get(ctx, int(id))
		if err != nil {
			return http_error.NewApiInternalError("failed to get data")
		}
		if data.CompanyID != companyId {
			return http_error.NewApiForbiddenError("no permission to access this data")
		}
	default:
		// 未知表，直接放行
		return nil
	}

	return nil
}
