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
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"go.uber.org/zap"

	ld "raku-redmine/lib/dialog"
	"raku-redmine/lib/theme/icons"
	"raku-redmine/share"
	"raku-redmine/utils/database/models"
	"raku-redmine/utils/log"
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
			ld.NewIssueTimeEntry(t._topWindow, func(s string, isLast bool) {
				if len(s) == 0 {
					return
				}
				raws := strings.Split(s, "/")
				inputId := raws[len(raws)-1]
				issueId, err := strconv.Atoi(inputId)
				if err != nil {
					log.Error("SetToTimeEntry.ContentAdd", err.Error(), zap.String("input", inputId))
					share.UI.InfoBar.SendError(err)
					return
				}
				// get issue data
				share.UI.InfoBar.SendInfo("now get redmine issue title...")
				res, err := share.UI.Client.Issue(issueId)
				if err != nil {
					log.Error("SetToTimeEntry.ContentAdd.Client", err.Error(), zap.String("input", inputId))
					share.UI.InfoBar.SendError(err)
					return
				}
				// add to ui
				teUI := models.MakeTimeEntryUI(
					res.Project.Id, issueId, res.Subject, share.UI.TimeEntry.LastCustomFields(),
				)
				if isLast {
					share.UI.TimeEntry.Append(teUI)
				} else {
					share.UI.TimeEntry.Prepend(teUI)
				}
				share.UI.InfoBar.SendInfo("added time entry item.")
			})
		}),
		widget.NewToolbarAction(theme.DeleteIcon(), func() {
			dialog.ShowConfirm("WARNING", "Are you sure delete not checked items?", func(b bool) {
				if b {
					share.UI.TimeEntry.DeleteNoChecked(false)
				}
			}, t._topWindow)
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.UploadIcon(), func() {
			dialog.ShowConfirm("WARNING", "Are you sure post checked items?", func(b bool) {
				if b {
					share.UI.TimeEntry.PostChecked()
				}
			}, t._topWindow)
		}),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
			dialog.ShowConfirm("WARNING", "Are you sure reload all?", func(b bool) {
				if b {
					err := share.UI.TimeEntry.ReloadAll()
					if err != nil {
						log.Error("SendError.ViewRefresh", err.Error(), zap.Error(err))
						share.UI.InfoBar.SendError(err)
					}
				}
			}, t._topWindow)
		}),
		widget.NewToolbarAction(theme.DocumentSaveIcon(), func() {
			err := share.UI.TimeEntry.SaveAll()
			if err != nil {
				// TODO: show error log window
			}
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(icons.LoggerIconRes, func() {
			dialog.ShowInformation("Sorry", "Coming soon...", t._topWindow)
		}),
		widget.NewToolbarAction(theme.HistoryIcon(), func() {
			dialog.ShowInformation("Sorry", "Coming soon...", t._topWindow)
		}),
	}
	t.Refresh()
}
