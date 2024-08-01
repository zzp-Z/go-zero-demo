package main

import (
	"flag"
	"fmt"

	"user_server/internal/config"
	roleserverServer "user_server/internal/server/roleserver"
	userroleserverServer "user_server/internal/server/userroleserver"
	userserverServer "user_server/internal/server/userserver"
	"user_server/internal/svc"
	"user_server/user_server"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/userServer.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user_server.RegisterUserServerServer(grpcServer, userserverServer.NewUserServerServer(ctx))
		user_server.RegisterRoleServerServer(grpcServer, roleserverServer.NewRoleServerServer(ctx))
		user_server.RegisterUserRoleServerServer(grpcServer, userroleserverServer.NewUserRoleServerServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
