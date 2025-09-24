package file

import (
	"context"
	"os"

	"github.com/suyuan32/simple-admin-common/utils/uuidx"

	file2 "upgradelink-admin-file/server/ent/file"

	"github.com/suyuan32/simple-admin-common/i18n"

	"upgradelink-admin-file/server/internal/utils/dberrorhandler"
	"upgradelink-admin-file/server/internal/utils/entx"

	"upgradelink-admin-file/server/ent"
	"upgradelink-admin-file/server/internal/svc"
	"upgradelink-admin-file/server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewDeleteFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFileLogic {
	return &DeleteFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteFileLogic) DeleteFile(req *types.UUIDsReq) (resp *types.BaseMsgResp, err error) {
	err = entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		files, err := tx.File.Query().Where(file2.IDIn(uuidx.ParseUUIDSlice(req.Ids)...)).All(l.ctx)

		if err != nil {
			return err
		}

		_, err = tx.File.Delete().Where(file2.IDIn(uuidx.ParseUUIDSlice(req.Ids)...)).Exec(l.ctx)

		if err != nil {
			return err
		}

		for _, v := range files {
			if v.Status == 1 {
				err = os.RemoveAll(l.svcCtx.Config.UploadConf.PublicStorePath + v.Path)
				if err != nil {
					return err
				}
			} else {
				err = os.RemoveAll(l.svcCtx.Config.UploadConf.PrivateStorePath + v.Path)
				if err != nil {
					return err
				}
			}
		}

		return nil
	})

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{
		Code: 0,
		Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.DeleteSuccess),
	}, nil
}
