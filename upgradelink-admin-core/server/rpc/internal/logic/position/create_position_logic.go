package position

import (
	"context"

	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"upgradelink-admin-core/server/rpc/internal/svc"
	"upgradelink-admin-core/server/rpc/internal/utils/dberrorhandler"
	"upgradelink-admin-core/server/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-common/i18n"
)

type CreatePositionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePositionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePositionLogic {
	return &CreatePositionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatePositionLogic) CreatePosition(in *core.PositionInfo) (*core.BaseIDResp, error) {
	result, err := l.svcCtx.DB.Position.Create().
		SetNotNilStatus(pointy.GetStatusPointer(in.Status)).
		SetNotNilSort(in.Sort).
		SetNotNilName(in.Name).
		SetNotNilCode(in.Code).
		SetNotNilRemark(in.Remark).
		Save(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil
}
