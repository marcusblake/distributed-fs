package master

import (
	"context"
	"errors"
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
	UnimplementedMasterServer
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
		rpctype.NewRPCServer(nil),
		UnimplementedMasterServer{},
	}
	RegisterMasterServer(master.Server, master)
	return master
}

// ClientOperationRequest is the function to be called when a request is made
func (mstr *Master) ClientOperationRequest(context context.Context, req *ClientRequest) (*ClientResponse, error) {

	// switch operation {
	// case cmn.Open:
	// 	fmt.
	// case cmn.Close:
	// case cmn.Read:
	// case cmn.Append:
	// case cmn.Delete:
	// case cmn.Snapshot:
	// }
	res := &ClientResponse{}

	appId := req.ApplicationId
	filename := req.Filename

	var allowedPermissions common.PermissionType = 0
	if file, ok := mstr.namespace.GetFileInformation(filename); ok {
		appIdAsString := appId
		if file.owner == appIdAsString {
			allowedPermissions = file.permissions[common.PermissionGroup_Application]
		} else if file.group[appIdAsString] {
			allowedPermissions = file.permissions[common.PermissionGroup_ApplicationGroup]
		} else {
			allowedPermissions = file.permissions[common.PermissionGroup_All]
		}

		requestedPermission := common.OperationToPermissionType(req.Operation)
		if allowedPermissions&requestedPermission == 0 {
			res.Success = false
			return res, errors.New("application does not have permission to perform requested operation on this file")
		}
	}

	token, err := security.CreateToken(appId, filename, allowedPermissions)
	if err != nil {
		res.Success = false
	}

	res.Token = token
	res.Success = true
	return res, nil
}

// RegisterChunkserver is a function which handles registering chunkservers
func (mstr *Master) RegisterChunkserver(
	context context.Context,
	req *RegisterRequest,
) (*RegisterResponse, error) {
	res := &RegisterResponse{}

	chunkserverAddr := req.ServerAddress
	info := ChunkserverInfo{
		time.Now(),
		0,
		make(map[string]bool),
	}
	mstr.Chunkservers[chunkserverAddr] = info
	logger.Successf("successfully registered chunkserver %v", chunkserverAddr)
	res.Success = true
	return res, nil
}
