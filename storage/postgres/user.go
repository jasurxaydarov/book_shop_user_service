package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/book_shop_user_service/genproto/user_service"
	"github.com/jasurxaydarov/book_shop_user_service/pkg/db/helpers"
	"github.com/saidamir98/udevs_pkg/logger"
)

type UserRepo struct {
	db  *pgx.Conn
	log logger.LoggerI
}

func NewUserRepo(db *pgx.Conn, log logger.LoggerI) UserRepoI {

	return &UserRepo{db: db, log: log}
}

func (u *UserRepo) CreateUser(ctx context.Context, req *user_service.UserCreateReq) (*user_service.User, error) {
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

	resp, err := u.GetUserById(context.Background(), &user_service.GetByIdReq{Id: id.String()})

	if err != nil {

		u.log.Error("err on db GetUserById", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (u *UserRepo) GetUserById(ctx context.Context, req *user_service.GetByIdReq) (*user_service.User, error) {

	var resp user_service.User
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

func (u *UserRepo) IsExists(ctx context.Context, req *user_service.Common) (*user_service.CommonResp, error) {
	var isExists bool

	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM %s WHERE %s = '%s')", req.TableName, req.ColumnName, req.Expvalue)

	err := u.db.QueryRow(ctx, query).Scan(&isExists)

	if err != nil {
		u.log.Error("error on CheckExists", logger.Error(err))
		return &user_service.CommonResp{IsExists: false}, nil
	}

	return &user_service.CommonResp{IsExists: isExists}, nil
}

func (u *UserRepo) UserLogin(ctx context.Context, req *user_service.UserLogIn) (*user_service.Clamis, error) {

	var viwerId, gmail, hashPassword, userRole string

	query := `
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

	err := u.db.QueryRow(ctx, query, req.Username).Scan(&viwerId, &gmail, &hashPassword, &userRole)

	if err != nil {
		return nil, err
	}

	if !helpers.CompareHashPassword(hashPassword, req.Password) {
		return nil, errors.New("password is incorrect")
	}

	return &user_service.Clamis{UserId: viwerId, UserRole: userRole}, nil

}
