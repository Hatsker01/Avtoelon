package postgres

import (
	"database/sql"
	"time"

	"github.com/Avtoelon/storage/repo"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"

	pb "github.com/Avtoelon/pkg/structs"
)

type bodyRepasitory struct {
	db *sqlx.DB
}

func NewBodyRepo(db *sqlx.DB) repo.BodyRepoInterface {
	return &bodyRepasitory{
		db: db,
	}
}

func (r *bodyRepasitory) CreateBody(body *pb.CreateBody) (*pb.Body, error) {
	newBody := pb.Body{}
	query := `INSERT INTO body (name,created_at) VALUES($1,$2) RETURNING id,name,created_at`
	err := r.db.QueryRow(query, body.Name, time.Now().UTC()).Scan(
		&newBody.Id,
		&newBody.Name,
		&newBody.Created_at,
	)
	if err != nil {
		return nil, err
	}
	return &newBody, nil
}

func (r *bodyRepasitory) UpdateBody(upBody *pb.UpdateBody) (*pb.Body, error) {
	body := pb.Body{}
	query := `UPDATE body SET name=$2,updated_at=$3 where deleted_at is null and id=$1 RETURNING id,name,created_at,updated_at`
	err := r.db.QueryRow(query, upBody.Id, upBody.Name, time.Now().UTC()).Scan(
		&body.Id,
		&body.Name,
		&body.Created_at,
		&body.Updated_at,
	)
	if err != nil {
		return nil, err
	}
	return &body, nil
}

func (r *bodyRepasitory) GetBody(id string) (*pb.Body, error) {
	body := pb.Body{}
	var updated_at sql.NullTime
	query := `SELECT id,name,created_at,updated_at from body where deleted_at is null and id=$1`
	err := r.db.QueryRow(query, id).Scan(
		&body.Id,
		&body.Name,
		&body.Created_at,
		&updated_at,
	)
	if err != nil {
		return nil, err
	}
	if updated_at.Valid {
		body.Updated_at = updated_at.Time.String()
	}
	return &body, nil
}

func (r *bodyRepasitory) GetAllBody() ([]*pb.Body, error) {
	var bodies []*pb.Body
	query := `SELECT id,name,created_at,updated_at from body where deleted_at is null`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var updated_at sql.NullTime
		body := pb.Body{}
		err := rows.Scan(
			&body.Id,
			&body.Name,
			&body.Created_at,
			&updated_at,
		)
		if err != nil {
			return nil, err
		}
		if updated_at.Valid {
			body.Updated_at = updated_at.Time.String()
		}
		bodies = append(bodies, &body)
	}
	return bodies, nil
}

func (r *bodyRepasitory) DeleteBody(id string) (*pb.Body, error) {
	body, err := r.GetBody(id)
	if err != nil {
		return nil, err
	}
	query := `UPDATE body SET deleted_at=$2 where deleted_at is null and id=$1`
	_, err = r.db.Exec(query, id)
	if err != nil {
		return nil, err
	}
	return body, err
}

func (r *bodyRepasitory) GetCarByBody(id string) (*pb.Car, error) {
	newCar := pb.Car{}
	var updated_at sql.NullTime
	query := `SELECT id,user_id,category_id,marc_id,model_id,position_id,body_id,date,price,auction,enginee,oil_id,transmission_id,milage,
	color_id,drive_unit_id,outside_id,optic_id,salon_id,media_id,options_id,additionally_id,add_info,region_id,city_id,
	phone,created_at,updated_at from cars JOIN body ON cars.body_id=body.id and cars.deleted_at is null and body.id=$1`
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
