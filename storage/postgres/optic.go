package postgres

import (
	"database/sql"
	"time"

	pb "github.com/Avtoelon/pkg/structs"
	"github.com/Avtoelon/storage/repo"
	"github.com/jmoiron/sqlx"
)

type opticRepasitory struct {
	db *sqlx.DB
}

func NewOpticRepasitory(db *sqlx.DB) repo.OpticRepoInterface {
	return &opticRepasitory{
		db: db,
	}
}

func (r *opticRepasitory) Create(optic *pb.CreateOptic) (*pb.Optic, error) {
	query := `INSERT INTO optic(name,created_at) VALUES ($1,$2) RETURNING id,name,created_at`
	newOptic := pb.Optic{}
	err := r.db.QueryRow(query, optic.Name, time.Now().UTC()).Scan(
		&newOptic.Id,
		&newOptic.Name,
		&newOptic.Created_at,
	)
	if err != nil {
		return nil, err
	}
	return &newOptic, nil
}

func (r *opticRepasitory) Update(upOptic *pb.UpdateOpticReq) (*pb.Optic, error) {
	query := `UPDATE optic SET name=$2,updated_at=$3 where id=$1 and deleted_at is null RETURNING id,name,created_at,updated_at`
	optic := pb.Optic{}
	err := r.db.QueryRow(query, upOptic.Id, upOptic.Name, time.Now().UTC()).Scan(
		&optic.Id,
		&optic.Name,
		&optic.Created_at,
		&optic.Updated_at,
	)
	if err != nil {
		return nil, err
	}
	return &optic, nil
}

func (r *opticRepasitory) Get(id string) (*pb.Optic, error) {
	optic := pb.Optic{}
	var updated_at sql.NullTime
	query := `SELECT id,name,created_at,updated_at from optic where id=$1 and deleted_at is null`
	err := r.db.QueryRow(query, id).Scan(
		&optic.Id,
		&optic.Name,
		&optic.Created_at,
		&updated_at,
	)
	if err != nil {
		return nil, err
	}
	if updated_at.Valid {
		optic.Updated_at = updated_at.Time.String()
	}
	return &optic, nil
}

func (r *opticRepasitory) GetAll() ([]*pb.Optic, error) {
	var optics []*pb.Optic
	query := `SELECT id,name,created_at,updated_at from optic where deleted_at is null`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		optic := pb.Optic{}
		var updated_at sql.NullTime
		err := rows.Scan(
			&optic.Id,
			&optic.Name,
			&optic.Created_at,
			&updated_at,
		)
		if err != nil {
			return nil, err
		}
		if updated_at.Valid {
			optic.Updated_at = updated_at.Time.String()
		}
		optics = append(optics, &optic)
	}
	return optics, nil
}

func (r *opticRepasitory) Delete(id string) (*pb.Optic, error) {
	optic, err := r.Get(id)
	if err != nil {
		return nil, err
	}

	query := `UPDATE optic SET deleted_at = $2 where id=$1 and deleted_at is null`
	_, err = r.db.Exec(query, id, time.Now().UTC())
	if err != nil {
		return nil, err
	}
	return optic, nil

}
