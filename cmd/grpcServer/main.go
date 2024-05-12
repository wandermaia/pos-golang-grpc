package main

import (
	"database/sql"
	"net"

	_ "github.com/mattn/go-sqlite3"
	"github.com/wandermaia/pos-golang-grpc/internal/database"
	"github.com/wandermaia/pos-golang-grpc/internal/pb"
	"github.com/wandermaia/pos-golang-grpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := database.NewCategory(db)
	catgoryService := service.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, catgoryService)
	reflection.Register(grpcServer) // Necessário para utilização com o Evans

	lis, err := net.Listen("tcp", ":50051") // Porta padrão do grpc
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}

}
