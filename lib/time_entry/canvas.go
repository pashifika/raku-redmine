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
	"errors"
	"strconv"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"github.com/goccy/go-json"
	"go.uber.org/zap"

	lt "raku-redmine/lib/types"
	"raku-redmine/share"
	db "raku-redmine/utils/database"
	"raku-redmine/utils/database/models"
	"raku-redmine/utils/database/types"
	"raku-redmine/utils/log"
)

type ScrollList struct {
	_mu       sync.RWMutex
	scroll    *container.Scroll
	Last      types.CustomFields
	vbox      *fyne.Container
	timeEntry []*models.TimeEntry
}

func NewScrollList() *ScrollList {
	vbox := container.NewVBox()
	list := &ScrollList{
		_mu:       sync.RWMutex{},
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
	s._mu.Lock()
	d.Order = len(s.timeEntry)
	s.timeEntry = append(s.timeEntry, d)
	s.vbox.Add(NewPowerCard().Make(d))
	s._mu.Unlock()
}

// Prepend a new TimeEntry data to the start of this TimeEntry scroll ui
func (s *ScrollList) Prepend(d *models.TimeEntry) {
	s._mu.Lock()
	d.Order = 0
	for i := 0; i < len(s.timeEntry); i++ {
		s.timeEntry[i].Order++
	}
	s.timeEntry = append([]*models.TimeEntry{d}, s.timeEntry...)
	s.vbox.Objects = append([]fyne.CanvasObject{NewPowerCard().Make(d)}, s.vbox.Objects...)
	s._mu.Unlock()
	s.vbox.Refresh()
}

func (s *ScrollList) GetAll() []*models.TimeEntry {
	s._mu.RLock()
	defer s._mu.RUnlock()
	return s.timeEntry
}

func (s *ScrollList) LastCustomFields() types.CustomFields {
	s._mu.RLock()
	defer s._mu.RUnlock()
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

	s._mu.Lock()
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
	s._mu.Unlock()
	return nil
}

// LoadActivities unmarshal the redmine /enumerations/time_entry_activities.json API data to ActivityFields
func (s *ScrollList) LoadActivities() error {
	activities, err := share.UI.Client.TimeEntryActivities()
	if err != nil {
		return err
	}
	listLen := len(activities)
	if listLen == 0 {
		return nil
	}
	_activityFields = make([]*PossibleList, len(activities))
	for i, activity := range activities {
		if !activity.Active {
			continue
		}
		_activityFields[i] = &PossibleList{
			IsDefault: activity.IsDefault,
			ValueData: strconv.Itoa(activity.Id),
			LabelData: activity.Name,
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

	s._mu.Lock()
	defer s._mu.Unlock()
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
	if share.UI.InfoBar != nil {
		share.UI.InfoBar.SendInfo("time entry data is reloaded")
	}
	return nil
}

// SaveAll time entry data and to database and return the last database error
func (s *ScrollList) SaveAll() error {
	s._mu.RLock()
	var last error
	for _, data := range s.timeEntry {
		var count int64
		err := db.Conn.Model(data).Where(`id = ?`, data.UID).Count(&count).Error
		if err != nil {
			log.Error("ScrollList.SaveAll.Count", err.Error(),
				zap.String("id", data.UID.String()),
				zap.Int("issue_id", data.IssueId),
			)
			share.UI.InfoBar.SendError(err)
			last = err
			continue
		}
		if count == 0 {
			err = db.Conn.Create(data).Error
		} else {
			err = db.Conn.Omit("id,issue_id").Updates(data).Error
		}
		if err != nil {
			log.Error("ScrollList.SaveAll.CreateOrUpdates", err.Error(),
				zap.String("id", data.UID.String()),
				zap.Int("issue_id", data.IssueId),
			)
			share.UI.InfoBar.SendError(err)
			last = err
		} else {
			err = db.Conn.Create(models.MakeTimeEntryHistory(data)).Error
			if err != nil {
				log.Error("ScrollList.SaveAll.CreateHistory", err.Error(),
					zap.String("id", data.UID.String()),
					zap.Int("issue_id", data.IssueId),
				)
				share.UI.InfoBar.SendError(err)
				last = err
			}
		}
	}
	s._mu.RUnlock()
	share.UI.InfoBar.SendInfo("time entry data is saved")
	return last
}

func (s *ScrollList) DeleteNoChecked(check bool) {
	s._mu.Lock()
	defer s._mu.Unlock()
	s.findCheck(check, func(index int) {
		err := s.timeEntry[index].Checked.Set(false)
		if err != nil {
			log.Error("ScrollList.DeleteNoChecked.Checked", err.Error())
			share.UI.InfoBar.SendError(err)
			return
		}
		err = db.Conn.Delete(s.timeEntry[index]).Error
		if err != nil {
			log.Error("ScrollList.DeleteNoChecked.Delete", err.Error())
			share.UI.InfoBar.SendError(err)
			return
		}
		s.vbox.Objects[index].Hide()
	})
	share.UI.InfoBar.SendInfo("no checked items is delete")
}

func (s *ScrollList) PostChecked() {
	my, err := share.UI.Client.MyAccount()
	if err != nil {
		log.Error("ScrollList.PostChecked.MyAccount", err.Error())
		share.UI.InfoBar.SendError(errors.New("can not get my account data"))
		return
	}
	s._mu.RLock()
	defer s._mu.RUnlock()
	s.findCheck(true, func(index int) {
		data, err := models.MakeClientResponse(s.timeEntry[index], my.Id)
		if err != nil {
			log.Error("ScrollList.PostChecked.MakeClientResponse", err.Error())
			share.UI.InfoBar.SendError(err)
			return
		}
		_, err = share.UI.Client.CreateTimeEntry(data)
		if err != nil {
			log.Error("ScrollList.PostChecked.ClientTimeEntry", err.Error())
			share.UI.InfoBar.SendError(err)
			return
		}
	})
	share.UI.InfoBar.SendInfo("post to redmine time entry item complete")
}

func (s *ScrollList) findCheck(check bool, callback func(index int)) {
	for i, data := range s.timeEntry {
		checked, err := data.Checked.Get()
		if err != nil {
			log.Error("ScrollList.findCheck.Checked", err.Error(), zap.String("id", data.UID.String()))
			share.UI.InfoBar.SendError(err)
			continue
		}
		if checked == check && callback != nil {
			callback(i)
		}
	}
}
