package main

import (
	"log"
	"net"
	"os"
	"strconv"

	"github.com/kanguki/go-grpc-mysql/internal/core-variant/db"
	"github.com/kanguki/go-grpc-mysql/internal/core-variant/movie"
	m "github.com/kanguki/go-grpc-mysql/internal/core/movie"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	servicePort := os.Getenv("SERVICE_PORT")
	if servicePort == "" {
		servicePort = ":2381"
	} else {
		_, err := strconv.Atoi(servicePort)
		if err != nil {
			log.Fatalln("Invalid service port")
		} else {
			servicePort = ":" + servicePort
		}
	}
	mysqlConnectStr := os.Getenv("MYSQL_URL")
	if mysqlConnectStr == "" {
		log.Fatalln("Empty mysql url")
	}
	database := &db.Mysql{
		ConnectString: mysqlConnectStr,
	}
	database.Connect()
	movieService := movie.Netflix{Db: database}

	lis, err := net.Listen("tcp", servicePort)
	if err != nil {
		log.Fatalf("Fail to listen to port %v: %v", servicePort, err)
	}

	log.Printf("Serving grpc server at localhost%s\n", servicePort)
	grpcServer := grpc.NewServer()
	m.RegisterMovieServiceServer(grpcServer, &movieService)
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Fail to serve gRpc server: %v", err)
	}
}
