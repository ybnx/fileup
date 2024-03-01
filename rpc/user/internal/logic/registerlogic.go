package logic

import (
	"context"
	"errors"
	"fileup/model"
	"fileup/util"
	"fmt"
	"gorm.io/gorm"
	"time"

	"fileup/rpc/user/internal/svc"
	"fileup/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	// todo: add your logic here and delete this line
	veCode, err := l.svcCtx.RDB.Get(l.ctx, fmt.Sprintf("%s:%s", util.VeCodeNS, in.Email)).Result()
	if err != nil {
		return nil, errors.New("fail to get vecode")
	}
	if veCode != in.VeCode {
		return nil, errors.New("vecode invalid")
	}
	info := model.User{}
	err = l.svcCtx.DB.Where("email = ?", in.Email).First(&info).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		//return nil, errors.New("fail to get user")
		return nil, err
	}
	if info.Email != "" {
		return nil, errors.New("user has exist")
	}
	err = l.svcCtx.RDB.Set(l.ctx, fmt.Sprintf("%s:%s", util.VersionNS, in.Email), 0, 0).Err()
	if err != nil {
		return nil, errors.New("fail to set token version")
	}
	info = model.User{
		Name:      in.Name,
		Password:  util.Hash(in.Password),
		Email:     in.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = l.svcCtx.DB.Create(&info).Error
	if err != nil {
		return nil, errors.New("fail to create user")
	}
	accessToken, refreshToken, err := util.GenerateToken(info.Id, int64(0), info.Name, info.Email, l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire)
	if err != nil {
		return nil, errors.New("fail to generate token")
	} // TODO bug 如果这里报错，那之前插入的user怎么算
	return &user.RegisterResponse{
		UserId:       info.Id,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
