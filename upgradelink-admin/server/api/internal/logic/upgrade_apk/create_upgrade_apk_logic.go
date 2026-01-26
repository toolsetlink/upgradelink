package upgrade_apk

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/http_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"

	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgradeapk"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUpgradeApkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUpgradeApkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUpgradeApkLogic {
	return &CreateUpgradeApkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUpgradeApkLogic) CreateUpgradeApk(req *types.UpgradeApkInfo) (*types.BaseMsgResp, error) {

	// 效验请求参数数据
	err := l.CheckCreateUpgrade(req)
	if err != nil {
		return nil, err
	}

	// 生成Access Key (16字节 -> 24字符Base64)
	apkBytes := make([]byte, 16)
	_, _ = rand.Read(apkBytes)
	apkKey := base64.RawURLEncoding.EncodeToString(apkBytes)

	isDel := int32(0)
	_, err = l.svcCtx.DB.UpgradeApk.Create().
		SetNotNilCompanyID(companyctx.GetCompanyIDPointerFromCtx(l.ctx)).
		SetNotNilKey(&apkKey).
		SetNotNilName(req.Name).
		SetNotNilPackageName(req.PackageName).
		SetNotNilDescription(req.Description).
		SetNotNilIsDel(&isDel).
		SetNotNilCreateAt(pointy.GetTimeMilliPointer(req.CreateAt)).
		SetNotNilUpdateAt(pointy.GetTimeMilliPointer(req.UpdateAt)).
		Save(l.ctx)

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.CreateSuccess)}, nil
}

func (l *CreateUpgradeApkLogic) CheckCreateUpgrade(req *types.UpgradeApkInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeApk
	predicates = append(predicates, upgradeapk.Name(*req.Name))
	predicates = append(predicates, upgradeapk.IsDelEQ(0))
	predicates = append(predicates, upgradeapk.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count, err := l.svcCtx.DB.UpgradeApk.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return http_error.NewCodeBadRequestError("应用名称重复")
	}

	// 判断是否重复
	var predicates1 []predicate.UpgradeApk
	predicates1 = append(predicates1, upgradeapk.PackageName(*req.PackageName))
	predicates1 = append(predicates1, upgradeapk.IsDelEQ(0))
	predicates1 = append(predicates1, upgradeapk.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count1, err := l.svcCtx.DB.UpgradeApk.Query().Where(predicates1...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count1 > 0 {
		return http_error.NewCodeBadRequestError("应用包名名称重复")
	}

	return nil
}
