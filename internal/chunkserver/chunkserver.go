package chunkserver

import (
	"fmt"
	"net/rpc"

	"github.com/distributed-fs/internal"
	"github.com/distributed-fs/internal/rpctype"
)

const (
	masterRegistrationMethod = "Master.ChunkserverRegistration"
)

// Chunkserver is a struct that represents the chunks
type Chunkserver struct {
	files map[string]bool
	*rpctype.RPCServer
}

// NewChunkserver allocates a new instances of a chunkserver and initializes all of its fields
func NewChunkserver() *Chunkserver {
	chunkserver := &Chunkserver{
		make(map[string]bool),
		rpctype.NewRPCServer(),
	}
	rpc.Register(chunkserver)
	return chunkserver
}

// RegisterChunkserver register the given chunkserver to the master server
func RegisterChunkserver(masterAddress string, chunkserverAddress string) error {
	internal.Info("making call to master")
	// Setup necessary arguments and parameters
	args := &rpctype.ChunkserverRegisterRequest{
		ServerAddress: chunkserverAddress,
	}
	var reply rpctype.ChunkserverRegisterResponse

	registerClient, err := rpc.Dial("tcp", masterAddress)
	if err != nil {
		return err
	}

	// Close the TCP connection when done
	defer registerClient.Close()

	err = registerClient.Call(masterRegistrationMethod, args, &reply)
	if err != nil {
		return err
	} else if !reply.Ok {
		errMsg := "rpc call to register chunkserver returned as unsuccessful"
		internal.Failure(errMsg)
		return fmt.Errorf(errMsg)
	}

	return nil
}

// FileIORequest handles request to perform operations on files from client
func FileIORequest(req *rpctype.FileIORequest, res *rpctype.FileIOResponse) error {
	// switch operation {
	// case cmn.Open:
	// 	fmt.
	// case cmn.Close:
	// case cmn.Read:
	// case cmn.Append:
	// case cmn.Delete:
	// case cmn.Snapshot:
	// }
	res.Ok = true
	return nil
}
