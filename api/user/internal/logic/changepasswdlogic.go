package logic

import (
	"context"
	"fileup/rpc/user/user"

	"fileup/api/user/internal/svc"
	"fileup/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePasswdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangePasswdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePasswdLogic {
	return &ChangePasswdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangePasswdLogic) ChangePasswd(req *types.ChangePasswdRequest) (resp *types.ChangePasswdResponse, err error) {
	// todo: add your logic here and delete this line
	cnt, err := l.svcCtx.Rpc.ChangePasswd(l.ctx, &user.ChangePasswdRequest{
		NewPasswd: req.NewPasswd,
	})
	if err != nil {
		return nil, err
	}
	return &types.ChangePasswdResponse{
		Message: cnt.Message,
	}, nil
}
