package main

import (
	"context"

	pb "github.com/vlasove/Lec12/shippyvessel/proto/vessel"
)

//Контроллер

type handler struct {
	repository
}

func (s *handler) FindAvailable(ctx context.Context, spec *pb.Specification, res *pb.Response) error {
	//Найти доступное место под загрузку (вызвать FindAvailable у модели)
	vessel, err := s.repository.FindAvailable(ctx, MarshallSpecification(spec))
	if err != nil {
		return err
	}

	res.Vessel = UnmarshallVessel(vessel)
	return nil
}

func (s *handler) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	if err := s.repository.Create(ctx, MarshallVessel(req)); err != nil {
		return err
	}
	res.Vessel = req
	res.Created = true
	return nil
}
