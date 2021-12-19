// Package widgets
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
package widgets

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/mobile"
	"fyne.io/fyne/v2/widget"
)

type LimitEntry struct {
	widget.Entry
	decimal bool
	symbols []rune
	max     int
}

func NewLimitEntry(decimal bool, symbols []rune, max int) *LimitEntry {
	if max < 0 {
		max = 0
	}
	entry := &LimitEntry{decimal: decimal, symbols: symbols, max: max}
	entry.ExtendBaseWidget(entry)
	return entry._init()
}

func (e *LimitEntry) _init() *LimitEntry {
	e.Validator = nil
	return e
}

func (e *LimitEntry) BindString(s binding.String) *LimitEntry {
	e.Bind(s)
	return e._init()
}

func (e *LimitEntry) TypedRune(r rune) {
	if e.max > 0 && e.max <= e.CursorColumn {
		return
	}

	switch r {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		e.Entry.TypedRune(r)
	case '.':
		if e.decimal {
			e.Entry.TypedRune(r)
		}
	default:
		if e.symbols != nil {
			for _, sr := range e.symbols {
				if r == sr {
					e.Entry.TypedRune(r)
					break
				}
			}
		}
	}
}

func (e *LimitEntry) TypedShortcut(shortcut fyne.Shortcut) {
	paste, ok := shortcut.(*fyne.ShortcutPaste)
	if !ok {
		e.Entry.TypedShortcut(shortcut)
		return
	}

	content := paste.Clipboard.Content()
	if _, err := strconv.ParseFloat(content, 64); err == nil {
		e.Entry.TypedShortcut(shortcut)
	}
}

func (e *LimitEntry) Keyboard() mobile.KeyboardType {
	if e.decimal {
		return mobile.SingleLineKeyboard
	} else {
		return mobile.NumberKeyboard
	}
}
