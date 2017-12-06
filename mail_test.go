// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package log

import (
	"testing"
)

func TestNewMailTarget(t *testing.T) {
	target := NewMailTarget()
	if target.MaxLevel != LevelDebug {
		t.Errorf("NewMailTarget.MaxLevel = %v, expected %v", target.MaxLevel, LevelDebug)
	}
}
