package postgres

import (
	"database/sql"
	"time"

	"github.com/Avtoelon/storage/repo"
	"github.com/jmoiron/sqlx"

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
