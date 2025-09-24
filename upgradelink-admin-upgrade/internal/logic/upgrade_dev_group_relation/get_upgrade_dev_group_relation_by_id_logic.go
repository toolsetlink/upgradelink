package upgrade_dev_group_relation

import (
	"context"

	"upgradelink-admin-upgrade/internal/svc"
	"upgradelink-admin-upgrade/internal/types"
	"upgradelink-admin-upgrade/internal/utils/dberrorhandler"

    "github.com/suyuan32/simple-admin-common/msg/errormsg"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeDevGroupRelationByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeDevGroupRelationByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeDevGroupRelationByIdLogic {
	return &GetUpgradeDevGroupRelationByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeDevGroupRelationByIdLogic) GetUpgradeDevGroupRelationById(req *types.IDReq) (*types.UpgradeDevGroupRelationInfoResp, error) {
	data, err := l.svcCtx.DB.UpgradeDevGroupRelation.Get(l.ctx, req.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.UpgradeDevGroupRelationInfoResp{
	    BaseDataInfo: types.BaseDataInfo{
            Code: 0,
            Msg:  errormsg.Success,
        },
        Data: types.UpgradeDevGroupRelationInfo{
			Id:  &data.ID, 
			DevId:	&data.DevID,
			DevGroupId:	&data.DevGroupID,
        },
	}, nil
}

