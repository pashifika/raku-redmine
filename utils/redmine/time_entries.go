// Package redmine
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
package redmine

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/goccy/go-json"
	"github.com/pashifika/util/conv"
)

type timeEntriesResult struct {
	TimeEntries []TimeEntry `json:"time_entries"`
}

type timeEntryResult struct {
	TimeEntry TimeEntry `json:"time_entry"`
}

type timeEntryRequest struct {
	TimeEntry TimeEntry `json:"time_entry"`
}

type TimeEntry struct {
	Id           int            `json:"id,omitempty"`
	Project      IdName         `json:"project,omitempty"`
	Issue        Id             `json:"issue"`
	IssueId      int            `json:"issue_id,omitempty"`
	User         IdName         `json:"user"`
	Activity     IdName         `json:"activity"`
	ActivityId   int            `json:"activity_id,omitempty"`
	Hours        float32        `json:"hours"`
	Comments     string         `json:"comments"`
	SpentOn      string         `json:"spent_on"`
	CreatedOn    string         `json:"created_on,omitempty"`
	UpdatedOn    string         `json:"updated_on,omitempty"`
	CustomFields []*CustomField `json:"custom_fields,omitempty"`
}

// TimeEntriesWithFilter send query and return parsed result
func (c *Client) TimeEntriesWithFilter(filter Filter) ([]TimeEntry, error) {
	uri, err := c.URLWithFilter("/time_entries.json", filter)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Redmine-API-Key", c.apikey)
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var r timeEntriesResult
	if res.StatusCode == 404 {
		return nil, errors.New("Not Found")
	}
	if res.StatusCode != 200 {
		var er errorsResult
		err = decoder.Decode(&er)
		if err == nil {
			err = errors.New(strings.Join(er.Errors, "\n"))
		}
	} else {
		err = decoder.Decode(&r)
	}
	if err != nil {
		return nil, err
	}
	return r.TimeEntries, nil
}

func (c *Client) TimeEntries(projectId int) ([]TimeEntry, error) {
	res, err := c.Get(c.endpoint + "/projects/" + strconv.Itoa(projectId) + "/time_entries.json?key=" + c.apikey + c.getPaginationClause())
	if err != nil {
		return nil, err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var r timeEntriesResult
	if res.StatusCode == 404 {
		return nil, errors.New("Not Found")
	}
	if res.StatusCode != 200 {
		var er errorsResult
		err = decoder.Decode(&er)
		if err == nil {
			err = errors.New(strings.Join(er.Errors, "\n"))
		}
	} else {
		err = decoder.Decode(&r)
	}
	if err != nil {
		return nil, err
	}
	return r.TimeEntries, nil
}

func (c *Client) TimeEntry(id int) (*TimeEntry, error) {
	res, err := c.Get(c.endpoint + "/time_entries/" + strconv.Itoa(id) + ".json?key=" + c.apikey)
	if err != nil {
		return nil, err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var r timeEntryResult
	if res.StatusCode == 404 {
		return nil, errors.New("Not Found")
	}
	if res.StatusCode != 200 {
		var er errorsResult
		err = decoder.Decode(&er)
		if err == nil {
			err = errors.New(strings.Join(er.Errors, "\n"))
		}
	} else {
		err = decoder.Decode(&r)
	}
	if err != nil {
		return nil, err
	}
	return &r.TimeEntry, nil
}

func (c *Client) CreateTimeEntry(timeEntry TimeEntry) (*TimeEntry, error) {
	var ir timeEntryRequest
	ir.TimeEntry = timeEntry
	if timeEntry.Issue.Id > 0 {
		ir.TimeEntry.IssueId = timeEntry.Issue.Id
		ir.TimeEntry.Project.Id = 0
	}
	if timeEntry.Activity.Id > 0 {
		ir.TimeEntry.ActivityId = timeEntry.Activity.Id
	}
	s, err := json.Marshal(ir)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST",
		c.endpoint+"/time_entries.json?key="+c.apikey,
		strings.NewReader(strings.Replace(conv.BytesToString(s), "\"project\":{},", "", 1)),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var r timeEntryResult
	if res.StatusCode != 201 {
		var er errorsResult
		err = decoder.Decode(&er)
		if err == nil {
			err = errors.New(strings.Join(er.Errors, "\n"))
		}
	} else {
		err = decoder.Decode(&r)
	}
	if err != nil {
		return nil, err
	}
	return &r.TimeEntry, nil
}

func (c *Client) UpdateTimeEntry(timeEntry TimeEntry) error {
	var ir timeEntryRequest
	ir.TimeEntry = timeEntry
	if timeEntry.Issue.Id > 0 {
		ir.TimeEntry.IssueId = timeEntry.Issue.Id
		ir.TimeEntry.Project.Id = 0
	}
	if timeEntry.Activity.Id > 0 {
		ir.TimeEntry.ActivityId = timeEntry.Activity.Id
	}
	s, err := json.Marshal(ir)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT",
		c.endpoint+"/time_entries/"+strconv.Itoa(timeEntry.Id)+".json?key="+c.apikey,
		strings.NewReader(strings.Replace(conv.BytesToString(s), "\"project\":{},", "", 1)),
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := c.Do(req)
	if err != nil {
		return err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return errors.New("Not Found")
	}
	if res.StatusCode != 200 {
		decoder := json.NewDecoder(res.Body)
		var er errorsResult
		err = decoder.Decode(&er)
		if err == nil {
			err = errors.New(strings.Join(er.Errors, "\n"))
		}
	}
	if err != nil {
		return err
	}
	return err
}

func (c *Client) DeleteTimeEntry(id int) error {
	req, err := http.NewRequest("DELETE", c.endpoint+"/time_entries/"+strconv.Itoa(id)+".json?key="+c.apikey, strings.NewReader(""))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := c.Do(req)
	if err != nil {
		return err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return errors.New("Not Found")
	}

	decoder := json.NewDecoder(res.Body)
	if res.StatusCode != 200 {
		var er errorsResult
		err = decoder.Decode(&er)
		if err == nil {
			err = errors.New(strings.Join(er.Errors, "\n"))
		}
	}
	return err
}
