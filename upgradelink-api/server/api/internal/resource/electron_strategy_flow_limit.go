package resource

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"upgradelink-api/server/api/internal/resource/model"
)

const (
	CacheKeyGetElectronFlowLimitStrategyListByIds              = PREFIX + "ELECTRON_FLOW_LIMIT_LIST:IDS:%v"
	CacheKeyGetElectronFlowLimitStrategyIncrementByIdAndPeriod = PREFIX + "ELECTRON_FLOW_LIMIT_STRATEGY_INCREMENT:ID:%v:PERIOD:%v"
)

// GetElectronFlowLimitStrategyListByIds
// 通过 ids 获取到流控策略list
func (c *Ctx) GetElectronFlowLimitStrategyListByIds(ctx context.Context, ids string) ([]*model.UpgradeElectronUpgradeStrategyFlowLimitStrategy, error) {

	cacheKey := fmt.Sprintf(CacheKeyGetElectronFlowLimitStrategyListByIds, ids)
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
		query := fmt.Sprintf("select * from upgrade_electron_upgrade_strategy_flow_limit_strategy "+
			"where id in ( %s ) "+
			"AND enable = 1 "+
			"AND is_del = 0 ", strings.Join(placeholders, ", "))

		// 构建参数列表
		args := make([]interface{}, len(idList))
		for i, id := range idList {
			args[i] = id
		}

		// sql 缓存查询
		var list []*model.UpgradeElectronUpgradeStrategyFlowLimitStrategy
		err := c.mysqlConn.QueryRowsCtx(ctx, &list, query, args...)
		if err != nil {
			return nil, err
		}

		return list, nil
	})

	if v != nil {
		list := v.([]*model.UpgradeElectronUpgradeStrategyFlowLimitStrategy)
		return list, err
	}

	return nil, err
}

// ElectronFlowLimitStrategyIncrement 判断当前请求在流控策略要求的数量范围内  使用 redis 计数器方式实现
func (c *Ctx) ElectronFlowLimitStrategyIncrement(ctx context.Context, flowLimitStrategy *model.UpgradeElectronUpgradeStrategyFlowLimitStrategy) (bool, error) {

	now := time.Now()
	startTime, err := time.Parse("15:04:05", flowLimitStrategy.BeginTime)
	if err != nil {
		return false, fmt.Errorf("解析开始时间失败: %w", err)
	}
	endTime, err := time.Parse("15:04:05", flowLimitStrategy.EndTime)
	if err != nil {
		return false, fmt.Errorf("解析结束时间失败: %w", err)
	}

	start := time.Date(now.Year(), now.Month(), now.Day(), startTime.Hour(), startTime.Minute(), startTime.Second(), 0, now.Location())
	end := time.Date(now.Year(), now.Month(), now.Day(), endTime.Hour(), endTime.Minute(), endTime.Second(), 0, now.Location())

	if now.Before(start) || now.After(end) {
		return true, nil
	}

	// 判断频控时间维度
	var period int64
	switch flowLimitStrategy.Dimension {
	case 1:
		period = now.Unix()
	case 2:
		period = now.Unix() / 60
	case 3:
		period = now.Unix() / 3600
	case 4:
		period = now.Unix() / 86400
	default:
		return false, fmt.Errorf("无效的频控维度: %d", flowLimitStrategy.Dimension)
	}

	// 构建 Redis 键名
	cacheKey := fmt.Sprintf(CacheKeyGetElectronFlowLimitStrategyIncrementByIdAndPeriod, flowLimitStrategy.Id, period)

	// 获取过期时间
	expiration := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location()).Sub(now)

	// 获取当前计数器的值
	currentCount, err := c.redisConn.Get(cacheKey)
	if err != nil {
		return false, err
	}

	// 如果 key 不存在，则设置初始值为 0
	if currentCount == "" {
		err = c.redisConn.SetexCtx(ctx, cacheKey, "1", int(expiration))
		if err != nil {
			return false, err
		}
		return true, nil
	}

	currentCountInt, err := strconv.ParseInt(currentCount, 10, 64)
	if err != nil {
		return false, err
	}

	// 判断是否超过限制
	if currentCountInt < flowLimitStrategy.Limit {
		_, err := c.redisConn.IncrCtx(ctx, cacheKey)
		if err != nil {
			return false, err
		}
		return true, nil
	}

	return false, nil
}
