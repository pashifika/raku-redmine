// Package main
package main

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"

	"github.com/pashifika/util/files"

	"raku-redmine/share"
	db "raku-redmine/utils/database"
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
	//goland:noinspection GoBoolExpressions
	if err == nil && isRelease {
		if !files.Exists(conf) {
			err = mkdirIfNotExist(conf)
			if err != nil {
				return
			}
		}
		db.DsnConf = filepath.Join(conf, "setting.db")
		confCustomFields = filepath.Join(conf, "custom_fields.json")
	}
	share.UI.UserDir = conf
	share.UI.OS = runtime.GOOS
	conf = filepath.Join(conf, "setting.ini")

	return
}

func mkdirIfNotExist(path string) error {
	if _, err := os.Stat(path); err != nil {
		if err = os.Mkdir(path, 0700); err != nil {
			return err
		}
	}
	return nil
}
