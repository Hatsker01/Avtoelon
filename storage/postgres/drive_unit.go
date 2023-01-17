package postgres

import (
	"database/sql"
	"time"

	"github.com/Avtoelon/storage/repo"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"

	pb "github.com/Avtoelon/pkg/structs"
)

type driveunitRepasitory struct {
	db *sqlx.DB
}

func NewDriveUnitRepasitory(db *sqlx.DB) repo.DriveUnitRepoInterface {
	return &driveunitRepasitory{
		db: db,
	}
}

func (r *driveunitRepasitory) Create(dr *pb.DriveUnitCreateReq) (*pb.Drive_Unit, error) {
	drive := pb.Drive_Unit{}
	query := `INSERT INTO drive_unit(name,created_at) VALUES($1,$2) RETURNING id,name,created_at`
	err := r.db.QueryRow(query, dr.Name, time.Now().UTC()).Scan(
		&drive.Id,
		&drive.Name,
		&drive.Created_at,
	)
	if err != nil {
		return nil, err
	}
	return &drive, nil
}

func (r *driveunitRepasitory) Update(dr *pb.UpdateDriveUnit) (*pb.Drive_Unit, error) {
	drive := pb.Drive_Unit{}
	query := `UPDATE drive_unit SET name=$2,updated_at=$3 where id=$1 and deleted_at is null RETURNING id,name,created_at,updated_at`
	err := r.db.QueryRow(query, dr.Id, &dr.Name, time.Now().UTC()).Scan(
		&drive.Id,
		&drive.Name,
		&drive.Created_at,
		&drive.Updated_at,
	)
	if err != nil {
		return nil, err
	}
	return &drive, nil
}

func (r *driveunitRepasitory) Get(id string) (*pb.Drive_Unit, error) {
	var updated_at sql.NullTime
	drive := pb.Drive_Unit{}
	query := `SELECT id,name,created_at,updated_at from drive_unit where id=$1 and deleted_at is null`
	err := r.db.QueryRow(query, id).Scan(
		&drive.Id,
		&drive.Name,
		&drive.Created_at,
		&updated_at,
	)
	if err != nil {
		return nil, err
	}
	if updated_at.Valid {
		drive.Updated_at = updated_at.Time.String()
	}
	return &drive, nil
}

func (r *driveunitRepasitory) GetAll() ([]*pb.Drive_Unit, error) {
	var drive_units []*pb.Drive_Unit
	query := `SELECT id,name,created_at,updated_at from drive_unit where deleted_at is null`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		drive_unit := pb.Drive_Unit{}
		var updated_at sql.NullTime
		err := rows.Scan(
			&drive_unit.Id,
			&drive_unit.Name,
			&drive_unit.Created_at,
			&updated_at,
		)
		if err != nil {
			return nil, err
		}
		if updated_at.Valid {
			drive_unit.Updated_at = updated_at.Time.String()
		}
		drive_units = append(drive_units, &drive_unit)
	}
	return drive_units, nil
}

func (r *driveunitRepasitory) Delete(id string) (*pb.Drive_Unit, error) {
	drive, err := r.Get(id)
	if err != nil {
		return nil, err
	}
	query := `UPDATE drive_unit deleted_at=$2 where id=$1 and deleted_at is null`
	_, err = r.db.Exec(query, id, time.Now().UTC())
	if err != nil {
		return nil, err
	}
	return drive, nil
}

func (r *driveunitRepasitory) GetCarByDriveUnit(id string) (*pb.Car, error) {
	newCar := pb.Car{}
	var updated_at sql.NullTime
	query := `SELECT id,user_id,category_id,marc_id,model_id,position_id,body_id,date,price,auction,enginee,oil_id,transmission_id,milage,
	color_id,drive_unit_id,outside_id,optic_id,salon_id,media_id,options_id,additionally_id,add_info,region_id,city_id,
	phone,created_at,updated_at from cars JOIN drive_unit ON cars.drive_unit_id=drive_unit.id and cars.deleted_at is null and drive_unit.id=$1`
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
