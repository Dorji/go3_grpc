package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	pb "github.com/vlasove/Lec12/shippyvessel/proto/vessel"
)

//Модельный файл

//Repository ...
type repository interface {
	FindAvailable(ctx context.Context, spec *Specification) (*Vessel, error)
	Create(ctx context.Context, vessel *Vessel) error
}

//VesselRepository ...
type MongoRepository struct {
	collection *mongo.Collection
}

// message Specification {
// 	int32 capacity = 1;
// 	int32 max_weight = 2;
//   }
//Локальная реализация объекта, понтяного для монги
type Specification struct {
	Capacity  int32
	MaxWeight int32
}

// message Vessel {
// 	string id = 1;
// 	int32 capacity = 2;
// 	int32 max_weight = 3;
// 	string name = 4;
// 	bool available = 5;
// 	string owner_id = 6;
//   }
//Локальная реализация объекта Vessel, понятного для монги
type Vessel struct {
	ID        string
	Capacity  int32
	MaxWeight int32
	Available bool
	OwnerID   string
	Name      string
}

//Превращаем *pb.Specification в *Specification
func MarshallSpecification(spec *pb.Specification) *Specification {
	return &Specification{
		Capacity:  spec.Capacity,
		MaxWeight: spec.MaxWeight,
	}
}

//Превращаем *Specification в *pb.Specification
func UnmarshallSpecification(spec *Specification) *pb.Specification {
	return &pb.Specification{
		Capacity:  spec.Capacity,
		MaxWeight: spec.MaxWeight,
	}
}

//Превращаем *pb.Vessel в *Vessel
func MarshallVessel(vessel *pb.Vessel) *Vessel {
	return &Vessel{
		ID:        vessel.Id,
		Capacity:  vessel.Capacity,
		MaxWeight: vessel.MaxWeight,
		Name:      vessel.Name,
		Available: vessel.Available,
		OwnerID:   vessel.OwnerId,
	}
}

//Превращаем *Vessel в *pb.Vessel
func UnmarshallVessel(vessel *Vessel) *pb.Vessel {
	return &pb.Vessel{
		Id:        vessel.ID,
		Capacity:  vessel.Capacity,
		MaxWeight: vessel.MaxWeight,
		Name:      vessel.Name,
		Available: vessel.Available,
		OwnerId:   vessel.OwnerID,
	}
}

//FindAvailable ...
func (repo *MongoRepository) FindAvailable(ctx context.Context, spec *Specification) (*Vessel, error) {
	// for _, vessel := range repo.vessels {
	// 	if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
	// 		return vessel, nil
	// 	}
	// }
	// return nil, errors.New("No vessel found by that spec")

	//Фильтр для отбора подходящего vessel
	filter := bson.D{{
		"capacity",
		bson.D{{
			"$lte",
			spec.Capacity,
		}, {
			"$lte",
			spec.MaxWeight,
		}},
	}}

	vessel := &Vessel{}
	if err := repo.collection.FindOne(ctx, filter).Decode(vessel); err != nil {
		return nil, err
	}
	return vessel, nil
}

//Create ...
func (repo *MongoRepository) Create(ctx context.Context, vessel *Vessel) error {
	_, err := repo.collection.InsertOne(ctx, vessel)
	return err
}
