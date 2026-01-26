package upgrade_lnx_version

import (
	"context"
	"fmt"
	"time"
	"upgradelink-admin/server/api/internal/common"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgradelnxversion"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"upgradelink-admin/server/api/internal/common/utils/pointy"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeLnxVersionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeLnxVersionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeLnxVersionListLogic {
	return &GetUpgradeLnxVersionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeLnxVersionListLogic) GetUpgradeLnxVersionList(req *types.UpgradeLnxVersionListReq) (*types.UpgradeLnxVersionListResp, error) {
	var predicates []predicate.UpgradeLnxVersion
	predicates = append(predicates, upgradelnxversion.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	predicates = append(predicates, upgradelnxversion.IsDelEQ(0))

	if req.LnxId != nil {
		predicates = append(predicates, upgradelnxversion.LnxIDEQ(*req.LnxId))
	}
	if req.CloudFileId != nil {
		predicates = append(predicates, upgradelnxversion.CloudFileIDContains(*req.CloudFileId))
	}
	if req.VersionName != nil {
		predicates = append(predicates, upgradelnxversion.VersionNameContains(*req.VersionName))
	}
	if req.VersionCode != nil {
		predicates = append(predicates, upgradelnxversion.VersionCodeEQ(*req.VersionCode))
	}
	if req.Description != nil {
		predicates = append(predicates, upgradelnxversion.DescriptionContains(*req.Description))
	}
	if req.IsDel != nil {
		predicates = append(predicates, upgradelnxversion.IsDelEQ(*req.IsDel))
	}
	if req.CreateAt != nil {
		predicates = append(predicates, upgradelnxversion.CreateAtGTE(time.UnixMilli(*req.CreateAt)))
	}
	if req.UpdateAt != nil {
		predicates = append(predicates, upgradelnxversion.UpdateAtGTE(time.UnixMilli(*req.UpdateAt)))
	}
	data, err := l.svcCtx.DB.UpgradeLnxVersion.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.UpgradeLnxVersionListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {

		fileData, err := l.svcCtx.DB.UpgradeLnx.Get(l.ctx, v.LnxID)
		if err != nil {
			return nil, db_error.DefaultEntError(l.Logger, err, req)
		}

		cloudFileData, err := l.svcCtx.DB.FmsCloudFile.Get(l.ctx, v.CloudFileID)
		if err != nil {
			fmt.Println("err: ", err)
			return nil, db_error.DefaultEntError(l.Logger, err, req)
		}

		// 文件大小
		fileSize := common.BytesToMBString(cloudFileData.Size)

		resp.Data.Data = append(resp.Data.Data,
			types.RespUpgradeLnxVersionInfo{
				Id:              &v.ID,
				LnxId:           &v.LnxID,
				LnxName:         &fileData.Name,
				CloudFileId:     &v.CloudFileID,
				CloudFileName:   &cloudFileData.Name,
				CloudFilePath:   &cloudFileData.URL,
				VersionName:     &v.VersionName,
				VersionCode:     &v.VersionCode,
				Arch:            &v.Arch,
				VersionFileSize: &fileSize,
				Description:     &v.Description,
				IsDel:           &v.IsDel,
				CreateAt:        pointy.GetUnixMilliPointer(v.CreateAt.UnixMilli()),
				UpdateAt:        pointy.GetUnixMilliPointer(v.UpdateAt.UnixMilli()),
			})
	}

	return resp, nil
}
