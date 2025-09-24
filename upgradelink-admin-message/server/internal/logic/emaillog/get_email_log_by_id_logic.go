package emaillog

import (
	"context"

	"upgradelink-admin-message/server/internal/svc"
	"upgradelink-admin-message/server/internal/utils/dberrorhandler"
	"upgradelink-admin-message/server/types/mcms"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetEmailLogByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEmailLogByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmailLogByIdLogic {
	return &GetEmailLogByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetEmailLogByIdLogic) GetEmailLogById(in *mcms.UUIDReq) (*mcms.EmailLogInfo, error) {
	result, err := l.svcCtx.DB.EmailLog.Get(l.ctx, uuidx.ParseUUIDString(in.Id))
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &mcms.EmailLogInfo{
		Id:         pointy.GetPointer(result.ID.String()),
		CreatedAt:  pointy.GetPointer(result.CreatedAt.UnixMilli()),
		UpdatedAt:  pointy.GetPointer(result.UpdatedAt.UnixMilli()),
		Target:     &result.Target,
		Subject:    &result.Subject,
		Content:    &result.Content,
		SendStatus: pointy.GetPointer(uint32(result.SendStatus)),
	}, nil
}
