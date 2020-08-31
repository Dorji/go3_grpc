package main

import (
	"context"
	"errors"
	"log"

	"github.com/micro/go-micro/v2"
	pb "github.com/vlasove/Lec10/dop/proto/dop"
)

//Repository ... Репозитарный интерфейс (обрабатывает хранилище)
type Repository interface {
	FindAvailable(*pb.Request) (*pb.Place, error)
}

//PlacesRepository ...
type PlacesRepository struct {
	places []*pb.Place
}

//FindAvailable ...
func (repo *PlacesRepository) FindAvailable(req *pb.Request) (*pb.Place, error) {
	for _, place := range repo.places {
		if place.Capacity >= req.Capacity && place.MaxWeight >= req.MaxWeight && place.Available {
			return place, nil
		}
	}
	return nil, errors.New("No places found by this request")
}

type dopService struct {
	repo Repository
}

//FindAvailable ...
func (dS *dopService) FindAvailable(ctx context.Context, req *pb.Request, res *pb.Response) error {
	place, err := dS.repo.FindAvailable(req)
	if err != nil {
		return err
	}
	res.Place = place
	return nil
}

func main() {
	//В нашей БД есть 2 места
	places := []*pb.Place{
		&pb.Place{Id: "place001", Name: "Шкаф", MaxWeight: 50, Capacity: 400, Available: true},
		&pb.Place{Id: "place002", Name: "table", MaxWeight: 100, Capacity: 200, Available: false},
	}

	//	Подготавливаем хранилище мест под погрузку
	repo := &PlacesRepository{places}

	//Инициализируем серверную часть
	service := micro.NewService(
		micro.Name("dop"),
	)
	service.Init()

	//Инициализируем СЕРВИСНУЮ часть
	ourService := &dopService{repo}
	//Сопоставляем сервисную часть и серверную
	if err := pb.RegisterDopServiceHandler(service.Server(), ourService); err != nil {
		log.Panic(err)
	}

	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
