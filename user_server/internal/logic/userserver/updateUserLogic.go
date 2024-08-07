package userserverlogic

import (
	"context"
	"user_server/internal/logic"
	"user_server/model"

	"user_server/internal/svc"
	"user_server/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	userModel model.UserModel // 模型接口定义为 UserModel
	Tools     *logic.Tools
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		Logger:    logx.WithContext(ctx),
		userModel: model.NewUserModel(svcCtx.SqlConn, svcCtx.CacheConf), // 初始化模型
		Tools:     logic.NewTools(),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *user_server.UpdateUserReqVo) (*user_server.JwtTokenRespVo, error) {
	user, err := l.userModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil {
		return nil, logic.NewAppError(l.ctx, "FC452", "用户不存在", err)
	}
	isPass := l.Tools.CheckPassword(in.Password, user.Password)
	if !isPass {
		return nil, logic.NewAppError(l.ctx, "FC492", "密码错误", nil)
	}
	user.Name = in.NewName
	err = l.userModel.Update(l.ctx, user)
	if err != nil {
		return nil, logic.NewAppError(l.ctx, "FC442", "更新用户失败", err)
	}
	token, err := l.Tools.GenerateJwtToken(user.Id)
	if err != nil {
		return nil, logic.NewAppError(l.ctx, "FC642", "生成Jwt失败", err)
	}
	return &user_server.JwtTokenRespVo{
		Token: token,
	}, nil
}
