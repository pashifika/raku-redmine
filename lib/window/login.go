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
	"errors"
	"net/url"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/pashifika/util/files"

	"raku-redmine/lib"
	"raku-redmine/share"
)

func Login(w fyne.Window, onProcessed func(masterUrl, apiKey, fontPath string) error) {
	w.CenterOnScreen()
	w.SetFixedSize(true)
	w.Resize(lib.LoginWindow)
	w.SetTitle(share.UI.AppName + " - Init")

	// input setting
	redmineURL := &widget.Entry{
		PlaceHolder: "https://(you Redmine or RedMica site)",
		Wrapping:    fyne.TextTruncate,
		Validator:   validation.NewRegexp(`^http[s]?://(?:[a-zA-Z]|[0-9]|[$-_@.&+]|[!*(),]){5,}$`, "not a valid url"),
	}
	redmineKey := &widget.Entry{
		PlaceHolder: "My account -> API Key",
		Wrapping:    fyne.TextTruncate,
		Validator:   validation.NewRegexp(`^[0-9a-zA-Z]{10,}$`, "not a valid api key"),
	}
	fontPath := &widget.Entry{
		Text:        share.UI.SystemFont,
		PlaceHolder: "(you ttf font file full path)",
		Wrapping:    fyne.TextTruncate,
		Validator: func(s string) error {
			if !files.Exists(s) {
				return errors.New("file does not exist")
			} else {
				return nil
			}
		},
	}

	bottom := container.NewVBox(
		widget.NewSeparator(),
		container.NewHBox(
			layout.NewSpacer(),
			widget.NewButtonWithIcon("Cancel", theme.CancelIcon(), func() {
				w.Close()
			}),
			&widget.Button{Text: "Apply", Icon: theme.ConfirmIcon(), Importance: widget.HighImportance, OnTapped: func() {
				if err := redmineURL.Validate(); err != nil {
					dialog.ShowError(err, w)
					return
				}
				if err := redmineKey.Validate(); err != nil {
					dialog.ShowError(err, w)
					return
				}
				if err := fontPath.Validate(); err != nil {
					dialog.ShowError(err, w)
					return
				}

				if onProcessed != nil {
					err := onProcessed(strings.TrimRight(redmineURL.Text, "/"), redmineKey.Text, fontPath.Text)
					if err != nil {
						dialog.ShowError(err, w)
						return
					}
				}
				w.Close()
			}},
		),
	)
	w.SetContent(container.NewBorder(
		nil, bottom, nil, nil,
		container.NewVBox(
			widget.NewSeparator(),
			widget.NewCard("Redmine setting", "", container.NewVBox(
				container.NewHBox(
					widget.NewLabel("Master Url:"),
					container.NewGridWrap(fyne.NewSize(340, 40), redmineURL),
				),
				container.NewHBox(
					widget.NewLabel("API key:"),
					container.NewGridWrap(fyne.NewSize(340, 40), redmineKey),
				),
			)),
			widget.NewCard("App setting", "", container.NewVBox(
				widget.NewCard("", "Recommended download URL:", container.NewVBox(
					container.NewHBox(widget.NewHyperlink("M PLUS Rounded 1c (Japanese)", urlMustParse(
						"https://fonts.google.com/specimen/M+PLUS+Rounded+1c?subset=japanese"))),
				)),
				container.NewHBox(
					widget.NewLabel("Font Path:"),
					container.NewGridWrap(fyne.NewSize(340, 40), fontPath),
				),
			)),
		),
	))
	w.Show()
}

func urlMustParse(rawURL string) *url.URL {
	_url, _ := url.Parse(rawURL)
	return _url
}
