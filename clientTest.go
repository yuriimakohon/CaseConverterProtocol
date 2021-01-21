package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	c, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	clientSession(c)
}

func clientSession(c net.Conn) {
	for {
		// Read string to convert
		fmt.Print(">: ")
		req, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		// Send upper-case conversion request
		fmt.Fprintf(c, req)
		// Fetching response with converted string
		resp := make([]byte, 255)
		c.Read(resp)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf(">> %s\n", resp)
	}
}
