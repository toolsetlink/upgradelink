package upgrade_url

import (
	"context"

	"upgradelink-admin-upgrade/server/ent/predicate"
	"upgradelink-admin-upgrade/server/ent/upgradeurl"
	"upgradelink-admin-upgrade/server/internal/logic/base"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
	"upgradelink-admin-upgrade/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeUrlListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeUrlListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeUrlListLogic {
	return &GetUpgradeUrlListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeUrlListLogic) GetUpgradeUrlList(req *types.UpgradeUrlListReq) (*types.UpgradeUrlListResp, error) {
	var predicates []predicate.UpgradeUrl

	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}
	predicates = append(predicates, upgradeurl.CompanyIDEQ(companyID))

	// 删除状态
	predicates = append(predicates, upgradeurl.IsDelEQ(0))

	if req.Key != nil {
		predicates = append(predicates, upgradeurl.KeyContains(*req.Key))
	}
	if req.Name != nil {
		predicates = append(predicates, upgradeurl.NameContains(*req.Name))
	}
	data, err := l.svcCtx.DB.UpgradeUrl.Query().Where(predicates...).Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.UpgradeUrlListResp{}
	resp.Msg = errormsg.Success
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {
		resp.Data.Data = append(resp.Data.Data,
			types.UpgradeUrlInfo{
				Id:          &v.ID,
				Key:         &v.Key,
				Name:        &v.Name,
				Description: &v.Description,
				IsDel:       &v.IsDel,
				CreateAt:    pointy.GetUnixMilliPointer(v.CreateAt.UnixMilli()),
				UpdateAt:    pointy.GetUnixMilliPointer(v.UpdateAt.UnixMilli()),
			})
	}

	return resp, nil
}
