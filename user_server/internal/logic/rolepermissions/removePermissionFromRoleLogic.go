package rolepermissionslogic

import (
	"context"

	"user_server/internal/svc"
	"user_server/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemovePermissionFromRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemovePermissionFromRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemovePermissionFromRoleLogic {
	return &RemovePermissionFromRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemovePermissionFromRoleLogic) RemovePermissionFromRole(in *user_server.RolePermissionReqVo) (*user_server.RolePermissionRespVo, error) {
	// todo: add your logic here and delete this line

	return &user_server.RolePermissionRespVo{}, nil
}
