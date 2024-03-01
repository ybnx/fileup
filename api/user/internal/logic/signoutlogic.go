package logic

import (
	"context"
	"errors"
	"fileup/rpc/user/user"

	"fileup/api/user/internal/svc"
	"fileup/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignoutLogic {
	return &SignoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignoutLogic) Signout(req *types.SignoutRequest) (resp *types.SignoutResponse, err error) {
	// todo: add your logic here and delete this line

	if req.Email == "" {
		return nil, errors.New("email cannot empty")
	}

	cnt, err := l.svcCtx.Rpc.Signout(l.ctx, &user.SignoutRequest{
		Email: req.Email,
	})
	if err != nil {
		return nil, err
	}
	return &types.SignoutResponse{
		Message: cnt.Message,
	}, nil
}
