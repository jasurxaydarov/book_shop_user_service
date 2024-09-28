package postgres

import (
	"context"

	"github.com/jasurxaydarov/book_shop_user_service/genproto/user_service"
)

type UserRepoI interface {
	CreateUser(ctx context.Context, req *user_service.UserCreateReq) (*user_service.User, error)
	GetUserById(ctx context.Context, req *user_service.GetByIdReq) (*user_service.User, error)
	GetUsers(ctx context.Context, req *user_service.GetListReq) (*user_service.UserGetListResp, error)
	UpdateUser(ctx context.Context, req *user_service.UserUpdateReq) (*user_service.User, error)
	DeleteUser(ctx context.Context, req *user_service.DeleteReq) (string, error)
	IsExists(ctx context.Context, req *user_service.Common) (*user_service.CommonResp, error)
	UserLogin(ctx context.Context, req *user_service.UserLogIn) (*user_service.Clamis, error)
}
