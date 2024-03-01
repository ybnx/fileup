package logic

import (
	"context"
	"errors"
	"fileup/rpc/user/user"

	"fileup/api/user/internal/svc"
	"fileup/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	// todo: add your logic here and delete this line

	if req.Name == "" || req.Password == "" || req.Email == "" || req.VeCode == "" {
		return nil, errors.New("arg cant be null")
	}

	cnt, err := l.svcCtx.Rpc.Register(l.ctx, &user.RegisterRequest{
		Name:     req.Name,
		Password: req.Password,
		Email:    req.Email,
		VeCode:   req.VeCode,
	})
	if err != nil {
		return nil, err
	}
	return &types.RegisterResponse{
		UserId:       cnt.UserId,
		AccessToken:  cnt.AccessToken,
		RefreshToken: cnt.RefreshToken,
	}, nil
}
