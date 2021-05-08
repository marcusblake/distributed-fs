package chunkserver

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/hashicorp/raft"
	boltdb "github.com/hashicorp/raft-boltdb"
)

const (
	NumberOfSnapshotsRetained = 2
)

// ChunkserverFSM implements the FSM interface in raft. It is used to perform operations on log data to
// to get chunkservers to reach the same state
type ChunkserverFSM struct {
}

// ChunkserverTCPStreamLayer
type ChunkserverTCPStreamLayer struct {
}

// Apply is called when a log entry is commited
func (fsm *ChunkserverFSM) Apply(log *raft.Log) interface{} {
	return nil
}

// Used to support log compaction which can be used to save a point in time snapshot of FSM
func (fsm *ChunkserverFSM) Snapshot() (raft.FSMSnapshot, error) {
	return nil, nil
}

// Restore is used to restore a fsm from a snapshot. It is not called concurrently
func (fsm *ChunkserverFSM) Restore(closer io.ReadCloser) error {
	return nil
}

func NewRaft() *raft.Raft {

	baseDirectory := fmt.Sprintf("chunkserver%v/", time.Now().String())

	config := raft.DefaultConfig()

	logStore, err := boltdb.NewBoltStore(filepath.Join(baseDirectory, "logstore.dat"))
	if err != nil {
		log.Fatalf("error creating log store %v", err)
	}

	stableStore, err := boltdb.NewBoltStore(filepath.Join(baseDirectory, "stableStore.dat"))
	if err != nil {
		log.Fatalf("error creating stable store %v", err)
	}

	// TODO: instead of os.Stderr, write to a log file that can be used for debugging
	snapshotStore, err := raft.NewFileSnapshotStore(filepath.Join("store.dat"), NumberOfSnapshotsRetained, os.Stderr)
	if err != nil {
		log.Fatalf("error creating snapshot store %v", err)
	}

	raft, err := raft.NewRaft(config, &ChunkserverFSM{}, logStore, stableStore, snapshotStore, nil)
	if err != nil {
		log.Fatalf("error creating raft struct %v", err)
	}

	return raft
}
