// Code generated by goctl. DO NOT EDIT.
// Source: user_server.proto

package server

import (
	"context"

	"user_server/internal/logic/userserver"
	"user_server/internal/svc"
	"user_server/user_server"
)

type UserServerServer struct {
	svcCtx *svc.ServiceContext
	user_server.UnimplementedUserServerServer
}

func NewUserServerServer(svcCtx *svc.ServiceContext) *UserServerServer {
	return &UserServerServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServerServer) GetUser(ctx context.Context, in *user_server.UserIdReqVo) (*user_server.UserInfoRespVo, error) {
	l := userserverlogic.NewGetUserLogic(ctx, s.svcCtx)
	return l.GetUser(in)
}

func (s *UserServerServer) UpdateUser(ctx context.Context, in *user_server.UpdateUserReqVo) (*user_server.JwtTokenRespVo, error) {
	l := userserverlogic.NewUpdateUserLogic(ctx, s.svcCtx)
	return l.UpdateUser(in)
}

func (s *UserServerServer) CreateUser(ctx context.Context, in *user_server.CreateUserReqVo) (*user_server.JwtTokenRespVo, error) {
	l := userserverlogic.NewCreateUserLogic(ctx, s.svcCtx)
	return l.CreateUser(in)
}

func (s *UserServerServer) Login(ctx context.Context, in *user_server.VerificationReqVo) (*user_server.JwtTokenRespVo, error) {
	l := userserverlogic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServerServer) DeleteUser(ctx context.Context, in *user_server.VerificationReqVo) (*user_server.JwtTokenRespVo, error) {
	l := userserverlogic.NewDeleteUserLogic(ctx, s.svcCtx)
	return l.DeleteUser(in)
}
