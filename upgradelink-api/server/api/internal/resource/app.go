package resource

import (
	"context"
	"errors"
	"fmt"

	"upgradelink-api/server/api/internal/resource/model"
)

const (
	CacheKeyAppTypeByAppKey = PREFIX + "APP_TYPE:APP_KEY:%v"
)

// GetAppTypeByAppKey 通过 apType 获取到对应的公司id
func (c *Ctx) GetAppTypeByAppKey(ctx context.Context, appKey string) (appType string, err error) {

	cacheKey := fmt.Sprintf(CacheKeyAppTypeByAppKey, appKey)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {

		// 通过唯一标识 获取到对应的应用信息
		urlInfo, err := c.GetUrlInfoByKey(ctx, appKey)
		if err != nil && errors.Is(err, model.ErrNotFound) {
		} else if err != nil {
			return "", err
		}
		if urlInfo != nil {
			return "url", nil
		}

		// 通过唯一标识 获取到对应的应用信息
		fileInfo, err := c.GetFileInfoByKey(ctx, appKey)
		if err != nil && errors.Is(err, model.ErrNotFound) {
		} else if err != nil {
			return "", err
		}
		if fileInfo != nil {
			return "file", nil
		}

		// 通过唯一标识 获取到对应的应用信息
		tauriInfo, err := c.GetTauriInfoByKey(ctx, appKey)
		if err != nil && errors.Is(err, model.ErrNotFound) {
		} else if err != nil {
			return "", err
		}
		if tauriInfo != nil {
			return "tauri", nil
		}

		// 通过唯一标识 获取到对应的应用信息
		apkInfo, err := c.GetApkInfoByKey(ctx, appKey)
		if err != nil && errors.Is(err, model.ErrNotFound) {
		} else if err != nil {
			return "", err
		}
		if apkInfo != nil {
			return "apk", nil
		}

		// 通过唯一标识 获取到对应的应用信息
		configurationInfo, err := c.GetConfigurationInfoByKey(ctx, appKey)
		if err != nil && errors.Is(err, model.ErrNotFound) {
		} else if err != nil {
			return "", err
		}
		if configurationInfo != nil {
			return "configuration", nil
		}

		return "", nil
	})

	if v != nil {
		appType = v.(string)
		return appType, err
	}

	return "", nil
}
