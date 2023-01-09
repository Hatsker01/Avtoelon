package postgres

import (
	"database/sql"
	"time"

	"github.com/Avtoelon/storage/repo"
	"github.com/jmoiron/sqlx"

	pb "github.com/Avtoelon/pkg/structs"
)

type transmissionRepasitory struct {
	db *sqlx.DB
}

func NewTransmissionRepasitory(db *sqlx.DB) repo.TransmissionRepoInterface {
	return &transmissionRepasitory{
		db: db,
	}
}

func (r *transmissionRepasitory) Create(trans *pb.CreateTrans) (*pb.Transmission, error) {
	newT := pb.Transmission{}
	query := `INSERT INTO transmission(name,created_at) VALUES ($1,$2) RETURNING id,name,created_at`
	err := r.db.QueryRow(query, trans.Name, time.Now().UTC()).Scan(
		&newT.Id,
		&newT.Name,
		&newT.Created_at,
	)
	if err != nil {
		return nil, err
	}
	return &newT, nil
}

func (r *transmissionRepasitory) Update(upTrans *pb.UpdateTrans) (*pb.Transmission, error) {
	trans := pb.Transmission{}
	query := `UPDATE transmission SET name=$2,updated_at=$3 where deleted at is null and id=$1 RETURNING id,name,created_at,updated_at`
	err := r.db.QueryRow(query, upTrans.Id, upTrans.Name, time.Now().UTC()).Scan(
		&trans.Id,
		&trans.Name,
		&trans.Created_at,
		&trans.Updated_at,
	)
	if err != nil {
		return nil, err
	}
	return &trans, nil
}

func (r *transmissionRepasitory) Get(id string) (*pb.Transmission, error) {
	trans := pb.Transmission{}
	var updated_at sql.NullTime
	query := `SELECT id,name,created_at,updated_at from oil where id=$1 and deleted_at is null`
	err := r.db.QueryRow(query, id).Scan(
		&trans.Id,
		&trans.Name,
		&trans.Created_at,
		&updated_at,
	)
	if err != nil {
		return nil, err
	}
	if updated_at.Valid {
		trans.Updated_at = updated_at.Time.String()
	}
	return &trans, nil
}

func (r *transmissionRepasitory) GetAll() ([]*pb.Transmission, error) {
	var transs []*pb.Transmission
	query := `SELECT id,name,created_at,updated_at from transmission where deleted_at is null`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var updated_at sql.NullTime
		trans := pb.Transmission{}
		err := rows.Scan(
			&trans.Id,
			&trans.Name,
			&trans.Created_at,
			&updated_at,
		)
		if err != nil {
			return nil, err
		}
		if updated_at.Valid {
			trans.Updated_at = updated_at.Time.String()
		}
		transs = append(transs, &trans)
	}
	return transs, nil
}

func (r *transmissionRepasitory) Delete(id string) (*pb.Transmission, error) {
	trans, err := r.Get(id)
	if err != nil {
		return nil, err
	}
	query := `UPDATE transmission SET deleted_at=$2 where id=$1 and deleted_at is null`
	_, err = r.db.Exec(query, id, time.Now().UTC())
	if err != nil {
		return nil, err
	}
	return trans, nil
}
