// Package configs
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
package configs

import (
	"gopkg.in/ini.v1"
)

var Config *Root

// Load and parses from INI data sources map to Config.
func Load(path string) error {
	cfg, err := ini.Load(path)
	if err != nil {
		return err
	}
	if Config == nil {
		Config = &Root{Font: Font{}, Redmine: Redmine{}}
	}
	return cfg.MapTo(Config)
}

// Save Config data to file system.
func Save(path string) error {
	cfg := ini.Empty()
	err := cfg.ReflectFrom(Config)
	if err != nil {
		return err
	}
	return cfg.SaveTo(path)
}
