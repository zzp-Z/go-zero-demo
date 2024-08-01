package userroleserverlogic

import (
	"context"

	"user_server/internal/svc"
	"user_server/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveRoleFromUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveRoleFromUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveRoleFromUserLogic {
	return &RemoveRoleFromUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveRoleFromUserLogic) RemoveRoleFromUser(in *user_server.UserRoleReqVo) (*user_server.UserRoleRespVo, error) {
	// todo: add your logic here and delete this line

	return &user_server.UserRoleRespVo{}, nil
}
