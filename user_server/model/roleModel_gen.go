// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	roleFieldNames          = builder.RawFieldNames(&Role{})
	roleRows                = strings.Join(roleFieldNames, ",")
	roleRowsExpectAutoSet   = strings.Join(stringx.Remove(roleFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	roleRowsWithPlaceHolder = strings.Join(stringx.Remove(roleFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheUserServerRoleIdPrefix   = "cache:userServer:role:id:"
	cacheUserServerRoleNamePrefix = "cache:userServer:role:name:"
)

type (
	roleModel interface {
		Insert(ctx context.Context, data *Role) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*Role, error)
		FindOneByName(ctx context.Context, name string) (*Role, error)
		Update(ctx context.Context, data *Role) error
		Delete(ctx context.Context, id uint64) error
	}

	defaultRoleModel struct {
		sqlc.CachedConn
		table string
	}

	Role struct {
		Id        uint64       `db:"id"`
		CreatedAt sql.NullTime `db:"created_at"`
		UpdatedAt sql.NullTime `db:"updated_at"`
		DeletedAt sql.NullTime `db:"deleted_at"`
		Name      string       `db:"name"`
		Status    string       `db:"status"`
	}
)

func newRoleModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultRoleModel {
	return &defaultRoleModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`role`",
	}
}

func (m *defaultRoleModel) Delete(ctx context.Context, id uint64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	userServerRoleIdKey := fmt.Sprintf("%s%v", cacheUserServerRoleIdPrefix, id)
	userServerRoleNameKey := fmt.Sprintf("%s%v", cacheUserServerRoleNamePrefix, data.Name)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, userServerRoleIdKey, userServerRoleNameKey)
	return err
}

func (m *defaultRoleModel) FindOne(ctx context.Context, id uint64) (*Role, error) {
	userServerRoleIdKey := fmt.Sprintf("%s%v", cacheUserServerRoleIdPrefix, id)
	var resp Role
	err := m.QueryRowCtx(ctx, &resp, userServerRoleIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", roleRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRoleModel) FindOneByName(ctx context.Context, name string) (*Role, error) {
	userServerRoleNameKey := fmt.Sprintf("%s%v", cacheUserServerRoleNamePrefix, name)
	var resp Role
	err := m.QueryRowIndexCtx(ctx, &resp, userServerRoleNameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", roleRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, name); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRoleModel) Insert(ctx context.Context, data *Role) (sql.Result, error) {
	userServerRoleIdKey := fmt.Sprintf("%s%v", cacheUserServerRoleIdPrefix, data.Id)
	userServerRoleNameKey := fmt.Sprintf("%s%v", cacheUserServerRoleNamePrefix, data.Name)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, roleRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.DeletedAt, data.Name, data.Status)
	}, userServerRoleIdKey, userServerRoleNameKey)
	return ret, err
}

func (m *defaultRoleModel) Update(ctx context.Context, newData *Role) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	userServerRoleIdKey := fmt.Sprintf("%s%v", cacheUserServerRoleIdPrefix, data.Id)
	userServerRoleNameKey := fmt.Sprintf("%s%v", cacheUserServerRoleNamePrefix, data.Name)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, roleRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.DeletedAt, newData.Name, newData.Status, newData.Id)
	}, userServerRoleIdKey, userServerRoleNameKey)
	return err
}

func (m *defaultRoleModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheUserServerRoleIdPrefix, primary)
}

func (m *defaultRoleModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", roleRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultRoleModel) tableName() string {
	return m.table
}
