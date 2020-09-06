package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/micro/go-micro"
	pb "github.com/vlasove/Lec12/shippyvessel/proto/vessel"
)

//Хост бд на случай если в процессе запуска не было обнаружено переменной окружения
//DB_HOST
const (
	defaultHost = "datastore:27017"
)

func main() {
	//Регистрируем наш сервер в MDNS
	service := micro.NewService(
		micro.Name("shippyvessel"),
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
	//Создаем коллекцию (аналог таблицы в РСУБД) в БД shippy с названием vessels (аналог названия таблицы)
	vesselCollection := client.Database("shippy").Collection("vessels")

	//Инициализация модели
	repo := &MongoRepository{vesselCollection}

	//Инициализация контроллера
	h := &handler{repo}
	//Регистрируем нащ сервер с контроллером
	pb.RegisterVesselServiceHandler(service.Server(), h)
	//Запускаем наш сервер
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
