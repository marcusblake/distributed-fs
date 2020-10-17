package rpctype

import (
	"github.com/distributed-fs/pkg/common"
)

// Server is an interface for RPC servers to use
type Server interface {
	StartServer()
	StopServer()
}

// ChunkserverRegisterRequest is a struct that will be used to register a chunkserver to master when it spawns
type ChunkserverRegisterRequest struct {
	ipAddress string
	port      uint16
}

// ChunkserverRegisterResponse is struct that will be sent to the chunkserver from master
type ChunkserverRegisterResponse struct {
	ok bool
}

// PollChunkserverRequest is a struct that will be used to poll the chunkservers periodically
type PollChunkserverRequest struct {
}

// PollChunkserverResponse is a struct that will be used when chunkservers respond to masters poll request
type PollChunkserverResponse struct {
	files []string
	ok    bool
}

// OperationRequest is a request to master from client to perform an operation
type OperationRequest struct {
	operation common.Operation
	offset    uint32
}

// OperationResponse is a reponse to the client indicate which chunkserver to use to make request
type OperationResponse struct {
	ipAddress string
	port      uint16
	ok        bool
}

// FileIORequest is a request to make file io operation
type FileIORequest struct {
	operation common.Operation
	data      []byte
}

// FileIOResponse is a response from a chunkserver when a request is made for a file operation
type FileIOResponse struct {
	data []byte
	ok   bool
}
