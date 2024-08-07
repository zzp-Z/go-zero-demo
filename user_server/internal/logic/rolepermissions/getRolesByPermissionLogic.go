package rolepermissionslogic

import (
	"context"

	"user_server/internal/svc"
	"user_server/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRolesByPermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRolesByPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRolesByPermissionLogic {
	return &GetRolesByPermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetRolesByPermission 权限获取角色列表
func (l *GetRolesByPermissionLogic) GetRolesByPermission(in *user_server.PermissionIdReqVo) (*user_server.PermissionRolesRespVo, error) {
	// todo: add your logic here and delete this line

	return &user_server.PermissionRolesRespVo{}, nil
}
