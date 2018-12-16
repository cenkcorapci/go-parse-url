package main

import (
	"fmt"
	"github.com/cenkcorapci/go-parse-url/api"
	"github.com/cenkcorapci/go-parse-url/utils"
	"google.golang.org/grpc"
	"log"
	"net"
)

// main start a gRPC server and waits for connection
func main() {
	port := utils.GetEnv("PORT", "7777")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// create a server instance
	s := api.Server{}
	// create a gRPC server object
	gRPCServer := grpc.NewServer()
	// attach the Ping service to the server
	api.RegisterPingServer(gRPCServer, &s)
	// start the server
	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
