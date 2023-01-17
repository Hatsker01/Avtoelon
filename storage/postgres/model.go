package postgres

import (
	"database/sql"
	"time"

	pb "github.com/Avtoelon/pkg/structs"
	"github.com/Avtoelon/storage/repo"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type modelRepasitory struct {
	db *sqlx.DB
}

func NewModelRepasitory(db *sqlx.DB) repo.ModelRepoInterface {
	return &modelRepasitory{
		db: db,
	}
}

func (r *modelRepasitory) CreateModel(model *pb.CreateModelReq) (*pb.Model, error) {
	newModel := pb.Model{}
	query := `INSERT INTO models(name,marc_id,created_at) VALUES($1,$2,$3) RETURNING id,name,marc_id,created_at`
	err := r.db.QueryRow(query, model.Name, model.Marc_Id, time.Now().UTC()).Scan(
		&newModel.Id,
		&newModel.Name,
		&newModel.Marc_Id,
		&newModel.Created_at,
	)
	if err != nil {
		return nil, err
	}
	return &newModel, nil
}

func (r *modelRepasitory) UpdateModel(upModel *pb.UpdateModel) (*pb.Model, error) {
	model := pb.Model{}
	query := `UPDATE models SET name=$2,marc_id=$3,updated_at=$4 where id=$1 and deleted_at is null returning id,name,marc_id,created_at,updated_at`
	err := r.db.QueryRow(query, upModel.Id, upModel.Name, upModel.Marc_Id, time.Now().UTC()).Scan(
		&model.Id,
		&model.Name,
		&model.Marc_Id,
		&model.Created_at,
		&model.Updated_at,
	)
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func (r *modelRepasitory) GetModel(id string) (*pb.Model, error) {
	model := pb.Model{}
	var updated_at sql.NullTime
	query := `SELECT id,name,marc_id,created_at,updated_at from models where deleted_at is null and id=$1`
	err := r.db.QueryRow(query, id).Scan(
		&model.Id,
		&model.Name,
		&model.Marc_Id,
		&model.Created_at,
		&updated_at,
	)
	if err != nil {
		return nil, err
	}
	if updated_at.Valid {
		model.Updated_at = updated_at.Time.String()
	}
	return &model, nil
}

func (r *modelRepasitory) GetAllModels() ([]*pb.Model, error) {
	var models []*pb.Model
	query := `SELECT id,name,marc_id,created_at,updated_at where deleted_at is null`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var updated_at sql.NullTime
		model := pb.Model{}
		err := rows.Scan(
			&model.Id,
			&model.Name,
			&model.Marc_Id,
			&model.Created_at,
			&updated_at,
		)
		if err != nil {
			return nil, err
		}
		if updated_at.Valid {
			model.Updated_at = updated_at.Time.String()
		}
		models = append(models, &model)
	}
	return models, err
}

func (r *modelRepasitory) DeleteModel(id string) (*pb.Model, error) {
	model, err := r.GetModel(id)
	if err != nil {
		return nil, err
	}
	query := `UPDATE models SET deleted_at=$2 where id=$1 and deleted_at is null`
	_, err = r.db.Exec(query, id)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (r *modelRepasitory) GetCarByModel(id string) (*pb.Car, error) {
	newCar := pb.Car{}
	var updated_at sql.NullTime
	query := `SELECT id,user_id,category_id,marc_id,model_id,position_id,body_id,date,price,auction,enginee,oil_id,transmission_id,milage,
	color_id,drive_unit_id,outside_id,optic_id,salon_id,media_id,options_id,additionally_id,add_info,region_id,city_id,
	phone,created_at,updated_at from cars JOIN models ON cars.model_id=models.id and cars.deleted_at is null and models.id=$1`
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
