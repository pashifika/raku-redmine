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
	"bytes"
	"database/sql/driver"
	"errors"

	"fyne.io/fyne/v2/data/binding"
	"github.com/goccy/go-json"
	"github.com/pashifika/util/conv"

	lt "raku-redmine/lib/types"
)

// TODO: switch interface

type CustomFields map[int]*lt.FieldList

// Value returns a driver must not panic.
func (j CustomFields) Value() (driver.Value, error) {
	b, err := json.Marshal(j)
	return conv.BytesToString(b), err
}

// Scan assigns a value from a database driver.
func (j *CustomFields) Scan(src interface{}) error {
	var err error
	if src != nil {
		switch src.(type) {
		case string:
			err = decode(&j, conv.StringToBytes(src.(string)))
		case []byte:
			err = decode(&j, src.([]byte))
		default:
			err = errors.New("can not scan value to CustomFields")
		}
	} else {
		err = errors.New("can not scan nil to CustomFields")
	}
	if err == nil {
		for _, val := range *j {
			val.BindingData = binding.NewString()
			err = val.BindingData.Set(val.Value)
		}
	}

	return err
}

func decode(res interface{}, value []byte) error {
	decoder := json.NewDecoder(bytes.NewReader(value))
	decoder.UseNumber()
	return decoder.Decode(res)
}
