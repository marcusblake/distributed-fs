package rpctype

import (
	"github.com/distributed-fs/pkg/common"
	"github.com/google/uuid"
)

// Server is an interface for RPC servers to use
type Server interface {
	Start(address string) error
	Shutdown()
}

// ChunkserverRegisterRequest is a struct that will be used to register a chunkserver to master when it spawns
type ChunkserverRegisterRequest struct {
	ServerAddress string
}

// ChunkserverRegisterResponse is struct that will be sent to the chunkserver from master
type ChunkserverRegisterResponse struct {
	Ok bool
}

// PollChunkserverRequest is a struct that will be used to poll the chunkservers periodically
type PollChunkserverRequest struct {
}

// PollChunkserverResponse is a struct that will be used when chunkservers respond to masters poll request
type PollChunkserverResponse struct {
	Files []string
	Ok    bool
}

// OperationRequest is a request to master from client to perform an operation
type OperationRequest struct {
	ApplicationId uuid.UUID
	Operation     common.FileOperation
	Filename      string
	Permissions   map[common.PermissionGroup]common.PermissionType
	Offset        uint32
}

// OperationResponse is a reponse to the client indicate which chunkserver to use to make request
type OperationResponse struct {
	IPAddress string
	Port      uint16
	Token     string
	Ok        bool
}

// FileIORequest is a request to make file io operation
type FileIORequest struct {
	ApplicationId uuid.UUID
	Token         string
	Operation     common.FileOperation
	Filename      string
	Bytes         int64
	Offset        int64
	Data          []byte
}

// FileIOResponse is a response from a chunkserver when a request is made for a file operation
type FileIOResponse struct {
	Data []byte
	Ok   bool
}
