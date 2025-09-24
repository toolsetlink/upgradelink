package resource

import (
	"context"
	"fmt"

	"upgradelink-api/server/api/internal/resource/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	CacheKeyUrlVersionListByUrlIdAndVersionCode = PREFIX + "URL_VERSION_LIST:URL_ID:%v:VERSION_CODE:%v"
	CacheKeyUrlVersionInfoById                  = PREFIX + "URL_VERSION_INFO:ID:%v"
	CacheKeyUrlVersionLastInfoByUrlId           = PREFIX + "URL_VERSION_LAST_INFO:URL_ID:%v"
	CacheKeyUrlVersionInfoByUrlIdAndVersionCode = PREFIX + "URL_VERSION_INFO:URL_ID:%v:VERSION_CODE:%v"
)

// GetUrlVersionListByUrlIdAndVersionCode
// 获取大于 versionCode 的版本列表
func (c *Ctx) GetUrlVersionListByUrlIdAndVersionCode(ctx context.Context, urlId int64, versionCode int64) ([]*model.UpgradeUrlVersion, error) {

	cacheKey := fmt.Sprintf(CacheKeyUrlVersionListByUrlIdAndVersionCode, urlId, versionCode)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var UrlVersionList []*model.UpgradeUrlVersion
		query := fmt.Sprintf("select * from upgrade_url_version where `url_id` = ? and is_del = 0 and `version_code` > ? order by version_code desc ")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &UrlVersionList, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowsCtx(ctx, &UrlVersionList, query, urlId, versionCode)
		})

		if err != nil {
			return nil, err
		}
		return UrlVersionList, nil
	})

	if v != nil {
		list := v.([]*model.UpgradeUrlVersion)
		return list, err
	}

	return nil, err
}

// GetUrlVersionLastInfoByUrlId
// 获取最新版本的信息
func (c *Ctx) GetUrlVersionLastInfoByUrlId(ctx context.Context, urlId int64) (*model.UpgradeUrlVersion, error) {

	cacheKey := fmt.Sprintf(CacheKeyUrlVersionLastInfoByUrlId, urlId)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info model.UpgradeUrlVersion
		query := fmt.Sprintf("select * from upgrade_url_version where `url_id` = ? and is_del = 0 order by version_code desc limit 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &info, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, &info, query, urlId)
		})

		if err != nil {
			return nil, err
		}
		return info, nil
	})

	if v != nil {
		info := v.(model.UpgradeUrlVersion)
		return &info, err
	}

	return nil, err
}

// 获取固定版本的信息
func (c *Ctx) GetUrlVersionInfoByUrlIdAndVersionCode(ctx context.Context, urlId int64, versionCode int64) (*model.UpgradeUrlVersion, error) {

	cacheKey := fmt.Sprintf(CacheKeyUrlVersionInfoByUrlIdAndVersionCode, urlId, versionCode)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info model.UpgradeUrlVersion
		query := fmt.Sprintf("select * from upgrade_url_version where `url_id` = ? and is_del = 0 and `version_code` = ? LIMIT 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &info, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, &info, query, urlId, versionCode)
		})

		if err != nil {
			return nil, err
		}
		return info, nil
	})

	if v != nil {
		info := v.(model.UpgradeUrlVersion)
		return &info, err
	}

	return nil, err
}

// GetUrlVersionInfoById
// 通过 url version id 获取信息
func (c *Ctx) GetUrlVersionInfoById(ctx context.Context,
	id int64) (*model.UpgradeUrlVersion, error) {

	cacheKey := fmt.Sprintf(CacheKeyUrlVersionInfoById, id)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info model.UpgradeUrlVersion
		query := fmt.Sprintf("select * from upgrade_url_version where `id` = ?  limit 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &info, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, v, query, id)
		})
		if err != nil {
			return nil, err
		}

		return info, nil
	})

	if v != nil {
		info := v.(model.UpgradeUrlVersion)
		return &info, err
	}

	return nil, err
}
