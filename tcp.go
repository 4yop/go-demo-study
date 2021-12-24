package main

import (
	"fmt"
	"net"
)

func handle(conn net.Conn) {
	defer conn.Close()
	fmt.Println(conn.RemoteAddr().Network())
	fmt.Println(conn.RemoteAddr().String())

	b := make([]byte,1024)
	for {
		len,err := conn.Read(b)
		if err != nil {
			fmt.Println(err)
			return
		}
		//len-1 不要\n
		msg := string(b[:len-1])
		fmt.Println(msg)

		if msg == "exit" {
			fmt.Println("退出")
			return
		}

		conn.Write([]byte("I amd server computer \n"))
	}


}


func main()  {
	//监听端口
	listener,err := net.Listen("tcp","127.0.0.1:7684")
	if err != nil {
		fmt.Println(err)
		return;
	}
	defer listener.Close()
	for  {
		conn,err :=listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handle(conn)
	}

}