package userserverlogic

import (
	"context"
	"user_server/internal/logic"
	"user_server/model"

	"user_server/internal/svc"
	"user_server/user_server"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	userModel model.UserModel // 模型接口定义为 UserModel
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		Logger:    logx.WithContext(ctx),
		userModel: model.NewUserModel(svcCtx.SqlConn, svcCtx.CacheConf), // 初始化模型
	}
}

func (l *GetUserLogic) GetUser(in *user_server.UserIdReqVo) (*user_server.UserInfoRespVo, error) {
	user, err := l.userModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		logx.WithContext(l.ctx).Errorf("GU301：用户不存在 err:%v", err)
		return nil, &logic.AppError{
			Code:    "GU301",
			Message: "用户不存在",
		}
	}

	return &user_server.UserInfoRespVo{
		Id:    in.Id,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
