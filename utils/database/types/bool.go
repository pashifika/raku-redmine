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

	"fyne.io/fyne/v2/data/binding"
)

type Bool struct {
	binding.Bool
}

// Value returns a driver must not panic.
func (b Bool) Value() (driver.Value, error) {
	if b.Bool == nil {
		return false, nil
	}
	str, err := b.Get()
	if err != nil {
		return nil, err
	}
	return str, err
}

// Scan assigns a value from a database driver.
func (b *Bool) Scan(src interface{}) error {
	var err error
	if src != nil {
		if val, ok := src.(bool); ok {
			if b.Bool == nil {
				b.Bool = binding.NewBool()
			}
			err = b.Set(val)
		} else {
			err = errors.New("can not scan value to binding.String")
		}
	}
	return err
}
