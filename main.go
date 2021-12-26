// Package raku_redmine
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
	"io/ioutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"github.com/pashifika/util/files"

	"raku-redmine/lib"
	"raku-redmine/lib/master"
	"raku-redmine/lib/theme"
	"raku-redmine/lib/time_entry"
	"raku-redmine/lib/window"
	"raku-redmine/resource"
	"raku-redmine/share"
	"raku-redmine/utils/configs"
	db "raku-redmine/utils/database"
	"raku-redmine/utils/database/models"
	ur "raku-redmine/utils/resource"
)

const appName = "Raku redmine"

var topWindow fyne.Window

func main() {
	a := app.New()
	a.SetIcon(resource.AppIconRes)
	share.UI = share.AppUI{AppName: appName}
	topWindow = a.NewWindow(appName)
	topWindow.SetFixedSize(true)
	topWindow.Resize(lib.MainWindow)
	topWindow.SetMainMenu(master.NewMenu(a, topWindow))

	// get app setting file
	var ready bool
	conf, err := getAppData()
	if err != nil {
		dialog.ShowError(err, topWindow)
		topWindow.Show()
	}
	if !files.Exists(conf) {
		window.Login(a.NewWindow(appName), func(masterUrl, apiKey, fontPath string) error {
			configs.Config = &configs.Root{
				Font: configs.Font{Path: fontPath},
				Redmine: configs.Redmine{
					MasterUrl: masterUrl,
					ApiKey:    apiKey,
				},
			}
			err := configs.Save(conf)
			if err != nil {
				return err
			}
			err = loadMaster(a)
			if err != nil {
				return err
			}
			topWindow.Show()
			ready = true
			return nil
		})
	} else {
		err = configs.Load(conf)
		if err != nil {
			dialog.ShowError(err, topWindow)
		} else {
			ready = true
		}
	}

	if ready {
		err = loadMaster(a)
		if err != nil {
			dialog.ShowError(err, topWindow)
		}
		topWindow.Show()
	}
	defer func() {
		if ready {
			share.UI.InfoBar.Close()
		}
	}()
	a.Run()
}

func loadMaster(a fyne.App) error {
	err := db.NewDatabase()
	if err != nil {
		return err
	}
	err = db.Conn.AutoMigrate(models.TimeEntry{})
	if err != nil {
		return err
	}
	file, err := files.FileOpen(confCustomFields, "r")
	if err != nil {
		return err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer file.Close()
	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	// time entry ui
	share.UI.TimeEntry = time_entry.NewScrollList()
	err = share.UI.TimeEntry.LoadCustomFields(buf)
	if err != nil {
		return err
	}
	err = share.UI.TimeEntry.ReloadAll()
	if err != nil {
		return err
	}

	// app ui theme and font
	font, err := ur.New(configs.Config.Font.Path)
	if err != nil {
		return err
	}
	a.Settings().SetTheme(theme.DarkMk2{
		Monospace:  font,
		Regular:    font,
		Bold:       font,
		Italic:     font,
		ItalicBold: font,
	})

	// master window settings
	topWindow.SetContent(masterBorder(topWindow))
	topWindow.SetMaster()
	return nil
}

func masterBorder(w fyne.Window) fyne.CanvasObject {
	toolbar := master.NewToolbar(w)
	toolbar.SetToTimeEntry()
	share.UI.Toolbar = toolbar
	infobar := master.NewInfoBar(w)
	share.UI.InfoBar = infobar
	return container.NewBorder(
		toolbar.Build(),
		infobar.Build(),
		nil,
		nil,
		share.UI.TimeEntry.Scroll(),
	)
}
