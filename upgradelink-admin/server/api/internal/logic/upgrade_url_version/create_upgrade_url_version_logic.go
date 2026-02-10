package upgrade_url_version

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/enum"
	"upgradelink-admin/server/api/internal/common/http_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"

	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgradeurlversion"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUpgradeUrlVersionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUpgradeUrlVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUpgradeUrlVersionLogic {
	return &CreateUpgradeUrlVersionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUpgradeUrlVersionLogic) CreateUpgradeUrlVersion(req *types.UpgradeUrlVersionInfo) (*types.BaseMsgResp, error) {

	// 效验请求参数数据
	err := l.CheckCreateUpgradeUrlVersion(req)
	if err != nil {
		return nil, err
	}

	intDelFalse := enum.IsDelFalse
	_, err = l.svcCtx.DB.UpgradeUrlVersion.Create().
		SetNotNilCompanyID(companyctx.GetCompanyIDPointerFromCtx(l.ctx)).
		SetNotNilURLID(req.UrlId).
		SetNotNilURLPath(req.UrlPath).
		SetNotNilVersionName(req.VersionName).
		SetNotNilVersionCode(req.VersionCode).
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

func (l *CreateUpgradeUrlVersionLogic) CheckCreateUpgradeUrlVersion(req *types.UpgradeUrlVersionInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeUrlVersion
	predicates = append(predicates, upgradeurlversion.URLID(*req.UrlId))
	predicates = append(predicates, upgradeurlversion.VersionName(*req.VersionName))
	predicates = append(predicates, upgradeurlversion.IsDelEQ(0))
	predicates = append(predicates, upgradeurlversion.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count, err := l.svcCtx.DB.UpgradeUrlVersion.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return http_error.NewCodeBadRequestError(l.svcCtx.Trans.Trans(l.ctx, i18n.AppVersionDuplicate))
	}

	// 判断是否重复
	var predicates1 []predicate.UpgradeUrlVersion
	predicates1 = append(predicates1, upgradeurlversion.URLID(*req.UrlId))
	predicates1 = append(predicates1, upgradeurlversion.VersionCode(*req.VersionCode))
	predicates1 = append(predicates1, upgradeurlversion.IsDelEQ(0))
	predicates1 = append(predicates1, upgradeurlversion.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count1, err := l.svcCtx.DB.UpgradeUrlVersion.Query().Where(predicates1...).Count(l.ctx)

	if err != nil {
		return err
	}
	if count1 > 0 {
		return http_error.NewCodeBadRequestError(l.svcCtx.Trans.Trans(l.ctx, i18n.AppVersionCodeDuplicate))
	}

	return nil
}
