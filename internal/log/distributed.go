package log

import (
	"os"
	"path/filepath"

	"github.com/hashicorp/raft"
)

type DistributedLog struct {
	config Config
	log    *Log
	raft   *raft.Raft
}

func (l *DistributedLog) setupLog(dataDir string) error {
	logDir := filepath.Join(dataDir, "log")
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return err
	}
	var err error
	l.log, err = NewLog(logDir, l.config)
	return err
}
