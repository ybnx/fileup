package logic

import (
	"context"

	"fileup/rpc/user/internal/svc"
	"fileup/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePasswdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangePasswdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswdLogic {
	return &ChangePasswdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChangePasswdLogic) ChangePasswd(in *user.ChangePasswdRequest) (*user.ChangePasswdResponse, error) {
	// todo: add your logic here and delete this line

	return &user.ChangePasswdResponse{}, nil
}
