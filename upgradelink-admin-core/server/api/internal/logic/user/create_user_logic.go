package user

import (
	"context"
	"crypto/rand"
	"encoding/base64"

	"upgradelink-admin-core/server/api/internal/svc"
	"upgradelink-admin-core/server/api/internal/types"
	"upgradelink-admin-core/server/rpc/types/core"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.UserInfo) (resp *types.BaseMsgResp, err error) {

	if req.Password == nil || *req.Password == "" {
		return nil, errorx.NewApiBadRequestError("password can not be empty")
	}

	data, err := l.svcCtx.CoreRpc.CreateUser(l.ctx,
		&core.UserInfo{
			Status:       req.Status,
			Username:     req.Username,
			Password:     req.Password,
			Nickname:     req.Nickname,
			Description:  req.Description,
			HomePath:     req.HomePath,
			RoleIds:      req.RoleIds,
			Mobile:       req.Mobile,
			Email:        req.Email,
			Avatar:       req.Avatar,
			DepartmentId: req.DepartmentId,
			PositionIds:  req.PositionIds,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}

// GenerateAPIKeys 生成安全随机的AccessKey和SecretKey
func GenerateAPIKeys() (accessKey string, secretKey string, err error) {
	// 生成Access Key (16字节 -> 24字符Base64)
	accessBytes := make([]byte, 16)
	if _, err := rand.Read(accessBytes); err != nil {
		return "", "", err
	}
	accessKey = base64.RawURLEncoding.EncodeToString(accessBytes)

	// 生成Secret Key (32字节 -> 43字符Base64)
	secretBytes := make([]byte, 32)
	if _, err := rand.Read(secretBytes); err != nil {
		return "", "", err
	}
	secretKey = base64.RawURLEncoding.EncodeToString(secretBytes)

	return accessKey, secretKey, nil
}
