package logic

import (
	"context"
	"errors"
	"fileup/model"
	"fileup/util"
	"fmt"
	"strconv"

	"fileup/rpc/user/internal/svc"
	"fileup/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type PasswdLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPasswdLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PasswdLoginLogic {
	return &PasswdLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PasswdLoginLogic) PasswdLogin(in *user.PasswdLoginRequest) (*user.PasswdLoginResponse, error) {
	// todo: add your logic here and delete this line
	info := model.User{}
	err := l.svcCtx.DB.Where("email = ? AND password = ?", in.Email, in.Password).First(&info).Error
	if err != nil {
		return nil, errors.New("fail to get user")
	}
	if info.Password == "" {
		return nil, errors.New("name or password error")
	}
	val, err := l.svcCtx.RDB.Get(l.ctx, fmt.Sprintf("%s:%s", util.VersionNS, in.Email)).Result()
	if err != nil {
		return nil, errors.New("fail to get token version")
	}
	version, err := strconv.Atoi(val)
	if err != nil {
		return nil, errors.New("fail to strconv")
	}
	accessToken, refreshToken, err := util.GenerateToken(info.Id, int64(version), info.Name, info.Email, l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire)
	if err != nil {
		return nil, errors.New("fail to generate token")
	}
	return &user.PasswdLoginResponse{
		UserId:       info.Id,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
