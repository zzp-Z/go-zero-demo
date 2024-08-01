package userserverlogic

import (
	"context"
	"user_server/internal/logic"
	"user_server/model"

	"user_server/internal/svc"
	"user_server/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	userModel model.UserModel // 模型接口定义为 UserModel
	Tools     *logic.Tools
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		Logger:    logx.WithContext(ctx),
		userModel: model.NewUserModel(svcCtx.SqlConn, svcCtx.CacheConf), // 初始化模型
		Tools:     logic.NewTools(),
	}
}

func (l *LoginLogic) Login(in *user_server.VerificationReqVo) (*user_server.JwtTokenRespVo, error) {
	user, err := l.userModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("LF036：用户不存在 err:%v", err)
		return nil, &logic.AppError{
			Code:    "LF036",
			Message: "用户不存在",
		}
	}
	if user.DeletedAt.Valid {
		logx.WithContext(l.ctx).Errorf("LF035：用户已被删除 err:%v", err)
		return nil, &logic.AppError{
			Code:    "LF035",
			Message: "用户已被删除",
		}

	}
	if !l.Tools.CheckPassword(in.Password, user.Password) {
		logx.WithContext(l.ctx).Errorf("LF037：密码错误 err:%v", err)
		return nil, &logic.AppError{
			Code:    "LF037",
			Message: "密码错误",
		}
	}
	token, err := l.Tools.GenerateJwtToken(user.Id)

	return &user_server.JwtTokenRespVo{
		Token: token,
	}, nil
}
