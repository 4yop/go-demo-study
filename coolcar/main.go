package main

import (
	trippd "coolcar/proto/gen/go"
	trip "coolcar/tripserver"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis,err := net.Listen("tcp","127.0.0.1:8081")
	if err != nil {
		log.Fatalf("监听出错:%v",err)
	}
	s := grpc.NewServer()
	trippd.RegisterTripServiceServer(s, trip.Service{})
	log.Fatal(s.Serve(lis))
}
