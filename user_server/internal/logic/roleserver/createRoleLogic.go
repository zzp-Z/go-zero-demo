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
	result, err := l.roleModel.Insert(l.ctx, &model.Role{
		Name:   in.Name,
		Status: "0",
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("CR362：创建角色失败, err：%v", err)
		return nil, &logic.AppError{
			Code:    "CR362",
			Message: "创建角色失败",
		}
	}
	roleId, err := result.LastInsertId()
	if err != nil {
		logx.WithContext(l.ctx).Errorf("CR363：创建角色成功，查询id失败, err：%v", err)
		return nil, &logic.AppError{
			Code:    "CR363",
			Message: "创建角色成功，查询id失败",
		}
	}
	role, err := l.roleModel.FindOne(l.ctx, uint64(roleId))
	if err != nil {
		logx.WithContext(l.ctx).Errorf("CR364：创建角色成功，查询角色失败, err：%v", err)
		return nil, &logic.AppError{
			Code:    "CR364",
			Message: "创建角色成功，查询角色失败",
		}
	}
	return &user_server.RoleInfoRespVo{
		Id:   roleId,
		Name: role.Name,
	}, nil
}
