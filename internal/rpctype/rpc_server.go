package rpctype

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/distributed-fs/pkg/logger"
)

// RPCServer is a struct that will act as an RpcServer for our distrubuted file system
type RPCServer struct {
	shutdown chan os.Signal
	server   http.Server
	ln       net.Listener
}

// NewRPCServer creates allocates and initializes an instance of RpcServer
func NewRPCServer() *RPCServer {
	srv := new(RPCServer)
	srv.shutdown = make(chan os.Signal, 1)
	signal.Notify(srv.shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	srv.server = http.Server{}
	rpc.HandleHTTP()
	return srv
}

// Start starts the rpc sever using the address specified
func (srv *RPCServer) Start(address string) error {
	var err error
	srv.ln, err = net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("error starting RPC server with address %v. Received the following error: %v", address, err)
	}

	go func() {
		go func() {
			if err := srv.server.Serve(srv.ln); err != nil {
				log.Fatal("http server couldn't be started")
			}
		}()

		<-srv.shutdown

		shutdownWithTime, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer func() {
			cancel()
		}()

		if err := srv.server.Shutdown(shutdownWithTime); err != nil {
			logger.Info("http server shutdown")
		}

		os.Exit(0)
	}()

	return nil
}
