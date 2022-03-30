package ccp

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"net"
)

type client struct {
	conn     net.Conn
	commands chan<- command
}

func (s *server) listen() {
	for {
		msg, err := bufio.NewReader(s.conn).ReadBytes('\n')
		if err == io.EOF {
			return
		}

		if err != nil {
			panic(err)
		}

		s.handle(msg)
	}
}

func (s *server) handle(message []byte) {
	cmd := bytes.TrimSpace(bytes.Split(message, []byte(" "))[0])
	body := bytes.TrimSpace(bytes.TrimPrefix(message, cmd))
	cmd = bytes.ToUpper(cmd)

	if len(body) == 0 {
		s.err(errors.New("the string len must be greater than 0"))
		return
	}
	switch string(cmd) {
	case "UP":
		s.sendCMD(UP, body)
	case "LOW":
		s.sendCMD(LOW, body)
	case "CAMEL":
		s.sendCMD(CAMEL, body)
	default:
		s.err(errors.New("Unknown command: " + string(cmd)))
	}
}

func (s *server) sendCMD(cmdType TYPE, body []byte) {
	s.commands <- command{
		cmdType:   cmdType,
		body:      body,
		recipient: s.conn,
	}
}

func (s *server) err(e error) {
	s.conn.Write([]byte("Error: " + e.Error() + "\n"))
}
