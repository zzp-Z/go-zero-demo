// Code generated by goctl. DO NOT EDIT.
// Source: user_server.proto

package userserver

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

	UserServer interface {
		// 获取用户信息
		GetUser(ctx context.Context, in *UserIdReqVo, opts ...grpc.CallOption) (*UserInfoRespVo, error)
		// 更新用户信息
		UpdateUser(ctx context.Context, in *UpdateUserReqVo, opts ...grpc.CallOption) (*JwtTokenRespVo, error)
		// 创建新用户
		CreateUser(ctx context.Context, in *CreateUserReqVo, opts ...grpc.CallOption) (*JwtTokenRespVo, error)
		// 用户登录
		Login(ctx context.Context, in *VerificationReqVo, opts ...grpc.CallOption) (*JwtTokenRespVo, error)
		// 删除用户
		DeleteUser(ctx context.Context, in *VerificationReqVo, opts ...grpc.CallOption) (*JwtTokenRespVo, error)
	}

	defaultUserServer struct {
		cli zrpc.Client
	}
)

func NewUserServer(cli zrpc.Client) UserServer {
	return &defaultUserServer{
		cli: cli,
	}
}

// 获取用户信息
func (m *defaultUserServer) GetUser(ctx context.Context, in *UserIdReqVo, opts ...grpc.CallOption) (*UserInfoRespVo, error) {
	client := user_server.NewUserServerClient(m.cli.Conn())
	return client.GetUser(ctx, in, opts...)
}

// 更新用户信息
func (m *defaultUserServer) UpdateUser(ctx context.Context, in *UpdateUserReqVo, opts ...grpc.CallOption) (*JwtTokenRespVo, error) {
	client := user_server.NewUserServerClient(m.cli.Conn())
	return client.UpdateUser(ctx, in, opts...)
}

// 创建新用户
func (m *defaultUserServer) CreateUser(ctx context.Context, in *CreateUserReqVo, opts ...grpc.CallOption) (*JwtTokenRespVo, error) {
	client := user_server.NewUserServerClient(m.cli.Conn())
	return client.CreateUser(ctx, in, opts...)
}

// 用户登录
func (m *defaultUserServer) Login(ctx context.Context, in *VerificationReqVo, opts ...grpc.CallOption) (*JwtTokenRespVo, error) {
	client := user_server.NewUserServerClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

// 删除用户
func (m *defaultUserServer) DeleteUser(ctx context.Context, in *VerificationReqVo, opts ...grpc.CallOption) (*JwtTokenRespVo, error) {
	client := user_server.NewUserServerClient(m.cli.Conn())
	return client.DeleteUser(ctx, in, opts...)
}
