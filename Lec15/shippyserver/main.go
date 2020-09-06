package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/micro/go-micro"
	pb "github.com/vlasove/Lec15/shippyserver/proto/consignment"
	vesselProto "github.com/vlasove/Lec15/shippyserver/proto/vessel"
	userProto "github.com/vlasove/Lec15/userserver/proto/user"
)

//Хост бд на случай если в процессе запуска не было обнаружено переменной окружения
//DB_HOST
const (
	defaultHost = "datastore:27017"
)
//Высокоуровневая функция, которая принимает на вход RPC запрос
//и возвращает функцию, которая может принимать 1) контекст 2) запрос и ответ
//Извлекая токен из контекста проверяем, валиден ли он, если ок - возвращаем функцию,
//если нет - валимся с ошибкой
func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc{
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("No token found in request")
		}
		//Если не будет находить в контексте токен - сделайте с заглавной буквы
		token := meta["token"]
		log.Println("Auth token:", token)

		//Процедура аутентификации
		authClient, err := userProto.NewUserService("userserver", client.DefaultClient)
		_, err := authClient.ValidateToken(context.Background(), &userProto.Token{
			Token : token,
		})
		if err != nil {
			return err
		}
		err = fn(ctx, req, resp)
		return err 
	}
}

func main() {
	//Регистрируем наш сервер в MDNS
	service := micro.NewService(
		micro.Name("shippyserver"),
		micro.Version("latest"),
		//Добавим middleware для проверки токена
		micro.WrapHandler(AuthWrapper)
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
