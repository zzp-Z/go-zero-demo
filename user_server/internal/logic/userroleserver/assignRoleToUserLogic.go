package userroleserverlogic

import (
	"context"
	"database/sql"
	"user_server/internal/logic"
	"user_server/model"

	"user_server/internal/svc"
	"user_server/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssignRoleToUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	userRoleModel model.UserRoleModel
}

func NewAssignRoleToUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssignRoleToUserLogic {
	return &AssignRoleToUserLogic{
		ctx:           ctx,
		svcCtx:        svcCtx,
		Logger:        logx.WithContext(ctx),
		userRoleModel: model.NewUserRoleModel(svcCtx.SqlConn, svcCtx.CacheConf),
	}
}

func (l *AssignRoleToUserLogic) AssignRoleToUser(in *user_server.UserRoleReqVo) (*user_server.UserRoleRespVo, error) {
	//sql.NullInt64
	userId := sql.NullInt64{Int64: in.UserId, Valid: true}
	roleId := sql.NullInt64{Int64: in.RoleId, Valid: true}
	_, err := l.userRoleModel.Insert(l.ctx, &model.UserRole{
		UserId: userId,
		RoleId: roleId,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("ARU301：添加失败, err:%v", err)
		return nil, &logic.AppError{
			Code:    "ARU301",
			Message: "添加失败",
		}
	}

	return &user_server.UserRoleRespVo{
		UserId: in.UserId,
		RoleId: in.RoleId,
	}, nil
}
