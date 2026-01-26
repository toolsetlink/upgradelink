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

type CreateOrUpdateApiAuthorityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrUpdateApiAuthorityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateApiAuthorityLogic {
	return &CreateOrUpdateApiAuthorityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrUpdateApiAuthorityLogic) CreateOrUpdateApiAuthority(req *types.CreateOrUpdateApiAuthorityReq) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.DB.Role.Get(l.ctx, req.RoleId)
	if err != nil {
		return nil, db_error.DefaultEntError(l.Logger, err, req)
	}

	// clear old policies
	var oldPolicies [][]string
	oldPolicies, err = l.svcCtx.Casbin.GetFilteredPolicy(0, data.Code)
	if err != nil {
		logx.Error("failed to get old Casbin policy", logx.Field("detail", err))
		return nil, http_error.NewCodeBadRequestError("failed to get old Casbin policy")
	}

	if len(oldPolicies) != 0 {
		removeResult, err := l.svcCtx.Casbin.RemoveFilteredPolicy(0, data.Code)
		if err != nil {
			l.Logger.Errorw("failed to remove roles policy", logx.Field("roleCode", data.Code), logx.Field("detail", err.Error()))
			return nil, http_error.NewCodeBadRequestError("failed to remove roles policy")

		}
		if !removeResult {
			return nil, http_error.NewCodeBadRequestError("casbin.removeFailed")
		}
	}
	// add new policies
	var policies [][]string
	for _, v := range req.Data {
		policies = append(policies, []string{data.Code, v.Path, v.Method})
	}

	addResult, err := l.svcCtx.Casbin.AddPolicies(policies)
	if err != nil {
		return nil, http_error.NewCodeBadRequestError("casbin.addFailed")
	}

	if addResult {
		return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.UpdateSuccess)}, nil
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.UpdateFailed)}, nil
}
