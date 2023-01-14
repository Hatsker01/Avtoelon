package postgres

import (
	"database/sql"
	"fmt"
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
	if err != nil {
		return nil, err
	}
	newCar := pb.Car{}
	fmt.Print(car)
	query := `INSERT INTO cars(id,category_id,model_id,body_id,date,price,auction,enginee,oil_id,transmission_id,milage,
		color_id,drive_unit_id,outside_id,optic_id,salon_id,media_id,options_id,additionally_id,add_info,region_id,city_id,
		phone,created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23,$24)
		 RETURNING id,category_id,model_id,body_id,date,price,auction,enginee,oil_id,transmission_id,milage,
		 color_id,drive_unit_id,outside_id,optic_id,salon_id,media_id,options_id,additionally_id,add_info,region_id,city_id,
		 phone,created_at`
	err = r.db.QueryRow(query, id, car.Category_Id, car.Model_Id, car.Body_Id, car.Date, car.Price, car.Auction,
		car.Enginee, car.Oil_Id, car.Transmission_id, car.Milage, car.Color_id, car.Drive_unit_id, pq.Array(car.Outside_Id), pq.Array(car.Optic_Id),
		pq.Array(car.Salon_Id), pq.Array(car.Media_Id), pq.Array(car.Options_Id), pq.Array(car.Additionally_Id),
		car.Add_Info, car.Region_Id, car.City_Id, car.Phone, time.Now().UTC()).Scan(
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
		pq.Array(&newCar.Additionally_Id),
		&newCar.Add_Info,
		&newCar.Region_Id,
		&newCar.City_Id,
		&newCar.Phone,
		&newCar.Created_at,
	)
	fmt.Println(newCar)
	if err != nil {
		return nil, err
	}
	return &newCar, nil

}

func (r *carsRepasitory) UpdateCar(upCar *pb.Car) (*pb.Car, error) {
	query := `UPDATE cars SET category_id=$2,model_id=$3,body_id=$4,date=$5,price=$6,auction=$7,
	enginee=$8,oil_id=$9,transmission_id=$10,milage=$11,color_id=$12,drive_unit_id=$13,outside_id=$14,
	optic_id=$15,salon_id=$16,media_id=$17,options_id=$18,additionally_id=$19,add_info=$20,region_id=$21,
	city_id=$22,phone=$23,updated_at=$24 where deleted_at is null RETURNING id,category_id,model_id,body_id,date,price,auction,enginee,oil_id,transmission_id,milage,
	color_id,drive_unit_id,outside_id,optic_id,salon_id,media_id,additionally_id,add_info,region_id,city_id,
	phone,created_at,updated_at`
	var updated_at sql.NullTime
	newCar := pb.Car{}
	err := r.db.QueryRow(query, upCar.Id, upCar.Category_Id, upCar.Model_Id, upCar.Body_Id, upCar.Date, upCar.Price, upCar.Auction,
		upCar.Enginee, upCar.Oil_Id, upCar.Transmission_id, upCar.Milage, upCar.Color_id, upCar.Drive_unit_id, pq.Array(upCar.Outside_Id), pq.Array(upCar.Optic_Id),
		pq.Array(upCar.Salon_Id), pq.Array(upCar.Media_Id), pq.Array(upCar.Options_Id), pq.Array(upCar.Additionally_Id),
		upCar.Add_Info, upCar.Region_Id, upCar.City_Id, upCar.Phone, time.Now().UTC()).Scan(
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
		pq.Array(&newCar.Additionally_Id),
		&newCar.Add_Info,
		&newCar.Region_Id,
		&newCar.City_Id,
		&newCar.Phone,
		&newCar.Created_at,
		&updated_at,
	)
	if err != nil {
		return nil, err
	}
	if updated_at.Valid {
		newCar.Updated_at = updated_at.Time.String()
	}
	return &newCar, nil
}

func (r *carsRepasitory) GetCar(id string) (*pb.Car, error) {
	query := `SELECT id,category_id,model_id,body_id,date,price,auction,enginee,oil_id,transmission_id,milage,
	color_id,drive_unit_id,outside_id,optic_id,salon_id,media_id,options_id,additionally_id,add_info,region_id,city_id,
	phone,created_at,updated_at from cars where deleted_at is null and id = $1`
	newCar := pb.Car{}
	var updated_at sql.NullTime
	err := r.db.QueryRow(query, id).Scan(
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
		&newCar.Outside_Id,
		&newCar.Optic_Id,
		&newCar.Salon_Id,
		&newCar.Media_Id,
		&newCar.Options_Id,
		&newCar.Additionally_Id,
		&newCar.Add_Info,
		&newCar.Region_Id,
		&newCar.City_Id,
		&newCar.Phone,
		&newCar.Created_at,
		&updated_at,
	)
	if err != nil {
		return nil, err
	}
	if updated_at.Valid {
		newCar.Updated_at = updated_at.Time.String()
	}
	return &newCar, nil
}

func (r *carsRepasitory) GetAllCars() ([]*pb.Car, error) {
	cars := []*pb.Car{}
	query := `SELECT id,category_id,model_id,body_id,date,price,auction,enginee,oil_id,transmission_id,milage,
	color_id,drive_unit_id,outside_id,optic_id,salon_id,media_id,additionally_id,add_info,region_id,city_id,
	phone,created_at,updated_at from cars where deleted_at is null`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		car := pb.Car{}
		var updated_at sql.NullTime
		err := rows.Scan(
			&car.Id,
			&car.Category_Id,
			&car.Model_Id,
			&car.Body_Id,
			&car.Date,
			&car.Price,
			&car.Auction,
			&car.Enginee,
			&car.Oil_Id,
			&car.Transmission_id,
			&car.Milage,
			&car.Color_id,
			&car.Drive_unit_id,
			pq.Array(&car.Outside_Id),
			pq.Array(&car.Optic_Id),
			pq.Array(&car.Salon_Id),
			pq.Array(&car.Media_Id),
			pq.Array(&car.Options_Id),
			pq.Array(&car.Additionally_Id),
			&car.Add_Info,
			&car.Region_Id,
			&car.City_Id,
			&car.Phone,
			&car.Created_at,
			&updated_at,
		)
		if err != nil {
			return nil, err
		}
		if updated_at.Valid {
			car.Updated_at = updated_at.Time.String()
		}
		cars = append(cars, &car)
	}
	return cars, nil
}

func (r *carsRepasitory) DeleteCar(id string) (*pb.Car, error) {
	car, err := r.GetCar(id)
	if err != nil {
		return nil, err
	}
	query := `UPDATE cars SET deleted_at=$2 where id=$1`
	_, err = r.db.Exec(query, id, time.Now().UTC())
	if err != nil {
		return nil, err
	}
	car.Deleted_at = time.Now().UTC().String()
	return car, nil

}
