package upgrade_win_version

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/enum"
	"upgradelink-admin/server/api/internal/common/http_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"
	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgradewinversion"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUpgradeWinVersionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUpgradeWinVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUpgradeWinVersionLogic {
	return &CreateUpgradeWinVersionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUpgradeWinVersionLogic) CreateUpgradeWinVersion(req *types.UpgradeWinVersionInfo) (*types.BaseMsgResp, error) {

	// 效验请求参数数据
	err := l.CheckCreateUpgradeWinVersion(req)
	if err != nil {
		return nil, err
	}

	intDelFalse := enum.IsDelFalse
	_, err = l.svcCtx.DB.UpgradeWinVersion.Create().
		SetNotNilCompanyID(companyctx.GetCompanyIDPointerFromCtx(l.ctx)).
		SetNotNilWinID(req.WinId).
		SetNotNilCloudFileID(req.CloudFileId).
		SetNotNilVersionName(req.VersionName).
		SetNotNilVersionCode(req.VersionCode).
		SetNotNilArch(req.Arch).
		SetNotNilDescription(req.Description).
		SetNotNilIsDel(&intDelFalse).
		SetNotNilCreateAt(pointy.GetTimeMilliPointer(req.CreateAt)).
		SetNotNilUpdateAt(pointy.GetTimeMilliPointer(req.UpdateAt)).
		Save(l.ctx)

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.CreateSuccess)}, nil
}

func (l *CreateUpgradeWinVersionLogic) CheckCreateUpgradeWinVersion(req *types.UpgradeWinVersionInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeWinVersion
	predicates = append(predicates, upgradewinversion.WinID(*req.WinId))
	predicates = append(predicates, upgradewinversion.VersionName(*req.VersionName))
	predicates = append(predicates, upgradewinversion.IsDelEQ(0))
	predicates = append(predicates, upgradewinversion.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count, err := l.svcCtx.DB.UpgradeWinVersion.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return http_error.NewCodeBadRequestError(l.svcCtx.Trans.Trans(l.ctx, i18n.AppVersionDuplicate))
	}

	// 判断是否重复
	var predicates1 []predicate.UpgradeWinVersion
	predicates1 = append(predicates1, upgradewinversion.WinID(*req.WinId))
	predicates1 = append(predicates1, upgradewinversion.VersionCode(*req.VersionCode))
	predicates1 = append(predicates1, upgradewinversion.IsDelEQ(0))
	predicates1 = append(predicates1, upgradewinversion.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count1, err := l.svcCtx.DB.UpgradeWinVersion.Query().Where(predicates1...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count1 > 0 {
		return http_error.NewCodeBadRequestError(l.svcCtx.Trans.Trans(l.ctx, i18n.AppVersionCodeDuplicate))
	}

	return nil
}
