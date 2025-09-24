package resource

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"upgradelink-api/server/api/internal/common"
	"upgradelink-api/server/api/internal/resource/model"
)

const (
	CacheKeyGetTauriGrayStrategyListByIds     = PREFIX + "TAURI_GRAY_LIST:IDS:%v"
	CacheKeyGetTauriGrayStrategyIncrementById = PREFIX + "TAURI_GRAY_STRATEGY_INCREMENT:ID:%v"
)

// GetTauriGrayStrategyListByIds
// 通过 ids 获取到灰度策略list
func (c *Ctx) GetTauriGrayStrategyListByIds(ctx context.Context, ids string) ([]*model.UpgradeTauriUpgradeStrategyGrayStrategy, error) {

	cacheKey := fmt.Sprintf(CacheKeyGetTauriGrayStrategyListByIds, ids)
	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// 拆分 ids 字符串
		idList := strings.Split(ids, ",")
		// 构建占位符
		placeholders := make([]string, len(idList))
		for i := range placeholders {
			placeholders[i] = "?"
		}
		// 构建 SQL 查询语句
		query := fmt.Sprintf("select * from upgrade_tauri_upgrade_strategy_gray_strategy "+
			"where id in ( %s ) "+
			"AND enable = 1 "+
			"AND is_del = 0 ", strings.Join(placeholders, ", "))

		// 构建参数列表
		args := make([]interface{}, len(idList))
		for i, id := range idList {
			args[i] = id
		}

		// sql 缓存查询
		var grayStrategyList []*model.UpgradeTauriUpgradeStrategyGrayStrategy
		err := c.mysqlConn.QueryRowsCtx(ctx, &grayStrategyList, query, args...)
		if err != nil {
			return nil, err
		}

		return grayStrategyList, nil
	})

	if v != nil {
		strategyList := v.([]*model.UpgradeTauriUpgradeStrategyGrayStrategy)
		return strategyList, err
	}

	return nil, err
}

// TauriGrayStrategyIncrement 判断当前请求在灰度策略要求的数量范围内  使用 redis 计数器方式实现
func (c *Ctx) TauriGrayStrategyIncrement(ctx context.Context, grayStrategy *model.UpgradeTauriUpgradeStrategyGrayStrategy) (bool, error) {

	// 构建 Redis 键名
	cacheKey := fmt.Sprintf(CacheKeyGetTauriGrayStrategyIncrementById, grayStrategy.Id)

	// 获取过期时间
	expiration := common.CalculateTimeDifferenceSeconds(common.GetCurrentTime(), grayStrategy.EndDatetime)

	// 获取当前计数器的值
	currentCount, err := c.redisConn.Get(cacheKey)
	if err != nil {
		return false, err
	}

	if currentCount == "" {
		// 如果 key 不存在，则设置初始值为 0
		err = c.redisConn.SetexCtx(ctx, cacheKey, "0", expiration)
		if err != nil {
			return false, err
		}
		return true, nil
	}

	currentCountInt, err := strconv.ParseInt(currentCount, 10, 64)
	if err != nil {
		return false, err
	}
	if currentCountInt < grayStrategy.Limit {
		_, err := c.redisConn.IncrCtx(ctx, cacheKey)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}
