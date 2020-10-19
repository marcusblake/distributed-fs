package chunkserver

import (
	"os"
	"testing"

	"github.com/distributed-fs/internal/master"
	"github.com/stretchr/testify/assert"
)

var chunkserver *Chunkserver
var chunkserverAddress string

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	chunkserverAddress = "localhost:8080"

}

func teardown() {

}

func TestRegistrationToMaster(t *testing.T) {
	// Arrange
	masterAddress := ":8081"
	master := master.NewMaster()
	if err := master.Start(masterAddress); err != nil {
		t.Fatal("master server failed to start")
	}

	// Gracefully shutdown master server once test method is complete
	defer func() {
		if err := master.Stop(); err != nil {
			t.Fatalf("master server failed to stop %v", err.Error())
		}
	}()

	// Act
	if err := RegisterChunkserver(masterAddress, chunkserverAddress); err != nil {
		t.Fatalf("ChunkserverRegistration failed with %v", err.Error())
	}

	// Assert
	assert.Contains(t, master.Chunkservers, chunkserverAddress)
}
