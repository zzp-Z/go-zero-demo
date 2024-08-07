package roleserverlogic

import (
	"context"
	"user_server/internal/logic"
	"user_server/model"

	"user_server/internal/svc"
	"user_server/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	roleModel model.RoleModel // 角色model
	tools     *logic.Tools
}

func NewDeleteRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoleLogic {
	return &DeleteRoleLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		Logger:    logx.WithContext(ctx),
		roleModel: model.NewRoleModel(svcCtx.SqlConn, svcCtx.CacheConf), // 初始化角色model
		tools:     logic.NewTools(),
	}
}

func (l *DeleteRoleLogic) DeleteRole(in *user_server.RoleIdReqVo) (*user_server.RoleInfoRespVo, error) {
	/*
		1. 查找角色
		2. 删除角色
		3. 返回角色信息
	*/
	role, err := l.roleModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		return nil, logic.NewAppError(l.ctx, "DR304", "未找到该角色", err)
	}
	role.DeletedAt = l.tools.GetNowTime()
	err = l.roleModel.Update(l.ctx, role)
	if err != nil {
		return nil, logic.NewAppError(l.ctx, "DR305", "删除角色失败", err)
	}

	return &user_server.RoleInfoRespVo{
		Id:   in.Id,
		Name: role.Name,
	}, nil
}
