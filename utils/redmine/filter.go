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

import "strings"

type Filter struct {
	filters map[string]string
}

func NewFilter(args ...string) *Filter {
	f := &Filter{}
	if len(args)%2 == 0 {
		for i := 0; i < len(args); i += 2 {
			f.AddPair(args[i], args[i+1])
		}
	}
	return f
}

func (f *Filter) AddPair(key, value string) {
	if f.filters == nil {
		f.filters = make(map[string]string)
	}
	f.filters[key] = encode4Redmine(value)
}

func (f *Filter) ToURLParams() string {
	params := ""
	for k, v := range f.filters {
		params += "&" + k + "=" + v
	}
	return params
}

func encode4Redmine(s string) string {
	a := strings.Replace(s, ">", "%3E", -1)
	a = strings.Replace(a, "<", "%3C", -1)
	a = strings.Replace(a, "=", "%3D", -1)
	return a
}
