// Package time_entry
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
package time_entry

import (
	"image/color"
	"net/url"
	"sort"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"go.uber.org/zap"

	lt "raku-redmine/lib/types"
	"raku-redmine/lib/widgets"
	"raku-redmine/utils/database/models"
	"raku-redmine/utils/database/types"
	"raku-redmine/utils/log"
)

var (
	_inactiveColor = color.NRGBA{R: 0x55, G: 0x61, B: 0x78, A: 0xb9}
	_activeColor   = color.NRGBA{R: 0x72, G: 0xac, B: 0x87, A: 0xa3}
)

type PowerCard struct {
	label        *fyne.Container
	custom       *fyne.Container
	mainObj      *fyne.Container
	modeSwitch   *widget.Button
	SwitchAction func()
}

func NewPowerCard() *PowerCard {
	return &PowerCard{}
}

func (pc *PowerCard) Make(d *models.TimeEntry) fyne.CanvasObject {
	date := widgets.NewLimitEntry(false, []rune("-"), 10).BindString(d.Date)
	date.Validator = validation.NewRegexp(
		`^([2-9][0-9]{3})-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[0-1])$`,
		"not a valid date time",
	)

	pc.modeSwitch = &widget.Button{
		Text: "",
		Icon: theme.MenuDropUpIcon(),
		OnTapped: func() {
			pc.actionSwitch()
		},
	}
	pc.label = container.NewGridWithRows(2,
		container.NewHBox(
			container.NewGridWrap(fyne.NewSize(28, 40), widget.NewCheckWithData("", d.Checked)),
			makeIssuesHyperlink(d.IssueId),
			container.NewGridWrap(fyne.NewSize(310, 40), widgets.NewEntryWithData(d.IssueTitle)),
			layout.NewSpacer(),
			widget.NewLabel("Date:"),
			container.NewGridWrap(fyne.NewSize(134, 40), date),
			pc.modeSwitch,
		),
	)
	pc.label.Add(container.NewHBox(
		widget.NewLabel("Comment:"),
		container.NewGridWrap(fyne.NewSize(pc.label.Size().Width-224, 40), widgets.NewEntryWithData(d.Comment)),
		layout.NewSpacer(),
		widget.NewLabel("Time:"),
		container.NewGridWrap(
			fyne.NewSize(54, 40),
			widgets.NewLimitEntry(true, nil, 4).BindString(d.Time),
		),
	))

	pc.custom = container.NewGridWithColumns(2)
	_activityLen := len(_activityFields)
	if _activityLen > 0 {
		var height float32 = 0
		if _activityLen >= 10 {
			height = date.Size().Height
		}
		value, _ := d.Activity.Get()
		if value == "" {
			for _, field := range _activityFields {
				if field.IsDefault {
					value = field.ValueData
					break
				}
			}
			_ = d.Activity.Set(value)
		}
		pc.custom.Add(makeSelectBox(height, _activityFields, &lt.FieldList{
			BindingData: d.Activity,
			Value:       value,
			Name:        "Activities",
			Default:     value,
			Required:    false,
		}))
	}
	// TODO: switch interface
	for _, id := range getCustomFieldKeys(d.CustomFields) {
		var height float32 = 0
		if len(_customFields[id]) >= 10 {
			height = date.Size().Height
		}
		pc.custom.Add(makeSelectBox(height, _customFields[id], d.CustomFields[id]))
	}
	pc.custom.Hide()

	pc.mainObj = container.NewBorder(
		_rectangle(),
		_rectangle(),
		_rectangle(),
		_rectangle(),
		container.NewVBox(pc.label, pc.custom),
	)
	return pc.mainObj
}

func (pc *PowerCard) actionSwitch() {
	if pc.SwitchAction != nil {
		pc.SwitchAction()
	}
	if pc.custom.Visible() {
		pc.modeSwitch.SetIcon(theme.MenuDropUpIcon())
		pc.custom.Hide()
	} else {
		pc.modeSwitch.SetIcon(theme.MenuDropDownIcon())
		pc.custom.Show()
	}
}

func _rectangle() *canvas.Rectangle {
	rect := canvas.NewRectangle(_inactiveColor)
	rect.SetMinSize(fyne.NewSize(1, 1))
	return rect
}

func makeSelectBox(height float32, items []*PossibleList, field *lt.FieldList) *fyne.Container {
	box := container.NewHBox()
	space := " "
	if field.Required {
		space = "※"
	}
	box.Add(container.NewVBox(container.NewGridWrap(
		fyne.NewSize(26, 36),
		widget.NewLabelWithStyle(space, fyne.TextAlignTrailing, fyne.TextStyle{}),
	)))
	box.Add(widget.NewLabel(field.Name + ":"))
	box.Add(widget.NewSelectEx(
		makeSelectOptions(items),
		"--- Select one ---",
		field, height,
		func(opt widget.SelectOption) {
			log.Debug("NewSelectEx.OnChanged", "opt",
				zap.String("Label", opt.Label()),
				zap.String("Value", opt.Value()),
			)
		},
	).SyncSelected())
	return box
}

func getCustomFieldKeys(cf types.CustomFields) []int {
	ids := make([]int, 0, len(cf))
	for id := range cf {
		ids = append(ids, id)
	}
	sort.Ints(ids)
	return ids
}

func makeSelectOptions(items []*PossibleList) []widget.SelectOption {
	opts := make([]widget.SelectOption, len(items))
	for i, item := range items {
		opts[i] = item
	}
	return opts
}

func makeIssuesHyperlink(id int) *widget.Hyperlink {
	idStr := strconv.Itoa(id)
	link, err := url.Parse("http://localhost/issues/" + idStr)
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}
	return widget.NewHyperlink("#"+idStr, link)
}
