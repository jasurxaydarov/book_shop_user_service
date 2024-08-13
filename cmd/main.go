package main

import (
	"context"
	"fmt"
	"net"

	"github.com/jasurxaydarov/book_shop/genproto/book_shop"
	"github.com/jasurxaydarov/book_shop/pkg/db"
	"github.com/jasurxaydarov/book_shop/services"
	"github.com/jasurxaydarov/book_shop/storage"
	"github.com/saidamir98/udevs_pkg/logger"
	"google.golang.org/grpc"
)

func main() {
	log := logger.NewLogger("", logger.LevelDebug)
	PgxConn, err := db.ConnDB(context.Background())

	if err != nil {
		log.Error(err.Error())
		return
	}
	fmt.Println(PgxConn)

	storage:=storage.NewUserRepo(PgxConn,log)
	fmt.Println(storage)

	service:=services.NewUserService(storage)

	listen,err:=net.Listen("tcp","localhost:8000")


	if err != nil {
		log.Error(err.Error())
		return
	}

	server:=grpc.NewServer()

	book_shop.RegisterUserServiceServer(server,service)

	log.Debug("server serve on :8000")

	if err =server.Serve(listen);err!=nil{
		log.Error(err.Error())
		return 

	}
}
