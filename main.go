package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/lutece-awesome/osiris-next/codegen"
	constant "github.com/lutece-awesome/osiris-next/constant"
	osiris "github.com/lutece-awesome/osiris-next/server"
	"google.golang.org/grpc"
)


func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *constant.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("osiris grpc server listen to :%d\n", *constant.Port)
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterOsirisServer(grpcServer, osiris.NewServer())
	grpcServer.Serve(lis)
}
