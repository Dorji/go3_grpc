package main

import (
	"fmt"

	"github.com/micro/go-micro"
	pb "github.com/vlasove/Lec12/shippyserver/proto/consignment"
	vesselProto "github.com/vlasove/Lec12/shippyserver/proto/vessel"
)

func main() {

	repo := &Repository{}

	srv := micro.NewService(micro.Name("shippyserver"))
	srv.Init()
	vesselClient := vesselProto.NewVesselServiceClient("shippyvessel", srv.Client())
	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo, vesselClient})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

}
