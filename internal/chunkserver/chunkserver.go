package chunkserver

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/distributed-fs/internal/master"
	"github.com/distributed-fs/internal/rpctype"
	"github.com/distributed-fs/internal/security"
	"github.com/distributed-fs/pkg/common"
	"github.com/distributed-fs/pkg/logger"
	"google.golang.org/grpc"
)

const (
	masterRegistrationMethod = "Master.ChunkserverRegistration"
)

var (
	dialOptions = []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithTimeout(1 * time.Minute),
	}
)

// Chunkserver is a struct that represents the chunks
type Chunkserver struct {
	files   map[string]bool
	handler *FileHandler
	*rpctype.RPCServer
	UnimplementedChunkserverServer
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
		rpctype.NewRPCServer(nil),
		UnimplementedChunkserverServer{},
	}
	RegisterChunkserverServer(chunkserver.Server, chunkserver)
	return chunkserver
}

// RegisterChunkserver register the given chunkserver to the master server
func RegisterChunkserver(masterAddress string, chunkserverAddress string) error {
	// Setup necessary arguments and parameters
	args := &master.RegisterRequest{
		ServerAddress: chunkserverAddress,
	}

	conn, err := grpc.Dial(masterAddress, dialOptions...)
	if err != nil {
		return err
	}

	// Close the TCP connection when done
	defer conn.Close()

	masterClient := master.NewMasterClient(conn)
	reply, err := masterClient.RegisterChunkserver(context.Background(), args)
	if err != nil {
		return err
	} else if !reply.Success {
		errMsg := "rpc call to register chunkserver returned as unsuccessful"
		logger.Failure(errMsg)
		return fmt.Errorf(errMsg)
	}

	return nil
}

// ServeClientFileRequest handles request to perform operations on files from client
func (chunk *Chunkserver) ServeClientFileRequest(
	context context.Context,
	req *FileIORequest,
) (*FileIOResponse, error) {
	var err error
	var responseData []byte
	var res *FileIOResponse

	operation := req.Operation
	filename := req.Filename
	bytes := req.Bytes
	offset := req.Offset
	data := req.Data
	token := req.Token
	appId := req.ApplicationId

	err = security.VerifyToken(token, appId, filename, operation)
	if err != nil {
		res.Success = false
		return res, err
	}

	switch operation {
	case common.OperationType_OPEN:
		err = chunk.handler.Open(filename)
	case common.OperationType_READ:
		responseData, err = chunk.handler.Read(filename, bytes, offset)
	case common.OperationType_APPEND:
		err = chunk.handler.Append(filename, data)
	case common.OperationType_CLOSE:
		err = chunk.handler.Close(filename)
	default:
	}

	if err != nil {
		res.Success = false
		return res, err
	}

	res.Data = responseData
	res.Success = true
	return res, nil
}
