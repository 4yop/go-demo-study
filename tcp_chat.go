package main

import (
	"fmt"
	"net"
)

type Client struct {
	Address string
	C chan string
	Conn net.Conn
}

var addressMap map[string]Client = make(map[string]Client)

var m = make(chan string)

func main() {

	listener,err := net.Listen("tcp","127.0.0.1:7812")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	for  {
		conn,err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go addClient(conn)
		go checkStatus(conn)

	}

}

//检查状态
func checkStatus(conn net.Conn){
	addr := conn.RemoteAddr().String()
	b := make([]byte,1024)
	for  {
		_,err := conn.Read(b)
		if err != nil{
			fmt.Println(addr,"下线")
			delete(addressMap,addr)
			go sendMsgToAll(addr,"下线了")
			break;
		}
	}
}

func addClient(conn net.Conn){
	addr := conn.RemoteAddr().String()
	fmt.Println(addr)
	var cli Client
	cli = Client{
		Address: addr,
		C:       make(chan string),
		Conn: conn,
	}
	addressMap[addr] = cli
	go sendMsgToAll(conn.RemoteAddr().String(),"上线了\n")
}

func sendMsgToAll(address string,msg string) {

	for _, cli := range addressMap {
		fmt.Println("给"+cli.Address+":"+address + msg)
		go func(c Client) {
			c.Conn.Write([]byte(address + msg))
		}(cli)

	}

}
