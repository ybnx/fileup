package logic

import (
	"context"
	"errors"
	"fileup/rpc/user/user"

	"fileup/api/user/internal/svc"
	"fileup/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line

	if req.UserId == "" {
		return nil, errors.New("userid cannot empty")
	}

	cnt, err := l.svcCtx.Rpc.UserInfo(l.ctx, &user.UserInfoRequest{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	return &types.UserInfoResponse{
		UserId: cnt.UserId,
		Name:   cnt.Name,
		Email:  cnt.Email,
	}, nil
}
