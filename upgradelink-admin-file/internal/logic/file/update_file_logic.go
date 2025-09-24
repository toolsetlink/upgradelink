package file

import (
	"context"

	"github.com/suyuan32/simple-admin-common/utils/uuidx"

	"github.com/suyuan32/simple-admin-common/i18n"

	"upgradelink-admin-file/internal/utils/dberrorhandler"

	"upgradelink-admin-file/internal/svc"
	"upgradelink-admin-file/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewUpdateFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFileLogic {
	return &UpdateFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateFileLogic) UpdateFile(req *types.UpdateFileReq) (resp *types.BaseMsgResp, err error) {
	query := l.svcCtx.DB.File.UpdateOneID(uuidx.ParseUUIDString(req.ID)).SetNotNilName(req.Name)

	if req.FileTagIds != nil {
		query.AddTagIDs(req.FileTagIds...)
	} else {
		query.ClearTags()
	}

	err = query.Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{
		Code: 0,
		Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
	}, nil
}
