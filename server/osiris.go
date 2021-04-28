package server

import (
	pb "github.com/lutece-awesome/osiris-next/codegen"
)

type OsirisServer struct {
	pb.UnimplementedOsirisServer
}

func NewServer() *OsirisServer {
	return &OsirisServer{}
}

func (OsirisServer) RunProgram(request *pb.RunProgramRequest, stream pb.Osiris_RunProgramServer) error {
	return nil
}
