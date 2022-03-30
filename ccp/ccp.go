package ccp

import "net"

type TYPE int

const (
	UP TYPE = iota
	LOW
	CAMEL
)

type command struct {
	cmdType   TYPE
	body      []byte
	recipient net.Conn
}
