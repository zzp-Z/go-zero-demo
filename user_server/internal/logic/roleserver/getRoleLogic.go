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
	role, err := l.roleModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		logx.WithContext(l.ctx).Errorf("GR301：角色未找到, err: %v", err)
		return nil, &logic.AppError{
			Code:    "GR301",
			Message: "角色未找到",
		}
	}

	return &user_server.RoleInfoRespVo{
		Id:   in.Id,
		Name: role.Name,
	}, nil
}
