//Сопряжение gRPC вызовов с моделью данных
package main

import (
	"context"
	"log"

	pb "github.com/vlasove/Lec12/shippyserver/proto/consignment"
	vesselProto "github.com/vlasove/Lec12/shippyserver/proto/vessel"
)

type handler struct {
	repository
	vesselClient vesselProto.VesselServiceClient
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)
	if err != nil {
		return err
	}

	req.VesselId = vesselResponse.Vessel.Id

	if err := s.repo.Create(ctx, MarshalConsignment(req)); err != nil {
		return err
	}

	res.Created = true
	res.Consignment = req
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments, err := s.repo.GetAll(ctx)
	if err != nil {
		return err
	}
	res.Consignments = UnmarshalConsignmentCollection(consignments)
	return nil
}
