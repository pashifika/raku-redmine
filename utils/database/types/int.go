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
	"database/sql/driver"
	"errors"
	"strconv"

	"fyne.io/fyne/v2/data/binding"
)

type Int struct {
	binding.String
}

// Value returns a driver must not panic.
func (b Int) Value() (driver.Value, error) {
	if b.String == nil {
		return false, nil
	}
	str, err := b.Get()
	if err != nil {
		return nil, err
	}
	val, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return nil, err
	}
	return val, err
}

// Scan assigns a value from a database driver.
func (b *Int) Scan(src interface{}) error {
	var err error
	if b.String == nil {
		b.String = binding.NewString()
	}
	if src != nil {
		switch src.(type) {
		case int:
			err = b.Set(strconv.Itoa(src.(int)))
		case int32:
			err = b.Set(strconv.Itoa(int(src.(int32))))
		case int64:
			err = b.Set(strconv.FormatInt(src.(int64), 10))
		default:
			err = errors.New("can not scan value to binding.Int")
		}
	}
	return err
}
