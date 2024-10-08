// Code generated by goctl. DO NOT EDIT.
// Source: user_server.proto

package userroleserver

import (
	"context"

	"user_server/user_server"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BoolRespVo                 = user_server.BoolRespVo
	CreatePermissionReqVo      = user_server.CreatePermissionReqVo
	CreateRoleReqVo            = user_server.CreateRoleReqVo
	CreateUserReqVo            = user_server.CreateUserReqVo
	Empty                      = user_server.Empty
	JwtTokenRespVo             = user_server.JwtTokenRespVo
	PermissionIdReqVo          = user_server.PermissionIdReqVo
	PermissionInfoRespVo       = user_server.PermissionInfoRespVo
	PermissionNameReqVo        = user_server.PermissionNameReqVo
	PermissionRolesRespVo      = user_server.PermissionRolesRespVo
	RoleIdPermissionIdReqVo    = user_server.RoleIdPermissionIdReqVo
	RoleIdReqVo                = user_server.RoleIdReqVo
	RoleInfoListRespVo         = user_server.RoleInfoListRespVo
	RoleInfoRespVo             = user_server.RoleInfoRespVo
	RolePermissionReqVo        = user_server.RolePermissionReqVo
	RolePermissionRespVo       = user_server.RolePermissionRespVo
	RolePermissionsRespVo      = user_server.RolePermissionsRespVo
	RoleUsersRespVo            = user_server.RoleUsersRespVo
	UpdateUserReqVo            = user_server.UpdateUserReqVo
	UserIdReqVo                = user_server.UserIdReqVo
	UserInfoRespVo             = user_server.UserInfoRespVo
	UserRoleReqVo              = user_server.UserRoleReqVo
	UserRoleRespVo             = user_server.UserRoleRespVo
	UserRolesRespVo            = user_server.UserRolesRespVo
	UserTokenPermissionIdReqVo = user_server.UserTokenPermissionIdReqVo
	VerificationReqVo          = user_server.VerificationReqVo

	UserRoleServer interface {
		// 分配角色给用户
		AssignRoleToUser(ctx context.Context, in *UserRoleReqVo, opts ...grpc.CallOption) (*UserRoleRespVo, error)
		// 移除用户的角色
		RemoveRoleFromUser(ctx context.Context, in *UserRoleReqVo, opts ...grpc.CallOption) (*UserRoleRespVo, error)
		// 获取用户角色
		GetRolesByUser(ctx context.Context, in *UserIdReqVo, opts ...grpc.CallOption) (*UserRolesRespVo, error)
		// 获取角色用户
		GetUsersByRole(ctx context.Context, in *RoleIdReqVo, opts ...grpc.CallOption) (*RoleUsersRespVo, error)
	}

	defaultUserRoleServer struct {
		cli zrpc.Client
	}
)

func NewUserRoleServer(cli zrpc.Client) UserRoleServer {
	return &defaultUserRoleServer{
		cli: cli,
	}
}

// 分配角色给用户
func (m *defaultUserRoleServer) AssignRoleToUser(ctx context.Context, in *UserRoleReqVo, opts ...grpc.CallOption) (*UserRoleRespVo, error) {
	client := user_server.NewUserRoleServerClient(m.cli.Conn())
	return client.AssignRoleToUser(ctx, in, opts...)
}

// 移除用户的角色
func (m *defaultUserRoleServer) RemoveRoleFromUser(ctx context.Context, in *UserRoleReqVo, opts ...grpc.CallOption) (*UserRoleRespVo, error) {
	client := user_server.NewUserRoleServerClient(m.cli.Conn())
	return client.RemoveRoleFromUser(ctx, in, opts...)
}

// 获取用户角色
func (m *defaultUserRoleServer) GetRolesByUser(ctx context.Context, in *UserIdReqVo, opts ...grpc.CallOption) (*UserRolesRespVo, error) {
	client := user_server.NewUserRoleServerClient(m.cli.Conn())
	return client.GetRolesByUser(ctx, in, opts...)
}

// 获取角色用户
func (m *defaultUserRoleServer) GetUsersByRole(ctx context.Context, in *RoleIdReqVo, opts ...grpc.CallOption) (*RoleUsersRespVo, error) {
	client := user_server.NewUserRoleServerClient(m.cli.Conn())
	return client.GetUsersByRole(ctx, in, opts...)
}
