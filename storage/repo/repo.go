package repo

import pb "github.com/Avtoelon/pkg/structs"

type CarsRepoInterface interface{
	CreateCar(car *pb.Car)(*pb.Car,error)
	UpdateCar(car *pb.Car)(*pb.Car,error)
	GetCar(id string)(*pb.Car,error)
	GetAllCars()([]*pb.Car,error)
	DeleteCar(id string)(*pb.Car,error)
}