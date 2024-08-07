package permissionslogic

import (
	"context"
	"user_server/internal/logic"
	"user_server/model"

	"user_server/internal/svc"
	"user_server/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	PermissionsModel model.PermissionsModel
}

func NewGetPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPermissionLogic {
	return &GetPermissionLogic{
		ctx:              ctx,
		svcCtx:           svcCtx,
		Logger:           logx.WithContext(ctx),
		PermissionsModel: model.NewPermissionsModel(svcCtx.SqlConn, svcCtx.CacheConf),
	}
}

func (l *GetPermissionLogic) GetPermission(in *user_server.PermissionIdReqVo) (*user_server.PermissionInfoRespVo, error) {
	// 根据id查询权限
	permissions, err := l.PermissionsModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		return nil, logic.NewAppError(l.ctx, "GP336", "查询失败", err)
	}
	// 返回权限信息

	return &user_server.PermissionInfoRespVo{
		Id:          in.Id,
		Name:        permissions.Name,
		Description: permissions.Description.String,
	}, nil
}
