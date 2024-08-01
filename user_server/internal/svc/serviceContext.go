// Package svc user_server/internal/svc/serviceContext.go
package svc

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"user_server/internal/config"
	"user_server/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
	SqlConn   sqlx.SqlConn
	CacheConf cache.CacheConf
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:    c,
		SqlConn:   conn,
		UserModel: model.NewUserModel(conn, c.Cache),
		CacheConf: c.Cache, // 确保 CacheConf 被正确传递
	}
}
