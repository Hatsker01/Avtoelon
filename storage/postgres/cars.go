package postgres

import (
	"time"

	"github.com/Avtoelon/storage/repo"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"

	pb "github.com/Avtoelon/pkg/structs"
)

type carsRepasitory struct {
	db *sqlx.DB
}

func NewCarsRepo(db *sqlx.DB) repo.CarsRepoInterface {
	return &carsRepasitory{
		db: db,
	}
}

func (r *carsRepasitory) CreateCar(car *pb.Car) (*pb.Car, error) {

	id, err := uuid.NewV4()
	if err!=nil{
		return nil,err
	}
	newCar:=pb.Car{}
	query := `INSERT INTO cars(id,category_id,model_id,body_id,date,price,auction,enginee,oil_id,transmission_id,milage,
		color_id,drive_unit_id,outside_id,optic_id,salon_id,media_id,additionally_id,add_info,region_id,city_id,
		phone,created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24)
		 RETURNING id,category_id,model_id,body_id,date,price,auction,enginee,oil_id,transmission_id,milage,
		 color_id,drive_unit_id,outside_id,optic_id,salon_id,media_id,additionally_id,add_info,region_id,city_id,
		 phone,created_at`
	err = r.db.QueryRow(query,id,car.Category_Id,car.Model_Id,car.Body_Id,car.Date,car.Price,car.Auction,
	car.Enginee,car.Oil_Id,car.Transmission_id,car.Milage,car.Color_id,car.Drive_unit_id,pq.Array(car.Outside_Id),pq.Array(car.Optic_Id),
	pq.Array(car.Salon_Id),pq.Array(car.Media_Id),pq.Array(car.Options_Id),pq.Array(car.Additionally_Id),
	car.Add_Info,car.Region_Id,car.City_Id,car.Phone,time.Now().UTC()).Scan(
		&newCar.Id,
		&newCar.Category_Id,
		&newCar.Model_Id,
		&newCar.Body_Id,
		&newCar.Date,
		&newCar.Price,
		&newCar.Auction,
		&newCar.Enginee,
		&newCar.Oil_Id,
		&newCar.Transmission_id,
		&newCar.Milage,
		&newCar.Color_id,
		&newCar.Drive_unit_id,
		pq.Array(&newCar.Outside_Id),
		pq.Array(&newCar.Optic_Id),
		pq.Array(&newCar.Salon_Id),
		pq.Array(&newCar.Media_Id),
		pq.Array(&newCar.Options_Id),
		pq.Array(&car.Additionally_Id),
		&newCar.Add_Info,
		&newCar.Region_Id,
		&newCar.City_Id,
		&newCar.Phone,
		&newCar.Created_at,
	)
	return &newCar,nil

}
