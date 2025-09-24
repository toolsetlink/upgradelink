package resource

import (
	"context"
	"errors"
	"fmt"

	"upgradelink-api/server/api/internal/common"
	"upgradelink-api/server/api/internal/resource/model"
)

const (
	CacheKeyAppVersionIdByReport = PREFIX + "APP_VERSION_ID:REPORT:%v"
)

type GetAppVersionIdByReportReq struct {
	AppKey             string `db:"app_key"`              // 应用Key
	AppVersionCode     int64  `db:"app_version_code"`     // 应用版本号
	DevModelKey        string `db:"dev_model_key"`        // 设备机型唯一标识
	DevKey             string `db:"dev_key"`              // 设备唯一标识
	AppVersionPlatform string `db:"app_version_platform"` // 操作系统
	AppVersionTarget   string `db:"app_version_target"`   // 操作系统
	AppVersionArch     string `db:"app_version_arch"`     // 机器架构
}

// GetAppVersionIdByReport
// 通过 上报接口提供的信息，获取出来具体的应用版本号
func (c *Ctx) GetAppVersionIdByReport(ctx context.Context, req GetAppVersionIdByReportReq) (versionId int64, err error) {

	versionId = int64(0)

	// 计算 MD5 值
	md5Hash, err := common.CalculateMD5(req)
	if err != nil {
		return 0, err
	}

	cacheKey := fmt.Sprintf(CacheKeyAppVersionIdByReport, md5Hash)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {

		// 通过唯一标识 获取到对应的应用信息
		urlInfo, err := c.GetUrlInfoByKey(ctx, req.AppKey)
		if err != nil && errors.Is(err, model.ErrNotFound) {
		} else if err != nil {
			return versionId, err
		}
		if urlInfo != nil {
			// 查询出版本号 id
			versionInfo, err := c.GetUrlVersionInfoByUrlIdAndVersionCode(ctx, urlInfo.Id, req.AppVersionCode)
			if err != nil && errors.Is(err, model.ErrNotFound) {
			} else if err != nil {
				return versionId, err
			}
			if versionInfo == nil {
				return versionId, nil
			}
			return versionInfo.Id, nil
		}

		// 通过唯一标识 获取到对应的应用信息
		fileInfo, err := c.GetFileInfoByKey(ctx, req.AppKey)
		if err != nil && errors.Is(err, model.ErrNotFound) {
		} else if err != nil {
			return versionId, err
		}
		if fileInfo != nil {
			// 查询出版本号 id
			versionInfo, err := c.GetFileVersionInfoByFileIdAndVersionCode(ctx, fileInfo.Id, req.AppVersionCode)
			if err != nil && errors.Is(err, model.ErrNotFound) {
			} else if err != nil {
				return versionId, err
			}
			if versionInfo == nil {
				return versionId, nil
			}
			return versionInfo.Id, nil
		}

		// 通过唯一标识 获取到对应的应用信息
		tauriInfo, err := c.GetTauriInfoByKey(ctx, req.AppKey)
		if err != nil && errors.Is(err, model.ErrNotFound) {
		} else if err != nil {
			return versionId, err
		}
		if tauriInfo != nil {
			// 查询出版本号 id
			versionInfo, err := c.GetTauriVersionInfoByTauriIdAndVersionCodeAndTargetAndArch(ctx, tauriInfo.Id, req.AppVersionCode, req.AppVersionTarget, req.AppVersionArch)
			if err != nil && errors.Is(err, model.ErrNotFound) {
			} else if err != nil {
				return versionId, err
			}
			if versionInfo == nil {
				return versionId, nil
			}
			return versionInfo.Id, nil
		}

		// 通过唯一标识 获取到对应的应用信息
		electronInfo, err := c.GetElectronInfoByKey(ctx, req.AppKey)
		if err != nil && errors.Is(err, model.ErrNotFound) {
		} else if err != nil {
			return versionId, err
		}
		if electronInfo != nil {
			// 查询出版本号 id
			versionInfo, err := c.GetElectronVersionInfoByElectronIdAndVersionCodeAndPlatformAndArch(ctx, electronInfo.Id, req.AppVersionCode, req.AppVersionPlatform, req.AppVersionArch)
			if err != nil && errors.Is(err, model.ErrNotFound) {
			} else if err != nil {
				return versionId, err
			}
			if versionInfo == nil {
				return versionId, nil
			}
			return versionInfo.Id, nil
		}

		// 通过唯一标识 获取到对应的应用信息
		apkInfo, err := c.GetApkInfoByKey(ctx, req.AppKey)
		if err != nil && errors.Is(err, model.ErrNotFound) {
		} else if err != nil {
			return versionId, err
		}
		if apkInfo != nil {
			// 查询出版本号 id
			versionInfo, err := c.GetApkVersionInfoByApkIdAndVersionCode(ctx, apkInfo.Id, req.AppVersionCode)
			if err != nil && errors.Is(err, model.ErrNotFound) {
			} else if err != nil {
				return versionId, err
			}
			if versionInfo == nil {
				return versionId, nil
			}
			return versionInfo.Id, nil
		}

		// 通过唯一标识 获取到对应的应用信息
		configurationInfo, err := c.GetConfigurationInfoByKey(ctx, req.AppKey)
		if err != nil && errors.Is(err, model.ErrNotFound) {
		} else if err != nil {
			return versionId, err
		}
		if configurationInfo != nil {
			// 查询出版本号 id
			versionInfo, err := c.GetConfigurationVersionInfoByConfigurationIdAndVersionCode(ctx, configurationInfo.Id, req.AppVersionCode)
			if err != nil && errors.Is(err, model.ErrNotFound) {
			} else if err != nil {
				return versionId, err
			}
			if versionInfo == nil {
				return versionId, nil
			}
			return versionInfo.Id, nil
		}

		return versionId, nil
	})

	if v != nil {
		versionId = v.(int64)
		return versionId, err
	}

	return versionId, nil
}
