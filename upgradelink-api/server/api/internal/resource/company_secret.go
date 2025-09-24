package resource

import (
	"context"
	"fmt"

	"upgradelink-api/server/api/internal/resource/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	CacheKeyCompanySecretByAccessKey = PREFIX + "COMPANY_SECRET_INFO:ACCESS_KEY:%v"
)

func (c *Ctx) GetCompanySecretByAccessKey(ctx context.Context,
	key string) (*model.SysCompanySecret, error) {

	cacheKey := fmt.Sprintf(CacheKeyCompanySecretByAccessKey, key)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info model.SysCompanySecret
		query := fmt.Sprintf("select * from sys_company_secret where `access_key` = ? limit 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &info, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, v, query, key)
		})

		if err != nil {
			return nil, err
		}

		return info, nil
	})

	if v != nil {
		info := v.(model.SysCompanySecret)
		return &info, err
	}

	return nil, err
}
