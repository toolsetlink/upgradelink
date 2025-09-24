package file

import (
	"context"
	"os"
	"path"

	"github.com/suyuan32/simple-admin-common/utils/uuidx"

	"github.com/suyuan32/simple-admin-common/i18n"

	"upgradelink-admin-file/internal/utils/dberrorhandler"
	"upgradelink-admin-file/internal/utils/entx"

	"upgradelink-admin-file/ent"
	"upgradelink-admin-file/internal/svc"
	"upgradelink-admin-file/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePublicStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangePublicStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePublicStatusLogic {
	return &ChangePublicStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangePublicStatusLogic) ChangePublicStatus(req *types.StatusCodeReq) (resp *types.BaseMsgResp, err error) {
	err = entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		file, err := tx.File.UpdateOneID(uuidx.ParseUUIDString(req.Id)).SetStatus(uint8(req.Status)).Save(l.ctx)

		if err != nil {
			return err
		}

		if req.Status == 1 {
			err = os.Rename(path.Join(l.svcCtx.Config.UploadConf.PrivateStorePath, file.Path),
				path.Join(l.svcCtx.Config.UploadConf.PublicStorePath, file.Path))
			if err != nil {
				logx.Errorw("fail to change the path of file", logx.Field("detail", err.Error()))
				return err
			}
		} else {
			err = os.Rename(path.Join(l.svcCtx.Config.UploadConf.PublicStorePath, file.Path),
				path.Join(l.svcCtx.Config.UploadConf.PrivateStorePath, file.Path))
			if err != nil {
				logx.Errorw("fail to change the path of file", logx.Field("detail", err.Error()))
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{
		Code: 0,
		Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
	}, nil
}
