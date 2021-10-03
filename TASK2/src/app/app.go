package app

import (
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"github.com/rampo0/bibit/layer/application/services"
	"github.com/rampo0/bibit/layer/infrastructure/db"
	"github.com/rampo0/bibit/layer/infrastructure/rest"
	"github.com/rampo0/bibit/layer/presentation/controllers"
	"github.com/rampo0/bibit/layer/presentation/rpc"
	"github.com/rampo0/bibit/layer/presentation/rpc/moviepb"
	"google.golang.org/grpc"
)

var (
	router       = gin.Default()
	movieService = services.NewMovieService(rest.NewMovieRepository(), db.NewLoggerRepository())
)

func runHttpServer() {
	// REST
	movieHandler := controllers.NewMovieHandler(movieService)

	router.GET("/detail", movieHandler.Detail)
	router.GET("/search", movieHandler.Search)

	if err := router.Run(); err != nil {
		log.Fatal("Failed to listen: %v", err)
	}
}

func runGRPCServer() {
	//gRPC
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal("Failed to listen: %v", err)
	}

	defer lis.Close()

	tls := false
	opts := []grpc.ServerOption{}

	if tls {
	}

	gRPCServer := grpc.NewServer(opts...)
	moviepb.RegisterMovieServiceServer(gRPCServer, rpc.NewMovieHandler(movieService))

	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatal("Failed to server : %v", err)
	}
}

func Run() {
	go runHttpServer()
	go runGRPCServer()

	// wait for control + c to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// block until a signal received
	<-ch
}
