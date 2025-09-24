package resource

import (
	"context"
	"fmt"
	"time"

	"upgradelink-api/server/api/internal/resource/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	CacheKeyTauriVersionCodeListByTauriIdAndVersionCode = PREFIX + "TAURI_VERSION_CODE_LIST:TAURI_ID:%v:VERSION_CODE:%v"

	CacheKeyTauriVersionListByTauriIdAndVersionCodeAndTargetAndArch = PREFIX + "TAURI_VERSION_LIST:TAURI_ID:%v:VERSION_CODE:%v:target:%v:arch:%v"
	CacheKeyTauriVersionInfoById                                    = PREFIX + "TAURI_VERSION_INFO:ID:%v"
	CacheKeyTauriVersionLastInfoByTauriIdAndTargetAndArch           = PREFIX + "TAURI_VERSION_LAST_INFO:TAURI_ID:%v:target:%v:arch:%v"
	CacheKeyTauriVersionInfoByTauriIdAndVersionCodeAndTargetAndArch = PREFIX + "TAURI_VERSION_INFO:TAURI_ID:%v:VERSION_CODE:%v:target:%v:arch:%v"

	CacheKeyTauriVersionListByTauriIdAndVersionCode = PREFIX + "TAURI_VERSION_LIST:TAURI_ID:%v:VERSION_CODE:%v"
)

// GetTauriVersionListByTauriIdAndVersionCode
// 获取对应版本 versionCode 的版本列表
func (c *Ctx) GetTauriVersionListByTauriIdAndVersionCode(ctx context.Context, tauriId int64, versionCode int64) ([]*model.UpgradeTauriVersion, error) {

	cacheKey := fmt.Sprintf(CacheKeyTauriVersionListByTauriIdAndVersionCode, tauriId, versionCode)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var TauriVersionList []*model.UpgradeTauriVersion
		query := fmt.Sprintf("select * from upgrade_tauri_version where `tauri_id` = ? and is_del = 0 and `version_code` = ? order by target desc ")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &TauriVersionList, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowsCtx(ctx, &TauriVersionList, query, tauriId, versionCode)
		})

		if err != nil {
			return nil, err
		}
		return TauriVersionList, nil
	})

	if v != nil {
		list := v.([]*model.UpgradeTauriVersion)
		return list, err
	}

	return nil, err
}

// GetTauriVersionListByKeyAndTargetAndArch
// 获取大于 versionCode 的版本列表
func (c *Ctx) GetTauriVersionListByKeyAndTargetAndArch(ctx context.Context, tauriId int64, versionCode int64, target string, arch string) ([]*model.UpgradeTauriVersion, error) {

	cacheKey := fmt.Sprintf(CacheKeyTauriVersionListByTauriIdAndVersionCodeAndTargetAndArch, tauriId, versionCode, target, arch)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var TauriVersionList []*model.UpgradeTauriVersion
		query := fmt.Sprintf("select * from upgrade_tauri_version where `tauri_id` = ? and `target` = ? and `arch` = ? and is_del = 0 and `version_code` > ? order by version_code desc ")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &TauriVersionList, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowsCtx(ctx, &TauriVersionList, query, tauriId, target, arch, versionCode)
		})

		if err != nil {
			return nil, err
		}
		return TauriVersionList, nil
	})

	if v != nil {
		list := v.([]*model.UpgradeTauriVersion)
		return list, err
	}

	return nil, err
}

// GetTauriVersionLastInfoByTauriIdAndTargetAndArch
// 获取最新版本的信息
func (c *Ctx) GetTauriVersionLastInfoByTauriIdAndTargetAndArch(ctx context.Context, tauriId int64, target string, arch string) (*model.UpgradeTauriVersion, error) {

	cacheKey := fmt.Sprintf(CacheKeyTauriVersionLastInfoByTauriIdAndTargetAndArch, tauriId, target, arch)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info model.UpgradeTauriVersion
		query := fmt.Sprintf("select * from upgrade_tauri_version where `tauri_id` = ? and is_del = 0 and `target` = ? and `arch` = ? order by version_code desc limit 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &info, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, &info, query, tauriId, target, arch)
		})

		if err != nil {
			return nil, err
		}
		return info, nil
	})

	if v != nil {
		info := v.(model.UpgradeTauriVersion)
		return &info, err
	}

	return nil, err
}

// GetTauriVersionInfoByTauriIdAndVersionCodeAndTargetAndArch 获取固定版本的信息
func (c *Ctx) GetTauriVersionInfoByTauriIdAndVersionCodeAndTargetAndArch(ctx context.Context, tauriId int64, versionCode int64, target string, arch string) (*model.UpgradeTauriVersion, error) {

	cacheKey := fmt.Sprintf(CacheKeyTauriVersionInfoByTauriIdAndVersionCodeAndTargetAndArch, tauriId, versionCode, target, arch)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info model.UpgradeTauriVersion
		query := fmt.Sprintf("select * from upgrade_tauri_version where `tauri_id` = ? and is_del = 0 and `version_code` = ? and `target` = ? and `arch` = ? LIMIT 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &info, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, &info, query, tauriId, versionCode, target, arch)
		})

		if err != nil {
			return nil, err
		}
		return info, nil
	})

	if v != nil {
		info := v.(model.UpgradeTauriVersion)
		return &info, err
	}

	return nil, err
}

