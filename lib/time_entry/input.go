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

var (
	_customFields   map[int][]*PossibleList // TODO: switch interface
	_activityFields []*PossibleList
)

type CustomField struct {
	Id             int             `json:"id"`
	Name           string          `json:"name"`
	CustomizedType string          `json:"customized_type"` // = time_entry
	FieldFormat    string          `json:"field_format"`    // = list
	Regexp         string          `json:"regexp"`
	IsRequired     bool            `json:"is_required"`
	IsFilter       bool            `json:"is_filter"`
	Searchable     bool            `json:"searchable"`
	Multiple       bool            `json:"multiple"`
	DefaultValue   string          `json:"default_value"`
	Visible        bool            `json:"visible"` // = true
	PossibleValues []*PossibleList `json:"possible_values"`
}

type PossibleList struct {
	IsDefault bool   `json:"-"`
	ValueData string `json:"value"`
	LabelData string `json:"label"`
}

func (p *PossibleList) Label() string {
	return p.LabelData
}

func (p *PossibleList) Value() string {
	return p.ValueData
}
