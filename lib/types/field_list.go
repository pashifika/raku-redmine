// Package types
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
package types

import (
	"sync"

	"fyne.io/fyne/v2/data/binding"
)

type FieldList struct {
	_mu         sync.Mutex
	BindingData binding.String `json:"-"`
	Value       string         `json:"value"`
	Name        string         `json:"-"`
	Default     string         `json:"-"`
	Required    bool           `json:"-"`
}

// AddListener attaches a new change listener to this DataItem.
// Listeners are called each time the data inside this DataItem changes.
// Additionally the listener will be triggered upon successful connection to get the current value.
func (fl *FieldList) AddListener(listener binding.DataListener) {
	fl.BindingData.AddListener(listener)
}

// RemoveListener will detach the specified change listener from the DataItem.
// Disconnected listener will no longer be triggered when changes occur.
func (fl *FieldList) RemoveListener(listener binding.DataListener) {
	fl.BindingData.RemoveListener(listener)
}

func (fl *FieldList) Get() (string, error) {
	return fl.BindingData.Get()
}

func (fl *FieldList) Set(val string) error {
	err := fl.BindingData.Set(val)
	if err == nil {
		fl._mu.Lock()
		fl.Value = val
		fl._mu.Unlock()
	}
	return err
}
