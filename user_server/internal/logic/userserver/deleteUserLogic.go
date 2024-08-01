package userserverlogic

import (
	"context"
	"user_server/internal/logic"
	"user_server/model"

	"user_server/internal/svc"
	"user_server/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	userModel model.UserModel // 模型接口定义为 UserModel
	Tools     *logic.Tools
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		Logger:    logx.WithContext(ctx),
		userModel: model.NewUserModel(svcCtx.SqlConn, svcCtx.CacheConf), // 初始化模型
		Tools:     logic.NewTools(),
	}
}

func (l *DeleteUserLogic) DeleteUser(in *user_server.VerificationReqVo) (*user_server.JwtTokenRespVo, error) {
	user, err := l.userModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("DU036：账号未找到 err:%v", err)
		return nil, &logic.AppError{
			Code:    "DU036",
			Message: "账号未找到",
		}
	}
	if user.DeletedAt.Valid {
		logx.WithContext(l.ctx).Errorf("DU035：用户已被删除 err:%v", err)
		return nil, &logic.AppError{
			Code:    "DU035",
			Message: "用户已被删除",
		}

	}
	if !l.Tools.CheckPassword(in.Password, user.Password) {
		logx.WithContext(l.ctx).Errorf("DU037：密码错误 err:%v", err)
		return nil, &logic.AppError{
			Code:    "DU037",
			Message: "密码错误",
		}
	}
	err = l.userModel.Delete(l.ctx, user.Id)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("DU038：删除用户失败 err:%v", err)
		return nil, &logic.AppError{
			Code:    "DU038",
			Message: "删除用户失败",
		}
	}
	token, err := l.Tools.GenerateJwtToken(user.Id)

	return &user_server.JwtTokenRespVo{
		Token: token,
	}, nil
}
