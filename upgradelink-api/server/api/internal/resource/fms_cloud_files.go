package resource

import (
	"context"
	"fmt"
	"time"

	"upgradelink-api/server/api/internal/resource/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	CacheKeyCloudFileInfoById = PREFIX + "CLOUD_FILE_INFO:ID:%v"
)

type AddFmsCloudFilesReq struct {
	Id                        string    `db:"id"`         // UUID
	CreatedAt                 time.Time `db:"created_at"` // Create Time | 创建日期
	UpdatedAt                 time.Time `db:"updated_at"` // Update Time | 修改日期
	State                     int64     `db:"state"`      // State true: normal false: ban | 状态 true 正常 false 禁用
	Name                      string    `db:"name"`       // The file''s name | 文件名
	Url                       string    `db:"url"`        // The file''s url | 文件地址
	Size                      uint64    `db:"size"`       // The file''s size | 文件大小
	FileType                  uint64    `db:"file_type"`  // The file''s type | 文件类型
	UserId                    string    `db:"user_id"`    // The user who upload the file | 上传用户的 ID
	CloudFileStorageProviders int64     `db:"cloud_file_storage_providers"`
}

func (c *Ctx) AddFmsCloudFiles(ctx context.Context, req AddFmsCloudFilesReq) (*model.FmsCloudFiles, error) {

	query := fmt.Sprintf("insert into %s (id, created_at, updated_at, state, name, url, size, file_type, user_id, cloud_file_storage_providers) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", "fms_cloud_files")
	_, err := c.mysqlConn.ExecCtx(ctx, query, req.Id, req.CreatedAt, req.UpdatedAt, req.State, req.Name, req.Url, req.Size, req.FileType, req.UserId, req.CloudFileStorageProviders)
	//fmt.Println(ret.RowsAffected())

	return nil, err
}

func (c *Ctx) GetCloudFileInfoById(ctx context.Context,
	id string) (*model.FmsCloudFiles, error) {

	cacheKey := fmt.Sprintf(CacheKeyCloudFileInfoById, id)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var fileInfo model.FmsCloudFiles
		query := fmt.Sprintf("select * from fms_cloud_files where `id` = ? limit 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &fileInfo, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, v, query, id)
		})

		if err != nil {
			return nil, err
		}

		return fileInfo, nil
	})

	if v != nil {
		info := v.(model.FmsCloudFiles)
		return &info, err
	}

	return nil, err
}
