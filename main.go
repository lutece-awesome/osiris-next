package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/lutece-awesome/osiris-next/codegen"
	osiris "github.com/lutece-awesome/osiris-next/server"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 9876, "The gRPC server port")
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("osiris gRPC server listen to :%d\n", *port)
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterOsirisServer(grpcServer, osiris.NewServer())
	grpcServer.Serve(lis)
}
