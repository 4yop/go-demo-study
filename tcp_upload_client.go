package main

import (
	"fmt"
	"net"
)



func conn() net.Conn {
	tcpConn,err := net.Dial("tcp","127.0.0.1:789")
	if err != nil {
		fmt.Println(err)
		panic("tcp连接失败")
	}
	return tcpConn
}

func main() {
	tcpConn := conn()
	fmt.Println(tcpConn)

	fmt.Scan()

	defer tcpConn.Close()
}
