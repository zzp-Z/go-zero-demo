// Package model user_server/model/userModel.go
package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c, opts...),
	}
}

// Delete 重写Delete方法
func (m *customUserModel) Delete(ctx context.Context, id uint64) error {
	// 先查询数据
	data, err := m.FindOne(ctx, id)
	if err != nil {
		fmt.Println("errTX034:", err)
		return err
	}

	// 获取缓存相关的key
	userServerUserEmailKey := fmt.Sprintf("%s%v", cacheUserServerUserEmailPrefix, data.Email)
	userServerUserIdKey := fmt.Sprintf("%s%v", cacheUserServerUserIdPrefix, id)
	userServerUserNameKey := fmt.Sprintf("%s%v", cacheUserServerUserNamePrefix, data.Name)

	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `deleted_at` = now() where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, userServerUserEmailKey, userServerUserIdKey, userServerUserNameKey)

	return err
}
