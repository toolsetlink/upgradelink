package upgrade_configuration

import (
	"context"
	"crypto/rand"
	"encoding/base64"

	"upgradelink-admin-upgrade/ent/predicate"
	"upgradelink-admin-upgrade/ent/upgradeconfiguration"
	"upgradelink-admin-upgrade/internal/logic/base"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
	"upgradelink-admin-upgrade/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUpgradeConfigurationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger

	companyID int
}

func NewCreateUpgradeConfigurationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUpgradeConfigurationLogic {
	return &CreateUpgradeConfigurationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUpgradeConfigurationLogic) CreateUpgradeConfiguration(req *types.UpgradeConfigurationInfo) (*types.BaseMsgResp, error) {
	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}
	l.companyID = companyID

	// 效验请求参数数据
	err = l.CheckCreateUpgradeConfiguration(req)
	if err != nil {
		return nil, errorx.NewCodeInvalidArgumentError(err.Error())
	}

	// 生成Access Key (16字节 -> 24字符Base64)
	configurationBytes := make([]byte, 16)
	_, _ = rand.Read(configurationBytes)
	configurationKey := base64.RawURLEncoding.EncodeToString(configurationBytes)

	isDel := int32(0)
	_, err = l.svcCtx.DB.UpgradeConfiguration.Create().
		SetNotNilCompanyID(&companyID).
		SetNotNilKey(&configurationKey).
		SetNotNilName(req.Name).
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

func (l *CreateUpgradeConfigurationLogic) CheckCreateUpgradeConfiguration(req *types.UpgradeConfigurationInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeConfiguration
	predicates = append(predicates, upgradeconfiguration.Name(*req.Name))
	predicates = append(predicates, upgradeconfiguration.IsDelEQ(0))
	predicates = append(predicates, upgradeconfiguration.CompanyIDEQ(l.companyID))
	count, err := l.svcCtx.DB.UpgradeConfiguration.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return errorx.NewCodeInvalidArgumentError("配置名称重复")
	}

	return nil
}
