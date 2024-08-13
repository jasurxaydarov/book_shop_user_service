package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/book_shop_user_service/config"
)

func ConnDB(ctx context.Context)(*pgx.Conn,error){

	PgxConn,err:=pgx.Connect(ctx,config.GetPgURL())

	if err!= nil{

		log.Println(err)
		return nil,err
	}
	return PgxConn,err
}
