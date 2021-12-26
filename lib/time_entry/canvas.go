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
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"github.com/goccy/go-json"

	lt "raku-redmine/lib/types"
	"raku-redmine/share"
	db "raku-redmine/utils/database"
	"raku-redmine/utils/database/models"
	"raku-redmine/utils/database/types"
)

type ScrollList struct {
	scroll    *container.Scroll
	Last      types.CustomFields
	vbox      *fyne.Container
	timeEntry []*models.TimeEntry
}

func NewScrollList() *ScrollList {
	vbox := container.NewVBox()
	list := &ScrollList{
		scroll:    container.NewVScroll(vbox),
		vbox:      vbox,
		timeEntry: []*models.TimeEntry{},
	}
	return list
}

// Scroll return time entry items scroll container
func (s *ScrollList) Scroll() *container.Scroll {
	return s.scroll
}

// Append a new TimeEntry data to the end of this TimeEntry scroll ui
func (s *ScrollList) Append(d *models.TimeEntry) {
	d.Order = len(s.timeEntry)
	s.timeEntry = append(s.timeEntry, d)
	s.vbox.Add(NewPowerCard().Make(d))
}

// Prepend a new TimeEntry data to the start of this TimeEntry scroll ui
func (s *ScrollList) Prepend(d *models.TimeEntry) {
	// TODO: fix the models.TimeEntry.Order
	d.Order = 0
	s.timeEntry = append([]*models.TimeEntry{d}, s.timeEntry...)
	s.vbox.Objects = append([]fyne.CanvasObject{NewPowerCard().Make(d)}, s.vbox.Objects...)
	s.vbox.Refresh()
}

func (s *ScrollList) GetAll() []*models.TimeEntry {
	return s.timeEntry
}

func (s *ScrollList) LastCustomFields() types.CustomFields {
	return s.Last
}

// LoadCustomFields unmarshal the redmine /custom_fields.json API data to default CustomFields
func (s *ScrollList) LoadCustomFields(data []byte) error {
	_customFields = map[int][]*PossibleList{}
	var fields struct {
		Data []*CustomField `json:"custom_fields"`
	}
	err := json.Unmarshal(data, &fields)
	if err != nil {
		return err
	}

	s.Last = make(types.CustomFields)
	for _, field := range fields.Data {
		if field.CustomizedType == "time_entry" && field.Visible {
			// TODO: switch interface
			switch field.FieldFormat {
			case "list":
				s.Last[field.Id] = &lt.FieldList{
					Name:     field.Name,
					Default:  field.DefaultValue,
					Required: field.IsRequired,
				}
				_customFields[field.Id] = field.PossibleValues
			}
		}
	}
	return nil
}

func (s *ScrollList) ReloadAll() error {
	var datas []*models.TimeEntry
	err := db.Conn.Order("order_id").Find(&datas).Error
	if err != nil {
		return err
	}

	dLen := len(datas)
	if dLen > 0 {
		s.timeEntry = datas
		s.vbox.Objects = make([]fyne.CanvasObject, dLen)
		for i, data := range datas {
			// update database saved data to last api data
			for id, field := range s.Last {
				if val, ok := data.CustomFields[id]; ok {
					val.Name = field.Name
					val.Default = field.Default
					val.Required = field.Required
				} else {
					data.CustomFields[id] = &lt.FieldList{
						BindingData: binding.NewString(),
						Name:        field.Name,
						Default:     field.Default,
						Required:    field.Required,
					}
					err = data.CustomFields[id].Set(field.Default)
					if err != nil {
						return err
					}
				}
			}
			// remove database old data
			for id := range data.CustomFields {
				if _, ok := s.Last[id]; !ok {
					delete(data.CustomFields, id)
				}
			}
			s.vbox.Objects[i] = NewPowerCard().Make(data)
		}
		s.vbox.Refresh()
	}
	return nil
}

func (s *ScrollList) SaveAll() error {
	for _, data := range s.timeEntry {
		var count int64
		err := db.Conn.Model(data).Where(`id = ?`, data.UID).Count(&count).Error
		if err != nil {
			share.UI.InfoBar.SendError(err)
		}
		if count == 0 {
			err = db.Conn.Create(data).Error
		} else {
			err = db.Conn.Debug().Omit("id,issue_id").Updates(data).Error
		}
		if err != nil {
			share.UI.InfoBar.SendError(err)
		}
	}
	return nil
}
