// Package master
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
package master

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"raku-redmine/lib/types"
)

type InfoBarBuilder struct {
	_icon    *widget.Icon
	_text    *widget.TextGrid
	_msg     chan<- types.MsgData
	_control chan<- int
}

func NewInfoBar(_ fyne.Window) *InfoBarBuilder {
	return &InfoBarBuilder{
		_icon: widget.NewIcon(theme.InfoIcon()),
		_text: widget.NewTextGridFromString("init complete..."),
	}
}

func (b *InfoBarBuilder) Build() fyne.CanvasObject {
	rect := canvas.NewRectangle(&color.NRGBA{R: 128, G: 128, B: 128, A: 255})
	rect.SetMinSize(fyne.NewSize(2, 2))

	go b.listener()
	return container.NewVBox(rect, container.NewHBox(b._icon, b._text))
}

func (b *InfoBarBuilder) SendDebug(msg string) {
	b._msg <- types.MsgData{
		Type: types.MsgTypeDebug,
		Text: msg,
	}
}

func (b *InfoBarBuilder) SendWarning(msg string) {
	b._msg <- types.MsgData{
		Type: types.MsgTypeWarning,
		Text: msg,
	}
}

func (b *InfoBarBuilder) SendError(err error) {
	b._msg <- types.MsgData{
		Type: types.MsgTypeError,
		Text: err.Error(),
	}
}

func (b *InfoBarBuilder) SendInfo(msg string) {
	b._msg <- types.MsgData{
		Type: types.MsgTypeInfo,
		Text: msg,
	}
}

func (b *InfoBarBuilder) Close() {
	b._control <- 0
}

func (b *InfoBarBuilder) listener() {
	msg := make(chan types.MsgData)
	control := make(chan int)
	b._msg = msg
	b._control = control

	for {
	loop:
		select {
		case data, ok := <-msg:
			if ok && data.Text != "" {
				var icon fyne.Resource
				switch data.Type {
				case types.MsgTypeDebug, types.MsgTypeInfo:
					icon = theme.InfoIcon()
				case types.MsgTypeWarning:
					icon = theme.WarningIcon()
				case types.MsgTypeError:
					icon = theme.ErrorIcon()
				default:
					break loop
				}
				text := data.Text
				if len(text) > 48 {
					text = text[:45] + "..."
				}
				b._icon.SetResource(icon)
				b._text.SetText(text)
			}
		case signal, ok := <-control:
			if ok {
				switch signal {
				case 0:
					close(msg)
					close(control)
					println("InfoBar listener exit.")
					return
				}
			}
		}
	}
}
