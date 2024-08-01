package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserRoleModel = (*customUserRoleModel)(nil)

type (
	// UserRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserRoleModel.
	UserRoleModel interface {
		userRoleModel
		FindByUserId(ctx context.Context, userId int64) ([]*RoleIdOnly, error)
		FindByRoleId(ctx context.Context, roleId int64) ([]*UserIdRole, error)
	}

	customUserRoleModel struct {
		*defaultUserRoleModel
	}
	RoleIdOnly struct {
		RoleId sql.NullInt64 `db:"role_id"`
	}
	UserIdRole struct {
		UserId sql.NullInt64 `db:"user_id"`
	}
)

// NewUserRoleModel returns a model for the database table.
func NewUserRoleModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserRoleModel {
	return &customUserRoleModel{
		defaultUserRoleModel: newUserRoleModel(conn, c, opts...),
	}
}

// FindByUserId 根据 userId 查找roleId
func (m *customUserRoleModel) FindByUserId(ctx context.Context, userId int64) ([]*RoleIdOnly, error) {
	// 构造查询语句
	query := fmt.Sprintf("select `role_id` from %s where `user_id` = ? and `deleted_at` IS NULL", m.table)

	var userRoles []*RoleIdOnly

	// 查询无缓存的数据
	err := m.QueryRowsNoCacheCtx(ctx, &userRoles, query, userId)

	if err != nil {
		return nil, err
	}

	return userRoles, nil
}

// FindByRoleId 根据roleId 查找 userId
func (m *customUserRoleModel) FindByRoleId(ctx context.Context, roleId int64) ([]*UserIdRole, error) {
	// 构造查询语句
	query := fmt.Sprintf("select `user_id` from %s where `role_id` = ? and `deleted_at` IS NULL", m.table)

	var userRoles []*UserIdRole

	// 查询无缓存的数据
	err := m.QueryRowsNoCacheCtx(ctx, &userRoles, query, roleId)

	if err != nil {
		return nil, err
	}

	return userRoles, nil
}
