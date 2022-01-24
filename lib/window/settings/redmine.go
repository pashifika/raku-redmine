// Package settings
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
package settings

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"raku-redmine/utils/configs"
)

func makeRedmine(_ fyne.Window, conf *configs.Root) (fyne.CanvasObject, func() (bool, error)) {
	redmineURL := EntryOfRedmineURL()
	redmineURL.SetText(conf.Redmine.MasterUrl)
	redmineKey := EntryOfRedmineKey()
	redmineKey.SetText(conf.Redmine.ApiKey)
	save := func() (bool, error) {
		var changed bool
		if err := redmineURL.Validate(); err != nil {
			return changed, err
		}
		if err := redmineKey.Validate(); err != nil {
			return changed, err
		}
		if conf.Redmine.MasterUrl != redmineURL.Text {
			conf.Redmine.MasterUrl = redmineURL.Text
			changed = true
		}
		if conf.Redmine.ApiKey != redmineKey.Text {
			conf.Redmine.ApiKey = redmineKey.Text
			changed = true
		}
		return changed, nil
	}

	return container.NewVBox(
		widget.NewLabel("Master Url:"),
		container.NewHBox(widget.NewLabel(" "), container.NewGridWrap(fyne.NewSize(435, 40), redmineURL)),
		widget.NewLabel("API key:"),
		container.NewHBox(widget.NewLabel(" "), container.NewGridWrap(fyne.NewSize(435, 40), redmineKey)),
	), save
}
