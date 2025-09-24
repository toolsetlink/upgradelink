package upgrade_apk

import (
	"context"
	"crypto/rand"
	"encoding/base64"

	"upgradelink-admin-upgrade/ent/predicate"
	"upgradelink-admin-upgrade/ent/upgradeapk"
	"upgradelink-admin-upgrade/internal/logic/base"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
	"upgradelink-admin-upgrade/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUpgradeApkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger

	companyID int
}

func NewCreateUpgradeApkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUpgradeApkLogic {
	return &CreateUpgradeApkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUpgradeApkLogic) CreateUpgradeApk(req *types.UpgradeApkInfo) (*types.BaseMsgResp, error) {

	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}
	l.companyID = companyID

	// 效验请求参数数据
	err = l.CheckCreateUpgrade(req)
	if err != nil {
		return nil, errorx.NewCodeInvalidArgumentError(err.Error())
	}

	// 生成Access Key (16字节 -> 24字符Base64)
	apkBytes := make([]byte, 16)
	_, _ = rand.Read(apkBytes)
	apkKey := base64.RawURLEncoding.EncodeToString(apkBytes)

	isDel := int32(0)
	_, err = l.svcCtx.DB.UpgradeApk.Create().
		SetNotNilCompanyID(&companyID).
		SetNotNilKey(&apkKey).
		SetNotNilName(req.Name).
		SetNotNilPackageName(req.PackageName).
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

func (l *CreateUpgradeApkLogic) CheckCreateUpgrade(req *types.UpgradeApkInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeApk
	predicates = append(predicates, upgradeapk.Name(*req.Name))
	predicates = append(predicates, upgradeapk.IsDelEQ(0))
	predicates = append(predicates, upgradeapk.CompanyIDEQ(l.companyID))
	count, err := l.svcCtx.DB.UpgradeApk.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return errorx.NewCodeInvalidArgumentError("应用名称重复")
	}

	// 判断是否重复
	var predicates1 []predicate.UpgradeApk
	predicates1 = append(predicates1, upgradeapk.PackageName(*req.PackageName))
	predicates1 = append(predicates1, upgradeapk.IsDelEQ(0))
	predicates1 = append(predicates1, upgradeapk.CompanyIDEQ(l.companyID))
	count1, err := l.svcCtx.DB.UpgradeApk.Query().Where(predicates1...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count1 > 0 {
		return errorx.NewCodeInvalidArgumentError("应用包名名称重复")
	}

	return nil
}
