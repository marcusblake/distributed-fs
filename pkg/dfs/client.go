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
	defaultTimeout          = 10 * time.Second
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
func (client *Client) IssueOperationRequest(op common.FileOperation, filename string, offset uint32) (string, error) {

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
func issueFileIORequest(client *Client, op common.FileOperation, filename string, data []byte, bytes, offset int64, chunkserverAddress string) ([]byte, error) {
	args := &rpctype.FileIORequest{
		Operation: op,
		Bytes:     bytes,
		Offset:    offset,
		Filename:  filename,
		Data:      data,
	}

	var reply rpctype.FileIOResponse

	conn, err := net.DialTimeout("tcp", chunkserverAddress, client.ConnTimeout)
	if err != nil {
		return nil, err
	}

	// Close TCP connection when done
	defer conn.Close()

	rpcClient := rpc.NewClient(conn)
	if err := rpcClient.Call(chunkserverFileOpMethod, args, &reply); err != nil {
		return nil, err
	} else if !reply.Ok {
		return nil, fmt.Errorf("rpc client failed to make a request")
	}
	fmt.Println(reply)

	return reply.Data, nil
}

// Open opens a file on the chunkserver
func (client *Client) Open(filename string, chunkserver string) error {
	_, err := issueFileIORequest(client, common.Operation.Open, filename, nil, 0, 0, chunkserver)
	return err
}

// Close closes a file on the chunkserver
func (client *Client) Close(filename string, chunkserver string) error {
	_, err := issueFileIORequest(client, common.Operation.Close, filename, nil, 0, 0, chunkserver)
	return err
}

// Read reads a file on the chunkserver
func (client *Client) Read(filename string, bytes, offset int64, chunkserver string) ([]byte, error) {
	return issueFileIORequest(client, common.Operation.Read, filename, nil, bytes, offset, chunkserver)
}

// Append reads a file on the chunkserver
func (client *Client) Append(filename string, data []byte, chunkserver string) error {
	_, err := issueFileIORequest(client, common.Operation.Append, filename, data, 0, 0, chunkserver)
	return err
}
