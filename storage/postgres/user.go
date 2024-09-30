package postgres

import (
	"context"
	"errors"
	"fmt"
	"time"

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
			user_role,
			created_at,
			updated_at
		FROM 
			users 
		WHERE
			user_id = $1 
		AND
			deleted_at is null
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
		&resp.CreatedAt,
		&resp.UpdatedAt,
	)

	if err != nil {

		u.log.Error("err on db GetUserById", logger.Error(err))
		return nil, err
	}

	return &resp, nil
}

func (u *UserRepo) GetUsers(ctx context.Context, req *user_service.GetListReq) (*user_service.UserGetListResp, error) {

	offset := (req.Page - 1) * req.Limit

	var resp user_service.User

	var res user_service.UserGetListResp
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
    		deleted_at IS NULL
		LIMIT $1 OFFSET $2;
	`

	row, err := u.db.Query(
		ctx,
		qury,
		req.Limit,
		offset,
	)

	if err != nil {

		u.log.Error("err on db GetUsers", logger.Error(err))
		return nil, err
	}

	for row.Next() {

		row.Scan(
			&resp.UserId,
			&resp.Username,
			&resp.Email,
			&resp.Password,
			&resp.Fullname,
			&resp.UserRole,
		)
		if err != nil {

			u.log.Error("err on db GetUsers", logger.Error(err))
			return nil, err
		}
		res.Count++
		res.Users = append(res.Users, &resp)

	}

	return &res, nil
}

func (u *UserRepo) UpdateUser(ctx context.Context, req *user_service.UserUpdateReq) (*user_service.User, error) {

	time := time.Now()
	query := `
		UPDATE
			users
		SET
				username = $1,
				email = $2, 
				password = $3,
				full_name = $4,
				user_role = $5,
				updated_at = $6
		WHERE 
				user_id = $7 
		AND  
				deleted_at is null
			`

	_, err := u.db.Exec(
		ctx,
		query,

		req.Username,
		req.Email,
		req.Password,
		req.FullName,
		req.UserRole,
		time,
		req.UserId,
	)
	if err != nil {

		u.log.Error("err on db update", logger.Error(err))
		return nil, err
	}

	resp, err := u.GetUserById(context.Background(), &user_service.GetByIdReq{Id: req.UserId})

	if err != nil {

		u.log.Error("err on db GetUserById", logger.Error(err))
		return nil, err
	}

	return resp, nil
}

func (u *UserRepo) DeleteUser(ctx context.Context, req *user_service.DeleteReq) (string, error) {

	delete := time.Now()
	query := `
		UPDATE
			users
		SET
			deleted_at = $1
		WHERE user_id = $2
			`

	_, err := u.db.Exec(
		ctx,
		query,
		delete,
		req.Id,
	)
	if err != nil {

		u.log.Error("err on db CreateUser", logger.Error(err))
		return "", err
	}

	return "succesfully deleted", nil
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

	up := `
		UPDATE
			users
		SET
			deleted_at =Null
		WHERE
			user_id = $1
			`

	_, err = u.db.Exec(
		ctx,
		up,
		viwerId,
	)

	if err != nil {
		return nil, err
	}

	if !helpers.CompareHashPassword(hashPassword, req.Password) {
		return nil, errors.New("password is incorrect")
	}

	return &user_service.Clamis{UserId: viwerId, UserRole: userRole}, nil

}
