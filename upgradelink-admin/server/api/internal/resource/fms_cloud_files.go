package resource

import (
	"context"
	"fmt"
	"log"
	"time"
	"upgradelink-admin/server/api/internal/resource/model"
	"upgradelink-admin/server/api/internal/svc"
)

type AddFmsCloudFilesReq struct {
	Id                        string    `db:"id"`         // UUID
	CreatedAt                 time.Time `db:"created_at"` // Create Time | 创建日期
	UpdatedAt                 time.Time `db:"updated_at"` // Update Time | 修改日期
	State                     int64     `db:"state"`      // State true: normal false: ban | 状态 true 正常 false 禁用
	Name                      string    `db:"name"`       // The file''s name | 文件名
	Url                       string    `db:"url"`        // The file''s url | 文件地址
	Size                      uint64    `db:"size"`       // The file''s size | 文件大小
	Md5                       string    `db:"md5"`        // The file''s md5 | 文件md5
	FileType                  uint64    `db:"file_type"`  // The file''s type | 文件类型
	UserId                    string    `db:"user_id"`    // The user who upload the file | 上传用户的 ID
	CloudFileStorageProviders int64     `db:"cloud_file_storage_providers"`
}

func AddFmsCloudFiles(ctx context.Context, svcCtx *svc.ServiceContext, req AddFmsCloudFilesReq) (*model.FmsCloudFiles, error) {

	query := fmt.Sprintf("insert into %s (id, created_at, updated_at, state, name, url, size, file_type, user_id, cloud_file_storage_providers) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", "fms_cloud_files")
	_, err := svcCtx.DB.ExecContext(ctx, query, req.Id, req.CreatedAt, req.UpdatedAt, req.State, req.Name, req.Url, req.Size, req.FileType, req.UserId, req.CloudFileStorageProviders)
	//fmt.Println(ret.RowsAffected())

	return nil, err
}

func GetCloudFileInfoById(ctx context.Context, svcCtx *svc.ServiceContext, id string) (*model.FmsCloudFiles, error) {

	query := fmt.Sprintf("select * from fms_cloud_files where `id` = ? limit 1")
	queryContext, err := svcCtx.DB.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	var result []model.FmsCloudFiles
	// 解析结构体
	for queryContext.Next() {
		var info model.FmsCloudFiles
		if err := queryContext.Scan(&info); err != nil {
			log.Fatal(err)
		}
		result = append(result, info)
	}

	if len(result) == 0 {
		return nil, nil
	}

	return &result[0], nil
}
