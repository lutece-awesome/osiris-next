package main

import (
	"fmt"
	"net"

	pb "github.com/lutece-awesome/osiris-next/codegen"
	constant "github.com/lutece-awesome/osiris-next/constant"
	osiris "github.com/lutece-awesome/osiris-next/server"
	util "github.com/lutece-awesome/osiris-next/util"
	"google.golang.org/grpc"

	log "github.com/sirupsen/logrus"
)

func main() {
	util.MakeDirectoryIfNotExists(*constant.Work)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *constant.Port))
	if err != nil {
		log.Error("failed to listen: %v", err)
	}
	log.Info("osiris gRPC server listen to ", *constant.Port)
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterOsirisServer(grpcServer, osiris.NewServer())
	grpcServer.Serve(lis)
}
