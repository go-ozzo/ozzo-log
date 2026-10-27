// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package log_test

import (
	"errors"
	"io"
	"testing"

	"github.com/go-ozzo/ozzo-log"
)

func TestNewLogger(t *testing.T) {
	logger := log.NewLogger()
	if logger.MaxLevel != log.LevelDebug {
		t.Errorf("NewLogger().MaxLevel = %v, expected %v", logger.MaxLevel, log.LevelDebug)
	}
	if logger.Category != "app" {
		t.Errorf("NewLogger().Category = %v, expected %v", logger.Category, "app")
	}
	if logger.CallStackDepth != 0 {
		t.Errorf("NewLogger().CallStackDepth = %v, expected %v", logger.CallStackDepth, 0)
	}
}

func TestGetLogger(t *testing.T) {
	formatter := func(*log.Logger, *log.Entry) string {
		return "test"
	}
	logger := log.NewLogger()
	logger1 := logger.GetLogger("testing")
	if logger1.Category != "testing" {
		t.Errorf("logger1.Category = %v, expected %v", logger1.Category, "testing")
	}
	logger2 := logger.GetLogger("routing", formatter)
	if logger2.Category != "routing" {
		t.Errorf("logger2.Category = %v, expected %v", logger2.Category, "routing")
	}
	if logger2.Formatter(logger2, nil) != "test" {
		t.Errorf("logger2.Formatter has an unexpected value")
	}
}

type MemoryTarget struct {
	entries []*log.Entry
	open    bool
	ready   chan bool
}

type failingTarget struct {
	err error
}

func (f failingTarget) Open(io.Writer) error { return f.err }
func (failingTarget) Process(*log.Entry)     {}
func (failingTarget) Close()                 {}

func TestLoggerOpenReturnsTargetError(t *testing.T) {
	want := errors.New("target open failed")
	logger := log.NewLogger()
	defer logger.Close()
	logger.Targets = append(logger.Targets, failingTarget{err: want})

	if err := logger.Open(); !errors.Is(err, want) {
		t.Fatalf("Logger.Open() error = %v, want error wrapping %v", err, want)
	}
}

func (m *MemoryTarget) Open(io.Writer) error {
	m.open = true
	m.entries = make([]*log.Entry, 0)
	return nil
}

func (m *MemoryTarget) Process(e *log.Entry) {
	if e == nil {
		m.ready <- true
	} else {
		m.entries = append(m.entries, e)
	}
}

func (m *MemoryTarget) Close() {
	<-m.ready
}

func TestLoggerLog(t *testing.T) {
	logger := log.NewLogger()
	target := &MemoryTarget{
		ready: make(chan bool),
	}
	logger.Targets = append(logger.Targets, target)

	if target.open {
		t.Errorf("target.open = %v, expected %v", target.open, false)
	}
	logger.Open()
	if !target.open {
		t.Errorf("target.open = %v, expected %v", target.open, true)
	}

	logger.Log(log.LevelInfo, "t0: %v", 1)
	logger.Debug("t1: %v", 2)
	logger.Info("t2")
	logger.Warning("t3")
	logger.Notice("t4")
	logger.Error("t5")
	logger.Critical("t6")
	logger.Alert("t7")
	logger.Emergency("t8")

	logger.Close()

	if len(target.entries) != 9 {
		t.Errorf("len(target.entries) = %v, expected %v", len(target.entries), 9)
	}
	levels := ""
	messages := ""
	for i := 0; i < 9; i++ {
		levels += target.entries[i].Level.String() + ","
		messages += target.entries[i].Message + ","
	}
	expectedLevels := "Info,Debug,Info,Warning,Notice,Error,Critical,Alert,Emergency,"
	expectedMessages := "t0: 1,t1: 2,t2,t3,t4,t5,t6,t7,t8,"
	if levels != expectedLevels {
		t.Errorf("levels = %v, expected %v", levels, expectedLevels)
	}
	if messages != expectedMessages {
		t.Errorf("messages = %v, expected %v", messages, expectedMessages)
	}
}
