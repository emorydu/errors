// Copyright 2023 Emory.Du <orangeduxiaocheng@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

import (
	"encoding/json"
	"regexp"
	"testing"
)

func TestFrameMarshalText(t *testing.T) {
	var tests = []struct {
		Frame
		want string
	}{{
		Frame: initpc,
		want:  `^github.com/emorydu/errors\.init(\.ializers)? .+/github\.com/emorydu/errors/stack_test.go:\d+$`,
	}, {
		Frame: 0,
		want:  "^unknown$",
	}}

	for i, tt := range tests {
		got, err := tt.Frame.MarshalText()
		if err != nil {
			t.Fatal(err)
		}
		if !regexp.MustCompile(tt.want).Match(got) {
			t.Errorf("test %d: MarshalJSON:\n got %q\n want %q", i+1, string(got), tt.want)
		}
	}
}

func TestFrameMarshalJSON(t *testing.T) {
	var tests = []struct {
		Frame
		want string
	}{{
		initpc,
		`^"github\.com/emorydu/errors\.init(\.ializers)? .+/github\.com/emorydu/errors/stack_test.go:\d+"$`,
	}, {
		0,
		`^"unknown"$`,
	}}
	for i, tt := range tests {
		got, err := json.Marshal(tt.Frame)
		if err != nil {
			t.Fatal(err)
		}

		if !regexp.MustCompile(tt.want).Match(got) {
			t.Errorf("test %d: MarshalJSON:\n got %q\n want %q", i+1, string(got), tt.want)
		}
	}

}
