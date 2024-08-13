package storage

import (
	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/book_shop/storage/postgres"
	"github.com/saidamir98/udevs_pkg/logger"
)

type StorageRepoI interface{
	GetUserRepo()postgres.UserRepoI
}

type storageRepo struct{
	userRepo postgres.UserRepoI
}

func NewUserRepo(db *pgx.Conn,log logger.LoggerI)StorageRepoI{

	return &storageRepo{userRepo: postgres.NewUserRepo(db,log)}
}


func (s *storageRepo)GetUserRepo()postgres.UserRepoI{

	return s.userRepo
}
