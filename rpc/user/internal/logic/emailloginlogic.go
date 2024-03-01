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

type EmailLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEmailLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailLoginLogic {
	return &EmailLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EmailLoginLogic) EmailLogin(in *user.EmailLoginRequest) (*user.EmailLoginResponse, error) {
	// todo: add your logic here and delete this line
	veCode, err := l.svcCtx.RDB.Get(l.ctx, fmt.Sprintf("%s:%s", util.VeCodeNS, in.Email)).Result()
	if err != nil {
		return nil, errors.New("vecode not exist")
	}
	if veCode != in.VeCode {
		return nil, errors.New("vecode invalid")
	}
	info := model.User{}
	err = l.svcCtx.DB.Where("email = ?", in.Email).First(&info).Error
	if err != nil {
		return nil, errors.New("fail to get user")
	}
	if info.Password == "" {
		return nil, errors.New("user not exist")
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
	return &user.EmailLoginResponse{
		UserId:       info.Id,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
