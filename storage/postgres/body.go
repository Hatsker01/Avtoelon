package postgres

import (
	"database/sql"
	"time"

	"github.com/Avtoelon/storage/repo"
	"github.com/jmoiron/sqlx"

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
	query := `SELECT id,name,created_at,updated_at where deleted_at is null`
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
