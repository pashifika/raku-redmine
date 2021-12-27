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
	"go.uber.org/zap"
)

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// ** fn is format to "func_name.why", please set it to debug.
func Debug(fn, msg string, fields ...zap.Field) {
	fs := []zap.Field{zap.String("func", fn)}
	if len(fields) != 0 {
		fs = append(fs, fields...)
	}
	std.logger.Debug(msg, fs...)
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// ** fn is format to "func_name.why", please set it to debug.
func Info(fn, msg string, fields ...zap.Field) {
	fs := []zap.Field{zap.String("func", fn)}
	if len(fields) != 0 {
		fs = append(fs, fields...)
	}
	std.logger.Info(msg, fs...)
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// ** fn is format to "func_name.why", please set it to debug.
func Warn(fn, msg string, fields ...zap.Field) {
	fs := []zap.Field{zap.String("func", fn)}
	if len(fields) != 0 {
		fs = append(fs, fields...)
	}
	std.logger.Warn(msg, fs...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// ** fn is format to "func_name.why", please set it to debug.
func Error(fn, msg string, fields ...zap.Field) {
	fs := []zap.Field{zap.String("func", fn)}
	if len(fields) != 0 {
		fs = append(fs, fields...)
	}
	std.logger.Error(msg, fs...)
}

// Panic logs a message at PanicLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then panics, even if logging at PanicLevel is disabled.
//
// ** fn is format to "func_name.why", please set it to debug.
func Panic(fn, msg string, fields ...zap.Field) {
	fs := []zap.Field{zap.String("func", fn)}
	if len(fields) != 0 {
		fs = append(fs, fields...)
	}
	std.logger.Panic(msg, fs...)
}

// DPanic logs a message at DPanicLevel. The message includes any fields
// passed at the log site, as well as any fields accumulated on the logger.
//
// If the logger is in development mode, it then panics (DPanic means
// "development panic"). This is useful for catching errors that are
// recoverable, but shouldn't ever happen.
//goland:noinspection GoUnusedExportedFunction
func DPanic(fn, msg string, fields ...zap.Field) {
	fs := []zap.Field{zap.String("func", fn)}
	if len(fields) != 0 {
		fs = append(fs, fields...)
	}
	std.logger.DPanic(msg, fs...)
}

// Fatal logs a message at FatalLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then calls os.Exit(1), even if logging at FatalLevel is
// disabled.
//
// ** fn is format to "func_name.why", please set it to debug.
func Fatal(fn, msg string, fields ...zap.Field) {
	fs := []zap.Field{zap.String("func", fn)}
	if len(fields) != 0 {
		fs = append(fs, fields...)
	}
	std.logger.Fatal(msg, fs...)
}
