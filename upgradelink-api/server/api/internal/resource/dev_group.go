package resource

import (
	"context"
	"fmt"

	"upgradelink-api/server/api/internal/resource/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	CacheKeyDevGroupInfoById     = PREFIX + "DEV_GROUP_INFO:ID:%v"
	CacheKeyDevGroupListByDevKey = PREFIX + "DEV_GROUP_LIST:DEV_KEY:%v"
)

// GetDevGroupInfoById
// 通过获取详细信息
func (c *Ctx) GetDevGroupInfoById(ctx context.Context, id int64) (*model.UpgradeDevGroup, error) {

	cacheKey := fmt.Sprintf(CacheKeyDevGroupInfoById, id)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info model.UpgradeDevGroup
		query := fmt.Sprintf("select * from upgrade_dev_group where `id` = ?  limit 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &info, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, v, query, id)
		})

		if err != nil {
			return nil, err
		}

		return info, nil
	})

	if v != nil {
		info := v.(model.UpgradeDevGroup)
		return &info, err
	}

	return nil, err
}

// GetDevGroupListByDevKey
// 通过 设备唯一标识获取 设备的分组列表
func (c *Ctx) GetDevGroupListByDevKey(ctx context.Context, devKey string) ([]*model.UpgradeDevGroup, error) {

	// 查询设备分组列表数据
	cacheKey := fmt.Sprintf(CacheKeyDevGroupListByDevKey, devKey)
	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var list []*model.UpgradeDevGroup
		//query := fmt.Sprintf("select * from upgrade_dev_group_relation where `dev_id` = ? and is_del = 0")
		query := fmt.Sprintf("select upgrade_dev_group.* from upgrade_devs " +
			" left join upgrade_dev_group_relation on upgrade_devs.id = upgrade_dev_group_relation.dev_id " +
			" left join upgrade_dev_group on upgrade_dev_group_relation.dev_group_id = upgrade_dev_group.id " +
			" where upgrade_devs.is_del = 0 " +
			" and upgrade_dev_group_relation.is_del = 0" +
			" and upgrade_dev_group.is_del = 0 " +
			" and upgrade_devs.key = ? ")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &list, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowsCtx(ctx, &list, query, devKey)
		})

		if err != nil {
			return nil, err
		}
		return list, nil
	})

	if v != nil {
		list := v.([]*model.UpgradeDevGroup)
		return list, err
	}

	return nil, nil
}
