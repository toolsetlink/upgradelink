package resource

import (
	"context"
	"fmt"

	"upgradelink-api/server/api/internal/resource/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	CacheKeyMacInfoByKey             = PREFIX + "MAC_INFO:KEY:%v"
	CacheKeyMacInfoByCompanyIdAndKey = PREFIX + "MAC_INFO:COMPANY_ID:%v:KEY:%v"
)

func (c *Ctx) GetMacInfoByKey(ctx context.Context,
	key string) (*model.UpgradeMac, error) {

	cacheKey := fmt.Sprintf(CacheKeyMacInfoByKey, key)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info model.UpgradeMac
		query := fmt.Sprintf("select * from upgrade_mac where `key` = ? and is_del = 0 limit 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &info, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, v, query, key)
		})

		if err != nil {
			return nil, err
		}

		return info, nil
	})

	if v != nil {
		info := v.(model.UpgradeMac)
		return &info, err
	}

	return nil, err
}

func (c *Ctx) GetMacInfoByCompanyIdAndKey(ctx context.Context,
	companyId int64, key string) (*model.UpgradeMac, error) {

	cacheKey := fmt.Sprintf(CacheKeyMacInfoByCompanyIdAndKey, companyId, key)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info model.UpgradeMac
		query := fmt.Sprintf("select * from upgrade_mac where `company_id` = ? and `key` = ? and is_del = 0 limit 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &info, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, v, query, companyId, key)
		})

		if err != nil {
			return nil, err
		}

		return info, nil
	})

	if v != nil {
		info := v.(model.UpgradeMac)
		return &info, err
	}

	return nil, err
}
