package main

import (
	"context"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
	"log"
	"user_server/user_server"
)

func main() {
	conn := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{ // 通过 etcd 服务发现时，只需要给 Etcd 配置即可
			Hosts: []string{"AiCodeStudio.top:2379"},
			Key:   "userServer.rpc",
		},
	})
	client := user_server.NewUserServerClient(conn.Conn())
	resp, err := client.GetUser(context.Background(), &user_server.UserIdReqVo{
		Id: 2,
	})
	if err != nil {
		log.Println("err：", err)
		return
	}
	log.Println("resp：", resp)
	log.Println(resp.Name)
}
