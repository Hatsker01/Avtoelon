package postgres

import (
	"database/sql"
	"time"

	"github.com/Avtoelon/storage/repo"
	"github.com/jmoiron/sqlx"

	pb "github.com/Avtoelon/pkg/structs"
)

type colorRepasitory struct {
	db *sqlx.DB
}

func NewColorRepasitory(db *sqlx.DB) repo.ColorRepoInterface {
	return &colorRepasitory{
		db: db,
	}
}

func (r *colorRepasitory) Create(color *pb.CreateColorReq) (*pb.Color, error) {
	query := `INSERT INTO color(name,created_at) VALUES ($1,$2) RETURNING id,name,color`
	newColor := pb.Color{}
	err := r.db.QueryRow(query, color.Name, time.Now().UTC()).Scan(
		&newColor.Id,
		&newColor.Name,
		&newColor.Created_at,
	)
	if err != nil {
		return nil, err
	}
	return &newColor, nil
}

func (r *colorRepasitory) Update(upColor *pb.UpdateColor) (*pb.Color, error) {
	query := `UPDATE color SET name=$2,updated_at=$3 where id=$1 and deleted at is null RETURNING id,name,created_at,updated_at`
	color := pb.Color{}
	err := r.db.QueryRow(query, upColor.Id, upColor.Name, time.Now().UTC()).Scan(
		&color.Id,
		&color.Name,
		&color.Created_at,
		&color.Updated_at,
	)
	if err != nil {
		return nil, err
	}
	return &color, err
}

func (r *colorRepasitory) Get(id string) (*pb.Color, error) {
	var updated_at sql.NullTime
	color := pb.Color{}
	query := `SELECT id,name,created_at,updated_at from color where id=$1 and deleted_at is null`
	err := r.db.QueryRow(query, id).Scan(
		&color.Id,
		&color.Name,
		&color.Created_at,
		&updated_at,
	)
	if err != nil {
		return nil, err
	}
	if updated_at.Valid {
		color.Updated_at = updated_at.Time.String()
	}

	return &color, nil
}

func (r *colorRepasitory) GetAll() ([]*pb.Color, error) {
	var colors []*pb.Color
	query := `SELECT id,name,created_at,updated_at from color where deleted_at is null`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		color := pb.Color{}
		var updated_at sql.NullTime
		err := rows.Scan(
			&color.Id,
			&color.Name,
			&color.Created_at,
			&updated_at,
		)
		if err != nil {
			return nil, err
		}
		if updated_at.Valid {
			color.Updated_at = updated_at.Time.String()
		}
		colors = append(colors, &color)
	}
	return colors, nil
}

func (r *colorRepasitory) Delete(id string) (*pb.Color, error) {
	color, err := r.Get(id)
	if err != nil {
		return nil, err
	}
	query := `UPDATE color deleted_at=$2 where id=$1 and deleted_at is null`
	_, err = r.db.Exec(query, id, time.Now().UTC())
	if err != nil {
		return nil, err
	}
	return color, err
}
