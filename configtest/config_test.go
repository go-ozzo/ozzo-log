// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package configtest_test

import (
	"io"
	"testing"

	"github.com/go-ozzo/ozzo-config"
	"github.com/go-ozzo/ozzo-log"
)

// optionTarget is a no-op target whose exported options ozzo-config fills in.
type optionTarget struct {
	Option1 string
	Option2 bool
}

func (t *optionTarget) Open(io.Writer) error { return nil }
func (t *optionTarget) Process(*log.Entry)   {}
func (t *optionTarget) Close()               {}

// TestLoggerConfig covers the README's "configure the logger with ozzo-config" path.
func TestLoggerConfig(t *testing.T) {
	c := config.New()
	err := c.LoadJSON([]byte(`{
		"Logger": {
			"MaxLevel": 2,
			"Category": "app2",
			"Targets": [
				{
					"type": "memory1",
					"Option1": "abc",
					"Option2": true
				},
				{
					"type": "memory2",
					"Option1": "xyz"
				}
			]
		}
	}`))
	if err != nil {
		t.Errorf("config.LoadJSON(): %v", err)
	}
	c.Register("memory1", func() *optionTarget {
		return &optionTarget{}
	})
	c.Register("memory2", func() *optionTarget {
		return &optionTarget{Option2: true}
	})
	logger := log.NewLogger()

	if err := c.Configure(logger, "Logger"); err != nil {
		t.Errorf("config.Configure(logger): %v", err)
	}
	if logger.MaxLevel != log.LevelCritical {
		t.Errorf("logger.MaxLevel = %v, expected %v", logger.MaxLevel, log.LevelCritical)
	}
	if logger.Category != "app2" {
		t.Errorf("logger.Category = %v, expected %v", logger.Category, "app2")
	}

	if len(logger.Targets) != 2 {
		t.Errorf("len(logger.Targets) = %v, expected %v", len(logger.Targets), 2)
		return
	}
	m1 := logger.Targets[0].(*optionTarget)
	m2 := logger.Targets[1].(*optionTarget)
	if m1.Option1 != "abc" || m1.Option2 != true {
		t.Errorf("m1.Option1 = %v, Option2 = %v, expected %v and %v", m1.Option1, m1.Option2, "abc", true)
	}
	if m2.Option1 != "xyz" || m2.Option2 != true {
		t.Errorf("m2.Option1 = %v, Option2 = %v, expected %v and %v", m2.Option1, m2.Option2, "xyz", true)
	}
}
