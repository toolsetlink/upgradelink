package upgrade_dev

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/ent/upgradedevgrouprelation"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"upgradelink-admin/server/api/internal/common/utils/pointy"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeDevByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeDevByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeDevByIdLogic {
	return &GetUpgradeDevByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeDevByIdLogic) GetUpgradeDevById(req *types.IDReq) (*types.UpgradeDevInfoResp, error) {

	// 获取数据
	data, err := l.svcCtx.DB.UpgradeDev.Get(l.ctx, int(req.Id))
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	// 查询出关联的设备组
	var predicates []predicate.UpgradeDevGroupRelation
	predicates = append(predicates, upgradedevgrouprelation.DevID(int(req.Id)))
	predicates = append(predicates, upgradedevgrouprelation.IsDelEQ(0))
	devGroupList, err := l.svcCtx.DB.UpgradeDevGroupRelation.Query().Where(predicates...).All(l.ctx)
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}
	var devGroupIds []uint64
	for _, group := range devGroupList {
		devGroupIds = append(devGroupIds, uint64(group.DevGroupID))
	}

	return &types.UpgradeDevInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.UpgradeDevInfo{
			Id:          &data.ID,
			Key:         &data.Key,
			IsDel:       &data.IsDel,
			DevGroupIds: devGroupIds,
			CreateAt:    pointy.GetUnixMilliPointer(data.CreateAt.UnixMilli()),
			UpdateAt:    pointy.GetUnixMilliPointer(data.UpdateAt.UnixMilli()),
		},
	}, nil
}
