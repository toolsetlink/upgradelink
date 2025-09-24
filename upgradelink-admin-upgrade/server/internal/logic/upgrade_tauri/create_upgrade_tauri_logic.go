package upgrade_tauri

import (
	"context"
	"crypto/rand"
	"encoding/base64"

	"upgradelink-admin-upgrade/server/ent/predicate"
	"upgradelink-admin-upgrade/server/ent/upgradetauri"
	"upgradelink-admin-upgrade/server/internal/logic/base"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
	"upgradelink-admin-upgrade/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUpgradeTauriLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger

	companyID int
}

func NewCreateUpgradeTauriLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUpgradeTauriLogic {
	return &CreateUpgradeTauriLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUpgradeTauriLogic) CreateUpgradeTauri(req *types.UpgradeTauriInfo) (*types.BaseMsgResp, error) {

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
	fileBytes := make([]byte, 16)
	_, _ = rand.Read(fileBytes)
	fileKey := base64.RawURLEncoding.EncodeToString(fileBytes)

	isDel := int32(0)
	_, err = l.svcCtx.DB.UpgradeTauri.Create().
		SetNotNilCompanyID(&companyID).
		SetNotNilKey(&fileKey).
		SetNotNilName(req.Name).
		SetNotNilDescription(req.Description).
		SetNotNilGithubURL(req.GithubUrl).
		SetNotNilIsDel(&isDel).
		SetNotNilCreateAt(pointy.GetTimeMilliPointer(req.CreateAt)).
		SetNotNilUpdateAt(pointy.GetTimeMilliPointer(req.UpdateAt)).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: errormsg.CreateSuccess}, nil
}

func (l *CreateUpgradeTauriLogic) CheckCreateUpgrade(req *types.UpgradeTauriInfo) error {
	// 判断是否重复
	var predicates []predicate.UpgradeTauri
	predicates = append(predicates, upgradetauri.Name(*req.Name))
	predicates = append(predicates, upgradetauri.IsDelEQ(0))
	predicates = append(predicates, upgradetauri.CompanyIDEQ(l.companyID))
	count, err := l.svcCtx.DB.UpgradeTauri.Query().Where(predicates...).Count(l.ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return errorx.NewCodeInvalidArgumentError("应用名称重复")
	}

	return nil
}
