package main

import (
	"log"

	"github.com/SaidovZohid/medium_api_gateway/api"
	_ "github.com/SaidovZohid/medium_api_gateway/api/docs"
	"github.com/SaidovZohid/medium_api_gateway/config"
	grpcPkg "github.com/SaidovZohid/medium_api_gateway/pkg/grpc_client"
	"github.com/SaidovZohid/medium_api_gateway/pkg/logger"
)

func main() {
	cfg := config.Load(".")

	grpcConn, err := grpcPkg.New(cfg)
	if err != nil {
		log.Fatalf("failed to get grpc connettion: %v", err)
	}

	logger := logger.New()

	apiServer := api.New(&api.RoutetOptions{
		Cfg:        &cfg,
		GrpcClient: grpcConn,
		Logger:     logger,
	})

	err = apiServer.Run(cfg.HttpPort)
	if err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
