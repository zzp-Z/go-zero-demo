package main

import (
	"context"
	"github.com/zeromicro/go-zero/zrpc"
	"log"
	"user_server/user_server"
)

func main() {
	conn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Target: "dns:///127.0.0.1:8080",
	})
	client := user_server.NewUserServerClient(conn.Conn())
	resp, err := client.GetUser(context.Background(), &user_server.UserIdReqVo{
		Id: 1,
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(resp)
}
