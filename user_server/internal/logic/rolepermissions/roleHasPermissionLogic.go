package rolepermissionslogic

import (
	"context"
	"database/sql"
	"user_server/model"

	"user_server/internal/svc"
	"user_server/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleHasPermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rolePermissionsModel model.RolePermissionsModel
}

func NewRoleHasPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleHasPermissionLogic {
	return &RoleHasPermissionLogic{
		ctx:                  ctx,
		svcCtx:               svcCtx,
		Logger:               logx.WithContext(ctx),
		rolePermissionsModel: model.NewRolePermissionsModel(svcCtx.SqlConn, svcCtx.CacheConf),
	}
}

// RoleHasPermission 检查角色是否具有某权限
func (l *RoleHasPermissionLogic) RoleHasPermission(in *user_server.RoleIdPermissionIdReqVo) (*user_server.BoolRespVo, error) {
	// 被权限校验依赖，不能添加权限校验
	// 查询角色权限表，判断是否存在
	_, err := l.rolePermissionsModel.FindOneByPermissionsIdRoleId(
		l.ctx,
		sql.NullInt64{Int64: in.PermissionId, Valid: true},
		sql.NullInt64{Int64: in.RoleId, Valid: true},
	)
	if err != nil {
		// 查询失败返回结果false
		return &user_server.BoolRespVo{
			Value: false,
		}, nil
	}
	// 返回结果true
	return &user_server.BoolRespVo{
		Value: true,
	}, nil
}
