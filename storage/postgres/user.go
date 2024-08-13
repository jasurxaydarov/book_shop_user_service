package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/book_shop/genproto/book_shop"
	"github.com/jasurxaydarov/book_shop/pkg/db/helpers"
	"github.com/saidamir98/udevs_pkg/logger"
)

type UserRepo struct {
	db  *pgx.Conn
	log logger.LoggerI
}

func NewUserRepo(db *pgx.Conn, log logger.LoggerI) UserRepoI {

	return &UserRepo{db: db, log: log}
}

func (u *UserRepo) CreateUser(ctx context.Context, req *book_shop.UserCreateReq) (*book_shop.User, error) {
	u.log.Debug("aaaaaaaaaaaaaaaaa")
	id := uuid.New()
	query := `
		INSERT INTO
			users (
				user_id,
				username,
				email, 
				password,
				full_name,
				user_role
			)VALUES(
				$1,$2,$3,$4,$5,$6
			)
			`

	_, err := u.db.Exec(
		ctx,
		query,
		id,
		req.Username,
		req.Email,
		req.Password,
		req.Fullname,
		req.UserRole,
	)
	if err != nil {

		u.log.Debug("errrrrr")
		u.log.Error("err on db CreateUser", logger.Error(err))
		return nil, err
	}

	resp, err := u.GetUserById(context.Background(), &book_shop.GetByIdReq{Id: id.String()})

	if err != nil {

		u.log.Error("err on db GetUserById", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (u *UserRepo) GetUserById(ctx context.Context, req *book_shop.GetByIdReq) (*book_shop.User, error) {

	var resp book_shop.User
	qury := `
		SELECT 
			user_id,
				username,
				email, 
				password,
				full_name,
				user_role
		FROM 
			users 
		WHERE
			user_id = $1
	`

	err := u.db.QueryRow(
		ctx,
		qury,
		req.Id,
	).Scan(
		&resp.UserId,
		&resp.Username,
		&resp.Email,
		&resp.Password,
		&resp.Fullname,
		&resp.UserRole,
	)

	if err != nil {

		u.log.Error("err on db GetUserById", logger.Error(err))
		return nil, err
	}

	return &resp, nil
}

func (u *UserRepo) IsExists(ctx context.Context, req *book_shop.Common) (*book_shop.CommonResp, error) {
	var isExists bool

	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM %s WHERE %s = '%s')", req.TableName, req.ColumnName, req.Expvalue)

	err := u.db.QueryRow(ctx, query).Scan(&isExists)

	if err != nil {
		u.log.Error("error on CheckExists", logger.Error(err))
		return &book_shop.CommonResp{IsExists: false}, nil
	}

	return &book_shop.CommonResp{IsExists: isExists}, nil
}



func (u *UserRepo)UserLogin(ctx context.Context, req *book_shop.UserLogIn)(*book_shop.Clamis,error){

	var viwerId,gmail,hashPassword,userRole string

	query:=`
		SELECT 
			user_id,
			email,
			password,
			user_role
		FROM
			users
		WHERE	
			username =$1
	`

	err:=u.db.QueryRow(ctx,query,req.Username).Scan(&viwerId,&gmail,&hashPassword,&userRole)

	if err != nil{
		return nil,err
	}

	if !helpers.CompareHashPassword(hashPassword,req.Password){
		return nil,errors.New("password is incorrect")
	}


	return &book_shop.Clamis{UserId: viwerId,UserRole:userRole}, nil

}
