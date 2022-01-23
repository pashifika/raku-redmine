// Package dialog
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
package dialog

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func NewInputSingle(title, label string, high float32, w fyne.Window, callback func(s string)) {
	input := widget.NewEntry()
	if high > 0 {
		input.MultiLine = true
	} else {
		high = 40
	}
	dialog.ShowForm(title, "OK", "Cancel",
		[]*widget.FormItem{
			widget.NewFormItem(label, container.NewGridWrap(fyne.NewSize(w.Canvas().Size().Width/2, high), input)),
		},
		func(b bool) {
			if !b {
				return
			}
			callback(input.Text)
		}, w,
	)
}

func NewIssueTimeEntry(w fyne.Window, callback func(s string, isLast bool)) {
	input := widget.NewEntry()
	var isLast bool
	check := widget.NewCheck("", func(b bool) { isLast = b })
	dialog.ShowForm("New issue time entry", "OK", "Cancel",
		[]*widget.FormItem{
			widget.NewFormItem("ID or URL:", container.NewGridWrap(fyne.NewSize(w.Canvas().Size().Width/2, 40), input)),
			widget.NewFormItem("Add to last:", container.NewGridWrap(fyne.NewSize(w.Canvas().Size().Width/2, 40), check)),
		},
		func(b bool) {
			if !b {
				return
			}
			callback(input.Text, isLast)
		}, w,
	)
}
