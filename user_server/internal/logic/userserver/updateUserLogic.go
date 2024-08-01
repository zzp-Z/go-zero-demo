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
		logx.WithContext(l.ctx).Errorf("FC452：用户不存在 err:%v", err)
		return nil, &logic.AppError{
			Code:    "FC452",
			Message: "用户不存在",
		}
	}
	isPass := l.Tools.CheckPassword(in.Password, user.Password)
	if !isPass {
		logx.WithContext(l.ctx).Errorf("FC492：密码错误 err:%v", err)
		return nil, &logic.AppError{
			Code:    "FC492",
			Message: "密码错误",
		}
	}
	user.Name = in.NewName
	err = l.userModel.Update(l.ctx, user)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("FC442：更新用户失败 err:%v", err)
		return nil, &logic.AppError{
			Code:    "FC442",
			Message: "更新用户失败",
		}
	}
	token, err := l.Tools.GenerateJwtToken(user.Id)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("FC642：生成Jwt失败 err:%v", err)
		return nil, &logic.AppError{
			Code:    "FC642",
			Message: "生成Jwt失败",
		}
	}
	return &user_server.JwtTokenRespVo{
		Token: token,
	}, nil
}
