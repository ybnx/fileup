package logic

import (
	"context"
	"errors"
	"fileup/util"
	"fmt"
	"strconv"

	"fileup/rpc/user/internal/svc"
	"fileup/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSignoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignoutLogic {
	return &SignoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SignoutLogic) Signout(in *user.SignoutRequest) (*user.SignoutResponse, error) {
	// todo: add your logic here and delete this line
	val, err := l.svcCtx.RDB.Get(l.ctx, fmt.Sprintf("%s:%s", util.VersionNS, in.Email)).Result()
	if err != nil {
		return nil, errors.New("fail to get token version")
	}
	version, err := strconv.Atoi(val)
	if err != nil {
		return nil, errors.New("fail to strconv")
	}
	err = l.svcCtx.RDB.Set(l.ctx, fmt.Sprintf("%s:%s", util.VersionNS, in.Email), version+1, 0).Err()
	if err != nil {
		return nil, errors.New("fail to set token version")
	}
	return &user.SignoutResponse{
		Message: "ok",
	}, nil
}
