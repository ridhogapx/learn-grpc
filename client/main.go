package main

import (
	"flag"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "Address gRPC")
)

type Movie struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Genre string `json:"genre"`
}

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic("failed to connect grpc server")
	}

	defer conn.Close()
}
