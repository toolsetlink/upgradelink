package upgrade_configuration

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
	"upgradelink-admin/server/api/internal/ent/upgradeconfiguration"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUpgradeConfigurationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUpgradeConfigurationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUpgradeConfigurationLogic {
	return &CreateUpgradeConfigurationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUpgradeConfigurationLogic) CreateUpgradeConfiguration(req *types.UpgradeConfigurationInfo) (*types.BaseMsgResp, error) {
	// 效验请求参数数据
	err := l.CheckCreateUpgradeConfiguration(req)
	if err != nil {
		return nil, err
	}

	// 生成Access Key (16字节 -> 24字符Base64)
	configurationBytes := make([]byte, 16)
	_, _ = rand.Read(configurationBytes)
	configurationKey := base64.RawURLEncoding.EncodeToString(configurationBytes)

	isDel := int32(0)
	_, err = l.svcCtx.DB.UpgradeConfiguration.Create().
		SetNotNilCompanyID(companyctx.GetCompanyIDPointerFromCtx(l.ctx)).
		SetNotNilKey(&configurationKey).
		SetNotNilName(req.Name).
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

func (l *CreateUpgradeConfigurationLogic) CheckCreateUpgradeConfiguration(req *types.UpgradeConfigurationInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeConfiguration
	predicates = append(predicates, upgradeconfiguration.Name(*req.Name))
	predicates = append(predicates, upgradeconfiguration.IsDelEQ(0))
	predicates = append(predicates, upgradeconfiguration.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count, err := l.svcCtx.DB.UpgradeConfiguration.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return http_error.NewCodeBadRequestError("配置名称重复")
	}

	return nil
}
