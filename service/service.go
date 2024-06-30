package service

import (
	"context"
	pb "garden-managment-service/generated"
	postgres "garden-managment-service/storage/postgres"
)

type Server struct {
	pb.UnimplementedGardenManagementServer
	Garden *postgres.GardenManagementRepo
}

func (s *Server) CreateGarden(ctx context.Context, in *pb.CreateGardenRequest) (*pb.CreateGardenResponces, error) {
	resp, err := s.Garden.CreateGarden(in)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *Server) ViewGarden(ctx context.Context, in *pb.ViewGardenRequest) (*pb.ViewGardenResponces, error) {
	resp, err := s.Garden.ViewGarden(in)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *Server) UpdateGarden(ctx context.Context, in *pb.UpdateGardenRequest) (*pb.UpdateGardenResponces, error) {
	resp, err := s.Garden.UpdateGarden(in)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *Server) DeleteGarden(ctx context.Context,in *pb.DeleteGardenRequest) (*pb.DeleteGardenResponces,error) {
	resp,err:=s.Garden.DeleteGarden(in)

	if err != nil {
		return nil, err
	}

	return resp,nil
}

func (s *Server) ViewUserGardens(ctx context.Context,in *pb.ViewUserGardensRequest) (*pb.ViewUserGardensResponces,error)  {
	resp,err:=s.Garden.ViewUserGardens(in)

	if err != nil {
		return nil, err
	}

	return resp,nil	
}


