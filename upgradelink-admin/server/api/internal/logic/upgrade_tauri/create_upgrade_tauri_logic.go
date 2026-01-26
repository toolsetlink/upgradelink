package upgrade_tauri

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
	"upgradelink-admin/server/api/internal/ent/upgradetauri"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUpgradeTauriLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUpgradeTauriLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUpgradeTauriLogic {
	return &CreateUpgradeTauriLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUpgradeTauriLogic) CreateUpgradeTauri(req *types.UpgradeTauriInfo) (*types.BaseMsgResp, error) {

	// 效验请求参数数据
	err := l.CheckCreateUpgrade(req)
	if err != nil {
		return nil, err
	}

	// 生成Access Key (16字节 -> 24字符Base64)
	fileBytes := make([]byte, 16)
	_, _ = rand.Read(fileBytes)
	fileKey := base64.RawURLEncoding.EncodeToString(fileBytes)

	isDel := int32(0)
	_, err = l.svcCtx.DB.UpgradeTauri.Create().
		SetNotNilCompanyID(companyctx.GetCompanyIDPointerFromCtx(l.ctx)).
		SetNotNilKey(&fileKey).
		SetNotNilName(req.Name).
		SetNotNilDescription(req.Description).
		SetNotNilGithubURL(req.GithubUrl).
		SetNotNilIsDel(&isDel).
		SetNotNilCreateAt(pointy.GetTimeMilliPointer(req.CreateAt)).
		SetNotNilUpdateAt(pointy.GetTimeMilliPointer(req.UpdateAt)).
		Save(l.ctx)

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.CreateSuccess)}, nil
}

func (l *CreateUpgradeTauriLogic) CheckCreateUpgrade(req *types.UpgradeTauriInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeTauri
	predicates = append(predicates, upgradetauri.Name(*req.Name))
	predicates = append(predicates, upgradetauri.IsDelEQ(0))
	predicates = append(predicates, upgradetauri.CompanyIDEQ(companyctx.GetCompanyIDFromCtx(l.ctx)))
	count, err := l.svcCtx.DB.UpgradeTauri.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return http_error.NewCodeBadRequestError("应用名称重复")
	}

	return nil
}
