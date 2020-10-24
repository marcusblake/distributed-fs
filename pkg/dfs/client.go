package dfs

import (
	"fmt"
	"net"
	"net/rpc"
	"time"

	"github.com/distributed-fs/internal/rpctype"
	"github.com/distributed-fs/pkg/common"
)

const (
	masterRequestMethod     = "Master.OperationRequest"
	chunkserverFileOpMethod = "Chunkserver.FileIORequest"
	defaultTimeout          = 20 * time.Second
)

// Client is a struct that represents the DFS client
type Client struct {
	MasterAddress string
	ConnTimeout   time.Duration
}

// NewClient allocates a Client struct
func NewClient(address string) *Client {
	return &Client{
		MasterAddress: address,
		ConnTimeout:   defaultTimeout,
	}
}

// IssueOperationRequest issues a request to the master server for an operation
func (client *Client) IssueOperationRequest(op common.Operation, filename string, offset uint32) (string, error) {

	args := &rpctype.OperationRequest{
		Operation: op,
		Offset:    offset,
	}

	var reply rpctype.OperationResponse

	conn, err := net.DialTimeout("tcp", client.MasterAddress, client.ConnTimeout)
	if err != nil {
		return "", err
	}

	// Close TCP connection when done
	defer conn.Close()

	rpcClient := rpc.NewClient(conn)
	if err := rpcClient.Call(masterRequestMethod, args, &reply); err != nil {
		return "", err
	} else if !reply.Ok {
		return "", fmt.Errorf("rpc client failed to make a request")
	}

	return reply.IPAddress, nil
}

// IssueFileIORequest issues a request to a chunkserver to get file data or
func (client *Client) IssueFileIORequest(op common.Operation, filename string, offset uint32, chunkserverAddress string) ([]byte, error) {
	return nil, nil
}
