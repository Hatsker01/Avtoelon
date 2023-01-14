package postgres

import (
	"database/sql"
	"time"

	"github.com/Avtoelon/storage/repo"
	"github.com/jmoiron/sqlx"

	pb "github.com/Avtoelon/pkg/structs"
)

type additionalsRepasitory struct {
	db *sqlx.DB
}

func NewAdditionalsRepo(db *sqlx.DB) repo.AdditionalsRepoInterface {
	return &additionalsRepasitory{
		db: db,
	}
}

func (r additionalsRepasitory) Create(additional pb.CreateAdditional) (pb.Additional, error) {
	add := pb.Additional{}
	query := `INSERT INTO additionally(name,created_at) VALUES ($1,$2) RETURNING id,name,created_at`
	err := r.db.QueryRow(query, additional.Name, time.Now().UTC()).Scan(
		&add.Id,
		&add.Name,
		&add.Created_at,
	)
	if err != nil {
		return pb.Additional{}, err
	}
	return add, nil
}

func (r *additionalsRepasitory) Update(upAdd *pb.UpdateAdditionalReq) (*pb.Additional, error) {
	add := pb.Additional{}
	query := `UPDATE additionally SET name=$2,updated_at=$3 where id=$1 and deleted_at is null RETURNING id,name,created_at,updated_at`
	err := r.db.QueryRow(query, upAdd.Id, upAdd.Name, time.Now().UTC()).Scan(
		&add.Id,
		&add.Name,
		&add.Created_at,
		&add.Updated_at,
	)
	if err != nil {
		return nil, err
	}
	return &add, nil
}

func (r *additionalsRepasitory) Get(id string) (*pb.Additional, error) {
	add := pb.Additional{}
	var updated_at sql.NullTime
	query := `SELECT id,name,created_at,updated_at from additionally where id=$1 and deleted_at is null`
	err := r.db.QueryRow(query, id).Scan(
		&add.Id,
		&add.Name,
		&add.Created_at,
		&updated_at,
	)
	if err != nil {
		return nil, err
	}
	if updated_at.Valid {
		add.Updated_at = updated_at.Time.String()
	}
	return &add, nil
}

func (r *additionalsRepasitory) GetAll() ([]*pb.Additional, error) {
	var adds []*pb.Additional
	query := `SELECT id,name,created_at,updated_at from additionally where deleted_at is null`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var updated_at sql.NullTime
		add := pb.Additional{}
		err := rows.Scan(
			&add.Id,
			&add.Name,
			&add.Created_at,
			&updated_at,
		)
		if err != nil {
			return nil, err
		}
		if updated_at.Valid {
			add.Updated_at = updated_at.Time.String()
		}
		adds = append(adds, &add)
	}
	return adds, nil
}

func (r *additionalsRepasitory) Delete(id string) (*pb.Additional, error) {
	add, err := r.Get(id)
	if err != nil {
		return nil, err
	}
	query := `UPDATE additionally SET deleted_at=$2 where deleted_at is null and id=$1`
	_, err = r.db.Exec(query, id)
	if err != nil {
		return nil, err
	}
	return add, nil
}
