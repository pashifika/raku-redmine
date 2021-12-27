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
	"fmt"

	"go.uber.org/zap"
)

// Create a log core from the setting
//
// layout: see the https://golang.org/pkg/time/#pkg-constants
//
// layout: 2006-01-02 15:04:05
func (l *Logger) Create(debug bool, outPath, layout string, maxSize, maxAge, maxBackups, callerSkip int) {
	l.debug = debug
	l.forStdlib = true
	l.build(outPath, layout, maxSize, maxAge, maxBackups, callerSkip, DefaultEncoder)
}

// Rotate causes Logger to close the existing log file and immediately create a
// new one.  This is a helper function for applications that want to initiate
// rotations outside of the normal rotation rules, such as in response to
// SIGHUP.  After rotating, this initiates compression and removal of old log
// files according to the configuration.
func (l *Logger) Rotate() error {
	return l.writer.Rotate()
}

// Close Sync calls the underlying Core's Sync method, flushing any buffered log entries.
// Applications should take care to call Sync before exiting.
// If mode is debug close implements io.Closer, and closes the current logfile.
func (l *Logger) Close() error {
	err := l.logger.Sync()
	if err != nil {
		return err
	}
	if l.debug {
		err = l.writer.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

// Log A Logger provides fast, leveled, structured logging. All methods are safe
// for concurrent use.
//
// The Logger is designed for contexts in which every microsecond and every
// allocation matters, so its API intentionally favors performance and type
// safety over brevity. For most applications, the SugaredLogger strikes a
// better balance between performance and ergonomics.
func (l *Logger) Log() *zap.Logger {
	return l.logger
}

// Sugar wraps the Logger to provide a more ergonomic, but slightly slower,
// API. Sugaring a Logger is quite inexpensive, so it's reasonable for a
// single application to use both Loggers and SugaredLoggers, converting
// between them on the boundaries of performance-sensitive code.
func (l *Logger) Sugar() *zap.SugaredLogger {
	return l.logger.Sugar()
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// ** fn is format to "func_name.why", please set it to debug.
func (l *Logger) Info(fn, msg string, fields ...zap.Field) {
	fs := []zap.Field{zap.String("func", fn)}
	if len(fields) != 0 {
		fs = append(fs, fields...)
	}
	l.logger.Info(msg, fs...)
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// ** fn is format to "func_name.why", please set it to debug.
func (l *Logger) Debug(fn, msg string, fields ...zap.Field) {
	fs := []zap.Field{zap.String("func", fn)}
	if len(fields) != 0 {
		fs = append(fs, fields...)
	}
	l.logger.Debug(msg, fs...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// ** fn is format to "func_name.why", please set it to debug.
func (l *Logger) Error(fn, msg string, fields ...zap.Field) {
	fs := []zap.Field{zap.String("func", fn)}
	if len(fields) != 0 {
		fs = append(fs, fields...)
	}
	l.logger.Error(msg, fs...)
}

// Panic logs a message at PanicLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then panics, even if logging at PanicLevel is disabled.
//
// ** fn is format to "func_name.why", please set it to debug.
func (l *Logger) Panic(fn, msg string, fields ...zap.Field) {
	fs := []zap.Field{zap.String("func", fn)}
	if len(fields) != 0 {
		fs = append(fs, fields...)
	}
	l.logger.Panic(msg, fs...)
}

// Fatal logs a message at FatalLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then calls os.Exit(1), even if logging at FatalLevel is
// disabled.
//
// ** fn is format to "func_name.why", please set it to debug.
func (l *Logger) Fatal(fn, msg string, fields ...zap.Field) {
	fs := []zap.Field{zap.String("func", fn)}
	if len(fields) != 0 {
		fs = append(fs, fields...)
	}
	l.logger.Fatal(msg, fs...)
}

// Print Debug to database access log used gorm
func (l *Logger) Print(v ...interface{}) {
	str := ""
	for index := 0; index < len(v); index++ {
		str1 := fmt.Sprintf("%v", v[index])
		if index == 0 {
			str += str1
		} else {
			str += "\t" + str1
		}
	}
	l.logger.Sugar().Debug(str)
}
