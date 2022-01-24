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
	"errors"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"raku-redmine/share"
	"raku-redmine/utils/configs"
)

type Containers struct {
	app   fyne.App
	label *widget.Label
	save  func() (bool, error)

	Conf *configs.Root
	TopW fyne.Window

	SaveToDisk func(conf *configs.Root) error
}

// pages defines the data structure
type pages struct {
	name  string
	title string
	view  func(w fyne.Window, conf *configs.Root) (fyne.CanvasObject, func() (bool, error))
}

func New() *Containers {
	return &Containers{
		Conf: &configs.Root{},
	}
}

func (c *Containers) Build() fyne.CanvasObject {
	c.app = fyne.CurrentApp()
	c.label = widget.NewLabel(_entryLabel)
	content := container.NewMax()
	setPages := func(t pages) {
		c.label.SetText(t.title)
		var obj fyne.CanvasObject
		obj, c.save = t.view(c.TopW, c.Conf)
		content.Objects = []fyne.CanvasObject{obj}
		content.Refresh()
	}
	apply := &widget.Button{
		Text:       "Apply",
		Icon:       theme.ConfirmIcon(),
		Importance: widget.HighImportance,
		OnTapped:   c.apply,
	}

	// Initialize setting items for the application
	entries := &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return pageIndex[uid]
		},
		IsBranch: func(uid string) bool {
			children, ok := pageIndex[uid]
			return ok && len(children) > 0
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel(_entryLabel)
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			t, ok := pageDatas[uid]
			if !ok {
				share.UI.InfoBar.SendError(errors.New("missing setting panel: " + uid))
				return
			}
			obj.(*widget.Label).SetText(t.name)
		},
		OnSelected: func(uid string) {
			if t, ok := pageDatas[uid]; ok {
				setPages(t)
			}
		},
	}
	entries.Select(_appGeneral)

	rgb := &color.NRGBA{R: 128, G: 128, B: 128, A: 255}
	rectV := canvas.NewRectangle(rgb)
	rectV.SetMinSize(fyne.NewSize(2, 2))
	rectH := canvas.NewRectangle(rgb)
	rectH.SetMinSize(fyne.NewSize(2, 2))
	return container.NewBorder(
		container.NewVBox(container.NewHBox(c.label, layout.NewSpacer(), apply), rectV),
		nil,
		container.NewHBox(entries, rectH),
		nil,
		container.NewScroll(content),
	)
}

func (c *Containers) apply() {
	if c.save != nil {
		changed, err := c.save()
		if err != nil {
			dialog.ShowError(err, c.TopW)
			return
		}
		if !changed {
			return
		}
		err = c.SaveToDisk(c.Conf)
		if err != nil {
			dialog.ShowError(err, c.TopW)
		} else {
			msg := `Restart the application to change the settings.`
			dialog.ShowInformation("No applied", msg, c.TopW)
		}
	}
}
