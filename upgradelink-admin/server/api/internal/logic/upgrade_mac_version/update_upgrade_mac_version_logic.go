package upgrade_mac_version

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/http_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"

	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/ent"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgrademacversion"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUpgradeMacVersionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUpgradeMacVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUpgradeMacVersionLogic {
	return &UpdateUpgradeMacVersionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUpgradeMacVersionLogic) UpdateUpgradeMacVersion(req *types.UpgradeMacVersionInfo) (*types.BaseMsgResp, error) {

	// 效验请求参数数据
	err := l.CheckUpdateUpgradeMacVersion(req)
	if err != nil {
		return nil, err
	}

	reqModel := ent.UpgradeMacVersion{
		ID: *req.Id,
	}
	err = l.svcCtx.DB.UpgradeMacVersion.UpdateOne(&reqModel).
		SetNotNilMACID(req.MacId).
		SetNotNilCloudFileID(req.CloudFileId).
		SetNotNilVersionName(req.VersionName).
		SetNotNilVersionCode(req.VersionCode).
		SetNotNilArch(req.Arch).
		SetNotNilDescription(req.Description).
		SetNotNilIsDel(req.IsDel).
		SetNotNilCreateAt(pointy.GetTimeMilliPointer(req.CreateAt)).
		SetNotNilUpdateAt(pointy.GetTimeMilliPointer(req.UpdateAt)).
		Exec(l.ctx)

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.UpdateSuccess)}, nil
}

func (l *UpdateUpgradeMacVersionLogic) CheckUpdateUpgradeMacVersion(req *types.UpgradeMacVersionInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeMacVersion
	predicates = append(predicates, upgrademacversion.MACID(*req.MacId))
	predicates = append(predicates, upgrademacversion.IDNEQ(*req.Id))
	predicates = append(predicates, upgrademacversion.VersionName(*req.VersionName))
	predicates = append(predicates, upgrademacversion.IsDelEQ(0))
	predicates = append(predicates, upgrademacversion.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count, err := l.svcCtx.DB.UpgradeMacVersion.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return http_error.NewCodeBadRequestError("当前应用版本名重复")
	}

	// 判断是否重复
	var predicates1 []predicate.UpgradeMacVersion
	predicates1 = append(predicates1, upgrademacversion.MACID(*req.MacId))
	predicates1 = append(predicates1, upgrademacversion.IDNEQ(*req.Id))
	predicates1 = append(predicates1, upgrademacversion.VersionCode(*req.VersionCode))
	predicates1 = append(predicates1, upgrademacversion.IsDelEQ(0))
	predicates1 = append(predicates1, upgrademacversion.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count1, err := l.svcCtx.DB.UpgradeMacVersion.Query().Where(predicates1...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count1 > 0 {
		return http_error.NewCodeBadRequestError("当前应用版本号重复")
	}

	return nil
}
