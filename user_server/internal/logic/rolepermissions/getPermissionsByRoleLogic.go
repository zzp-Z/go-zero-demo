package rolepermissionslogic

import (
	"context"

	"user_server/internal/svc"
	"user_server/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPermissionsByRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPermissionsByRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPermissionsByRoleLogic {
	return &GetPermissionsByRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetPermissionsByRole 获取角色权限
func (l *GetPermissionsByRoleLogic) GetPermissionsByRole(in *user_server.RoleIdReqVo) (*user_server.RolePermissionsRespVo, error) {
	// TODO：传入角色id，获取角色权限
	// TODO：获取角色权限列表
	// TODO：封装返回值

	return &user_server.RolePermissionsRespVo{}, nil
}
