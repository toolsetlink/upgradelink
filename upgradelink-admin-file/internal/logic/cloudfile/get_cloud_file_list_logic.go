package cloudfile

import (
	"context"

	"upgradelink-admin-file/ent"
	"upgradelink-admin-file/ent/cloudfiletag"
	"upgradelink-admin-file/ent/storageprovider"

	"upgradelink-admin-file/ent/cloudfile"
	"upgradelink-admin-file/ent/predicate"
	"upgradelink-admin-file/internal/svc"
	"upgradelink-admin-file/internal/types"
	"upgradelink-admin-file/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCloudFileListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCloudFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCloudFileListLogic {
	return &GetCloudFileListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCloudFileListLogic) GetCloudFileList(req *types.CloudFileListReq) (*types.CloudFileListResp, error) {
	var predicates []predicate.CloudFile
	if req.Name != nil {
		predicates = append(predicates, cloudfile.NameContains(*req.Name))
	}
	if req.ProviderId != nil {
		predicates = append(predicates, cloudfile.HasStorageProvidersWith(storageprovider.IDEQ(*req.ProviderId)))
	}
	if req.TagIds != nil {
		predicates = append(predicates, cloudfile.HasTagsWith(cloudfiletag.IDIn(req.TagIds...)))
	}
	if req.FileType != nil && *req.FileType != 0 {
		predicates = append(predicates, cloudfile.FileTypeEQ(*req.FileType))
	}
	data, err := l.svcCtx.DB.CloudFile.Query().Where(predicates...).WithStorageProviders().WithTags().
		Page(l.ctx, req.Page, req.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	resp := &types.CloudFileListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.PageDetails.Total

	for _, v := range data.List {
		resp.Data.Data = append(resp.Data.Data,
			types.CloudFileInfo{
				BaseUUIDInfo: types.BaseUUIDInfo{
					Id:        pointy.GetPointer(v.ID.String()),
					CreatedAt: pointy.GetPointer(v.CreatedAt.UnixMilli()),
					UpdatedAt: pointy.GetPointer(v.UpdatedAt.UnixMilli()),
				},
				State:       &v.State,
				Name:        &v.Name,
				Url:         &v.URL,
				Size:        &v.Size,
				FileType:    &v.FileType,
				UserId:      &v.UserID,
				ProviderId:  &v.Edges.StorageProviders.ID,
				TagIds:      l.getFileTagIds(v.Edges.Tags),
				RelativeSrc: &v.URL,
			})
	}

	return resp, nil
}

func (l *GetCloudFileListLogic) getFileTagIds(tags []*ent.CloudFileTag) (result []uint64) {
	for _, v := range tags {
		result = append(result, v.ID)
	}
	return result
}
