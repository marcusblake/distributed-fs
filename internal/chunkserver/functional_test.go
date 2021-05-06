// +build functional

package chunkserver

import (
	"testing"

	"github.com/distributed-fs/internal/master"
	"github.com/stretchr/testify/assert"
)

const (
	ChunkserverAddress = "localhost:8080"
)

var chunkserver *Chunkserver
var chunkserverAddress string

func TestRegistrationToMaster(t *testing.T) {
	// Arrange
	masterAddress := ":8081"
	master := master.NewMaster()
	if err := master.Start(masterAddress); err != nil {
		t.Fatal("master server failed to start")
	}

	//defer master.Shutdown()

	// Act
	if err := RegisterChunkserver(masterAddress, chunkserverAddress); err != nil {
		t.Fatalf("ChunkserverRegistration failed with %v", err.Error())
	}

	// Assert
	assert.Contains(t, master.Chunkservers, chunkserverAddress)
}

func TestLeaderElection(t *testing.T) {

}
