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

type CategoryRepoInterface interface{
	CreateCategory(category *pb.CategoryCreateReq)(*pb.Category,error)
	UpdateCategory(category *pb.Category)(*pb.Category,error)
	GetCategory(id string)(*pb.Category,error)
	GetAllCategory()([]*pb.Category,error)
	DeleteCategory(id string)(*pb.Category,error)
}

type ModelRepoInterface interface{
	CreateModel(model *pb.CreateModelReq)(*pb.Model,error)
	UpdateModel(upModel *pb.UpdateModel)(*pb.Model,error)
	GetModel(id string)(*pb.Model,error)
	GetAllModels()([]*pb.Model,error)
	DeleteModel(id string)(*pb.Model,error)
}