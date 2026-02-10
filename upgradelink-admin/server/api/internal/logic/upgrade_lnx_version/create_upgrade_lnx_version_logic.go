package upgrade_lnx_version

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/enum"
	"upgradelink-admin/server/api/internal/common/http_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"

	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgradelnxversion"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUpgradeLnxVersionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUpgradeLnxVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUpgradeLnxVersionLogic {
	return &CreateUpgradeLnxVersionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUpgradeLnxVersionLogic) CreateUpgradeLnxVersion(req *types.UpgradeLnxVersionInfo) (*types.BaseMsgResp, error) {

	// 效验请求参数数据
	err := l.CheckCreateUpgradeLnxVersion(req)
	if err != nil {
		return nil, err
	}

	intDelFalse := enum.IsDelFalse
	_, err = l.svcCtx.DB.UpgradeLnxVersion.Create().
		SetNotNilCompanyID(companyctx.GetCompanyIDPointerFromCtx(l.ctx)).
		SetNotNilLnxID(req.LnxId).
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

func (l *CreateUpgradeLnxVersionLogic) CheckCreateUpgradeLnxVersion(req *types.UpgradeLnxVersionInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeLnxVersion
	predicates = append(predicates, upgradelnxversion.LnxID(*req.LnxId))
	predicates = append(predicates, upgradelnxversion.VersionName(*req.VersionName))
	predicates = append(predicates, upgradelnxversion.IsDelEQ(0))
	predicates = append(predicates, upgradelnxversion.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count, err := l.svcCtx.DB.UpgradeLnxVersion.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return http_error.NewCodeBadRequestError(l.svcCtx.Trans.Trans(l.ctx, i18n.AppVersionDuplicate))
	}

	// 判断是否重复
	var predicates1 []predicate.UpgradeLnxVersion
	predicates1 = append(predicates1, upgradelnxversion.LnxID(*req.LnxId))
	predicates1 = append(predicates1, upgradelnxversion.VersionCode(*req.VersionCode))
	predicates1 = append(predicates1, upgradelnxversion.IsDelEQ(0))
	predicates1 = append(predicates1, upgradelnxversion.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count1, err := l.svcCtx.DB.UpgradeLnxVersion.Query().Where(predicates1...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count1 > 0 {
		return http_error.NewCodeBadRequestError(l.svcCtx.Trans.Trans(l.ctx, i18n.AppVersionCodeDuplicate))
	}

	return nil
}
