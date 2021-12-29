package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	listener,err := net.Listen("tcp","127.0.0.1:789")
	if err != nil {
		fmt.Println(err)
		return;
	}
	defer listener.Close()
	b := make([]byte,1024)


	for {
		conn,err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		n,err := conn.Read(b)
		if err != nil {
			fmt.Println(err)
			return
		}
		msg := string(b[:n])
		fmt.Println(msg)
		recvFile(msg,conn)
	}
}



func recvFile (filename string,conn net.Conn) {

	file,err := os.Create("upload/"+filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	b := make([]byte,1024)

	for  {
		n,err := conn.Read(b)

		if err != nil {
			if err == io.EOF {
				fmt.Println("FINISH")
			}
			fmt.Println(err)
			return
		}
		file.Write(b[:n])
	}
	
}