package authority

import (
	"context"
	"upgradelink-admin/server/api/internal/common/db_error"
	"upgradelink-admin/server/api/internal/common/http_error"
	"upgradelink-admin/server/api/internal/common/i18n"

	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetApiAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetApiAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApiAuthorityLogic {
	return &GetApiAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApiAuthorityLogic) GetApiAuthority(req *types.IDReq) (resp *types.ApiAuthorityListResp, err error) {
	roleData, err := l.svcCtx.DB.Role.Get(l.ctx, req.Id)
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	data, err := l.svcCtx.Casbin.GetFilteredPolicy(0, roleData.Code)
	if err != nil {
		logx.Error("failed to get old Casbin policy", logx.Field("detail", err))
		return nil, http_error.NewCodeBadRequestError("failed to get old Casbin policy")
	}

	resp = &types.ApiAuthorityListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = uint64(len(data))
	for _, v := range data {
		resp.Data.Data = append(resp.Data.Data, types.ApiAuthorityInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return resp, nil
}
