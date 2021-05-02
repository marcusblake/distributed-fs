package chunkserver

import (
	"fmt"
	"net/rpc"
	"sync"

	"github.com/distributed-fs/internal/rpctype"
	"github.com/distributed-fs/pkg/common"
	"github.com/distributed-fs/pkg/logger"
)

const (
	masterRegistrationMethod = "Master.ChunkserverRegistration"
)

// Chunkserver is a struct that represents the chunks
type Chunkserver struct {
	files   map[string]bool
	handler *FileHandler
	*rpctype.RPCServer
}

// NewChunkserver allocates a new instances of a chunkserver and initializes all of its fields
func NewChunkserver() *Chunkserver {
	fileHandler := &FileHandler{
		OpenFiles: make(map[string]*File),
		lck:       sync.Mutex{},
	}

	chunkserver := &Chunkserver{
		make(map[string]bool),
		fileHandler,
		rpctype.NewRPCServer(),
	}
	rpc.Register(chunkserver)
	return chunkserver
}

// RegisterChunkserver register the given chunkserver to the master server
func RegisterChunkserver(masterAddress string, chunkserverAddress string) error {
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
		logger.Failure(errMsg)
		return fmt.Errorf(errMsg)
	}

	return nil
}

// FileIORequest handles request to perform operations on files from client
func (chunk *Chunkserver) FileIORequest(req *rpctype.FileIORequest, res *rpctype.FileIOResponse) error {
	var err error
	var responseData []byte

	operation := req.Operation
	filename := req.Filename
	bytes := req.Bytes
	offset := req.Offset
	data := req.Data

	switch operation {
	case common.Operation.Open:
		err = chunk.handler.Open(filename)
	case common.Operation.Read:
		responseData, err = chunk.handler.Read(filename, bytes, offset)
	case common.Operation.Append:
		err = chunk.handler.Append(filename, data)
	case common.Operation.Close:
		err = chunk.handler.Close(filename)
	default:
	}

	if err != nil {
		res.Ok = false
		return err
	}

	res.Data = responseData
	res.Ok = true
	return nil
}
