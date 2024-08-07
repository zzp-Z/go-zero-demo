package permissionslogic

import (
	"context"
	"database/sql"
	"user_server/internal/logic"
	"user_server/model"

	"user_server/internal/svc"
	"user_server/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	PermissionsModel model.PermissionsModel
}

func NewCreatePermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePermissionLogic {
	return &CreatePermissionLogic{
		ctx:              ctx,
		svcCtx:           svcCtx,
		Logger:           logx.WithContext(ctx),
		PermissionsModel: model.NewPermissionsModel(svcCtx.SqlConn, svcCtx.CacheConf),
	}
}

func (l *CreatePermissionLogic) CreatePermission(in *user_server.CreatePermissionReqVo) (*user_server.PermissionInfoRespVo, error) {
	/*
		1. 创建权限
		2. 返回权限
	*/
	description := sql.NullString{ // 描述转换为sql.NullString
		String: in.Description,
		Valid:  true,
	}
	insert, err := l.PermissionsModel.Insert(l.ctx, &model.Permissions{ // 创建权限
		Description: description,
		Name:        in.Name,
		Status:      "0",
	})
	if err != nil {
		return nil, logic.NewAppError(l.ctx, "CP305", "创建权限错误", err)
	}
	id, _ := insert.LastInsertId()

	return &user_server.PermissionInfoRespVo{
		Id:          id,
		Name:        in.Name,
		Description: in.Description,
	}, nil
}
