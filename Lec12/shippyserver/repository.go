//Работа с моделью данных

package main

import (
	"context"

	pb "github.com/vlasove/Lec12/shippyserver/proto/consignment"
	"go.mongodb.org/mongo-driver/mongo"
)

//Container ...
type Container struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	UserID     string `json:"user_id"`
}

type Containers []*Container

//Consignment ...
type Consignment struct {
	ID          string     `json:"id"`
	Weight      int32      `json:"weight"`
	Description string     `json:"description"`
	Containers  Containers `json:"containers"`
	VesselID    string     `json:"vessel_id"`
}

func MarshalContainer(container *pb.Container) *Container {
	return &Container{
		ID:         container.Id,
		CustomerID: container.CustomerId,
		UserID:     container.UserId,
	}
}

func MarshalContainerCollection(containers []*pb.Container) []*Container {
	//К каждому контейнеру из RPC запроса будем применить некую функцию-преобразователь
	collection := make([]*Container, 0)
	for _, container := range containers {
		collection = append(collection, MarshalContainer(container))
	}

	return collection

}

func UnmarshalContainer(container *Container) *pb.Container {
	return &pb.Container{
		Id:         container.ID,
		CustomerId: container.CustomerID,
		UserId:     container.UserID,
	}
}

func UnmarshalContainerCollection(containers []*Container) []*pb.Container {
	collection := make([]*pb.Container, 0)
	for _, container := range containers {
		collection = append(collection, UnmarshalContainer(container))
	}

	return collection
}

func MarshalConsignment(consignment *pb.Consignment) *Consignment {
	containers := MarshalContainerCollection(consignment.Containers)
	return &Consignment{
		ID:          consignment.Id,
		Weight:      consignment.Weight,
		Description: consignment.Description,
		Containers:  containers,
		VesselID:    consignment.VesselId,
	}
}

func UnmarshalConsignment(consignment *Consignment) *pb.Consignment {
	return &pb.Consignment{
		Id:          consignment.ID,
		Weight:      consignment.Weight,
		Description: consignment.Description,
		Containers:  UnmarshalContainerCollection(consignment.Containers),
		VesselId:    consignment.VesselID,
	}
}

func UnmarshalConsignmentCollection(consignments []*Consignment) []*pb.Consignment {
	collection := make([]*pb.Consignment, 0)
	for _, consignment := range consignments {
		collection = append(collection, UnmarshalConsignment(consignment))
	}

	return collection
}

func MarshalConsignmentCollection(consignments []*pb.Consignment) []*Consignment {
	collection := make([]*Consignment, 0)
	for _, consignment := range consignments {
		collection = append(collection, MarshalConsignment(consignment))
	}

	return collection

}

// message Consignment {
// 	string id = 1;
// 	string description = 2;
// 	int32 weight = 3;
// 	repeated Container containers = 4;
// 	string vessel_id = 5;
//   }

//   message Container {
// 	string id = 1;
// 	string customer_id = 2;
// 	string origin = 3;
// 	string user_id = 4;
//   }

//Интерфейс МОДЕЛИ данных
type repository interface {
	Create(ctx context.Context, consignment *Consignment) error
	GetAll(ctx context.Context) ([]*Consignment, error)
}

type MongoRepository struct {
	collection *mongo.Collection
}

func (repo *MongoRepository) Create(ctx context.Context, consignment *Consignment) error {
	_, err := repo.collection.InsertOne(ctx, consignment)
	return err
}

func (repo *MongtoRepository) GetAll(ctx context.Context) ([]*Consignment, error) {
	cur, err := repo.collection.Find(ctx, nil, nil)
	var sampleConsignments []*Consignment

	for cur.Next(ctx) {
		var cons *Consignment
		if err := cur.Decode(&cons); err != nil {
			return nil, err
		}
	}

	return sampleConsignments, err
}
