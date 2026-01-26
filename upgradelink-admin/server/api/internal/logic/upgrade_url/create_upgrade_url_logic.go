package upgrade_url

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
	"upgradelink-admin/server/api/internal/ent/upgradeurl"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUpgradeUrlLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUpgradeUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUpgradeUrlLogic {
	return &CreateUpgradeUrlLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUpgradeUrlLogic) CreateUpgradeUrl(req *types.UpgradeUrlInfo) (*types.BaseMsgResp, error) {

	// 效验请求参数数据
	err := l.CheckCreateUpgradeUrl(req)
	if err != nil {
		return nil, err
	}

	// 生成Access Key (16字节 -> 24字符Base64)
	urlBytes := make([]byte, 16)
	_, _ = rand.Read(urlBytes)
	urlKey := base64.RawURLEncoding.EncodeToString(urlBytes)

	isDel := int32(0)
	_, err = l.svcCtx.DB.UpgradeUrl.Create().
		SetNotNilCompanyID(companyctx.GetCompanyIDPointerFromCtx(l.ctx)).
		SetNotNilKey(&urlKey).
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

func (l *CreateUpgradeUrlLogic) CheckCreateUpgradeUrl(req *types.UpgradeUrlInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeUrl
	predicates = append(predicates, upgradeurl.Name(*req.Name))
	predicates = append(predicates, upgradeurl.IsDelEQ(0))
	predicates = append(predicates, upgradeurl.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count, err := l.svcCtx.DB.UpgradeUrl.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return http_error.NewCodeBadRequestError("应用名称重复")
	}

	return nil
}
