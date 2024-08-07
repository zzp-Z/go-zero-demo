// Code generated by goctl. DO NOT EDIT.
// Source: user_server.proto

package server

import (
	"context"

	"user_server/internal/logic/roleserver"
	"user_server/internal/svc"
	"user_server/user_server"
)

type RoleServerServer struct {
	svcCtx *svc.ServiceContext
	user_server.UnimplementedRoleServerServer
}

func NewRoleServerServer(svcCtx *svc.ServiceContext) *RoleServerServer {
	return &RoleServerServer{
		svcCtx: svcCtx,
	}
}

// 创建新角色
func (s *RoleServerServer) CreateRole(ctx context.Context, in *user_server.CreateRoleReqVo) (*user_server.RoleInfoRespVo, error) {
	l := roleserverlogic.NewCreateRoleLogic(ctx, s.svcCtx)
	return l.CreateRole(in)
}

// 获取角色信息
func (s *RoleServerServer) GetRole(ctx context.Context, in *user_server.RoleIdReqVo) (*user_server.RoleInfoRespVo, error) {
	l := roleserverlogic.NewGetRoleLogic(ctx, s.svcCtx)
	return l.GetRole(in)
}

// 删除角色
func (s *RoleServerServer) DeleteRole(ctx context.Context, in *user_server.RoleIdReqVo) (*user_server.RoleInfoRespVo, error) {
	l := roleserverlogic.NewDeleteRoleLogic(ctx, s.svcCtx)
	return l.DeleteRole(in)
}

// 获取角色列表
func (s *RoleServerServer) GetRoleList(ctx context.Context, in *user_server.Empty) (*user_server.RoleInfoListRespVo, error) {
	l := roleserverlogic.NewGetRoleListLogic(ctx, s.svcCtx)
	return l.GetRoleList(in)
}
