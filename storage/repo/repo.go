package repo

import pb "github.com/Avtoelon/pkg/structs"

type CarsRepoInterface interface{
	CreateCar(car *pb.Car)(*pb.Car,error)
	UpdateCar(car *pb.Car)(*pb.Car,error)
	GetCar(id string)(*pb.Car,error)
	GetAllCars()([]*pb.Car,error)
	DeleteCar(id string)(*pb.Car,error)
}

type OutsideRepoInterface interface{
	CreateOutside(outside *pb.CreateOutside)(*pb.Outside,error)
	UpdateOutside(upOut *pb.Outside)(*pb.Outside,error)
	GetOutside(id string) (*pb.Outside, error)
	GetAllOutside() ([]*pb.Outside, error)
	DeletedOutside(id string) (*pb.Outside, error)
}