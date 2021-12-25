package main

import (
	"fmt"
	"net"
)

var conn net.Conn

func init() {
	conn,err := net.Dial("tcp","127.0.0.1:789")
	if err != nil {
		panic("tcp连接失败")
	}
}

func main() {

	fmt.Println(conn)

	defer conn.Close()
}
