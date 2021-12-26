// Package resource
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
package resource

import (
	"errors"
	"io/ioutil"
	"path/filepath"

	"github.com/pashifika/util/files"
)

type Fake struct {
	path  string
	name  string
	cache []byte
}

func New(path string) (*Fake, error) {
	if !files.Exists(path) {
		return nil, errors.New("resource file do not exist")
	}
	f, err := files.FileOpen(path, "r")
	if err != nil {
		return nil, err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return &Fake{
		path:  filepath.Dir(path),
		name:  filepath.Base(path),
		cache: buf,
	}, nil
}

func (f *Fake) Name() string {
	return f.name
}

func (f *Fake) Content() []byte {
	return f.cache
}
