package main

import (
	"fmt"
	"net"
	"os"
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


	filename := getInputFileName()
	fmt.Println(filename)
	fileInfo,err :=os.Stat(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	if fileInfo.IsDir() {
		fmt.Println("这是目录")
	}

}


func getInputFileName() (string) {
	fmt.Println("请输入文件完整路径:")
	var filename string
	_,err := fmt.Scanln(&filename)
	if err != nil {
		fmt.Println("输入出错")
		panic(err)
	}
	return filename
}