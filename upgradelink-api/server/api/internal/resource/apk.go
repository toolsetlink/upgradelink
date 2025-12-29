package resource

import (
	"context"
	"fmt"

	"upgradelink-api/server/api/internal/resource/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	CacheKeyApkInfoByKey             = PREFIX + "APK_INFO:KEY:%v"
	CacheKeyApkInfoByCompanyIdAndKey = PREFIX + "APK_INFO:COMPANY_ID:%v:KEY:%v"
)

// GetApkInfoByKey
// 获取应用信息
func (c *Ctx) GetApkInfoByKey(ctx context.Context,
	key string) (*model.UpgradeApk, error) {

	cacheKey := fmt.Sprintf(CacheKeyApkInfoByKey, key)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info model.UpgradeApk
		query := fmt.Sprintf("select * from upgrade_apk where `key` = ? and is_del = 0 limit 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &info, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, v, query, key)
		})

		if err != nil {
			return nil, err
		}

		return info, nil
	})

	if v != nil {
		info := v.(model.UpgradeApk)
		return &info, err
	}

	return nil, err
}

// GetApkInfoByCompanyIdAndKey
// 获取应用信息
func (c *Ctx) GetApkInfoByCompanyIdAndKey(ctx context.Context,
	companyId int64, key string) (*model.UpgradeApk, error) {

	cacheKey := fmt.Sprintf(CacheKeyApkInfoByCompanyIdAndKey, companyId, key)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info model.UpgradeApk
		query := fmt.Sprintf("select * from upgrade_apk where `company_id` = ? and `key` = ? and is_del = 0 limit 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &info, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, v, query, companyId, key)
		})

		if err != nil {
			return nil, err
		}

		return info, nil
	})

	if v != nil {
		info := v.(model.UpgradeApk)
		return &info, err
	}

	return nil, err
}
