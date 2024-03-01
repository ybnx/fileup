package logic

import (
	"context"
	"fileup/rpc/user/user"

	"fileup/api/user/internal/svc"
	"fileup/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailLoginLogic {
	return &EmailLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmailLoginLogic) EmailLogin(req *types.EmailLoginRequest) (resp *types.EmailLoginResponse, err error) {
	// todo: add your logic here and delete this line
	cnt, err := l.svcCtx.Rpc.EmailLogin(l.ctx, &user.EmailLoginRequest{
		Email:  req.Email,
		VeCode: req.VeCode,
	})
	if err != nil {
		return nil, err
	}
	return &types.EmailLoginResponse{
		UserId:       cnt.UserId,
		AccessToken:  cnt.AccessToken,
		RefreshToken: cnt.RefreshToken,
	}, nil
}
