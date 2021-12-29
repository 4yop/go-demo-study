package main

import (
	"fmt"
	"io"
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
		return
	}
	fmt.Println(fileInfo.Name())


	tcpConn := conn()
	defer tcpConn.Close()


	tcpConn.Write([]byte(fileInfo.Name()))
	sendFile(filename,tcpConn)

}



func sendFile(filename string,conn net.Conn) {
	file,err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	b := make([]byte,1024)
	for  {
		n,err := file.Read(b)
		if err == io.EOF {
			fmt.Println("读完了")
			return
		}
		fmt.Println(string(b[:n]))
		_,err = conn.Write(b[:n])
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

//获取输入的文件名
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