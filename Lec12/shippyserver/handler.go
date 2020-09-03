//Сопряжение gRPC вызовов с моделью данных
//файл для конфигурации контроллера
package main

import (
	"context"
	"log"

	pb "github.com/vlasove/Lec12/shippyserver/proto/consignment"
	vesselProto "github.com/vlasove/Lec12/shippyserver/proto/vessel"
)

//Создаем объект удовлетворяющий интерфейсу контроллера
type handler struct {
	//связывает с моделью данных (repository.go)
	repository
	//связывает этот сервер с сервисом vessel
	vesselClient vesselProto.VesselServiceClient
}

//CreateConsignment ...
//Реализация метода CreateConsignment для удовлетворения gRPC метода
//1) Запрос на сервис vessel - можно ли сохранить Consignment с параметрами Weight и набором Контейнеров
//2) Если можно - пробиваем id нашего Consignment
//3) Сохраняем НАШ CONSIGNMENT в БД (вызывая модельный метод Create())
func (s *handler) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	//1) Запрос на сервис vessel - можно ли сохранить Consignment с параметрами Weight и набором Контейнеров
	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)
	if err != nil {
		return err
	}
	//2) Если можно - пробиваем id нашего Consignment
	req.VesselId = vesselResponse.Vessel.Id
	//3) Сохраняем НАШ CONSIGNMENT в БД (вызывая модельный метод Create())
	//предварительно преобразовав *pb.Consignment в модельный *Consignment
	if err := s.repository.Create(ctx, MarshalConsignment(req)); err != nil {
		return err
	}

	res.Created = true
	res.Consignment = req
	return nil
}

//Реализация GetConsignments - gRPC метод (для удовлетворения интерфейса контроллера)
//1) Задаем запрос БД - вернуть все что там есть
//2) БД возвращает модельные []*Consignment \
//3) А gRPC нужны []*pb.Consignment
//4) Нужно превратить []*Consignment в []*pb.Consignment
func (s *handler) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	//1) Задаем запрос БД - вернуть все что там есть
	//2) БД возвращает модельные []*Consignment
	consignments, err := s.repository.GetAll(ctx)
	if err != nil {
		return err
	}
	//3) А gRPC нужны []*pb.Consignment
	//4) Нужно превратить []*Consignment в []*pb.Consignment
	res.Consignments = UnmarshalConsignmentCollection(consignments)
	return nil
}
