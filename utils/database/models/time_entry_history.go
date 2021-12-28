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

	"github.com/pashifika/util/xid"

	"raku-redmine/utils/database/types"
)

type TimeEntryHistory struct {
	ID uint `gorm:"primarykey"`

	TID          xid.ID             `gorm:"column:t_id;type:varchar(20);index"`
	IssueId      int                `gorm:"column:issue_id;type:integer;index;not null"`
	Date         types.StrToDate    `gorm:"column:date;type:timestamp;not null"`
	Time         types.StrToFloat   `gorm:"column:time;type:numeric;not null"`
	Comment      types.String       `gorm:"column:comment;type:varchar(512)"`
	CustomFields types.CustomFields `gorm:"column:custom_fields;type:json;default:{};not null"`

	CreatedAt time.Time
}

// TableName table rename func
func (TimeEntryHistory) TableName() string {
	return "issue_time_entry_history"
}

//
// ------ to ui
//

func MakeTimeEntryHistory(timeEntryId int, te *TimeEntry) *TimeEntryHistory {
	return &TimeEntryHistory{
		ID:           uint(timeEntryId),
		TID:          te.UID,
		IssueId:      te.IssueId,
		Date:         te.Date,
		Time:         te.Time,
		Comment:      te.Comment,
		CustomFields: te.CustomFields,
	}
}
