package roleserverlogic

import (
	"context"
	"user_server/internal/logic"
	"user_server/model"

	"user_server/internal/svc"
	"user_server/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	roleModel model.RoleModel
}

func NewGetRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleLogic {
	return &GetRoleLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		Logger:    logx.WithContext(ctx),
		roleModel: model.NewRoleModel(svcCtx.SqlConn, svcCtx.CacheConf),
	}
}

func (l *GetRoleLogic) GetRole(in *user_server.RoleIdReqVo) (*user_server.RoleInfoRespVo, error) {
	/*
		1. 获取角色
		2. 返回角色
	*/
	role, err := l.roleModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		return nil, logic.NewAppError(l.ctx, "GR301", "角色未找到", err)
	}

	return &user_server.RoleInfoRespVo{
		Id:   in.Id,
		Name: role.Name,
	}, nil
}
