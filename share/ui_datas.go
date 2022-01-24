// Package share
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
package share

import (
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"

	"raku-redmine/utils/database/models"
	"raku-redmine/utils/database/types"
	"raku-redmine/utils/redmine"
)

var UI AppUI

type AppUI struct {
	AppName    string
	AppVer     string
	UserDir    string
	SystemFont string
	OS         string
	ARCH       string
	Debug      bool
	Client     *redmine.Client
	Window     Window
	TimeEntry  TimeEntry
	Toolbar    Toolbar
	InfoBar    InfoBar
}

type Window interface {
	SetMaster(w fyne.Window)
	About()
	Setting()
}

type Toolbar interface {
	SetToTimeEntry()
}

type InfoBar interface {
	SendDebug(msg string)
	SendWarning(msg string)
	SendError(err error)
	SendInfo(msg string)
	Close()
}

type TimeEntry interface {
	Append(d *models.TimeEntry)
	Prepend(d *models.TimeEntry)
	LoadActivities() error
	LoadCustomFields(r io.Reader) error
	LastCustomFields() types.CustomFields
	Scroll() *container.Scroll
	ReloadAll() error
	SaveAll() error
	PostChecked()
	DeleteNoChecked(check bool)
}
