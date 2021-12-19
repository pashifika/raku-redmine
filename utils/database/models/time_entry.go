// Package models
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
package models

import (
	"time"

	"fyne.io/fyne/v2/data/binding"
	"github.com/pashifika/util/xid"

	"raku-redmine/lib"
	lt "raku-redmine/lib/types"
	"raku-redmine/utils/database/types"
)

type TimeEntry struct {
	UID          xid.ID             `gorm:"column:id;type:varchar(20);index;primaryKey;autoIncrement:false"`
	Checked      types.Bool         `gorm:"column:checked;type:boolean;default:false;not null"`
	IssueId      int                `gorm:"column:issue_id;type:integer;index;not null"`
	IssueTitle   types.String       `gorm:"column:issue_title;type:varchar(255);not null"`
	Date         types.StrToDate    `gorm:"column:date;type:timestamp;not null"`
	Time         types.StrToFloat   `gorm:"column:time;type:numeric;not null"`
	Comment      types.String       `gorm:"column:comment;type:varchar(512)"`
	Order        int                `gorm:"column:order_id;type:integer;not null"`
	CustomFields types.CustomFields `gorm:"column:custom_fields;type:json;default:{};not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

// TableName table rename func
func (TimeEntry) TableName() string {
	return "issue_time_entry"
}

// MakeTimeEntryUI make TimeEntry to time entry ui use data
func MakeTimeEntryUI(issueId int, last types.CustomFields) *TimeEntry {
	cf := make(types.CustomFields)
	for id, field := range last {
		cf[id] = &lt.FieldList{
			BindingData: binding.NewString(),
			Name:        field.Name,
			Default:     field.Default,
			Required:    field.Required,
		}
		if field.Default != "" {
			_ = cf[id].Set(field.Default)
		}
	}
	date := binding.NewString()
	_ = date.Set(time.Now().Format(lib.DateLayout))
	return &TimeEntry{
		UID:          xid.New(),
		Checked:      types.Bool{Bool: binding.NewBool()},
		IssueId:      issueId,
		IssueTitle:   types.String{String: binding.NewString()},
		Date:         types.StrToDate{String: date},
		Time:         types.StrToFloat{String: binding.NewString()},
		Comment:      types.String{String: binding.NewString()},
		Order:        0,
		CustomFields: cf,
	}
}
