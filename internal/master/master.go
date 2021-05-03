package master

import (
	"net/rpc"
	"time"

	"github.com/distributed-fs/internal/rpctype"
	"github.com/distributed-fs/internal/security"
	"github.com/distributed-fs/pkg/common"
	"github.com/distributed-fs/pkg/logger"
)

// Master is a struct representing master server
type Master struct {
	namespace    *Namespace
	Chunkservers map[string]ChunkserverInfo
	*rpctype.RPCServer
}

// ChunkserverInfo stores information regarding the state and properties of the server
type ChunkserverInfo struct {
	LastHeartbeat time.Time
	Capacity      uint64
	Files         map[string]bool
}

// NewMaster creates a new instqnces of the master server
func NewMaster() *Master {
	master := &Master{
		NewNamespace(),
		make(map[string]ChunkserverInfo),
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

	appId := req.ApplicationId
	filename := req.Filename

	var permissions common.PermissionType = 0
	if file, ok := mstr.namespace.GetFileInformation(filename); ok {
		appIdAsString := appId.String()
		if file.owner == appIdAsString {
			permissions = file.permissions[common.GroupPermissions.Application]
		} else if file.group[appIdAsString] {
			permissions = file.permissions[common.GroupPermissions.ApplicationGroup]
		} else {
			permissions = file.permissions[common.GroupPermissions.All]
		}
	}

	token, err := security.CreateToken(appId, filename, permissions)
	if err != nil {
		res.Ok = false
	}

	res.Token = token
	res.Ok = true
	return nil
}

// ChunkserverRegistration is a function which handles registering chunkservers
func (mstr *Master) ChunkserverRegistration(req *rpctype.ChunkserverRegisterRequest,
	res *rpctype.ChunkserverRegisterResponse) error {
	chunkserverAddr := req.ServerAddress
	info := ChunkserverInfo{
		time.Now(),
		0,
		make(map[string]bool),
	}
	mstr.Chunkservers[chunkserverAddr] = info
	logger.Successf("successfully registered chunkserver %v", chunkserverAddr)
	res.Ok = true
	return nil
}