// GetTauriVersionInfoById
// 通过 Tauri version id 获取信息
func (c *Ctx) GetTauriVersionInfoById(ctx context.Context,
	id int64) (*model.UpgradeTauriVersion, error) {

	cacheKey := fmt.Sprintf(CacheKeyTauriVersionInfoById, id)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info model.UpgradeTauriVersion
		query := fmt.Sprintf("select * from upgrade_tauri_version where `id` = ?  limit 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &info, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, v, query, id)
		})
		if err != nil {
			return nil, err
		}

		return info, nil
	})

	if v != nil {
		info := v.(model.UpgradeTauriVersion)
		return &info, err
	}

	return nil, err
}

type AddTauriVersionReq struct {
	Id                 int64     `db:"id"`                    // ID
	CompanyId          int64     `db:"company_id"`            // 公司ID
	TauriId            int64     `db:"tauri_id"`              // tauri应用ID
	CloudFileId        string    `db:"cloud_file_id"`         // 云文件id
	InstallCloudFileId string    `db:"install_cloud_file_id"` // 云文件id
	VersionName        string    `db:"version_name"`          // 版本名
	VersionCode        int64     `db:"version_code"`          // 版本号
	Target             string    `db:"target"`                // 操作系统:linux、darwin、windows
	Arch               string    `db:"arch"`                  // 机器架构:x86_64、i686、aarch64、armv7
	Signature          string    `db:"signature"`             // 生成的 .sig 文件的内容
	Description        string    `db:"description"`           // 描述信息
	IsDel              int64     `db:"is_del"`                // 是否删除 0：正常；1：已删除
	CreateAt           time.Time `db:"create_at"`             // 创建时间
	UpdateAt           time.Time `db:"update_at"`             // 修改时间
}

func (c *Ctx) AddTauriVersion(ctx context.Context, req AddTauriVersionReq) (int64, error) {

	query := fmt.Sprintf("insert into %s (company_id, tauri_id, cloud_file_id, install_cloud_file_id, version_name, version_code, target, arch, signature) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", "upgrade_tauri_version")
	ret, err := c.mysqlConn.ExecCtx(ctx, query, req.CompanyId, req.TauriId, req.CloudFileId, req.InstallCloudFileId, req.VersionName, req.VersionCode, req.Target, req.Arch, req.Signature)
	if err != nil {
		return 0, err
	}
	id, _ := ret.LastInsertId()
	return id, err
}

type GetTauriVersionCodeListByTauriIdAndVersionCodeResp struct {
	VersionName string    `db:"version_name"` // 版本名
	VersionCode int64     `db:"version_code"` // 版本号
	CreateAt    time.Time `db:"create_at"`    // 创建时间
}

// GetTauriVersionCodeListByTauriIdAndVersionCode
// 获取大于 versionCode 的版本code列表
func (c *Ctx) GetTauriVersionCodeListByTauriIdAndVersionCode(ctx context.Context, tauriId int64, versionCode int64) ([]*GetTauriVersionCodeListByTauriIdAndVersionCodeResp, error) {

	cacheKey := fmt.Sprintf(CacheKeyTauriVersionCodeListByTauriIdAndVersionCode, tauriId, versionCode)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var TauriVersionList []*GetTauriVersionCodeListByTauriIdAndVersionCodeResp
		query := fmt.Sprintf("SELECT version_code, MIN(version_name) AS version_name, MIN(create_at) AS create_at FROM upgrade_tauri_version WHERE tauri_id = ? AND is_del = 0 and `version_code` > ? GROUP BY version_code ORDER BY version_code DESC ")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &TauriVersionList, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowsCtx(ctx, &TauriVersionList, query, tauriId, versionCode)
		})

		if err != nil {
			return nil, err
		}
		return TauriVersionList, nil
	})

	if v != nil {
		list := v.([]*GetTauriVersionCodeListByTauriIdAndVersionCodeResp)
		return list, err
	}

	return nil, err
}

// DelTauriVersionById
// 删除 tauri 版本
func (c *Ctx) DelTauriVersionById(ctx context.Context, id int64) (int64, error) {

	query := fmt.Sprintf("update %s set `is_del` = 1 where `id` = ?", "upgrade_tauri_version")
	ret, err := c.mysqlConn.ExecCtx(ctx, query, id)
	if err != nil {
		return 0, err
	}
	id, _ = ret.LastInsertId()

	return id, err
}
