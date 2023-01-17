package postgres

import (
	"database/sql"
	"time"

	"github.com/Avtoelon/storage/repo"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"

	pb "github.com/Avtoelon/pkg/structs"
)

type marcsRepasitory struct {
	db *sqlx.DB
}

func NewMarcsRepo(db *sqlx.DB) repo.MarcsRepoInterface {
	return &marcsRepasitory{
		db: db,
	}
}
func (r *marcsRepasitory) Create(marc *pb.CreateMarc) (*pb.Marc, error) {
	newMarc := pb.Marc{}
	query := `INSERT INTO marc(name,created_at) VALUES ($1,$2) RETURNING id,name,created_at`
	err := r.db.QueryRow(query, marc.Name, time.Now().UTC()).Scan(
		&newMarc.Id,
		&newMarc.Name,
		&newMarc.Created_at,
	)
	if err != nil {
		return nil, err
	}
	return &newMarc, nil
}

func (r *marcsRepasitory) Update(upMarc *pb.UpdateMarcReq) (*pb.Marc, error) {
	marc := pb.Marc{}
	query := `UPDATE marc SET name=$2,updated_at=$3 where deleted at is null and id=$1 RETURNING id,name,created_at,updated_at`
	err := r.db.QueryRow(query, upMarc.Id, &upMarc.Name, time.Now().UTC()).Scan(
		&marc.Id,
		&marc.Name,
		&marc.Created_at,
		&marc.Updated_at,
	)
	if err != nil {
		return nil, err
	}
	return &marc, nil
}

func (r *marcsRepasitory) Get(id string) (*pb.Marc, error) {
	marc := pb.Marc{}

	var updated_at sql.NullTime
	query := `SELECT id,name,created_at,updated_at from marc where id=$1 and deleted_at is null`
	err := r.db.QueryRow(query, id).Scan(
		&marc.Id,
		&marc.Name,
		&marc.Created_at,
		&updated_at,
	)

	if err != nil {
		return nil, err
	}
	if updated_at.Valid {
		marc.Updated_at = updated_at.Time.String()
	}
	return &marc, nil
}

func (r *marcsRepasitory) GetAll() ([]*pb.Marc, error) {
	var marcs []*pb.Marc
	query := `SELECT id,name,created_at,updated_at from marc where deleted_at is null`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var updated_at sql.NullTime
		marc := pb.Marc{}
		err := rows.Scan(
			&marc.Id,
			&marc.Name,
			&marc.Created_at,
			&updated_at,
		)
		if err != nil {
			return nil, err
		}
		marcs = append(marcs, &marc)
	}
	return marcs, nil
}
func (r *marcsRepasitory) Delete(id string) (*pb.Marc, error) {
	marc, err := r.Get(id)
	if err != nil {
		return nil, err
	}
	query := `UPDATE marc SET deleted_at=$2 where id=$1 and deleted_at is null`
	_, err = r.db.Exec(query, id, time.Now().UTC())
	if err != nil {
		return nil, err
	}
	return marc, nil
}

func (r *marcsRepasitory) GetMarcModels(id string) ([]*pb.GetMarcModels, error) {
	var models []*pb.GetMarcModels
	query := `SELECT models.id,a.name,models.name from marc a join models on models.marc_id=a.id where a.deleted_at is null and models.deleted_at is null and a.id=$1`
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		model := pb.GetMarcModels{}
		err := rows.Scan(
			&model.Id,
			&model.Marc_Name,
			&model.Model_Name,
		)
		if err != nil {
			return nil, err
		}
		models = append(models, &model)
	}
	return models, nil
}

func (r *marcsRepasitory) GetCarByMarc(id string) (*pb.Car, error) {
	newCar := pb.Car{}
	var updated_at sql.NullTime
	query := `SELECT id,user_id,category_id,marc_id,model_id,position_id,body_id,date,price,auction,enginee,oil_id,transmission_id,milage,
	color_id,drive_unit_id,outside_id,optic_id,salon_id,media_id,options_id,additionally_id,add_info,region_id,city_id,
	phone,created_at,updated_at from cars JOIN marc ON cars.category_id=marc.id and cars.deleted_at is null and marc.id=$1`
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
