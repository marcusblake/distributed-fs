package master

import (
	"fmt"
	"net/rpc"

	"github.com/distributed-fs/internal/rpctype"
)

// Master is a struct representing master server
type Master struct {
	namespace *Namespace
	*rpctype.RPCServer
}

// NewMaster creates a new instqnces of the master server
func NewMaster() *Master {
	master := &Master{
		NewNamespace(),
		rpctype.NewRPCServer(),
	}
	rpc.Register(master)
	return master
}

// OperationRequest is the function to be called when a request is made
func (mstr *Master) OperationRequest(req *rpctype.OperationRequest, res *rpctype.OperationResponse) error {

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
	fmt.Println("I've been called!")

	return nil
}

// ChunkserverRegistration is a function which handles registering chunkservers
func (mstr *Master) ChunkserverRegistration(req *rpctype.ChunkserverRegisterRequest,
	res *rpctype.ChunkserverRegisterResponse) error {

	return nil
}
