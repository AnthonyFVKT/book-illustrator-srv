package main

import (
	"github.com/AnthonyFVKT/book-illustrator-srv/internal/client"
	"github.com/AnthonyFVKT/book-illustrator-srv/internal/config"
	handler "github.com/AnthonyFVKT/book-illustrator-srv/internal/rpc"
	"github.com/AnthonyFVKT/book-illustrator-srv/internal/service"
	pb "github.com/AnthonyFVKT/book-illustrator-srv/proto/illustrator"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	mjKey = "d7c73d8cc6a8341dc21198c78b663942ea7c207e2b9d12b95c71f1ca2d943c8c"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("error while reading config: %v", err)
	}

	mjClient := client.NewClient(mjKey)
	illustrator := service.NewIllustrator(mjClient)
	textProcessor := service.NewProcessor(cfg.GroupMaxSentences)

	illustartorHandler := handler.NewIllustrator(illustrator, textProcessor)

	lis, err := net.Listen("tcp", cfg.ServerAddress)
	if err != nil {
		log.Fatalf("error while listening port: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterIllustratorServiceServer(server, illustartorHandler)

	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("error while serving: %v", err)
	}
}
