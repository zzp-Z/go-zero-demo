package userroleserverlogic

import (
	"context"
	"user_server/internal/logic"
	"user_server/model"

	"user_server/internal/svc"
	"user_server/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRolesByUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	userModel     model.UserModel
	roleModel     model.RoleModel
	userRoleModel model.UserRoleModel
}

func NewGetRolesByUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRolesByUserLogic {
	return &GetRolesByUserLogic{
		ctx:           ctx,
		svcCtx:        svcCtx,
		Logger:        logx.WithContext(ctx),
		userModel:     model.NewUserModel(svcCtx.SqlConn, svcCtx.CacheConf),
		roleModel:     model.NewRoleModel(svcCtx.SqlConn, svcCtx.CacheConf),
		userRoleModel: model.NewUserRoleModel(svcCtx.SqlConn, svcCtx.CacheConf),
	}
}

func (l *GetRolesByUserLogic) GetRolesByUser(in *user_server.UserIdReqVo) (*user_server.UserRolesRespVo, error) {
	// 查询用户角色
	userRoles, err := l.userRoleModel.FindByUserId(l.ctx, in.Id)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("GRU303：查询失败, err: %v", err)
		return nil, &logic.AppError{
			Code:    "GRU303",
			Message: "查询失败",
		}
	}
	// 遍历用户角色id列表
	roles := make([]*user_server.RoleInfoRespVo, 0, len(userRoles))
	for _, userRole := range userRoles {
		// 查询角色
		role, err := l.roleModel.FindOne(l.ctx, uint64(userRole.RoleId.Int64))
		if err != nil {
			logx.WithContext(l.ctx).Errorf("GRU302：查询失败, err: %v", err)
			continue
		}
		// 封装角色信息
		roles = append(roles, &user_server.RoleInfoRespVo{
			Id:   int64(role.Id),
			Name: role.Name,
		})
	}

	return &user_server.UserRolesRespVo{
		Roles: roles,
	}, nil
}
