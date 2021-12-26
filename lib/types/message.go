// Package types
package types

type MsgData struct {
	Type MsgType
	Text string
}

type MsgType int

const (
	MsgTypeDebug MsgType = iota
	MsgTypeWarning
	MsgTypeError
	MsgTypeInfo
)
