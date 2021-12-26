// Package widgets
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
package widgets

import (
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"

	"raku-redmine/share"
)

func NewEntryWithData(data binding.String) *widget.Entry {
	entry := widget.NewEntry()
	oldData, err := data.Get()
	if err != nil {
		share.UI.InfoBar.SendError(err)
	} else {
		entry.SetText(oldData)
	}
	entry.OnChanged = func(s string) {
		err = data.Set(s)
		if err != nil {
			share.UI.InfoBar.SendError(err)
		}
	}
	return entry
}
