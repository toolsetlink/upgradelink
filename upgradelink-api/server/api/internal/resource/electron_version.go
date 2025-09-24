package resource

import (
	"context"
	"fmt"
	"time"

	"upgradelink-api/server/api/internal/resource/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const (
	CacheKeyElectronVersionCodeListByElectronIdAndVersionCode = PREFIX + "ELECTRON_VERSION_CODE_LIST:ELECTRON_ID:%v:VERSION_CODE:%v"

	CacheKeyElectronVersionListByElectronIdAndVersionCodeAndPlatformAndArch = PREFIX + "ELECTRON_VERSION_LIST:ELECTRON_ID:%v:VERSION_CODE:%v:platform:%v:arch:%v"
	CacheKeyElectronVersionInfoById                                         = PREFIX + "ELECTRON_VERSION_INFO:ID:%v"
	CacheKeyElectronVersionLastInfoByElectronIdAndPlatformAndArch           = PREFIX + "ELECTRON_VERSION_LAST_INFO:ELECTRON_ID:%v:platform:%v:arch:%v"
	CacheKeyElectronVersionInfoByElectronIdAndVersionCodeAndPlatformAndArch = PREFIX + "ELECTRON_VERSION_INFO:ELECTRON_ID:%v:VERSION_CODE:%v:platform:%v:arch:%v"

	CacheKeyElectronVersionListByElectronIdAndVersionCode = PREFIX + "ELECTRON_VERSION_LIST:ELECTRON_ID:%v:VERSION_CODE:%v"
)

// GetElectronVersionListByElectronIdAndVersionCode
// 获取对应版本 versionCode 的版本列表
func (c *Ctx) GetElectronVersionListByElectronIdAndVersionCode(ctx context.Context, electronId int64, versionCode int64) ([]*model.UpgradeElectronVersion, error) {

	cacheKey := fmt.Sprintf(CacheKeyElectronVersionListByElectronIdAndVersionCode, electronId, versionCode)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var ElectronVersionList []*model.UpgradeElectronVersion
		query := fmt.Sprintf("select * from upgrade_electron_version where `electron_id` = ? and is_del = 0 and `version_code` = ? order by platform desc ")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &ElectronVersionList, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowsCtx(ctx, &ElectronVersionList, query, electronId, versionCode)
		})

		if err != nil {
			return nil, err
		}
		return ElectronVersionList, nil
	})

	if v != nil {
		list := v.([]*model.UpgradeElectronVersion)
		return list, err
	}

	return nil, err
}

// GetElectronVersionListByKeyAndPlatformAndArch
// 获取大于 versionCode 的版本列表
func (c *Ctx) GetElectronVersionListByKeyAndPlatformAndArch(ctx context.Context, electronId int64, versionCode int64, platform string, arch string) ([]*model.UpgradeElectronVersion, error) {

	cacheKey := fmt.Sprintf(CacheKeyElectronVersionListByElectronIdAndVersionCodeAndPlatformAndArch, electronId, versionCode, platform, arch)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var ElectronVersionList []*model.UpgradeElectronVersion
		query := fmt.Sprintf("select * from upgrade_electron_version where `electron_id` = ? and `platform` = ? and `arch` = ? and is_del = 0 and `version_code` > ? order by version_code desc ")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &ElectronVersionList, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowsCtx(ctx, &ElectronVersionList, query, electronId, platform, arch, versionCode)
		})

		if err != nil {
			return nil, err
		}
		return ElectronVersionList, nil
	})

	if v != nil {
		list := v.([]*model.UpgradeElectronVersion)
		return list, err
	}

	return nil, err
}

// GetElectronVersionLastInfoByElectronIdAndPlatformAndArch
// 获取最新版本的信息
func (c *Ctx) GetElectronVersionLastInfoByElectronIdAndPlatformAndArch(ctx context.Context, electronId int64, platform string, arch string) (*model.UpgradeElectronVersion, error) {

	cacheKey := fmt.Sprintf(CacheKeyElectronVersionLastInfoByElectronIdAndPlatformAndArch, electronId, platform, arch)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info model.UpgradeElectronVersion
		query := fmt.Sprintf("select * from upgrade_electron_version where `electron_id` = ? and is_del = 0 and `platform` = ? and `arch` = ? order by version_code desc limit 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &info, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, &info, query, electronId, platform, arch)
		})

		if err != nil {
			return nil, err
		}
		return info, nil
	})

	if v != nil {
		info := v.(model.UpgradeElectronVersion)
		return &info, err
	}

	return nil, err
}

// GetElectronVersionInfoByElectronIdAndVersionCodeAndPlatformAndArch 获取固定版本的信息
func (c *Ctx) GetElectronVersionInfoByElectronIdAndVersionCodeAndPlatformAndArch(ctx context.Context, electronId int64, versionCode int64, platform string, arch string) (*model.UpgradeElectronVersion, error) {

	cacheKey := fmt.Sprintf(CacheKeyElectronVersionInfoByElectronIdAndVersionCodeAndPlatformAndArch, electronId, versionCode, platform, arch)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info model.UpgradeElectronVersion
		query := fmt.Sprintf("select * from upgrade_electron_version where `electron_id` = ? and is_del = 0 and `version_code` = ? and `platform` = ? and `arch` = ? LIMIT 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &info, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, &info, query, electronId, versionCode, platform, arch)
		})

		if err != nil {
			return nil, err
		}
		return info, nil
	})

	if v != nil {
		info := v.(model.UpgradeElectronVersion)
		return &info, err
	}

	return nil, err
}

