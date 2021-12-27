// Package log
/*
 * Version: 1.0.0
 * Copyright (c) 2021. Pashifika
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package log

import (
	"io"

	"go.uber.org/zap"
)

// std is the name of the standard zap log in stdlib `log`
var std = new(Logger)

// Init the custom zap log for package std out.
//
// layout: see the https://golang.org/pkg/time/#pkg-constants
//
// layout: 2006-01-02 15:04:05
func Init(debug bool, outPath, layout string, maxSize, maxAge, maxBackups int) {
	std.debug = debug
	std.forStdlib = true
	std.build(outPath, layout, maxSize, maxAge, maxBackups, 1, DefaultEncoder)
}

// Close Sync calls the underlying Core's Sync method, flushing any buffered log entries.
// Applications should take care to call Sync before exiting.
// If mode is debug close implements io.Closer, and closes the current logfile.
func Close() error {
	return std.Close()
}

// Sync calls the underlying Core's Sync method, flushing any buffered log
// entries. Applications should take care to call Sync before exiting.
func Sync() error {
	return std.logger.Sync()
}

// GetZapLogger output zap.Logger struct
func GetZapLogger() *zap.Logger {
	return std.logger
}

// GetLogger output zap.Logger struct
func GetLogger() *Logger {
	return std
}

// GetWriter output io.Writer struct
func GetWriter() io.Writer {
	return std.writer
}

// Rotate causes Logger to close the existing log file and immediately create a
// new one.  This is a helper function for applications that want to initiate
// rotations outside of the normal rotation rules, such as in response to
// SIGHUP.  After rotating, this initiates compression and removal of old log
// files according to the configuration.
func Rotate() error {
	return std.writer.Rotate()
}
