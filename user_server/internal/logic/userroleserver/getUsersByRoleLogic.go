package userroleserverlogic

import (
	"context"
	"user_server/internal/logic"
	"user_server/model"

	"user_server/internal/svc"
	"user_server/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUsersByRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	userRoleModel model.UserRoleModel
}

func NewGetUsersByRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUsersByRoleLogic {
	return &GetUsersByRoleLogic{
		ctx:           ctx,
		svcCtx:        svcCtx,
		Logger:        logx.WithContext(ctx),
		userRoleModel: model.NewUserRoleModel(svcCtx.SqlConn, svcCtx.CacheConf),
	}
}

func (l *GetUsersByRoleLogic) GetUsersByRole(in *user_server.RoleIdReqVo) (*user_server.RoleUsersRespVo, error) {
	// 查询用户角色
	roleUsers, err := l.userRoleModel.FindByRoleId(l.ctx, in.Id)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("GUR301：查询角色用户失败, err：%v", err)
		return nil, &logic.AppError{
			Code:    "GUR301",
			Message: "查询角色用户失败",
		}
	}
	users := make([]*user_server.UserInfoRespVo, 0, len(roleUsers))
	for _, roleUser := range roleUsers {
		user, err := l.svcCtx.UserModel.FindOne(l.ctx, uint64(roleUser.UserId.Int64))
		if err != nil {
			logx.WithContext(l.ctx).Errorf("GUR302：查询用户失败, err：%v", err)
			continue
		}
		users = append(users, &user_server.UserInfoRespVo{
			Id:    int64(user.Id),
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return &user_server.RoleUsersRespVo{
		Users: users,
	}, nil
}
