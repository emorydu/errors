// Copyright 2023 Emory.Du <orangeduxiaocheng@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package errors

import "fmt"

/*
WARNING - changing the line numbers in this file will break the
examples.
*/

const (
	// Error codes below 1000 are reserved future use by the
	// "github.com/bdlm/errors" package.
	ConfigurationNotValid int = iota + 1000
	ErrInvalidJSON
	ErrEOF
	ErrLoadConfigFailed
)

func init() {
	Register(defaultCoder{ConfigurationNotValid, 500, "ConfigurationNotValid error", ""})
	Register(defaultCoder{ErrInvalidJSON, 500, "Data is not valid JSON", ""})
	Register(defaultCoder{ErrEOF, 500, "End of input", ""})
	Register(defaultCoder{ErrLoadConfigFailed, 500, "Load configuration file failed", ""})
}

func loadConfig() error {
	err := decodeConfig()
	return WrapC(err, ConfigurationNotValid, "service configuration could not be loaded")
}

func decodeConfig() error {
	err := readConfig()
	return WrapC(err, ErrInvalidJSON, "could not decode configuration data")
}

func readConfig() error {
	err := fmt.Errorf("read: end of input")
	return WrapC(err, ErrEOF, "could not read configuration file")
}
