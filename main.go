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

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("error while reading config: %v", err)
	}

	mjClient := client.NewClient(cfg.MJKey)
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
