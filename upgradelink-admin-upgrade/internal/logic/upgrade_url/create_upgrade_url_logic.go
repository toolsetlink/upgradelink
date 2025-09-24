package upgrade_url

import (
	"context"
	"crypto/rand"
	"encoding/base64"

	"upgradelink-admin-upgrade/ent/predicate"
	"upgradelink-admin-upgrade/ent/upgradeurl"
	"upgradelink-admin-upgrade/internal/logic/base"
	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
	"upgradelink-admin-upgrade/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUpgradeUrlLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger

	companyID int
}

func NewCreateUpgradeUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUpgradeUrlLogic {
	return &CreateUpgradeUrlLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUpgradeUrlLogic) CreateUpgradeUrl(req *types.UpgradeUrlInfo) (*types.BaseMsgResp, error) {

	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}
	l.companyID = companyID

	// 效验请求参数数据
	err = l.CheckCreateUpgradeUrl(req)
	if err != nil {
		return nil, errorx.NewCodeInvalidArgumentError(err.Error())
	}

	// 生成Access Key (16字节 -> 24字符Base64)
	urlBytes := make([]byte, 16)
	_, _ = rand.Read(urlBytes)
	urlKey := base64.RawURLEncoding.EncodeToString(urlBytes)

	isDel := int32(0)
	_, err = l.svcCtx.DB.UpgradeUrl.Create().
		SetNotNilCompanyID(&companyID).
		SetNotNilKey(&urlKey).
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

func (l *CreateUpgradeUrlLogic) CheckCreateUpgradeUrl(req *types.UpgradeUrlInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeUrl
	predicates = append(predicates, upgradeurl.Name(*req.Name))
	predicates = append(predicates, upgradeurl.IsDelEQ(0))
	predicates = append(predicates, upgradeurl.CompanyIDEQ(l.companyID))
	count, err := l.svcCtx.DB.UpgradeUrl.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return errorx.NewCodeInvalidArgumentError("应用名称重复")
	}

	return nil
}
