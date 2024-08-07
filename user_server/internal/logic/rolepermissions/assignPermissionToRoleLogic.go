package rolepermissionslogic

import (
	"context"
	"database/sql"
	"user_server/internal/logic"
	"user_server/model"

	"user_server/internal/svc"
	"user_server/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssignPermissionToRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	rolePermissionsModel model.RolePermissionsModel
}

func NewAssignPermissionToRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssignPermissionToRoleLogic {
	return &AssignPermissionToRoleLogic{
		ctx:                  ctx,
		svcCtx:               svcCtx,
		Logger:               logx.WithContext(ctx),
		rolePermissionsModel: model.NewRolePermissionsModel(svcCtx.SqlConn, svcCtx.CacheConf),
	}
}

// AssignPermissionToRole 赋予角色权限
func (l *AssignPermissionToRoleLogic) AssignPermissionToRole(in *user_server.RolePermissionReqVo) (*user_server.RolePermissionRespVo, error) {
	/*
		1. 插入数据
		2. 返回数据
	*/
	_, err := l.rolePermissionsModel.Insert(l.ctx, &model.RolePermissions{
		PermissionsId: sql.NullInt64{
			Int64: in.PermissionId,
			Valid: true,
		},
		RoleId: sql.NullInt64{
			Int64: in.RoleId,
			Valid: true,
		},
	})
	if err != nil {
		return nil, logic.NewAppError(l.ctx, "APR301", "赋予权限失败", err)
	}

	return &user_server.RolePermissionRespVo{
		RoleId:       in.RoleId,
		PermissionId: in.PermissionId,
	}, nil
}
