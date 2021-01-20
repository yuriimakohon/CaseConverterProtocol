package main

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"net"
)

type ID int

type client struct {
	id       ID
	conn     net.Conn
	commands chan<- command
}

func (c *client) listen() error {
	for {
		msg, err := bufio.NewReader(c.conn).ReadBytes('\n')
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		c.handle(msg)
	}
}

func (c *client) handle(message []byte) {
	cmd := bytes.TrimSpace(bytes.Split(message, []byte(" "))[0])
	body := bytes.TrimSpace(bytes.TrimPrefix(message, cmd))
	cmd = bytes.ToUpper(cmd)

	switch string(cmd) {
	case "UP":
		c.convert(UP, body)
	case "LOW":
		c.convert(LOW, body)
	default:
		c.err(errors.New("Unknown command: " + string(cmd)))
	}
}

func (c *client) convert(cmdType TYPE, body []byte) {
	c.commands <- command{
		cmdType:   cmdType,
		body:      body,
		recipient: c.conn,
	}
}

func (c *client) err(e error) {
	c.conn.Write([]byte("Error: " + e.Error() + "\n"))
}

func newClient(id ID, conn net.Conn, commands chan<- command) *client {
	return &client{
		id:       id,
		conn:     conn,
		commands: commands,
	}
}
