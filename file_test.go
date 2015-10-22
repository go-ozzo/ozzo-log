package log_test

import (
	"testing"
	"github.com/go-ozzo/ozzo-log"
)

func TestNewFileTarget(t *testing.T) {
	target := log.NewFileTarget()
	if target.MaxLevel != log.LevelDebug {
		t.Errorf("NewFileTarget.MaxLevel = %v, expected %v", target.MaxLevel, log.LevelDebug)
	}
	if target.Rotate != true {
		t.Errorf("NewFileTarget.Rotate = %v, expected %v", target.Rotate, true)
	}
	if target.BackupCount != 10 {
		t.Errorf("NewFileTarget.BackupCount = %v, expected %v", target.BackupCount, 10)
	}
	if target.MaxBytes != (1 << 20) {
		t.Errorf("NewFileTarget.MaxBytes = %v, expected %v", target.MaxBytes, 1 << 20)
	}
}
