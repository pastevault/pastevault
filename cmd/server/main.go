package main

import (
	"net"

	"google.golang.org/grpc"
	"pastevault.com/internal/db"
	"pastevault.com/internal/expired"
	. "pastevault.com/internal/logger"
	"pastevault.com/internal/paste"
	pb "pastevault.com/internal/proto"
	"pastevault.com/internal/router"
)

type server struct {
	pb.UnimplementedServiceServer
}

func main() {
	NewLogger()
	err := db.Connect()
	if err != nil {
		Logger.Error("Error connecting to the database: ", "db.Connect", err)
		panic(err)
	}
	Logger.Info("Connected to the database")

	err = paste.Migrate()
	if err != nil {
		Logger.Error("Error migrating the database: ", "paste.Migrate", err)
		panic(err)
	}
	Logger.Info("Migrated the database")

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		Logger.Error("Failed to listen on gRPC port: ", "net.Listen", err)
		panic(err)
	}

	s := grpc.NewServer()

	pb.RegisterServiceServer(s, &server{})
	go func() {
		Logger.Info("gRPC server listening on", "port", "8081")
		if err := s.Serve(lis); err != nil {
			Logger.Error("Failed to serve gRPC server: ", "s.Serve", err)
			panic(err)
		}
	}()

	go expired.Delete()
	go router.Router()

	select {}
}
