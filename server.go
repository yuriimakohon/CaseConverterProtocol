package main

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

type server struct {
	commands chan command
}

func (s *server) start(address string) error {
	l, err := net.Listen("tcp", address)
	if err != nil {
		return errors.New("Server didn't start:\n\t" + err.Error())
	}
	go s.run()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
		} else {
			cli := newClient(0, conn, s.commands)
			go cli.listen()
		}
	}
}

func (s *server) run() {
	select {
	case cmd := <-s.commands:
		switch cmd.cmdType {
		case UP:
			s.convert(strings.ToUpper, cmd)
		case LOW:
			s.convert(strings.ToLower, cmd)
		}
	}
}

func (s *server) convert(f func(string) string, cmd command) {
	cmd.recipient.Write([]byte(f(string(cmd.body))))
}

func newServer() *server {
	return &server{
		commands: make(chan command),
	}
}
