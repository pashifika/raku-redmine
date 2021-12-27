// Package main
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
package main

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"

	"github.com/pashifika/util/files"

	"raku-redmine/share"
	db "raku-redmine/utils/database"
	"raku-redmine/utils/log"
)

func getAppData() (conf string, err error) {
	var home string
	home, err = os.UserHomeDir()
	if err != nil {
		return
	}

	switch runtime.GOOS {
	case "windows":
		conf, err = os.Getwd()
		if err != nil {
			return
		}
		conf = filepath.Join(conf, "configs")
		share.UI.SystemFont = "C:/Windows/Fonts/meiryo.ttc"
	case "darwin":
		conf = filepath.Join(home, ".raku-mine")
	case "linux":
		err = errors.New("not support this os")
		return
	}
	// logger init
	logPath := filepath.Join(conf, "logs")
	if !files.Exists(logPath) {
		err = mkdirIfNotExist(logPath)
		if err != nil {
			return
		}
	}
	//goland:noinspection GoBoolExpressions
	log.Init(!isRelease, filepath.Join(logPath, "app.log"), "20060102", 10, 5, 5)
	//goland:noinspection GoBoolExpressions
	if isRelease {
		db.DsnConf = filepath.Join(conf, "setting.db")
		confCustomFields = filepath.Join(conf, "custom_fields.json")
	}
	share.UI.UserDir = conf
	share.UI.OS = runtime.GOOS
	share.UI.ARCH = runtime.GOARCH
	//goland:noinspection GoBoolExpressions
	share.UI.Debug = !isRelease
	conf = filepath.Join(conf, "setting.ini")

	return
}

func mkdirIfNotExist(path string) error {
	if _, err := os.Stat(path); err != nil {
		if err = os.MkdirAll(path, 0700); err != nil {
			return err
		}
	}
	return nil
}
