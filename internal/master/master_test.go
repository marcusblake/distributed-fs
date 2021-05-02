package master

import (
	"fmt"
	"net/rpc"
	"os"
	"testing"

	"github.com/distributed-fs/internal/rpctype"
	cmn "github.com/distributed-fs/pkg/common"
	"github.com/stretchr/testify/assert"
)

var master *Master
var address string

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	address = "localhost:8080"
	master = NewMaster()
	if err := master.Start(address); err != nil {
		panic("master server failed to start")
	}
}

func teardown() {
	if err := master.Stop(); err != nil {
		panic(fmt.Sprintf("master server failed to stop %v", err.Error()))
	}
}

func TestOperationRequest(t *testing.T) {
	// Arrange
	method := "Master.OperationRequest"
	args := &rpctype.OperationRequest{
		Operation: cmn.Operation.Open,
		Offset:    0,
	}
	var reply rpctype.OperationResponse

	testClient, err := rpc.Dial("tcp", address)
	if err != nil {
		t.Fatal("client setup failed")
	}

	defer func() {
		if err := testClient.Close(); err != nil {
			t.Fatal("client couldn't close the connection")
		}
	}()

	// Act
	if err := testClient.Call(method, args, &reply); err != nil {
		t.Fatal("client call failed")
	}

	// Assert
	assert.True(t, reply.Ok)
}
