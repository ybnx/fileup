package logic

import (
	"context"
	"fileup/rpc/user/user"

	"fileup/api/user/internal/svc"
	"fileup/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PasswdLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPasswdLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PasswdLoginLogic {
	return &PasswdLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PasswdLoginLogic) PasswdLogin(req *types.PasswdLoginRequest) (resp *types.PasswdLoginResponse, err error) {
	// todo: add your logic here and delete this line
	cnt, err := l.svcCtx.Rpc.PasswdLogin(l.ctx, &user.PasswdLoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &types.PasswdLoginResponse{
		UserId:       cnt.UserId,
		AccessToken:  cnt.AccessToken,
		RefreshToken: cnt.RefreshToken,
	}, nil
}
