package postgres

import (
	"context"

	"github.com/jasurxaydarov/book_shop/genproto/book_shop"
)


type UserRepoI interface{

	CreateUser(ctx context.Context, req *book_shop.UserCreateReq)(*book_shop.User,error)
	GetUserById(ctx context.Context, req *book_shop.GetByIdReq)(*book_shop.User,error)
	IsExists(ctx context.Context, req *book_shop.Common)(*book_shop.CommonResp,error)
	UserLogin(ctx context.Context, req *book_shop.UserLogIn)(*book_shop.Clamis,error)
}