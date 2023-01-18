package postgres

import (
	"database/sql"
	"time"

	"github.com/Avtoelon/storage/repo"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"

	pb "github.com/Avtoelon/pkg/structs"
)

type citiesRepasitory struct {
	db *sqlx.DB
}

func NewCitiesRepo(db *sqlx.DB) repo.CitiesRepoInterface {
	return &citiesRepasitory{
		db: db,
	}
}

func (r *citiesRepasitory) Create(city *pb.CreateCity) (*pb.City, error) {
	query := `INSERT INTO city(name,created_at) VALUES ($1,$2) RETURNING id,name,created_at`
	newCity := pb.City{}
	err := r.db.QueryRow(query, city.Name, time.Now().UTC()).Scan(
		&newCity.Id,
		&newCity.Name,
		&newCity.Created_at,
	)
	if err != nil {
		return nil, err
	}
	return &newCity, nil
}

func (r *citiesRepasitory) Update(upCity *pb.UpdateCityReq) (*pb.City, error) {
	city := pb.City{}
	query := `UPDATE city SET name=$2,updated_at=$3 where id=$1 and deleted_at is null RETURNING id,name,created_at,updated_at`
	err := r.db.QueryRow(query, upCity.Id, upCity.Name, time.Now().UTC()).Scan(
		&city.Id,
		&city.Name,
		&city.Created_at,
		&city.Updated_at,
	)
	if err != nil {
		return nil, err
	}
	return &city, nil
}

func (r *citiesRepasitory) Get(id string) (*pb.City, error) {
	city := pb.City{}
	var updated_at sql.NullTime
	query := `SELECT id,name,created_at,updated_at FROM city where id=$1 and deleted_at is null`
	err := r.db.QueryRow(query, id).Scan(
		&city.Id,
		&city.Name,
		&city.Created_at,
		&updated_at,
	)
	if err != nil {
		return nil, err
	}
	if updated_at.Valid {
		city.Updated_at = updated_at.Time.String()
	}
	return &city, nil
}

func (r *citiesRepasitory) GetAll() ([]*pb.City, error) {
	var cities []*pb.City
	query := `SELECT id,name,created_at,updated_at from city where deleted_at is null`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var updated_at sql.NullTime
		city := pb.City{}
		err := rows.Scan(
			&city.Id,
			&city.Name,
			&city.Created_at,
			&updated_at,
		)
		if err != nil {
			return nil, err
		}
		if updated_at.Valid {
			city.Updated_at = updated_at.Time.String()
		}
		cities = append(cities, &city)
	}
	return cities, nil
}

func (r *citiesRepasitory) Delete(id string) (*pb.City, error) {
	city, err := r.Get(id)
	if err != nil {
		return nil, err
	}
	query := `UPDATE city SET deleted_at=$2 where id=$1 and deleted_at is null`
	_, err = r.db.Exec(query, id)
	if err != nil {
		return nil, err
	}
	return city, nil
}

func (r *citiesRepasitory) GetCarByCity(id string) (*pb.Car, error) {
	newCar := pb.Car{}
	var updated_at sql.NullTime
	query := `SELECT id,user_id,category_id,marc_id,model_id,position_id,body_id,date,price,auction,enginee,oil_id,transmission_id,milage,
	color_id,drive_unit_id,outside_id,optic_id,salon_id,media_id,options_id,additionally_id,add_info,region_id,city_id,
	phone,created_at,updated_at from cars JOIN city ON cars.city_id=city.id and cars.deleted_at is null and city.id=$1`
	err := r.db.QueryRow(query, id).Scan(
		&newCar.Id,
		&newCar.User_Id,
		&newCar.Category_Id,
		&newCar.Marc_Id,
		&newCar.Model_Id,
		&newCar.Position_Id,
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
