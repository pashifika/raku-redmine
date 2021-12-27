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
	"os"
	"time"

	"github.com/pashifika/rollingFiles"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	forStdlib bool
	runOne    bool
	debug     bool
	ops       []zap.Option         // config for new zapLog
	logger    *zap.Logger          // saved zap logger
	writer    *rollingFiles.Logger // write file io
}

var (
	DefaultEncoder = zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		CallerKey:      "caller",
		MessageKey:     "msg",
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
)

// ManualNew manual set the new zap log
func ManualNew(ec zapcore.EncoderConfig,
	debug bool, outPath, layout string, maxSize, maxAge, maxBackups, callerSkip int,
	ops ...zap.Option) *Logger {
	nl := &Logger{forStdlib: false, runOne: false, debug: false, ops: ops}
	nl.debug = debug
	nl.forStdlib = false
	nl.build(outPath, layout, maxSize, maxAge, maxBackups, callerSkip, ec)
	return nl
}

// Build a newLog Logger from the provided zapcore.Core and Options.
func (l *Logger) build(outPath, layout string, maxSize, maxAge, maxBackups, callerSkip int, ec zapcore.EncoderConfig) {
	if l.runOne {
		return
	} else {
		l.runOne = true
	}

	// zap core config
	level := zap.NewAtomicLevel()
	var (
		outFile io.Writer
		opts    []zap.Option
		zapEnc  zapcore.Encoder
	)
	if l.debug {
		outFile = os.Stdout
		level.SetLevel(zapcore.DebugLevel)
		zapEnc = zapcore.NewConsoleEncoder(ec)
		opts = append(opts, zap.ErrorOutput(zapcore.AddSync(os.Stderr)))
	} else {
		l.writer = &rollingFiles.Logger{
			Filename:   outPath,
			MaxSize:    maxSize,
			MaxAge:     maxAge,
			MaxBackups: maxBackups,
			TimeLayout: layout,
			LocalTime:  true,
			Compress:   false,
		}
		outFile = l.writer
		level.SetLevel(zapcore.InfoLevel)
		zapEnc = zapcore.NewJSONEncoder(ec)
	}
	if l.forStdlib {
		opts = append(opts, zap.AddCaller(), zap.AddCallerSkip(callerSkip))
		opts = append(opts, zap.WrapCore(func(core zapcore.Core) zapcore.Core {
			return zapcore.NewSamplerWithOptions(core, time.Second, 5, 50)
		}))
	} else {
		if len(l.ops) != 0 {
			opts = append(opts, l.ops...)
		}
	}

	// creates a log core that writes logs to a writer syncer.
	l.logger = zap.New(
		zapcore.NewCore(zapEnc, zapcore.AddSync(outFile), level),
		opts...,
	)
	l.ops = nil
}
