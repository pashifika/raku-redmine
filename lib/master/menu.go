// Package master
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
package master

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"go.uber.org/zap"

	"raku-redmine/share"
	"raku-redmine/utils/log"
)

func NewMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {
	openConfDir := fyne.NewMenuItem("User files...", func() {
		u, err := url.Parse(share.UI.UserDir)
		if err != nil {
			dialog.ShowError(err, w)
			log.Error("openConfDir", "url.Parse error", zap.Error(err))
			return
		}
		err = a.OpenURL(u)
		if err != nil {
			dialog.ShowError(err, w)
			log.Error("openConfDir", "OpenURL error", zap.Error(err))
			return
		}
	})
	settingsItem := fyne.NewMenuItem("Settings", func() {
		share.UI.Window.Setting()
	})

	helpMenu := fyne.NewMenu("Help",
		fyne.NewMenuItem("Documentation", func() {
			u, _ := url.Parse("https://github.com/pashifika/raku-redmine")
			_ = a.OpenURL(u)
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("About", func() {
			share.UI.Window.About()
		}),
	)

	// a quit item will be appended to our first (File) menu
	file := fyne.NewMenu("File", openConfDir)
	if !fyne.CurrentDevice().IsMobile() {
		file.Items = append(file.Items, settingsItem)
	}
	return fyne.NewMainMenu(
		file,
		helpMenu,
	)
}
