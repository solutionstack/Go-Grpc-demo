package main

import (
	logger "github.com/rs/zerolog"
	"os"
	handler "proto_test/handler"
	rpc_client "proto_test/rpc/client"
	"proto_test/server"
)

const (
	grpcAddress = "0.0.0.0:9000"
)

var Log = logger.New(os.Stdout)

func main() {
	//connect grpc
	userRPCClientManager := rpc_client.NewUserClientManager()
	rpc_conn, err := userRPCClientManager.Connect()
	if err != nil {
		Log.Fatal().Err(err).Msg("RPC client connection failed")
	}
	handler := handler.NewHandler(*(userRPCClientManager.GetClient()))

	//start local server
	if err := server.StartNew(handler); err != nil {
		Log.Fatal().Err(err).Msg("Failed starting local server")
	}
	defer rpc_conn.Close()
}
