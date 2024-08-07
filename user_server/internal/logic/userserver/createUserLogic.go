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
	Ctx    context.Context
	SvcCtx *svc.ServiceContext
	logx.Logger
	UserModel model.UserModel // 模型接口定义为 UserModel
	Tools     *logic.Tools
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Ctx:       ctx,
		SvcCtx:    svcCtx,
		Logger:    logx.WithContext(ctx),
		UserModel: model.NewUserModel(svcCtx.SqlConn, svcCtx.CacheConf), // 初始化模型
		Tools:     logic.NewTools(),
	}
}

func (l *CreateUserLogic) CreateUser(in *user_server.CreateUserReqVo) (*user_server.JwtTokenRespVo, error) {
	/*
		1. 验证密码长度
		2. 使用工具类生成哈希密码
		3. 插入用户
		4. 生成JWT Token
	*/
	// 如果密码为空，或长度小于六位数，则返回错误
	if len(in.Password) < 6 {
		return nil, logic.NewAppError(l.Ctx, "TX401", "密码长度小于六位", nil)
	}
	// 使用工具类生成哈希密码
	hashedPassword, err := l.Tools.HashPassword(in.Password)
	if err != nil {
		return nil, logic.NewAppError(l.Ctx, "TX982", "哈希密码失败", err)

	}
	// 插入用户
	userResult, err := l.UserModel.Insert(l.Ctx, &model.User{
		Name:     in.Name,
		Email:    in.Email,
		Password: hashedPassword,
		Status:   "0",
	})
	if err != nil {
		return nil, logic.NewAppError(l.Ctx, "TX983", "创建用户失败", err)
	}
	// 获取插入记录的ID
	userID, err := userResult.LastInsertId()
	if err != nil {
		return nil, logic.NewAppError(l.Ctx, "TX989", "创建成功，生成Token失败", err)
	}
	// 生成JWT Token
	token, err := l.Tools.GenerateJwtToken(uint64(userID))
	if err != nil {
		return nil, logic.NewAppError(l.Ctx, "TX990", "创建成功，生成Token失败", err)
	}
	return &user_server.JwtTokenRespVo{
		Token: token,
	}, nil
}
