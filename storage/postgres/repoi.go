package postgres

import (
	"context"

	"github.com/jasurxaydarov/book_shop_user_service/genproto/user_service"
)


type UserRepoI interface{

	CreateUser(ctx context.Context, req *user_service.UserCreateReq)(*user_service.User,error)
	GetUserById(ctx context.Context, req *user_service.GetByIdReq)(*user_service.User,error)
	IsExists(ctx context.Context, req *user_service.Common)(*user_service.CommonResp,error)
	UserLogin(ctx context.Context, req *user_service.UserLogIn)(*user_service.Clamis,error)
}