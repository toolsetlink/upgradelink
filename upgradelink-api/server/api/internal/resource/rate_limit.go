package resource

import (
	"context"
	"errors"
	"fmt"
	"time"
)

const (
	CacheKeyRateLimitSecondByIp = PREFIX + "RATE_LIMIT_SECOND:IP:%v:SECOND:%v"
	CacheKeyRateLimitMinuteByIp = PREFIX + "RATE_LIMIT_MINUTE:IP:%v:MINUTE:%v"
	CacheKeyRateLimitHourByIp   = PREFIX + "RATE_LIMIT_HOUR:IP:%v:HOUR:%v"
)

// GetRateLimitSecondByIp 通过 ip 获取到流量信息
// 检查秒级限制 (1秒最多2次)
func (c *Ctx) GetRateLimitSecondByIp(ctx context.Context, ip string) (err error) {

	secondKey := fmt.Sprintf(CacheKeyRateLimitSecondByIp, ip, time.Now().Unix())
	secondCount, err := c.redisConn.Incr(secondKey)
	if err != nil {
		return nil
	}

	if secondCount == 1 {
		err = c.redisConn.Expire(secondKey, 1)
		if err != nil {
			return nil
		}
	}

	if secondCount > 2 {
		return errors.New("too many requests")
	}

	return nil
}

// GetRateLimitMinuteByIp 通过 ip 获取到流量信息
// 检查分钟级限制 (1分钟最多20次)
func (c *Ctx) GetRateLimitMinuteByIp(ctx context.Context, ip string) (err error) {

	minuteKey := fmt.Sprintf(CacheKeyRateLimitMinuteByIp, ip, time.Now().Unix()/60)
	minuteCount, err := c.redisConn.Incr(minuteKey)
	if err != nil {
		return nil
	}
	if minuteCount == 1 {
		err = c.redisConn.Expire(minuteKey, 60)
		if err != nil {
			return nil
		}
	}

	if minuteCount > 10 {
		return errors.New("too many requests")
	}

	return nil
}

// GetRateLimitHourByIp 通过 ip 获取到流量信息
// 检查小时级限制 (1小时最多20次)
func (c *Ctx) GetRateLimitHourByIp(ctx context.Context, ip string) (err error) {

	hourKey := fmt.Sprintf(CacheKeyRateLimitHourByIp, ip, time.Now().Unix()/3600)
	hourCount, err := c.redisConn.Incr(hourKey)
	if err != nil {
		return nil
	}

	// 设置过期时间
	if hourCount == 1 {
		err = c.redisConn.Expire(hourKey, 3600)
		if err != nil {
			return nil
		}
	}

	if hourCount > 20 {
		return errors.New("too many requests")
	}

	return nil
}
