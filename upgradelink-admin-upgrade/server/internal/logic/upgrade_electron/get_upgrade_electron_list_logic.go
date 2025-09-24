package upgrade_electron

import (
	"context"
	"time"

	"upgradelink-admin-upgrade/server/ent/predicate"
	"upgradelink-admin-upgrade/server/ent/upgradeelectron"
	"upgradelink-admin-upgrade/server/internal/logic/base"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
	"upgradelink-admin-upgrade/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeElectronListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeElectronListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeElectronListLogic {
	return &GetUpgradeElectronListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeElectronListLogic) GetUpgradeElectronList(req *types.UpgradeElectronListReq) (*types.UpgradeElectronListResp, error) {
	var predicates []predicate.UpgradeElectron

	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}
	predicates = append(predicates, upgradeelectron.CompanyIDEQ(companyID))

	// 删除状态
	predicates = append(predicates, upgradeelectron.IsDelEQ(0))
	if req.Key != nil {
		predicates = append(predicates, upgradeelectron.KeyContains(*req.Key))
	}
	if req.Name != nil {
		predicates = append(predicates, upgradeelectron.NameContains(*req.Name))
	}
	if req.IsDel != nil {
		predicates = append(predicates, upgradeelectron.IsDelEQ(*req.IsDel))
	}
	if req.CreateAt != nil {
		predicates = append(predicates, upgradeelectron.CreateAtGTE(time.UnixMilli(*req.CreateAt)))
	}
	if req.UpdateAt != nil {
		predicates = append(predicates, upgradeelectron.UpdateAtGTE(time.UnixMilli(*req.UpdateAt)))
	}
	data, err := l.svcCtx.DB.UpgradeElectron.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.UpgradeElectronListResp{}
	resp.Msg = errormsg.Success
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {
		resp.Data.Data = append(resp.Data.Data,
			types.UpgradeElectronInfo{
				Id:          &v.ID,
				Key:         &v.Key,
				Name:        &v.Name,
				Description: &v.Description,
				GithubUrl:   &v.GithubURL,
				IsDel:       &v.IsDel,
				CreateAt:    pointy.GetUnixMilliPointer(v.CreateAt.UnixMilli()),
				UpdateAt:    pointy.GetUnixMilliPointer(v.UpdateAt.UnixMilli()),
			})
	}

	return resp, nil
}
