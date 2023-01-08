package postgres

import (
	"database/sql"
	"time"

	"github.com/Avtoelon/storage/repo"
	"github.com/jmoiron/sqlx"

	pb "github.com/Avtoelon/pkg/structs"
)

type oilRepasitory struct {
	db *sqlx.DB
}

func NewOilRepasitory(db *sqlx.DB) repo.OilRepoInterface {
	return &oilRepasitory{
		db: db,
	}
}

func (r *oilRepasitory) Create(oil *pb.CreateOil)(*pb.Oil,error){
	newOil:=pb.Oil{}
	query:=`INSERT INTO oil(name,created_at) VALUES ($1,$2) RETURNING id,name,created_at`
	err:=r.db.QueryRow(query,oil.Name,time.Now().UTC()).Scan(
		&newOil.Id,
		&newOil.Name,
		&newOil.Created_at,
	)
	if err!=nil{
		return nil,err
	}
	return &newOil,nil
}

func (r *oilRepasitory) Update(upOil *pb.UpdateOil)(*pb.Oil,error){
	oil:=pb.Oil{}
	query:=`UPDATE oil SET name=$2,updated_at=$3 where deleted at is null and id=$1 RETURNING id,name,created_at,updated_at`
	err:=r.db.QueryRow(query,upOil.Id,upOil.Name,time.Now().UTC()).Scan(
		&oil.Id,
		&oil.Name,
		&oil.Created_at,
		&oil.Updated_at,
	)
	if err!=nil{
		return nil,err
	}
	return &oil,nil
}

func (r *oilRepasitory) Get(id string)(*pb.Oil,error){
	oil:=pb.Oil{}
	var updated_at sql.NullTime
	query:=`SELECT id,name,created_at,updated_at from oil where id=$1 and deleted_at is null`
	err:=r.db.QueryRow(query,id).Scan(
		&oil.Id,
		&oil.Name,
		&oil.Created_at,
		&updated_at,
	)
	if err!=nil{
		return nil,err
	}
	if updated_at.Valid{
		oil.Updated_at=updated_at.Time.String()
	}
	return &oil,nil
}


func (r *oilRepasitory) GetAll()([]*pb.Oil,error){
	var oils []*pb.Oil
	query:=`SELECT id,name,created_at,updated_at from oil where deleted_at is null`
	rows,err:=r.db.Query(query)
	if err!=nil{
		return nil,err
	}
	for rows.Next(){
		var updated_at sql.NullTime
		oil:=pb.Oil{}
		err:=rows.Scan(
			&oil.Id,
			&oil.Name,
			&oil.Created_at,
			&updated_at,
		)
		if err!=nil{
			return nil,err
		}
		if updated_at.Valid{
			oil.Updated_at=updated_at.Time.String()
		}
		oils = append(oils, &oil)
	}
	return oils,nil
}

func (r *oilRepasitory) Delete(id string)(*pb.Oil,error){
	oil,err:=r.Get(id)
	if err!=nil{
		return nil,err
	}
	query:=`UPDATE oil SET deleted_at=$2 where id=$1 and deleted_at is null`
	_,err=r.db.Exec(query,id,time.Now().UTC())
	if err!=nil{
		return nil,err
	}
	return oil,nil
}
