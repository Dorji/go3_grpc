package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/micro/go-micro"
	pb "github.com/vlasove/Lec12/shippyserver/proto/consignment"
	vesselProto "github.com/vlasove/Lec12/shippyserver/proto/vessel"
)

//Хост бд на случай если в процессе запуска не было обнаружено переменной окружения
//DB_HOST
const (
	defaultHost = "datastore:27017"
)

func main() {
	//Регистрируем наш сервер в MDNS
	service := micro.NewService(
		micro.Name("shippyserver"),
	)
	//захват переменных окружения и инициализация
	service.Init()

	//Пытаемся считать переменную окружения, а если ее нет - то используем default
	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}
	//Пытаемся подключиться к бд по uri с retry=0
	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())
	//Создаем коллекцию (аналог таблицы в РСУБД) в БД shippy с названием consignments (аналог названия таблицы)
	consignmentCollection := client.Database("shippy").Collection("consignments")

	//Инициализация модели
	repo := &MongoRepository{consignmentCollection}
	//Инициализация клиента для второго сервиса
	vesselClient := vesselProto.NewVesselServiceClient("shippyvessel", service.Client())
	//Инициализация контроллера
	h := &handler{repo, vesselClient}
	//Регистрируем нащ сервер с контроллером
	pb.RegisterShippingServiceHandler(service.Server(), h)
	//Запускаем наш сервер
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
