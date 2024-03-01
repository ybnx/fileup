package logic

import (
	"context"
	"errors"
	"fileup/model"

	"fileup/rpc/user/internal/svc"
	"fileup/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	// todo: add your logic here and delete this line
	info := model.User{}
	err := l.svcCtx.DB.Where("id = ?", in.UserId).First(&info).Error
	if err != nil {
		return nil, errors.New("fail to get user")
	}
	if info.Email == "" {
		return nil, errors.New("user not exist")
	}
	return &user.UserInfoResponse{
		UserId: info.Id,
		Name:   info.Name,
		Email:  info.Email,
	}, nil
}
