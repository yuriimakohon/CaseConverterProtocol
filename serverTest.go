package main

import (
	"Case_Converter_Protocol/ccp"
	"fmt"
)

func main() {
	s := ccp.NewServer()
	err := s.Start(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
