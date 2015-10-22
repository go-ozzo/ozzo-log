package log_test

import (
	"testing"
	"github.com/go-ozzo/ozzo-log"
)

func TestNewMailTarget(t *testing.T) {
	target := log.NewMailTarget()
	if target.MaxLevel != log.LevelDebug {
		t.Errorf("NewMailTarget.MaxLevel = %v, expected %v", target.MaxLevel, log.LevelDebug)
	}
}
