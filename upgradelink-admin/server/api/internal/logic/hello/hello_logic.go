package hello

import (
	"context"
	"fmt"

	"upgradelink-admin/server/api/internal/svc"
	"upgradelink-admin/server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HelloLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HelloLogic {
	return &HelloLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HelloLogic) Hello(req *types.GetHelloReq) (resp *types.GetHelloResp, err error) {
	fmt.Println("hello")
	resp = &types.GetHelloResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  "hello world",
		},
	}
	l.Logger.Info("hello world")
	return resp, nil
}
