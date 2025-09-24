package resource

import (
	"context"
	"errors"
	"fmt"

	"upgradelink-api/server/api/internal/resource/model"
)

const (
	CacheKeyCompanyIdByAppKey = PREFIX + "COMPANY_ID:APP_KEY:%v"
)

// GetCompanyIdByAppKey 通过 appkey 获取到对应的公司id
func (c *Ctx) GetCompanyIdByAppKey(ctx context.Context, appKey string) (companyId int64, err error) {

	companyId = int64(0)

	cacheKey := fmt.Sprintf(CacheKeyCompanyIdByAppKey, appKey)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {

		// 通过唯一标识 获取到对应的应用信息
		urlInfo, err := c.GetUrlInfoByKey(ctx, appKey)
		if err != nil && errors.Is(err, model.ErrNotFound) {
		} else if err != nil {
			return companyId, err
		}
		if urlInfo != nil {
			return urlInfo.CompanyId, nil
		}

		// 通过唯一标识 获取到对应的应用信息
		fileInfo, err := c.GetFileInfoByKey(ctx, appKey)
		if err != nil && errors.Is(err, model.ErrNotFound) {
		} else if err != nil {
			return companyId, err
		}
		if fileInfo != nil {
			return fileInfo.CompanyId, nil
		}

		// 通过唯一标识 获取到对应的应用信息
		tauriInfo, err := c.GetTauriInfoByKey(ctx, appKey)
		if err != nil && errors.Is(err, model.ErrNotFound) {
		} else if err != nil {
			return companyId, err
		}
		if tauriInfo != nil {
			return tauriInfo.CompanyId, nil
		}

		// 通过唯一标识 获取到对应的应用信息
		apkInfo, err := c.GetApkInfoByKey(ctx, appKey)
		if err != nil && errors.Is(err, model.ErrNotFound) {
		} else if err != nil {
			return companyId, err
		}
		if apkInfo != nil {
			return apkInfo.CompanyId, nil
		}

		// 通过唯一标识 获取到对应的应用信息
		configurationInfo, err := c.GetConfigurationInfoByKey(ctx, appKey)
		if err != nil && errors.Is(err, model.ErrNotFound) {
		} else if err != nil {
			return companyId, err
		}
		if configurationInfo != nil {
			return configurationInfo.CompanyId, nil
		}

		return companyId, nil
	})

	if v != nil {
		companyId = v.(int64)
		return companyId, err
	}

	return companyId, nil
}
