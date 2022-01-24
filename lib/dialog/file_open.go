// Package dialog
/*
 * Version: 1.0.0
 * Copyright (c) 2022. Pashifika
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
package dialog

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"go.uber.org/zap"

	"raku-redmine/utils/log"
)

func ShowFileOpen(w fyne.Window, extensions []string, callCanReader func(reader fyne.URIReadCloser) error) {
	fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil {
			dialog.ShowError(err, w)
			log.Error("ShowFileOpen", "NewFileOpen error", zap.Error(err))
			return
		}
		if reader == nil {
			return
		}
		if callCanReader != nil {
			err = callCanReader(reader)
			if err != nil {
				log.Error("ShowFileOpen", "callCanReader error", zap.Error(err))
				dialog.ShowError(err, w)
			}
		}
	}, w)
	fd.SetFilter(storage.NewExtensionFileFilter(extensions))
	fd.Show()
}
