package userserverlogic

import (
	"context"
	"user_server/internal/logic"
	"user_server/internal/svc"
	"user_server/model"
	"user_server/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	userModel model.UserModel // 模型接口定义为 UserModel
	Tools     *logic.Tools
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		Logger:    logx.WithContext(ctx),
		userModel: model.NewUserModel(svcCtx.SqlConn, svcCtx.CacheConf), // 初始化模型
		Tools:     logic.NewTools(),
	}
}

func (l *CreateUserLogic) CreateUser(in *user_server.CreateUserReqVo) (*user_server.JwtTokenRespVo, error) {

	// 如果密码为空，或长度小于六位数，则返回错误
	if len(in.Password) < 6 {
		return nil, &logic.AppError{
			Code:    "TX401",
			Message: "无效密码",
		}
	}
	// 使用工具类生成哈希密码
	hashedPassword, err := l.Tools.HashPassword(in.Password)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("TX982：哈希密码失败, err：%v", err)
		return nil, &logic.AppError{
			Code:    "TX982",
			Message: "哈希密码失败",
		}
	}
	// 插入用户
	userResult, err := l.userModel.Insert(l.ctx, &model.User{
		Name:     in.Name,
		Email:    in.Email,
		Password: hashedPassword,
		Status:   "0",
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("TX983：创建用户失败, err：%v", err)
		return nil, &logic.AppError{
			Code:    "TX983",
			Message: "创建用户失败",
		}
	}
	// 获取插入记录的ID
	userID, err := userResult.LastInsertId()
	if err != nil {
		logx.WithContext(l.ctx).Errorf("TX989：创建成功，生成Token失败, err：%v", err)
		return nil, &logic.AppError{
			Code:    "TX989",
			Message: "创建成功，生成Token失败",
		}
	}
	// 生成JWT Token
	token, err := l.Tools.GenerateJwtToken(uint64(userID))
	if err != nil {
		logx.WithContext(l.ctx).Errorf("TX990：创建成功，生成Token失败, err：%v", err)
		return nil, &logic.AppError{
			Code:    "TX990",
			Message: "创建成功，生成Token失败",
		}
	}
	return &user_server.JwtTokenRespVo{
		Token: token,
	}, nil
}
