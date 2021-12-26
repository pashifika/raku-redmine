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
	"time"

	"fyne.io/fyne/v2/data/binding"

	"raku-redmine/lib"
)

type String struct {
	binding.String
}

// Value returns a driver must not panic.
func (s String) Value() (driver.Value, error) {
	if s.String == nil {
		return "", nil
	}
	str, err := s.Get()
	if err != nil {
		return nil, err
	}
	return str, err
}

// Scan assigns a value from a database driver.
func (s *String) Scan(src interface{}) error {
	var err error
	if s.String == nil {
		s.String = binding.NewString()
	}
	if src != nil {
		if val, ok := src.(string); ok {
			err = s.Set(val)
		} else {
			err = errors.New("can not scan value to binding.String")
		}
	}
	return err
}

//
// ------ StrToDate
//

type StrToDate struct {
	binding.String
}

// Value returns a driver must not panic.
func (s StrToDate) Value() (driver.Value, error) {
	if s.String == nil {
		return time.Now(), nil
	}
	str, err := s.Get()
	if err != nil {
		return nil, err
	}
	val, err := time.Parse(lib.DateLayout, str)
	if err != nil {
		return nil, err
	}
	return val, err
}

// Scan assigns a value from a database driver.
func (s *StrToDate) Scan(src interface{}) error {
	var err error
	if s.String == nil {
		s.String = binding.NewString()
	}
	if src != nil {
		if val, ok := src.(time.Time); ok {
			err = s.Set(val.Format(lib.DateLayout))
		} else {
			err = errors.New("can not scan value to binding.StrToDate")
		}
	}
	return err
}

//
// ------ StrToFloat
//

type StrToFloat struct {
	binding.String
}

// Value returns a driver must not panic.
func (s StrToFloat) Value() (driver.Value, error) {
	if s.String == nil {
		return 0, nil
	}
	str, err := s.Get()
	if err != nil {
		return nil, err
	}
	val, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return nil, err
	}
	return val, err
}

// Scan assigns a value from a database driver.
func (s *StrToFloat) Scan(src interface{}) error {
	var err error
	if s.String == nil {
		s.String = binding.NewString()
	}
	if src != nil {
		switch src.(type) {
		case float32:
			err = s.Set(strconv.FormatFloat(float64(src.(float32)), 'f', 2, 32))
		case float64:
			err = s.Set(strconv.FormatFloat(src.(float64), 'f', 2, 64))
		case int, int32:
			err = s.Set(strconv.Itoa(src.(int)))
		case int64:
			err = s.Set(strconv.FormatInt(src.(int64), 10))
		default:
			err = errors.New("can not scan value to binding.StrToFloat")
		}
	}
	return err
}
