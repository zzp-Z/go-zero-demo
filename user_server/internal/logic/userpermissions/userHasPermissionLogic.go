package userpermissionslogic

import (
	"context"
	"user_server/internal/logic"
	permissionslogic "user_server/internal/logic/permissions"
	rolepermissionslogic "user_server/internal/logic/rolepermissions"
	userroleserverlogic "user_server/internal/logic/userroleserver"
	"user_server/internal/svc"
	"user_server/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserHasPermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	tools *logic.Tools
}

func NewUserHasPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserHasPermissionLogic {
	return &UserHasPermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		tools:  logic.NewTools(),
	}
}

// UserHasPermission 检查用户是否具有某权限
func (l *UserHasPermissionLogic) UserHasPermission(in *user_server.UserTokenPermissionIdReqVo) (*user_server.BoolRespVo, error) {
	// 必须保证以下使用的方法并不需要校验权限
	// 校验token
	userId, err := l.tools.ParseJwtToken(in.Token)
	if err != nil {
		return nil, logic.NewAppError(l.ctx, "UHP306", "Token校验失败", err)
	}
	// 获取用户具备的角色列表（调用UserRoleServer的GetRolesByUser方法）
	getRolesByUserLogic := userroleserverlogic.NewGetRolesByUserLogic(l.ctx, l.svcCtx)
	roles, err := getRolesByUserLogic.GetRolesByUser(&user_server.UserIdReqVo{Id: int64(userId)})
	if err != nil {
		return nil, logic.NewAppError(l.ctx, "UHP307", "获取用户角色失败", err)
	}
	// 获取权限的id（调用PermissionServer的GetPermissionByName方法）
	getPermissionByNameLogic := permissionslogic.NewGetPermissionByNameLogic(l.ctx, l.svcCtx)
	permission, err := getPermissionByNameLogic.GetPermissionByName(&user_server.PermissionNameReqVo{
		Name: in.PermissionName,
	})
	if err != nil {
		return nil, logic.NewAppError(l.ctx, "UHP308", "未知权限，请检查权限名称是否正确", err)
	}
	// 判断用户是否具有权限（调用RolePermissionServer的RoleHasPermission方法）
	roleHasPermissionLogic := rolepermissionslogic.NewRoleHasPermissionLogic(l.ctx, l.svcCtx)
	for _, role := range roles.Roles {
		boolValue, _ := roleHasPermissionLogic.RoleHasPermission(&user_server.RoleIdPermissionIdReqVo{
			RoleId:       role.Id,
			PermissionId: permission.Id,
		})
		// 如果有，则返回true
		if boolValue.Value {
			return &user_server.BoolRespVo{
				Value: true,
			}, nil
		}
	}
	// 如果没有，则返回false
	return &user_server.BoolRespVo{
		Value: false,
	}, nil
}
