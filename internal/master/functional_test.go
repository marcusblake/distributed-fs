// +build functional

package master

import (
	"context"
	"os"
	"testing"

	cmn "github.com/distributed-fs/pkg/common"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

var master *Master

const address = ":8080"

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func setup() {
	master = NewMaster()
	if err := master.Start(address); err != nil {
		panic("master server failed to start")
	}
}

func TestOperationRequest(t *testing.T) {
	// Arrange
	args := ClientRequest{
		Operation: cmn.OperationType_OPEN,
		Offset:    0,
	}

	dialOptions := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(address, dialOptions...)
	if err != nil {
		t.Fatal("client setup failed", err)
	}

	defer conn.Close()

	masterClient := NewMasterClient(conn)

	// Act
	reply, err := masterClient.ClientOperationRequest(context.Background(), &args)

	// Assert
	assert.Nil(t, err)
	assert.True(t, reply.Success)
}
