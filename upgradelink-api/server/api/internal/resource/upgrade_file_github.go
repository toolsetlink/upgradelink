package resource

import (
	"context"
	"fmt"
	"time"

	"upgradelink-api/server/api/internal/resource/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	CacheKeyFileGithubInfoByFileId = PREFIX + "FILE_GITHUB_INFO:FILE_ID:%v"
)

func (c *Ctx) GetFileGithubInfoByFileId(ctx context.Context,
	fileId string) (*model.UpgradeFileGithub, error) {

	cacheKey := fmt.Sprintf(CacheKeyFileGithubInfoByFileId, fileId)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var fileInfo model.UpgradeFileGithub
		query := fmt.Sprintf("select * from upgrade_file_github where `file_id` = ? limit 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &fileInfo, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, v, query, fileId)
		})

		if err != nil {
			return nil, err
		}

		return fileInfo, nil
	})

	if v != nil {
		info := v.(model.UpgradeFileGithub)
		return &info, err
	}

	return nil, err
}

type AddFileGithubReq struct {
	Id        int64     `db:"id"`         //id
	CreatedAt time.Time `db:"created_at"` // Create Time | 创建日期
	UpdatedAt time.Time `db:"updated_at"` // Update Time | 修改日期
	FileId    string    `db:"file_id"`    // The file''s url | 文件地址
	Url       string    `db:"url"`        // The user who upload the file | 上传用户的 ID
}

func (c *Ctx) AddFileGithub(ctx context.Context, req AddFileGithubReq) (*model.FmsCloudFiles, error) {

	query := fmt.Sprintf("insert into %s (file_id, url) values (?, ?)", "upgrade_file_github")
	_, err := c.mysqlConn.ExecCtx(ctx, query, req.FileId, req.Url)
	//fmt.Println(ret.RowsAffected())

	return nil, err
}
