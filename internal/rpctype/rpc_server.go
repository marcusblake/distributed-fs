package rpctype

import (
	"fmt"
	"net"
	"net/rpc"
	"time"

	"github.com/distributed-fs/internal"
)

// RPCServer is a struct that will act as an RpcServer for our distrubuted file system
type RPCServer struct {
	server  *rpc.Server
	on      bool
	turnOff chan bool
}

// NewRPCServer creates allocates and initializes an instance of RpcServer
func NewRPCServer() *RPCServer {
	srv := new(RPCServer)
	srv.server = rpc.NewServer()
	srv.turnOff = make(chan bool)
	return srv
}

// Start starts the rpc sever using the address specified
func (srv *RPCServer) Start(address string) error {

	ln, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("Error starting RPC server with address %v. Received the following error: %v", address, err)
	}

	go func() {
		srv.on = true
		for srv.on {
			conn, err := ln.Accept()
			if err != nil {
				internal.Warning(err.Error())
				continue
			}
			go srv.server.ServeConn(conn)
		}
		srv.turnOff <- true
	}()

	return nil
}

func (srv *RPCServer) Stop() error {
	srv.on = false
	select {
	case <-srv.turnOff:
	case <-time.After(10 * time.Second):
		return fmt.Errorf("RPCServer failed to shut off after 10 seconds")
	}
	return nil
}

func (srv *RPCServer) RegisterFunction() {

}
