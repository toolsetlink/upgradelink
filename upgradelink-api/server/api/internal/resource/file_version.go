package resource

import (
	"context"
	"fmt"

	"upgradelink-api/server/api/internal/resource/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	CacheKeyFileVersionListByFileIdAndVersionCode = PREFIX + "FILE_VERSION_LIST:FILE_ID:%v:VERSION_CODE:%v"
	CacheKeyFileVersionInfoById                   = PREFIX + "FILE_VERSION_INFO:ID:%v"
	CacheKeyFileVersionLastInfoByFileId           = PREFIX + "FILE_VERSION_LAST_INFO:FILE_ID:%v"
	CacheKeyFileVersionInfoByFileIdAndVersionCode = PREFIX + "FILE_VERSION_INFO:FILE_ID:%v:VERSION_CODE:%v"
)

// GetFileVersionListByFileIdAndVersionCode
// 获取大于 versionCode 的版本列表
func (c *Ctx) GetFileVersionListByFileIdAndVersionCode(ctx context.Context, fileId int64, versionCode int64) ([]*model.UpgradeFileVersion, error) {

	cacheKey := fmt.Sprintf(CacheKeyFileVersionListByFileIdAndVersionCode, fileId, versionCode)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var FileVersionList []*model.UpgradeFileVersion
		query := fmt.Sprintf("select * from upgrade_file_version where `file_id` = ? and is_del = 0 and `version_code` > ? order by version_code desc ")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &FileVersionList, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowsCtx(ctx, &FileVersionList, query, fileId, versionCode)
		})

		if err != nil {
			return nil, err
		}
		return FileVersionList, nil
	})

	if v != nil {
		list := v.([]*model.UpgradeFileVersion)
		return list, err
	}

	return nil, err
}

// GetFileVersionLastInfoByFileId
// 获取最新版本的信息
func (c *Ctx) GetFileVersionLastInfoByFileId(ctx context.Context, fileId int64) (*model.UpgradeFileVersion, error) {

	cacheKey := fmt.Sprintf(CacheKeyFileVersionLastInfoByFileId, fileId)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info model.UpgradeFileVersion
		query := fmt.Sprintf("select * from upgrade_file_version where `file_id` = ? and is_del = 0 order by version_code desc limit 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &info, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, &info, query, fileId)
		})

		if err != nil {
			return nil, err
		}
		return info, nil
	})

	if v != nil {
		info := v.(model.UpgradeFileVersion)
		return &info, err
	}

	return nil, err
}

// 获取固定版本的信息
func (c *Ctx) GetFileVersionInfoByFileIdAndVersionCode(ctx context.Context, fileId int64, versionCode int64) (*model.UpgradeFileVersion, error) {

	cacheKey := fmt.Sprintf(CacheKeyFileVersionInfoByFileIdAndVersionCode, fileId, versionCode)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info model.UpgradeFileVersion
		query := fmt.Sprintf("select * from upgrade_file_version where `file_id` = ? and is_del = 0 and `version_code` = ? LIMIT 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &info, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, &info, query, fileId, versionCode)
		})

		if err != nil {
			return nil, err
		}
		return info, nil
	})

	if v != nil {
		info := v.(model.UpgradeFileVersion)
		return &info, err
	}

	return nil, err
}

// GetFileVersionInfoById
// 通过 file version id 获取信息
func (c *Ctx) GetFileVersionInfoById(ctx context.Context,
	id int64) (*model.UpgradeFileVersion, error) {

	cacheKey := fmt.Sprintf(CacheKeyFileVersionInfoById, id)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info model.UpgradeFileVersion
		query := fmt.Sprintf("select * from upgrade_file_version where `id` = ?  limit 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &info, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, v, query, id)
		})
		if err != nil {
			return nil, err
		}

		return info, nil
	})

	if v != nil {
		info := v.(model.UpgradeFileVersion)
		return &info, err
	}

	return nil, err
}
