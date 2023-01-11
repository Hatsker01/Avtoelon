package postgres

import (
	"database/sql"
	"time"

	"github.com/Avtoelon/storage/repo"
	"github.com/jmoiron/sqlx"

	pb "github.com/Avtoelon/pkg/structs"
)

type salonRepasitory struct {
	db *sqlx.DB
}

func NewSalonsRepo(db *sqlx.DB) repo.SalonRepoInterface {
	return &salonRepasitory{
		db: db,
	}
}

func (r *salonRepasitory) Create(salon *pb.CreateSalon) (*pb.Salon, error) {
	newSalon := pb.Salon{}
	query := `INSERT INTO salon(name,created_at) VALUES ($1,$2) RETURNING id,name,created_at`
	err := r.db.QueryRow(query, salon.Name, time.Now().UTC()).Scan(
		&newSalon.Id,
		&newSalon.Name,
		&newSalon.Created_at,
	)
	if err != nil {
		return nil, err
	}
	return &newSalon, nil
}

func (r *salonRepasitory) Update(upSalon *pb.UpdateSalonReq) (*pb.Salon, error) {
	salon := pb.Salon{}
	query := `UPDATE salon SET name=$2,updated_at=$3 where deleted_at is null and id=$1 RETURNING id,name,created_at,updated_at`
	err := r.db.QueryRow(query, upSalon.Id, upSalon.Name, time.Now().UTC()).Scan(
		&salon.Id,
		&salon.Name,
		&salon.Created_at,
		&salon.Updated_at,
	)
	if err != nil {
		return nil, err
	}
	return &salon, nil
}

func (r *salonRepasitory) Get(id string) (*pb.Salon, error) {
	salon := pb.Salon{}
	var updated_at sql.NullTime
	query := `SELECT id,name,created_at,updated_at from salon where deleted_at is null and id=$1`
	err := r.db.QueryRow(query, id).Scan(
		&salon.Id,
		&salon.Name,
		&salon.Created_at,
		&updated_at,
	)
	if err != nil {
		return nil, err
	}
	if updated_at.Valid {
		salon.Updated_at = updated_at.Time.String()
	}
	return &salon, nil
}

func (r *salonRepasitory) GetAll() ([]*pb.Salon, error) {
	var salons []*pb.Salon
	query := `SELECT id,name,created_at,updated_at from salon where deleted_at is null`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var updated_at sql.NullTime
		salon := pb.Salon{}
		err := rows.Scan(
			&salon.Id,
			&salon.Name,
			&salon.Created_at,
			&updated_at,
		)
		if err != nil {
			return nil, err
		}
		if updated_at.Valid {
			salon.Updated_at = updated_at.Time.String()
		}
		salons = append(salons, &salon)
	}
	return salons, nil
}

func (r *salonRepasitory) Delete(id string) (*pb.Salon, error) {
	salon, err := r.Get(id)
	if err != nil {
		return nil, err
	}
	query := `UPDATE salon SET deleted_at = $2 where id=$1 and deleted_at is null`
	_, err = r.db.Exec(query, id)
	if err != nil {
		return nil, err
	}
	return salon, nil
}
