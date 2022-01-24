// Package window
/*
 * Version: 1.0.0
 * Copyright (c) 2022. Pashifika
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
package window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"go.uber.org/zap"
	"gopkg.in/ini.v1"

	"raku-redmine/lib"
	"raku-redmine/lib/window/settings"
	"raku-redmine/utils/configs"
	"raku-redmine/utils/log"
)

func (win *Window) Setting() {
	w := fyne.CurrentApp().NewWindow("Settings")
	w.CenterOnScreen()
	w.SetFixedSize(true)
	w.Resize(lib.SettingWindow)

	cfg, err := ini.Load(win.ConfPath)
	if err != nil {
		dialog.ShowError(err, win.Master)
		log.Error("ini.Load", "path:"+win.ConfPath, zap.Error(err))
		return
	}
	setting := settings.New()
	err = cfg.MapTo(setting.Conf)
	if err != nil {
		dialog.ShowError(err, win.Master)
		log.Error("ini.MapTo", "path:"+win.ConfPath, zap.Error(err))
		return
	}
	setting.TopW = w
	setting.SaveToDisk = func(conf *configs.Root) error {
		cfg := ini.Empty()
		err := cfg.ReflectFrom(conf)
		if err != nil {
			log.Error("setting.SaveToDisk", "ReflectFrom error", zap.Error(err))
			return err
		}
		err = cfg.SaveTo(win.ConfPath)
		if err != nil {
			log.Error("setting.SaveToDisk", "SaveTo error", zap.Error(err))
		} else {
			log.Debug("setting.SaveToDisk", "configs now save to disk")
		}
		return err
	}
	w.SetContent(setting.Build())
	w.Show()
}
