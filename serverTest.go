package main

import (
	"CaseConverterProtocol/ccp"
	"fmt"
)

func main() {
	s := ccp.NewServer()
	err := s.Start(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