// GetElectronVersionInfoById
// 通过 Electron version id 获取信息
func (c *Ctx) GetElectronVersionInfoById(ctx context.Context,
	id int64) (*model.UpgradeElectronVersion, error) {

	cacheKey := fmt.Sprintf(CacheKeyElectronVersionInfoById, id)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var info model.UpgradeElectronVersion
		query := fmt.Sprintf("select * from upgrade_electron_version where `id` = ?  limit 1")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &info, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowCtx(ctx, v, query, id)
		})
		if err != nil {
			return nil, err
		}

		return info, nil
	})

	if v != nil {
		info := v.(model.UpgradeElectronVersion)
		return &info, err
	}

	return nil, err
}

type AddElectronVersionReq struct {
	Id                 int64     `db:"id"`                    // ID
	CompanyId          int64     `db:"company_id"`            // 公司ID
	ElectronId         int64     `db:"electron_id"`           // electron应用ID
	CloudFileId        string    `db:"cloud_file_id"`         // 云文件id
	Sha512             string    `db:"sha512"`                // 生成的 sha512
	InstallCloudFileId string    `db:"install_cloud_file_id"` // 云文件id
	InstallSha512      string    `db:"install_sha512"`        // 生成的 sha512
	VersionName        string    `db:"version_name"`          // 版本名
	VersionCode        int64     `db:"version_code"`          // 版本号
	Platform           string    `db:"platform"`              // 操作系统: linux、darwin、windows
	Arch               string    `db:"arch"`                  // 机器架构: x64、arm64
	Description        string    `db:"description"`           // 描述信息
	IsDel              int64     `db:"is_del"`                // 是否删除 0：正常；1：已删除
	CreateAt           time.Time `db:"create_at"`             // 创建时间
	UpdateAt           time.Time `db:"update_at"`             // 修改时间
}

func (c *Ctx) AddElectronVersion(ctx context.Context, req AddElectronVersionReq) (int64, error) {

	query := fmt.Sprintf("insert into %s (company_id, electron_id, cloud_file_id, sha512, install_cloud_file_id, install_sha512, version_name, version_code, platform, arch, description) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", "upgrade_electron_version")
	ret, err := c.mysqlConn.ExecCtx(ctx, query, req.CompanyId, req.ElectronId, req.CloudFileId, req.Sha512, req.InstallCloudFileId, req.InstallSha512, req.VersionName, req.VersionCode, req.Platform, req.Arch, req.Description)
	if err != nil {
		return 0, err
	}
	id, _ := ret.LastInsertId()
	return id, err
}

type GetElectronVersionCodeListByElectronIdAndVersionCodeResp struct {
	VersionName string    `db:"version_name"` // 版本名
	VersionCode int64     `db:"version_code"` // 版本号
	CreateAt    time.Time `db:"create_at"`    // 创建时间
}

// GetElectronVersionCodeListByElectronIdAndVersionCode
// 获取大于 versionCode 的版本code列表
func (c *Ctx) GetElectronVersionCodeListByElectronIdAndVersionCode(ctx context.Context, electronId int64, versionCode int64) ([]*GetElectronVersionCodeListByElectronIdAndVersionCodeResp, error) {

	cacheKey := fmt.Sprintf(CacheKeyElectronVersionCodeListByElectronIdAndVersionCode, electronId, versionCode)

	// 内存缓存
	v, err := c.localCache.Take(cacheKey, func() (interface{}, error) {
		// sql 缓存查询
		var ElectronVersionList []*GetElectronVersionCodeListByElectronIdAndVersionCodeResp
		query := fmt.Sprintf("SELECT version_code, MIN(version_name) AS version_name, MIN(create_at) AS create_at FROM upgrade_electron_version WHERE electron_id = ? AND is_del = 0 and `version_code` > ? GROUP BY version_code ORDER BY version_code DESC ")
		err := c.mysqlConnCache.QueryRowCtx(ctx, &ElectronVersionList, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
			return c.mysqlConn.QueryRowsCtx(ctx, &ElectronVersionList, query, electronId, versionCode)
		})

		if err != nil {
			return nil, err
		}
		return ElectronVersionList, nil
	})

	if v != nil {
		list := v.([]*GetElectronVersionCodeListByElectronIdAndVersionCodeResp)
		return list, err
	}

	return nil, err
}

// DelElectronVersionById
// 删除 electron 版本
func (c *Ctx) DelElectronVersionById(ctx context.Context, id int64) (int64, error) {

	query := fmt.Sprintf("update %s set `is_del` = 1 where `id` = ?", "upgrade_electron_version")
	ret, err := c.mysqlConn.ExecCtx(ctx, query, id)
	if err != nil {
		return 0, err
	}
	id, _ = ret.LastInsertId()

	return id, err
}
