package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net"

	"user_server/internal/config"
	permissionsServer "user_server/internal/server/permissions"
	rolepermissionsServer "user_server/internal/server/rolepermissions"
	roleserverServer "user_server/internal/server/roleserver"
	userpermissionsServer "user_server/internal/server/userpermissions"
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

// getRandomPort 生成一个在 min 和 max 之间的随机端口号。
func getRandomPort(min, max int) int {
	return rand.Intn(max-min+1) + min
}

// isPortAvailable 检查给定的端口是否可用。
func isPortAvailable(port int) bool {
	address := fmt.Sprintf(":%d", port)         // 创建带有端口号的地址字符串。
	listener, err := net.Listen("tcp", address) // 尝试在该端口上监听。
	if err != nil {                             // 如果发生错误，说明端口不可用。
		return false
	}
	_ = listener.Close()
	return true
}

// findAvailablePort 在 min 和 max 范围内找到一个可用的端口。
func findAvailablePort(min, max int) int {
	for {
		port := getRandomPort(min, max) // 生成一个随机端口号。
		if isPortAvailable(port) {      // 检查端口是否可用。
			return port // 返回可用的端口。
		}
	}
}
func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	// 如果端口号未设置，则使用 findAvailablePort 函数找到一个可用的端口。
	if c.Port == 0 {
		// 在 30000 到 40000 范围内找到一个可用的端口。
		c.Port = findAvailablePort(30000, 40000)
	}
	// 使用找到的端口号设置 ListenOn 字段。
	c.ListenOn = fmt.Sprintf("%s:%d", c.ListenOn, c.Port)

	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user_server.RegisterUserServerServer(grpcServer, userserverServer.NewUserServerServer(ctx))
		user_server.RegisterRoleServerServer(grpcServer, roleserverServer.NewRoleServerServer(ctx))
		user_server.RegisterUserRoleServerServer(grpcServer, userroleserverServer.NewUserRoleServerServer(ctx))
		user_server.RegisterPermissionsServer(grpcServer, permissionsServer.NewPermissionsServer(ctx))
		user_server.RegisterRolePermissionsServer(grpcServer, rolepermissionsServer.NewRolePermissionsServer(ctx))
		user_server.RegisterUserPermissionsServer(grpcServer, userpermissionsServer.NewUserPermissionsServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
