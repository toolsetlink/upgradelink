package upgrade_url_version

import (
	"context"
	"fmt"

	"upgradelink-admin-upgrade/server/ent/predicate"
	"upgradelink-admin-upgrade/server/ent/upgradeurlversion"
	"upgradelink-admin-upgrade/server/internal/logic/base"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
	"upgradelink-admin-upgrade/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUpgradeUrlVersionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger

	companyID int
}

func NewCreateUpgradeUrlVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUpgradeUrlVersionLogic {
	return &CreateUpgradeUrlVersionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUpgradeUrlVersionLogic) CreateUpgradeUrlVersion(req *types.UpgradeUrlVersionInfo) (*types.BaseMsgResp, error) {

	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}
	l.companyID = companyID

	// 效验请求参数数据
	err = l.CheckCreateUpgradeUrlVersion(req)
	if err != nil {
		fmt.Println(err)
		return nil, errorx.NewCodeInvalidArgumentError(err.Error())
	}

	isDel := int32(0)
	_, err = l.svcCtx.DB.UpgradeUrlVersion.Create().
		SetNotNilCompanyID(&companyID).
		SetNotNilURLID(req.UrlId).
		SetNotNilURLPath(req.UrlPath).
		SetNotNilVersionName(req.VersionName).
		SetNotNilVersionCode(req.VersionCode).
		SetNotNilDescription(req.Description).
		SetNotNilIsDel(&isDel).
		SetNotNilCreateAt(pointy.GetTimeMilliPointer(req.CreateAt)).
		SetNotNilUpdateAt(pointy.GetTimeMilliPointer(req.UpdateAt)).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: errormsg.CreateSuccess}, nil
}

func (l *CreateUpgradeUrlVersionLogic) CheckCreateUpgradeUrlVersion(req *types.UpgradeUrlVersionInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeUrlVersion
	predicates = append(predicates, upgradeurlversion.URLID(*req.UrlId))
	predicates = append(predicates, upgradeurlversion.VersionName(*req.VersionName))
	predicates = append(predicates, upgradeurlversion.IsDelEQ(0))
	predicates = append(predicates, upgradeurlversion.CompanyIDEQ(l.companyID))
	count, err := l.svcCtx.DB.UpgradeUrlVersion.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return errorx.NewCodeInvalidArgumentError("当前应用版本名重复")
	}

	// 判断是否重复
	var predicates1 []predicate.UpgradeUrlVersion
	predicates1 = append(predicates1, upgradeurlversion.URLID(*req.UrlId))
	predicates1 = append(predicates1, upgradeurlversion.VersionCode(*req.VersionCode))
	predicates1 = append(predicates1, upgradeurlversion.IsDelEQ(0))
	predicates1 = append(predicates1, upgradeurlversion.CompanyIDEQ(l.companyID))
	count1, err := l.svcCtx.DB.UpgradeUrlVersion.Query().Where(predicates1...).Count(l.ctx)

	if err != nil {
		return err
	}
	if count1 > 0 {
		return errorx.NewCodeInvalidArgumentError("当前应用版本号重复")
	}

	return nil
}
