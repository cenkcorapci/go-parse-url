package main

import (
	"flag"
	"fmt"
	"github.com/cenkcorapci/go-parse-url/api"
	"github.com/cenkcorapci/go-parse-url/fs"
	"github.com/cenkcorapci/go-parse-url/utils"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

/*
main initiates gRPC and web servers and waits for their termination. If one of them fails, shuts the app down.
*/
func main() {
	signals := make(chan int)
	go startWebServer(signals)
	go startRPCServer(signals)
	var shutdowns int
	select {
	case status := <-signals:
		if status == 0 {
			shutdowns += 1
			if shutdowns >= 2 {
				os.Exit(0)
			}
		} else {
			os.Exit(1)
		}

	}
}

/*
startWebServer starts a http server and waits for connection
*/
func startWebServer(signals chan int) {
	log.Println("Starting web server...")
	addr := fmt.Sprintf(":%s", utils.GetEnv("WEB_SERVER_PORT", "8080"))
	directory := flag.String("d", ".", "the directory of static file to host")

	flag.Parse()
	fileServer := http.FileServer(fs.FileSystem{http.Dir(*directory)})

	http.Handle("/statics/", http.StripPrefix(strings.TrimRight("/statics/", "/"), fileServer))

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Can't start http server! %s", err)
		signals <- 1
	}
	signals <- 0
}

/*
startRPCServer starts a gRPC server and waits for connection
*/
func startRPCServer(signals chan int) {
	log.Println("Starting gRPC server...")
	port := utils.GetEnv("RPC_PORT", "7777")
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
		log.Fatalf("Failed to start gRPC server! %s", err)
		signals <- 1
	}
	signals <- 0
}
