package main

import "net"

type TYPE int

const (
	UP TYPE = iota
	LOW
)

type command struct {
	cmdType   TYPE
	body      []byte
	recipient net.Conn
}
