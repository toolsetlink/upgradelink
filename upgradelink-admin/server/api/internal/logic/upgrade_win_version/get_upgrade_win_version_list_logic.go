package upgrade_win_version

import (
	"context"
	"fmt"
	"time"
	"upgradelink-admin/server/api/internal/common"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgradewinversion"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"upgradelink-admin/server/api/internal/common/utils/pointy"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeWinVersionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeWinVersionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeWinVersionListLogic {
	return &GetUpgradeWinVersionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeWinVersionListLogic) GetUpgradeWinVersionList(req *types.UpgradeWinVersionListReq) (*types.UpgradeWinVersionListResp, error) {
	var predicates []predicate.UpgradeWinVersion
	predicates = append(predicates, upgradewinversion.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	predicates = append(predicates, upgradewinversion.IsDelEQ(0))

	if req.WinId != nil {
		predicates = append(predicates, upgradewinversion.WinIDEQ(*req.WinId))
	}
	if req.CloudFileId != nil {
		predicates = append(predicates, upgradewinversion.CloudFileIDContains(*req.CloudFileId))
	}
	if req.VersionName != nil {
		predicates = append(predicates, upgradewinversion.VersionNameContains(*req.VersionName))
	}
	if req.VersionCode != nil {
		predicates = append(predicates, upgradewinversion.VersionCodeEQ(*req.VersionCode))
	}
	if req.Description != nil {
		predicates = append(predicates, upgradewinversion.DescriptionContains(*req.Description))
	}
	if req.IsDel != nil {
		predicates = append(predicates, upgradewinversion.IsDelEQ(*req.IsDel))
	}
	if req.CreateAt != nil {
		predicates = append(predicates, upgradewinversion.CreateAtGTE(time.UnixMilli(*req.CreateAt)))
	}
	if req.UpdateAt != nil {
		predicates = append(predicates, upgradewinversion.UpdateAtGTE(time.UnixMilli(*req.UpdateAt)))
	}
	data, err := l.svcCtx.DB.UpgradeWinVersion.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.UpgradeWinVersionListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {

		fileData, err := l.svcCtx.DB.UpgradeWin.Get(l.ctx, v.WinID)
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
			types.RespUpgradeWinVersionInfo{
				Id:              &v.ID,
				WinId:           &v.WinID,
				WinName:         &fileData.Name,
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
