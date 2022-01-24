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
	"io"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/widget"
	"github.com/pashifika/util/files"
	"github.com/pashifika/util/mem"

	ld "raku-redmine/lib/dialog"
	"raku-redmine/lib/time_entry"
	"raku-redmine/share"
)

func EntryOfRedmineURL() *widget.Entry {
	return &widget.Entry{
		PlaceHolder: "https://(you Redmine or RedMica site)",
		Wrapping:    fyne.TextTruncate,
		Validator:   validation.NewRegexp(`^http[s]?://(?:[a-zA-Z]|[0-9]|[$-_@.&+]|[!*(),]){5,}$`, "not a valid url"),
	}
}

func EntryOfRedmineKey() *widget.Entry {
	return &widget.Entry{
		PlaceHolder: "My account -> API Key",
		Wrapping:    fyne.TextTruncate,
		Validator:   validation.NewRegexp(`^[0-9a-zA-Z]{10,}$`, "not a valid api key"),
	}
}

func EntryOfFontPath() *widget.Entry {
	return &widget.Entry{
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
}

func EntriesOfCustomFieldsLoad(w fyne.Window) (*mem.FakeIO, *widget.Check, *widget.Button) {
	jsonData := new(mem.FakeIO)
	check := &widget.Check{
		Text:    "",
		Checked: false,
	}
	check.Disable()
	return jsonData, check, &widget.Button{
		Text: "Load JSON data",
		OnTapped: func() {
			ld.ShowFileOpen(w, []string{".json"}, func(reader fyne.URIReadCloser) error {
				check.SetChecked(false)
				jsonData.Reset()
				_, err := io.Copy(jsonData, reader)
				if err != nil {
					return err
				}
				data, err := time_entry.LoadCustomFieldJSON(jsonData)
				if err != nil {
					return err
				}
				check.Text = "Field(" + strconv.Itoa(len(data)) + ")"
				check.SetChecked(true)
				return nil
			})
		},
	}
}
