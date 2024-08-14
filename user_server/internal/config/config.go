package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	// 其他配置项
	Port       int    // 端口号
	DataSource string // 数据库连接字符串
	Cache      cache.CacheConf
}
