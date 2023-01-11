package postgres

import (
	"database/sql"
	"time"

	"github.com/Avtoelon/storage/repo"
	"github.com/jmoiron/sqlx"

	pb "github.com/Avtoelon/pkg/structs"
)

type driveunitRepasitory struct {
	db *sqlx.DB
}

func NewDriveUnitRepasitory(db *sqlx.DB) repo.DriveUnitRepoInterface {
	return &driveunitRepasitory{
		db: db,
	}
}

func (r *driveunitRepasitory) Create(dr *pb.DriveUnitCreateReq) (*pb.Drive_Unit, error) {
	drive := pb.Drive_Unit{}
	query := `INSERT INTO drive_unit(name,created_at) VALUES($1,$2) RETURNING id,name,created_at`
	err := r.db.QueryRow(query, dr.Name, time.Now().UTC()).Scan(
		&drive.Id,
		&drive.Name,
		&drive.Created_at,
	)
	if err != nil {
		return nil, err
	}
	return &drive, nil
}

func (r *driveunitRepasitory) Update(dr *pb.UpdateDriveUnit) (*pb.Drive_Unit, error) {
	drive := pb.Drive_Unit{}
	query := `UPDATE drive_unit SET name=$2,updated_at=$3 where id=$1 and deleted_at is null RETURNING id,name,created_at,updated_at`
	err := r.db.QueryRow(query, dr.Id, &dr.Name, time.Now().UTC()).Scan(
		&drive.Id,
		&drive.Name,
		&drive.Created_at,
		&drive.Updated_at,
	)
	if err != nil {
		return nil, err
	}
	return &drive, nil
}

func (r *driveunitRepasitory) Get(id string) (*pb.Drive_Unit, error) {
	var updated_at sql.NullTime
	drive := pb.Drive_Unit{}
	query := `SELECT id,name,created_at,updated_at from drive_unit where id=$1 and deleted_at is null`
	err := r.db.QueryRow(query, id).Scan(
		&drive.Id,
		&drive.Name,
		&drive.Created_at,
		&updated_at,
	)
	if err != nil {
		return nil, err
	}
	if updated_at.Valid {
		drive.Updated_at = updated_at.Time.String()
	}
	return &drive, nil
}

func (r *driveunitRepasitory) GetAll() ([]*pb.Drive_Unit, error) {
	var drive_units []*pb.Drive_Unit
	query := `SELECT id,name,created_at,updated_at from drive_unit where deleted_at is null`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		drive_unit := pb.Drive_Unit{}
		var updated_at sql.NullTime
		err := rows.Scan(
			&drive_unit.Id,
			&drive_unit.Name,
			&drive_unit.Created_at,
			&updated_at,
		)
		if err != nil {
			return nil, err
		}
		if updated_at.Valid {
			drive_unit.Updated_at = updated_at.Time.String()
		}
		drive_units = append(drive_units, &drive_unit)
	}
	return drive_units, nil
}

func (r *driveunitRepasitory) Delete(id string) (*pb.Drive_Unit, error) {
	drive, err := r.Get(id)
	if err != nil {
		return nil, err
	}
	query := `UPDATE drive_unit deleted_at=$2 where id=$1 and deleted_at is null`
	_, err = r.db.Exec(query, id, time.Now().UTC())
	if err != nil {
		return nil, err
	}
	return drive, nil
}
