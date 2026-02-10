package cloudfile

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/i18n"
	"upgradelink-admin/server/api/internal/common/utils/filex"
	"upgradelink-admin/server/api/internal/common/utils/pointy"
	"upgradelink-admin/server/api/internal/ent/fmscloudfile"
	"upgradelink-admin/server/api/internal/ent/predicate"
	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCloudFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCloudFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCloudFileListLogic {
	return &GetCloudFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCloudFileListLogic) GetCloudFileList(req *types.CloudFileListReq) (resp *types.CloudFileListResp, err error) {

	var predicates []predicate.FmsCloudFile
	if req.Name != nil {
		predicates = append(predicates, fmscloudfile.NameContains(*req.Name))
	}

	if req.FileType != nil && *req.FileType != 0 {
		predicates = append(predicates, fmscloudfile.FileTypeEQ(*req.FileType))
	}
	data, err := l.svcCtx.DB.FmsCloudFile.Query().Where(predicates...).
		Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	resp = &types.CloudFileListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {
		fileType := filex.ConvertUint8ToFileType(v.FileType)

		resp.Data.Data = append(resp.Data.Data,
			types.CloudFileInfo{
				BaseUUIDInfo: types.BaseUUIDInfo{
					Id:        pointy.GetPointer(v.ID),
					CreatedAt: pointy.GetPointer(v.CreatedAt.UnixMilli()),
					UpdatedAt: pointy.GetPointer(v.UpdatedAt.UnixMilli()),
				},
				State:    &v.State,
				Name:     &v.Name,
				Url:      &v.URL,
				Size:     &v.Size,
				Md5:      &v.Md5,
				FileType: &fileType,
				UserId:   &v.UserID,
			})
	}

	return resp, nil
}
