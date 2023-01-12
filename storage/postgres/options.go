package postgres

import (
	"database/sql"
	"time"

	"github.com/Avtoelon/storage/repo"
	"github.com/jmoiron/sqlx"

	pb "github.com/Avtoelon/pkg/structs"
)

type optionsRepasitory struct {
	db *sqlx.DB
}

func NewOptionsRepo(db *sqlx.DB) repo.OptionsRepoInterface {
	return &optionsRepasitory{
		db: db,
	}
}

func (r *optionsRepasitory) Create(option *pb.CreateOption) (*pb.Option, error) {
	newOption := pb.Option{}
	query := `INSERT INTO options(name,created_at) VALUES ($1,$2) RETURNING id,name,created_at`
	err := r.db.QueryRow(query, option.Name, time.Now().UTC()).Scan(
		&newOption.Id,
		&newOption.Name,
		&newOption.Created_at,
	)
	if err != nil {
		return nil, err
	}
	return &newOption, nil
}

func (r *optionsRepasitory) Update(upOption *pb.UpdateOptionReq) (*pb.Option, error) {
	option := pb.Option{}
	query := `UPDATE option SET name=$2,updated_at=$3 where id=$1 and deleted_at is null RETURNING id,name,created_at,updated_at`
	err := r.db.QueryRow(query, upOption.Id, upOption.Name, time.Now().UTC()).Scan(
		&option.Id,
		&option.Name,
		&option.Created_at,
		&option.Updated_at,
	)
	if err != nil {
		return nil, err
	}
	return &option, nil
}

func (r *optionsRepasitory) Get(id string) (*pb.Option, error) {
	option := pb.Option{}
	var updated_at sql.NullTime
	query := `SELECT id,name,created_at,updated_at from options where id=$1 and deleted_at is null`
	err := r.db.QueryRow(query, id).Scan(
		&option.Id,
		&option.Name,
		&option.Created_at,
		&updated_at,
	)
	if err != nil {
		return nil, err
	}
	if updated_at.Valid {
		option.Updated_at = updated_at.Time.String()
	}
	return &option, nil
}

func (r *optionsRepasitory) GetAll() ([]*pb.Option, error) {
	var options []*pb.Option
	query := `SELECT id,name,created_at,updated_at from options where deleted_at is null`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		option := pb.Option{}
		var updated_at sql.NullTime
		err := rows.Scan(
			&option.Id,
			&option.Name,
			&option.Created_at,
			&updated_at,
		)
		if err != nil {
			return nil, err
		}
		if updated_at.Valid {
			option.Updated_at = updated_at.Time.String()
		}
		options = append(options, &option)
	}
	return options, nil
}

func (r *optionsRepasitory) Delete(id string) (*pb.Option, error) {
	option, err := r.Get(id)
	if err != nil {
		return nil, err
	}
	query := `UPDATE options SET deleted_at = $2 where deleted_at is null and id=$1`
	_, err = r.db.Exec(query, id, time.Now().UTC())
	if err != nil {
		return nil, err
	}
	return option, nil
}
