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
	//tcpConn := conn()
	//defer tcpConn.Close()
	//fmt.Println(tcpConn)





}


func getInputFileName() (string) {
	var filename string
	_,err := fmt.Scanln(&filename)
	if err != nil {
		fmt.Println("输入出错")
		panic(err)
	}
	fmt.Println(filename)
	return filename
}