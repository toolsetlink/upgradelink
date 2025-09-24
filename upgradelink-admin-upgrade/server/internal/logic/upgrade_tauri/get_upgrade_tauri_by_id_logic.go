package upgrade_tauri

import (
	"context"

	"upgradelink-admin-upgrade/server/internal/logic/base"
	"upgradelink-admin-upgrade/server/internal/svc"
	"upgradelink-admin-upgrade/server/internal/types"
	"upgradelink-admin-upgrade/server/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgradeTauriByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUpgradeTauriByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgradeTauriByIdLogic {
	return &GetUpgradeTauriByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUpgradeTauriByIdLogic) GetUpgradeTauriById(req *types.IDReq) (*types.UpgradeTauriInfoResp, error) {
	// 获取公司 id
	companyID, err := base.GetCompanyId(l.ctx, l.svcCtx, l.Logger)
	if err != nil {
		return nil, err
	}

	data, err := l.svcCtx.DB.UpgradeTauri.Get(l.ctx, req.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}
	// 效验下请求数据是否属于当前操作者
	if data.CompanyID != companyID {
		return nil, errorx.NewCodeInvalidArgumentError(i18n.TargetNotFound)
	}

	return &types.UpgradeTauriInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  errormsg.Success,
		},
		Data: types.UpgradeTauriInfo{
			Id:          &data.ID,
			Key:         &data.Key,
			Name:        &data.Name,
			Description: &data.Description,
			GithubUrl:   &data.GithubURL,
			IsDel:       &data.IsDel,
			CreateAt:    pointy.GetUnixMilliPointer(data.CreateAt.UnixMilli()),
			UpdateAt:    pointy.GetUnixMilliPointer(data.UpdateAt.UnixMilli()),
		},
	}, nil
}
