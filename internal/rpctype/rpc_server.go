package rpctype

import (
	"fmt"
	"net"
	"net/rpc"
	"sync"
	"time"

	"github.com/distributed-fs/internal"
)

// RPCServer is a struct that will act as an RpcServer for our distrubuted file system
type RPCServer struct {
	shutdown chan bool
	ln       net.Listener
	wg       sync.WaitGroup
}

// NewRPCServer creates allocates and initializes an instance of RpcServer
func NewRPCServer() *RPCServer {
	srv := new(RPCServer)
	srv.shutdown = make(chan bool)
	srv.wg = sync.WaitGroup{}
	return srv
}

// Start starts the rpc sever using the address specified
func (srv *RPCServer) Start(address string) error {
	var err error
	srv.ln, err = net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("Error starting RPC server with address %v. Received the following error: %v", address, err)
	}

	go func() {
		for {
			conn, err := srv.ln.Accept()
			if err != nil {
				select {
				case <-srv.shutdown:
					fmt.Println("shutting down!")
					return
				default:
					internal.Warning(err.Error())
					continue
				}
			}

			srv.wg.Add(1)
			go func() {
				rpc.ServeConn(conn)
				srv.wg.Done()
			}()
		}
	}()

	internal.Success("server started!")

	return nil
}

// Stop function stops the RPCServer from running
func (srv *RPCServer) Stop() error {
	c := make(chan struct{})

	// Send signal to shutdown and wait until all running tasks are completed
	go func() {
		srv.shutdown <- true
		srv.wg.Wait()
		c <- struct{}{}
	}()

	srv.ln.Close()

	select {
	case <-c:
	case <-time.After(10 * time.Second):
		return fmt.Errorf("RPCServer failed to shut off after 10 seconds")
	}
	return nil
}
