package roleserverlogic

import (
	"context"
	"user_server/internal/logic"
	"user_server/model"

	"user_server/internal/svc"
	"user_server/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	roleModel model.RoleModel
}

func NewGetRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleListLogic {
	return &GetRoleListLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		Logger:    logx.WithContext(ctx),
		roleModel: model.NewRoleModel(svcCtx.SqlConn, svcCtx.CacheConf),
	}
}

// GetRoleList 获取角色列表
func (l *GetRoleListLogic) GetRoleList(in *user_server.Empty) (*user_server.RoleInfoListRespVo, error) {
	roles, err := l.roleModel.GetAllRoles(l.ctx)
	if err != nil {
		// 记录错误日志
		logx.WithContext(l.ctx).Errorf("GAR301：查询角色列表失败，err: %v", err)
		return nil, &logic.AppError{
			Code:    "GAR301",
			Message: "查询角色列表失败",
		}
	}

	// 预分配切片容量
	roleList := make([]*user_server.RoleInfoRespVo, 0, len(roles))
	for _, role := range roles {
		roleList = append(roleList, &user_server.RoleInfoRespVo{
			Id:   int64(role.Id),
			Name: role.Name,
		})
	}

	return &user_server.RoleInfoListRespVo{
		RoleList: roleList,
	}, nil
}
