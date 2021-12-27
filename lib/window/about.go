// Package window
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
package window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"raku-redmine/resource"
	"raku-redmine/share"
)

func (win *Window) About() {
	w := win.App.NewWindow("About")
	w.CenterOnScreen()
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(400, 355))

	bottom := container.NewVBox(
		widget.NewSeparator(),
		container.NewHBox(
			layout.NewSpacer(),
			&widget.Button{Text: "OK", Icon: theme.ConfirmIcon(), Importance: widget.HighImportance, OnTapped: func() {
				w.Close()
			}},
		),
	)
	w.SetContent(container.NewBorder(
		container.NewGridWrap(fyne.NewSize(10, 4)), bottom, nil, nil,
		container.NewHBox(
			container.NewVBox(
				container.NewGridWrap(fyne.NewSize(128, 20)),
				container.NewGridWrap(fyne.NewSize(128, 128), canvas.NewImageFromResource(resource.AppIconRes)),
			),
			widget.NewCard("Raku redmine", "easy add time entry to issues.", container.NewVBox(
				container.NewGridWrap(fyne.NewSize(248, 8)),
				container.NewHBox(
					widget.NewLabel("Home page:"),
					container.NewHBox(widget.NewHyperlink("Go to Github",
						urlMustParse("https://github.com/pashifika/raku-redmine"))),
				),
				container.NewHBox(
					widget.NewLabel("OS:"),
					container.NewHBox(widget.NewHyperlink(share.UI.OS+"/"+share.UI.ARCH, nil)),
				),
				container.NewHBox(
					widget.NewLabel("Version:"),
					container.NewHBox(widget.NewHyperlink(share.UI.AppVer, nil)),
				),
				container.NewGridWrap(fyne.NewSize(248, 12)),
				widget.NewLabel("   Copyright (c) 2021. Pashifika"),
			)),
		),
	))
	w.Show()
}
