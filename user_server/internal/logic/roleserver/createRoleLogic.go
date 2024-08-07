package roleserverlogic

import (
	"context"
	"user_server/internal/logic"
	"user_server/model"

	"user_server/internal/svc"
	"user_server/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	roleModel model.RoleModel // 模型接口定义为 RoleModel
}

func NewCreateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoleLogic {
	return &CreateRoleLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		Logger:    logx.WithContext(ctx),
		roleModel: model.NewRoleModel(svcCtx.SqlConn, svcCtx.CacheConf), // 初始化模型
	}
}

func (l *CreateRoleLogic) CreateRole(in *user_server.CreateRoleReqVo) (*user_server.RoleInfoRespVo, error) {
	/*
		1. 创建角色
	*/
	result, err := l.roleModel.Insert(l.ctx, &model.Role{
		Name:   in.Name,
		Status: "0",
	})
	if err != nil {
		return nil, logic.NewAppError(l.ctx, "CR362", "创建角色失败", err)
	}
	roleId, err := result.LastInsertId()
	if err != nil {
		return nil, logic.NewAppError(l.ctx, "CR363", "创建角色成功，查询id失败", err)
	}
	return &user_server.RoleInfoRespVo{
		Id:   roleId,
		Name: in.Name,
	}, nil
}
