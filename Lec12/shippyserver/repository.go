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

//MarshalContainer ...
//Принимает на вход объект *pb.Container и превращает его в *Container
//изложенный выше (для работы с mongodb)
func MarshalContainer(container *pb.Container) *Container {
	return &Container{
		ID:         container.Id,
		CustomerID: container.CustomerId,
		UserID:     container.UserId,
	}
}

//Применяет ко всем объектам из []*pb.Container функцию MarshallContainer
//возвращает []*Container который уже можно обработать модельными методами
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

//Объект, реализующий интерфейс модели
type MongoRepository struct {
	collection *mongo.Collection
}

//Create ...
//Метод принимает на вход ТО ЧТО НУЖНО СОХРАНИТЬ(В локальном модельном формате) и сохраняет его
func (repo *MongoRepository) Create(ctx context.Context, consignment *Consignment) error {
	_, err := repo.collection.InsertOne(ctx, consignment)
	return err
}

//GetAll ...
//Метод возвращает все ЛОКАЛЬНЫЕ МОДЕЛЬНЫЕ объекты + ошибку
func (repo *MongoRepository) GetAll(ctx context.Context) ([]*Consignment, error) {
	cur, err := repo.collection.Find(ctx, nil, nil)
	var sampleConsignments []*Consignment

	for cur.Next(ctx) {
		var cons *Consignment
		if err := cur.Decode(&cons); err != nil {
			//continue
			return nil, err
		}
		sampleConsignments = append(sampleConsignments, cons)
	}

	return sampleConsignments, err
}
