package postgres

import (
	"database/sql"
	"time"

	"github.com/Avtoelon/storage/repo"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"

	pb "github.com/Avtoelon/pkg/structs"
)

type transmissionRepasitory struct {
	db *sqlx.DB
}

func NewTransmissionRepasitory(db *sqlx.DB) repo.TransmissionRepoInterface {
	return &transmissionRepasitory{
		db: db,
	}
}

func (r *transmissionRepasitory) Create(trans *pb.CreateTrans) (*pb.Transmission, error) {
	newT := pb.Transmission{}
	query := `INSERT INTO transmission(name,created_at) VALUES ($1,$2) RETURNING id,name,created_at`
	err := r.db.QueryRow(query, trans.Name, time.Now().UTC()).Scan(
		&newT.Id,
		&newT.Name,
		&newT.Created_at,
	)
	if err != nil {
		return nil, err
	}
	return &newT, nil
}

func (r *transmissionRepasitory) Update(upTrans *pb.UpdateTrans) (*pb.Transmission, error) {
	trans := pb.Transmission{}
	query := `UPDATE transmission SET name=$2,updated_at=$3 where deleted at is null and id=$1 RETURNING id,name,created_at,updated_at`
	err := r.db.QueryRow(query, upTrans.Id, upTrans.Name, time.Now().UTC()).Scan(
		&trans.Id,
		&trans.Name,
		&trans.Created_at,
		&trans.Updated_at,
	)
	if err != nil {
		return nil, err
	}
	return &trans, nil
}

func (r *transmissionRepasitory) Get(id string) (*pb.Transmission, error) {
	trans := pb.Transmission{}
	var updated_at sql.NullTime
	query := `SELECT id,name,created_at,updated_at from oil where id=$1 and deleted_at is null`
	err := r.db.QueryRow(query, id).Scan(
		&trans.Id,
		&trans.Name,
		&trans.Created_at,
		&updated_at,
	)
	if err != nil {
		return nil, err
	}
	if updated_at.Valid {
		trans.Updated_at = updated_at.Time.String()
	}
	return &trans, nil
}

func (r *transmissionRepasitory) GetAll() ([]*pb.Transmission, error) {
	var transs []*pb.Transmission
	query := `SELECT id,name,created_at,updated_at from transmission where deleted_at is null`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var updated_at sql.NullTime
		trans := pb.Transmission{}
		err := rows.Scan(
			&trans.Id,
			&trans.Name,
			&trans.Created_at,
			&updated_at,
		)
		if err != nil {
			return nil, err
		}
		if updated_at.Valid {
			trans.Updated_at = updated_at.Time.String()
		}
		transs = append(transs, &trans)
	}
	return transs, nil
}

func (r *transmissionRepasitory) Delete(id string) (*pb.Transmission, error) {
	trans, err := r.Get(id)
	if err != nil {
		return nil, err
	}
	query := `UPDATE transmission SET deleted_at=$2 where id=$1 and deleted_at is null`
	_, err = r.db.Exec(query, id, time.Now().UTC())
	if err != nil {
		return nil, err
	}
	return trans, nil
}

func (r *transmissionRepasitory) GetCarByTrans(id string) (*pb.Car, error) {
	newCar := pb.Car{}
	var updated_at sql.NullTime
	query := `SELECT id,user_id,category_id,marc_id,model_id,position_id,body_id,date,price,auction,enginee,oil_id,transmission_id,milage,
	color_id,drive_unit_id,outside_id,optic_id,salon_id,media_id,options_id,additionally_id,add_info,region_id,city_id,
	phone,created_at,updated_at from cars JOIN transmission ON cars.transmission_id=transmission.id and cars.deleted_at is null and transmission.id=$1`
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
