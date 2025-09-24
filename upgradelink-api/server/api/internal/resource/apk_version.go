package resource

import (
	"context"
	"fmt"

	"upgradelink-api/server/api/internal/resource/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	CacheKeyApkVersionListByApkIdAndVersionCode = PREFIX + "APK_VERSION_LIST:APK_ID:%v:VERSION_CODE:%v"
	CacheKeyApkVersionInfoById                  = PREFIX + "APK_VERSION_INFO:ID:%v"
	CacheKeyApkVersionLastInfoByApkId           = PREFIX + "APK_VERSION_LAST_INFO:APK_ID:%v"
	CacheKeyApkVersionInfoByApkIdAndVersionCode = PREFIX + "APK_VERSION_INFO:APK_ID:%v:VERSION_CODE:%v"
)

// GetApkVersionListByApkIdAndVersionCode
// 获取大于 versionCode 的版本列表
func (c *Ctx) GetApkVersionListByApkIdAndVersionCode(ctx context.Context, apkId int64, versionCode int64) ([]*model.UpgradeApkVersion, error) {

	cacheKey := fmt.Sprintf(CacheKeyApkVersionListByApkIdAndVersionCode, apkId, versionCode)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var ApkVersionList []*model.UpgradeApkVersion
		query := fmt.Sprintf("select * from upgrade_apk_version where `apk_id` = ? and is_del = 0 and `version_code` > ? order by version_code desc ")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &ApkVersionList, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowsCtx(ctx, &ApkVersionList, query, apkId, versionCode)
		})

		if err != nil {
			return nil, err
		}
		return ApkVersionList, nil
	})

	if v != nil {
		list := v.([]*model.UpgradeApkVersion)
		return list, err
	}

	return nil, err
}

// GetApkVersionLastInfoByApkId
// 获取最新版本的信息
func (c *Ctx) GetApkVersionLastInfoByApkId(ctx context.Context, apkId int64) (*model.UpgradeApkVersion, error) {

	cacheKey := fmt.Sprintf(CacheKeyApkVersionLastInfoByApkId, apkId)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info model.UpgradeApkVersion
		query := fmt.Sprintf("select * from upgrade_apk_version where `apk_id` = ? and is_del = 0 order by version_code desc limit 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &info, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, &info, query, apkId)
		})

		if err != nil {
			return nil, err
		}
		return info, nil
	})

	if v != nil {
		info := v.(model.UpgradeApkVersion)
		return &info, err
	}

	return nil, err
}

// 获取固定版本的信息
func (c *Ctx) GetApkVersionInfoByApkIdAndVersionCode(ctx context.Context, apkId int64, versionCode int64) (*model.UpgradeApkVersion, error) {

	cacheKey := fmt.Sprintf(CacheKeyApkVersionInfoByApkIdAndVersionCode, apkId, versionCode)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info model.UpgradeApkVersion
		query := fmt.Sprintf("select * from upgrade_apk_version where `apk_id` = ? and is_del = 0 and `version_code` = ? LIMIT 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &info, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, &info, query, apkId, versionCode)
		})

		if err != nil {
			return nil, err
		}
		return info, nil
	})

	if v != nil {
		info := v.(model.UpgradeApkVersion)
		return &info, err
	}

	return nil, err
}

// GetApkVersionInfoById
// 通过 apk version id 获取信息
func (c *Ctx) GetApkVersionInfoById(ctx context.Context,
	id int64) (*model.UpgradeApkVersion, error) {

	cacheKey := fmt.Sprintf(CacheKeyApkVersionInfoById, id)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info model.UpgradeApkVersion
		query := fmt.Sprintf("select * from upgrade_apk_version where `id` = ?  limit 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &info, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, v, query, id)
		})
		if err != nil {
			return nil, err
		}

		return info, nil
	})

	if v != nil {
		info := v.(model.UpgradeApkVersion)
		return &info, err
	}

	return nil, err
}
