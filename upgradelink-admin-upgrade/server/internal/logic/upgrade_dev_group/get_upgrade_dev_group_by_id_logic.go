package upgrade_dev_group

import (
	"context"

	"upgradelink-admin-upgrade/server/internal/logic/base"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
	"upgradelink-admin-upgrade/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeDevGroupByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeDevGroupByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeDevGroupByIdLogic {
	return &GetUpgradeDevGroupByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeDevGroupByIdLogic) GetUpgradeDevGroupById(req *types.IDReq) (*types.UpgradeDevGroupInfoResp, error) {
	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}

	data, err := l.svcCtx.DB.UpgradeDevGroup.Get(l.ctx, req.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}
	// 效验下请求数据是否属于当前操作者
	if data.CompanyID != companyID {
		return nil, errorx.NewCodeInvalidArgumentError(i18n.TargetNotFound)
	}

	return &types.UpgradeDevGroupInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  errormsg.Success,
		},
		Data: types.UpgradeDevGroupInfo{
			Id:       &data.ID,
			Name:     &data.Name,
			IsDel:    &data.IsDel,
			CreateAt: pointy.GetUnixMilliPointer(data.CreateAt.UnixMilli()),
			UpdateAt: pointy.GetUnixMilliPointer(data.UpdateAt.UnixMilli()),
		},
	}, nil
}
