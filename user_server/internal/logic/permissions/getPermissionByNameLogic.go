package permissionslogic

import (
	"context"
	"user_server/internal/logic"
	"user_server/model"

	"user_server/internal/svc"
	"user_server/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPermissionByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	PermissionsModel model.PermissionsModel
}

func NewGetPermissionByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPermissionByNameLogic {
	return &GetPermissionByNameLogic{
		ctx:              ctx,
		svcCtx:           svcCtx,
		Logger:           logx.WithContext(ctx),
		PermissionsModel: model.NewPermissionsModel(svcCtx.SqlConn, svcCtx.CacheConf),
	}
}

// GetPermissionByName 从名称获取权限
func (l *GetPermissionByNameLogic) GetPermissionByName(in *user_server.PermissionNameReqVo) (*user_server.PermissionInfoRespVo, error) {
	// 被权限校验依赖，不能添加权限校验
	// 根据名称查询权限
	permissions, err := l.PermissionsModel.FindOneByName(l.ctx, in.Name)
	if err != nil {
		return nil, logic.NewAppError(l.ctx, "GP336", "查询失败", err)
	}
	return &user_server.PermissionInfoRespVo{
		Id:          int64(permissions.Id),
		Name:        permissions.Name,
		Description: permissions.Description.String,
	}, nil
}
