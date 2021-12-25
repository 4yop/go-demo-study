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
	defer recover()
	tcpConn := conn()
	fmt.Println(tcpConn)

	var filename string
	n,err := fmt.Scanln(&filename)
	if err != nil {
		fmt.Println(err)
		return
	}


	defer tcpConn.Close()
}
