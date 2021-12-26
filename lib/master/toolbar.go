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
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	ld "raku-redmine/lib/dialog"
	"raku-redmine/lib/theme/icons"
	"raku-redmine/share"
	"raku-redmine/utils/database/models"
)

type ToolbarBuilder struct {
	*widget.Toolbar
	_topWindow fyne.Window
}

func NewToolbar(w fyne.Window) *ToolbarBuilder {
	return &ToolbarBuilder{Toolbar: widget.NewToolbar(), _topWindow: w}
}

func (t *ToolbarBuilder) Build() fyne.CanvasObject {
	rect := canvas.NewRectangle(&color.NRGBA{R: 128, G: 128, B: 128, A: 255})
	rect.SetMinSize(fyne.NewSize(2, 2))
	return container.NewVBox(t, rect)
}

func (t *ToolbarBuilder) SetToTimeEntry() {
	t.Items = []widget.ToolbarItem{
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			ld.NewInputSingle("New issue time entry", "ID or URL:", t._topWindow, func(s string) {
				issueId, err := strconv.Atoi(s)
				if err != nil {
					share.UI.InfoBar.SendError(err)
					return
				}
				share.UI.TimeEntry.Append(models.MakeTimeEntryUI(issueId, share.UI.TimeEntry.LastCustomFields()))
			})
		}),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
			err := share.UI.TimeEntry.ReloadAll()
			if err != nil {
				share.UI.InfoBar.SendError(err)
			}
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.DocumentSaveIcon(), func() {
			err := share.UI.TimeEntry.SaveAll()
			if err != nil {
				share.UI.InfoBar.SendError(err)
			}
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(icons.LoggerIconRes, func() {
			share.UI.InfoBar.SendWarning("test warning...")
		}),
		widget.NewToolbarAction(theme.HistoryIcon(), func() {
			dialog.ShowInformation("Sorry", "Coming soon...", t._topWindow)
		}),
	}
	t.Refresh()
}
