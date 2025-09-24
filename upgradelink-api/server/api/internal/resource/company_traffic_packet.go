package resource

import (
	"context"
	"errors"
	"fmt"
	"time"
	"upgradelink-api/server/api/internal/resource/model"
)

const (
	CacheKeyCompanyTrafficPacketValidListByCompanyId = PREFIX + "CompanyTrafficPacketValidList:companyId:%v"
)

//// GetCompanyTrafficPacketValidListByCompanyId 查询当前公司符合条件的可用的流量包
//func (c *Ctx) GetCompanyTrafficPacketValidListByCompanyId(ctx context.Context, companyId int64) ([]*model.UpgradeCompanyTrafficPacket, error) {
//
//	cacheKey := fmt.Sprintf(CacheKeyCompanyTrafficPacketValidListByCompanyId, companyId)
//
//	now := time.Now()
//
//	// 内存缓存
//	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
//
//		// sql 缓存查询
//		var list []*model.UpgradeCompanyTrafficPacket
//		query := fmt.Sprintf("select * from %s "+
//			"where company_id = ? "+
//			"AND ? > start_time "+
//			"AND ? < end_time "+
//			"AND status = 1 "+
//			"AND remaining_size > 0 "+
//			"order by end_time desc", "upgrade_company_traffic_packet")
//
//		err := c.mysqlConn.QueryRowsCtx(context.Background(), &list, query, companyId, now, now)
//		if err != nil {
//			return nil, err
//		}
//
//		return list, nil
//	})
//
//	if v != nil {
//		list := v.([]*model.UpgradeCompanyTrafficPacket)
//		return list, err
//	}
//	return nil, err
//}

// GetCompanyTrafficPacketValidListByCompanyId 查询当前公司符合条件的可用的流量包  无缓存版本
func (c *Ctx) GetCompanyTrafficPacketValidListByCompanyId(ctx context.Context, companyId int64) ([]*model.UpgradeCompanyTrafficPacket, error) {

	now := time.Now()

	// sql 缓存查询
	var list []*model.UpgradeCompanyTrafficPacket
	query := fmt.Sprintf("select * from %s "+
		"where company_id = ? "+
		"AND ? > start_time "+
		"AND ? < end_time "+
		"AND status = 1 "+
		"AND remaining_size > 0 "+
		"order by end_time desc", "upgrade_company_traffic_packet")

	err := c.mysqlConn.QueryRowsCtx(context.Background(), &list, query, companyId, now, now)
	if err != nil {
		return nil, err
	}

	return list, nil
}

// ReduceCompanyTrafficPacketRemainingSize 减少流量包剩余流量
// ReduceCompanyTrafficPacketRemainingSize
func (c *Ctx) ReduceCompanyTrafficPacketRemainingSize(ctx context.Context, packetId int64, size int64) error {

	// 查询流量包
	var packet model.UpgradeCompanyTrafficPacket
	query := fmt.Sprintf("select * from %s where id = ?", "upgrade_company_traffic_packet")
	err := c.mysqlConn.QueryRowCtx(ctx, &packet, query, packetId)
	if err != nil {
		return err
	}
	fmt.Println("packet: ", packet)

	// 检查流量包是否存在
	if packet.Id == 0 {
		return errors.New("流量包不存在")
	}

	// 计算出剩余流量
	remainingSize := 0

	// 判断是否足够
	if packet.RemainingSize < size {
		remainingSize = 0
	} else {
		remainingSize = int(packet.RemainingSize - size)
	}

	// 减少流量包剩余流量
	query = fmt.Sprintf("update %s set remaining_size =  ? where id = ?", "upgrade_company_traffic_packet")
	_, err = c.mysqlConn.ExecCtx(ctx, query, remainingSize, packetId)
	if err != nil {
		return err
	}

	return nil
}
