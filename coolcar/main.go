package main

import (
	"context"
	trippd "coolcar/proto/gen/go"
	trip "coolcar/tripserver"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net"
	"net/http"
)

func main() {
	log.SetFlags(log.Lshortfile)
	go startGrpcGateway()
	lis, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatalf("监听出错:%v", err)
	}
	s := grpc.NewServer()
	fmt.Println(s)
	trippd.RegisterTripServiceServer(s, &trip.Service{})
	log.Fatal(s.Serve(lis))
}

func startGrpcGateway() {
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions:   protojson.MarshalOptions{
				UseProtoNames:     true,
				UseEnumNumbers:    true,
			},
		}))
	err := trippd.RegisterTripServiceHandlerFromEndpoint(
		c,
		mux,
		"127.0.0.1:8081",
		[]grpc.DialOption{grpc.WithInsecure()},
	)
	if err != nil {
		log.Fatalln("trippd.RegisterTripServiceHandlerFromEndpoint err: %v", err)
	}
	http.ListenAndServe(":8080", mux)
}
