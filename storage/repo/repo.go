package repo

import pb "github.com/Avtoelon/pkg/structs"

type CarsRepoInterface interface {
	CreateCar(car *pb.Car) (*pb.Car, error)
	UpdateCar(car *pb.Car) (*pb.Car, error)
	GetCar(id string) (*pb.Car, error)
	GetAllCars() ([]*pb.Car, error)
	DeleteCar(id string) (*pb.Car, error)
}

type OutsideRepoInterface interface {
	CreateOutside(outside *pb.CreateOutside) (*pb.Outside, error)
	UpdateOutside(upOut *pb.Outside) (*pb.Outside, error)
	GetOutside(id string) (*pb.Outside, error)
	GetAllOutside() ([]*pb.Outside, error)
	DeletedOutside(id string) (*pb.Outside, error)
}

type CategoryRepoInterface interface {
	CreateCategory(category *pb.CategoryCreateReq) (*pb.Category, error)
	UpdateCategory(category *pb.Category) (*pb.Category, error)
	GetCategory(id string) (*pb.Category, error)
	GetAllCategory() ([]*pb.Category, error)
	DeleteCategory(id string) (*pb.Category, error)
}

type ModelRepoInterface interface {
	CreateModel(model *pb.CreateModelReq) (*pb.Model, error)
	UpdateModel(upModel *pb.UpdateModel) (*pb.Model, error)
	GetModel(id string) (*pb.Model, error)
	GetAllModels() ([]*pb.Model, error)
	DeleteModel(id string) (*pb.Model, error)
}

type BodyRepoInterface interface {
	CreateBody(body *pb.CreateBody) (*pb.Body, error)
	UpdateBody(upBody *pb.UpdateBody) (*pb.Body, error)
	GetBody(id string) (*pb.Body, error)
	GetAllBody() ([]*pb.Body, error)
	DeleteBody(id string) (*pb.Body, error)
}

type OilRepoInterface interface {
	Create(oil *pb.CreateOil) (*pb.Oil, error)
	Update(upOil *pb.UpdateOil) (*pb.Oil, error)
	Get(id string) (*pb.Oil, error)
	GetAll() ([]*pb.Oil, error)
	Delete(id string) (*pb.Oil, error)
}

type TransmissionRepoInterface interface {
	Create(trans *pb.CreateTrans) (*pb.Transmission, error)
	Update(upTrans *pb.UpdateTrans) (*pb.Transmission, error)
	Get(id string) (*pb.Transmission, error)
	GetAll() ([]*pb.Transmission, error)
	Delete(id string) (*pb.Transmission, error)
}

type ColorRepoInterface interface {
	Create(color *pb.CreateColorReq) (*pb.Color, error)
	Update(upColor *pb.UpdateColor) (*pb.Color, error)
	Get(id string) (*pb.Color, error)
	GetAll() ([]*pb.Color, error)
	Delete(id string) (*pb.Color, error)
}

type DriveUnitRepoInterface interface {
	Create(dr *pb.DriveUnitCreateReq) (*pb.Drive_Unit, error)
	Update(dr *pb.UpdateDriveUnit) (*pb.Drive_Unit, error)
	Get(id string) (*pb.Drive_Unit, error)
	GetAll() ([]*pb.Drive_Unit, error)
	Delete(id string) (*pb.Drive_Unit, error)
}

type OpticRepoInterface interface {
	Create(optic *pb.CreateOptic) (*pb.Optic, error)
	Update(upOptic *pb.UpdateOpticReq) (*pb.Optic, error)
	Get(id string) (*pb.Optic, error)
	GetAll() ([]*pb.Optic, error)
	Delete(id string) (*pb.Optic, error)
}

type SalonRepoInterface interface {
	Create(salon *pb.CreateSalon) (*pb.Salon, error)
	Update(upSalon *pb.UpdateSalonReq) (*pb.Salon, error)
	Get(id string) (*pb.Salon, error)
	GetAll() ([]*pb.Salon, error)
	Delete(id string) (*pb.Salon, error)
}

type MediasRepoInterface interface {
	Create(media *pb.CreateMedia) (*pb.Media, error)
	Update(upMedia *pb.UpdateMediaReq) (*pb.Media, error)
	Get(id string) (*pb.Media, error)
	GetAll() ([]*pb.Media, error)
	Delete(id string) (*pb.Media, error)
}
