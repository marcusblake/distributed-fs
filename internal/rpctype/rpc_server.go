package rpctype

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
)

const (
	defaultTimeout = time.Duration(1 * time.Minute)
)

var (
	defaultServerOptions = []grpc.ServerOption{
		grpc.ConnectionTimeout(defaultTimeout),
	}
)

// RPCServer is a struct that will act as an RpcServer for our distrubuted file system
type RPCServer struct {
	shutdown chan os.Signal
	Server   *grpc.Server
	ln       net.Listener
}

// NewRPCServer creates allocates and initializes an instance of RpcServer
func NewRPCServer(serverOptions []grpc.ServerOption) *RPCServer {
	srv := new(RPCServer)
	srv.shutdown = make(chan os.Signal, 1)
	signal.Notify(srv.shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	if len(serverOptions) == 0 {
		serverOptions = defaultServerOptions
	}

	srv.Server = grpc.NewServer(serverOptions...)
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
			if err := srv.Server.Serve(srv.ln); err != nil {
				log.Fatal("http server couldn't be started")
			}
		}()

		<-srv.shutdown

		srv.Server.Stop()

		os.Exit(0)
	}()

	return nil
}
