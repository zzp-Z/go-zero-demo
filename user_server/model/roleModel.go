// Package model user_server/model/roleModel.go
package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RoleModel = (*customRoleModel)(nil)

type (
	// RoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRoleModel.
	RoleModel interface {
		roleModel
		GetAllRoles(ctx context.Context) ([]*Role, error)
	}

	customRoleModel struct {
		*defaultRoleModel
	}
)

// NewRoleModel returns a model for the database table.
func NewRoleModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) RoleModel {
	return &customRoleModel{
		defaultRoleModel: newRoleModel(conn, c, opts...),
	}
}

// GetAllRoles 获取全部未删除的角色方法
func (m *defaultRoleModel) GetAllRoles(ctx context.Context) ([]*Role, error) {
	// 获取全部角色（不包含删除）
	query := fmt.Sprintf("select %s from %s where `deleted_at` IS NULL", roleRows, m.table)
	var roles []*Role
	// 查询无缓存的数据
	err := m.QueryRowsNoCacheCtx(ctx, &roles, query)
	if err != nil {
		return nil, err
	}
	return roles, nil
}
