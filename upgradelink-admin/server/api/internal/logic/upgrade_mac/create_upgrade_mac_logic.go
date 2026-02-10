package upgrade_mac

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/enum"
	"upgradelink-admin/server/api/internal/common/http_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/jwtctx/companyctx"

	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgrademac"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUpgradeMacLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUpgradeMacLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUpgradeMacLogic {
	return &CreateUpgradeMacLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUpgradeMacLogic) CreateUpgradeMac(req *types.UpgradeMacInfo) (*types.BaseMsgResp, error) {

	// 效验请求参数数据
	err := l.CheckCreateUpgrade(req)
	if err != nil {
		return nil, err
	}

	// 生成Access Key (16字节 -> 24字符Base64)
	macBytes := make([]byte, 16)
	_, _ = rand.Read(macBytes)
	macKey := base64.RawURLEncoding.EncodeToString(macBytes)

	intDelFalse := enum.IsDelFalse
	_, err = l.svcCtx.DB.UpgradeMac.Create().
		SetNotNilCompanyID(companyctx.GetCompanyIDPointerFromCtx(l.ctx)).
		SetNotNilKey(&macKey).
		SetNotNilName(req.Name).
		SetNotNilPackageName(req.PackageName).
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

func (l *CreateUpgradeMacLogic) CheckCreateUpgrade(req *types.UpgradeMacInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeMac
	predicates = append(predicates, upgrademac.Name(*req.Name))
	predicates = append(predicates, upgrademac.IsDelEQ(0))
	predicates = append(predicates, upgrademac.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count, err := l.svcCtx.DB.UpgradeMac.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return http_error.NewCodeBadRequestError(l.svcCtx.Trans.Trans(l.ctx, i18n.AppNameDuplicate))
	}

	// 判断是否重复
	var predicates1 []predicate.UpgradeMac
	predicates1 = append(predicates1, upgrademac.PackageName(*req.PackageName))
	predicates1 = append(predicates1, upgrademac.IsDelEQ(0))
	predicates1 = append(predicates1, upgrademac.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count1, err := l.svcCtx.DB.UpgradeMac.Query().Where(predicates1...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count1 > 0 {
		return http_error.NewCodeBadRequestError(l.svcCtx.Trans.Trans(l.ctx, i18n.AppPackageNameDuplicate))
	}

	return nil
}
