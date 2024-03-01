package logic

import (
	"context"
	"errors"
	"fileup/rpc/user/internal/svc"
	"fileup/rpc/user/user"
	"fileup/util"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCodeLogic {
	return &SendCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendCodeLogic) SendCode(in *user.SendCodeRequest) (*user.SendCodeResponse, error) {
	// todo: add your logic here and delete this line
	veCode := util.GenerateCode()
	err := util.SendCode(l.svcCtx.Config.Email.UserName, l.svcCtx.Config.Email.Password, in.Email, veCode)
	if err != nil {
		// TODO log
		return nil, errors.New("send verify code error")
	}
	err = l.svcCtx.RDB.Set(l.ctx, fmt.Sprintf("%s:%s", util.VeCodeNS, in.Email), veCode, time.Duration(l.svcCtx.Config.Auth.AccessExpire)*time.Second).Err()
	if err != nil {
		return nil, errors.New("fail to set token version")
	}
	return &user.SendCodeResponse{
		Message: "ok",
	}, nil
}
