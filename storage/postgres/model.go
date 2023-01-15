package postgres

import (
	"database/sql"
	"time"

	pb "github.com/Avtoelon/pkg/structs"
	"github.com/Avtoelon/storage/repo"
	"github.com/jmoiron/sqlx"
)

type modelRepasitory struct {
	db *sqlx.DB
}

func NewModelRepasitory(db *sqlx.DB) repo.ModelRepoInterface {
	return &modelRepasitory{
		db: db,
	}
}

func (r *modelRepasitory) CreateModel(model *pb.CreateModelReq)(*pb.Model,error){
	newModel := pb.Model{}
	query:=`INSERT INTO models(name,marc_id,created_at) VALUES($1,$2,$3) RETURNING id,name,marc_id,created_at`
	err:=r.db.QueryRow(query,model.Name,model.Marc_Id,time.Now().UTC()).Scan(
		&newModel.Id,
		&newModel.Name,
		&newModel.Marc_Id,
		&newModel.Created_at,
	)
	if err!=nil{
		return nil,err
	}
	return &newModel,nil
}

func (r *modelRepasitory) UpdateModel(upModel *pb.UpdateModel)(*pb.Model,error){
	model:=pb.Model{}
	query:=`UPDATE models SET name=$2,marc_id=$3,updated_at=$4 where id=$1 and deleted_at is null returning id,name,marc_id,created_at,updated_at`
	err:=r.db.QueryRow(query,upModel.Id,upModel.Name,upModel.Marc_Id,time.Now().UTC()).Scan(
		&model.Id,
		&model.Name,
		&model.Marc_Id,
		&model.Created_at,
		&model.Updated_at,
	)
	if err!=nil{
		return nil,err
	}
	return &model,nil
}

func (r *modelRepasitory) GetModel(id string)(*pb.Model,error){
	model:=pb.Model{}
	var updated_at sql.NullTime 
	query:=`SELECT id,name,marc_id,created_at,updated_at from models where deleted_at is null and id=$1`
	err:=r.db.QueryRow(query,id).Scan(
		&model.Id,
		&model.Name,
		&model.Marc_Id,
		&model.Created_at,
		&updated_at,
	)
	if err!=nil{
		return nil,err
	}
	if updated_at.Valid{
		model.Updated_at=updated_at.Time.String()
	}
	return &model,nil
}

func (r *modelRepasitory) GetAllModels()([]*pb.Model,error){
	var models []*pb.Model
	query:=`SELECT id,name,marc_id,created_at,updated_at where deleted_at is null`
	rows,err:=r.db.Query(query)
	if err!=nil{
		return nil,err
	}
	for rows.Next(){
		var updated_at sql.NullTime
		model:=pb.Model{}
		err:=rows.Scan(
			&model.Id,
			&model.Name,
			&model.Marc_Id,
			&model.Created_at,
			&updated_at,
		)
		if err!=nil{
			return nil,err
		}
		if updated_at.Valid{
			model.Updated_at=updated_at.Time.String()
		}
		models = append(models, &model)
	}
	return models,err
}

func (r *modelRepasitory) DeleteModel(id string)(*pb.Model,error){
	model,err:=r.GetModel(id)
	if err!=nil{
		return nil,err
	}
	query:=`UPDATE models SET deleted_at=$2 where id=$1 and deleted_at is null`
	_,err=r.db.Exec(query,id)
	if err!=nil{
		return nil,err
	}
	return model,nil
}
