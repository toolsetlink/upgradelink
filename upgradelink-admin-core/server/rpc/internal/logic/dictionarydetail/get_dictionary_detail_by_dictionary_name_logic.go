package dictionarydetail

import (
	"context"

	"github.com/suyuan32/simple-admin-common/enum/common"

	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"upgradelink-admin-core/server/rpc/ent"
	"upgradelink-admin-core/server/rpc/ent/dictionary"
	"upgradelink-admin-core/server/rpc/ent/dictionarydetail"
	"upgradelink-admin-core/server/rpc/internal/svc"
	"upgradelink-admin-core/server/rpc/internal/utils/dberrorhandler"
	"upgradelink-admin-core/server/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDictionaryDetailByDictionaryNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDictionaryDetailByDictionaryNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDictionaryDetailByDictionaryNameLogic {
	return &GetDictionaryDetailByDictionaryNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDictionaryDetailByDictionaryNameLogic) GetDictionaryDetailByDictionaryName(in *core.BaseMsg) (*core.DictionaryDetailListResp, error) {
	dictionaryData, err := l.svcCtx.DB.Dictionary.Query().Where(dictionary.NameEQ(in.Msg)).First(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	result, err := l.svcCtx.DB.DictionaryDetail.Query().Where(dictionarydetail.DictionaryID(dictionaryData.ID), dictionarydetail.StatusEQ(common.StatusNormal)).
		Page(l.ctx, 1, 10000, func(pager *ent.DictionaryDetailPager) {
			pager.Order = ent.Asc(dictionarydetail.FieldSort)
		})
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.DictionaryDetailListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.DictionaryDetailInfo{
			Id:           &v.ID,
			CreatedAt:    pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:    pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Status:       pointy.GetPointer(uint32(v.Status)),
			Title:        &v.Title,
			Key:          &v.Key,
			Value:        &v.Value,
			DictionaryId: &v.DictionaryID,
			Sort:         &v.Sort,
		})
	}

	return resp, nil
}
