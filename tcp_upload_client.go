package main

import (
	"fmt"
	"net"
)



func main() {

	conn,err := net.Dial("tcp","127.0.0.1:789")
	if err != nil {
		fmt.Println(conn)
	}
	defer conn.Close()
}
