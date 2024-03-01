package logic

import (
	"context"
	"errors"
	"fileup/api/user/internal/svc"
	"fileup/api/user/internal/types"
	"fileup/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCodeLogic {
	return &SendCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendCodeLogic) SendCode(req *types.SendCodeRequest) (resp *types.SendCodeResponse, err error) {
	// todo: add your logic here and delete this line

	if req.Email == "" {
		return nil, errors.New("email cannot empty")
	}

	cnt, err := l.svcCtx.Rpc.SendCode(l.ctx, &user.SendCodeRequest{
		Email: req.Email,
	})
	if err != nil {
		return nil, err
	}
	return &types.SendCodeResponse{
		Message: cnt.Message,
	}, nil
}
