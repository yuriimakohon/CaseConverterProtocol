package main

import "fmt"

func main() {
	s := newServer()
	err := s.start(":8081")
	if err != nil {
		fmt.Println(err)
	}
}
