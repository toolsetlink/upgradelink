package upgrade_file

import (
	"context"
	"crypto/rand"
	"encoding/base64"

	"upgradelink-admin-upgrade/server/ent/predicate"
	"upgradelink-admin-upgrade/server/ent/upgradefile"
	"upgradelink-admin-upgrade/server/internal/logic/base"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
	"upgradelink-admin-upgrade/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUpgradeFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger

	companyID int
}

func NewCreateUpgradeFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUpgradeFileLogic {
	return &CreateUpgradeFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUpgradeFileLogic) CreateUpgradeFile(req *types.UpgradeFileInfo) (*types.BaseMsgResp, error) {

	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}
	l.companyID = companyID

	// 效验请求参数数据
	err = l.CheckCreateUpgradeFile(req)
	if err != nil {
		return nil, errorx.NewCodeInvalidArgumentError(err.Error())
	}

	// 生成Access Key (16字节 -> 24字符Base64)
	fileBytes := make([]byte, 16)
	_, _ = rand.Read(fileBytes)
	fileKey := base64.RawURLEncoding.EncodeToString(fileBytes)

	isDel := int32(0)
	_, err = l.svcCtx.DB.UpgradeFile.Create().
		SetNotNilCompanyID(&companyID).
		SetNotNilKey(&fileKey).
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

func (l *CreateUpgradeFileLogic) CheckCreateUpgradeFile(req *types.UpgradeFileInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeFile
	predicates = append(predicates, upgradefile.Name(*req.Name))
	predicates = append(predicates, upgradefile.IsDelEQ(0))
	predicates = append(predicates, upgradefile.CompanyIDEQ(l.companyID))
	count, err := l.svcCtx.DB.UpgradeFile.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return errorx.NewCodeInvalidArgumentError("应用名称重复")
	}

	return nil
}
