package services

import (
	"context"
	"fmt"

	"github.com/jasurxaydarov/book_shop/genproto/book_shop"
	"github.com/jasurxaydarov/book_shop/storage"
)

type UserService struct {
	storage storage.StorageRepoI
	book_shop.UnimplementedUserServiceServer
}

func NewUserService(storage storage.StorageRepoI) *UserService {

	return &UserService{storage: storage}
}

func (u *UserService) CreateUser(ctx context.Context, req *book_shop.UserCreateReq) (*book_shop.User, error) {

	resp, err := u.storage.GetUserRepo().CreateUser(ctx, req)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return resp, nil
}

func (u *UserService) GetUser(ctx context.Context, req *book_shop.GetByIdReq) (*book_shop.User, error) {

	resp, err := u.storage.GetUserRepo().GetUserById(ctx, req)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return resp, nil
}
func (u *UserService) GetUsers(context.Context, *book_shop.GetListReq) (*book_shop.UserGetListResp, error) {

	return nil, nil
}
func (u *UserService) UpdateUser(context.Context, *book_shop.UserUpdateReq) (*book_shop.User, error) {

	return nil, nil
}

func (u *UserService) DeleteUser(context.Context, *book_shop.DeleteReq) (*book_shop.Empty, error) {

	return nil, nil
}

func (u *UserService) CheckExists(ctx context.Context, req *book_shop.Common) (*book_shop.CommonResp, error) {

	resp, err := u.storage.GetUserRepo().IsExists(ctx,req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}
	return resp, nil
}


func (u *UserService)UserLogin(ctx context.Context,req *book_shop.UserLogIn) (*book_shop.Clamis, error){
	
	resp, err := u.storage.GetUserRepo().UserLogin(ctx,req)

	if err != nil {

		fmt.Println(err)
		return nil, err
	}
	return resp, nil
}