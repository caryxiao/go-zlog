package zlog

import (
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	SetLevel(DebugLevel)
	SetFormat("[%level%]: %time% - [%trace_id%] %msg%")
	//SetOutput("/tmp/go.log")
	SetOutput(os.Stdout)

	Logger.Warn("test warning")
	Logger.Info("test info")
	Logger.Error("test error")
	Logger.Debug("test debug")
}
