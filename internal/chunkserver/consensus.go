package chunkserver

import (
	"io"

	"github.com/hashicorp/raft"
)

// ChunkserverFSM implements the FSM interface in raft. It is used to perform operations on log data to
// to get chunkservers to reach the same state
type ChunkserverFSM struct {
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
	config := raft.DefaultConfig()

	raft, err := raft.NewRaft(config, nil, nil, nil, nil, nil)
	if err != nil {
		panic(err)
	}

	return raft
}
