package postgres

import (
	"database/sql"
	"fmt"
	"time"

	pb "github.com/Avtoelon/pkg/structs"
	"github.com/Avtoelon/storage/repo"
	"github.com/jmoiron/sqlx"
)

type outsideRepasitory struct {
	db *sqlx.DB
}

func NewOutsideRepasitory(db *sqlx.DB) repo.OutsideRepoInterface {
	return &outsideRepasitory{
		db: db,
	}
}

func (r *outsideRepasitory) CreateOutside(outside *pb.CreateOutside) (*pb.Outside, error) {
	query := `INSERT INTO outside(name,created_at) VALUES ($1,$2) RETURNING id,name,created_at`
	newOutside := pb.Outside{}
	err := r.db.QueryRow(query, outside.Name, time.Now().UTC()).Scan(
		&newOutside.Id,
		&newOutside.Name,
		&newOutside.Created_at,
	)
	if err != nil {
		return nil, err
	}
	return &newOutside, nil
}

func (r *outsideRepasitory) UpdateOutside(upOut *pb.Outside) (*pb.Outside, error) {
	outside := pb.Outside{}
	var updated_at sql.NullTime
	query := `UPDATE outside SET name=$2,updated_at=$3 where deleted at is null and id=$1 RETURNING id,name,created_at,updated_at`
	err := r.db.QueryRow(query, upOut.Id, upOut.Name, time.Now().UTC()).Scan(
		&outside.Id,
		&outside.Name,
		&outside.Created_at,
		&updated_at,
	)
	if err != nil {
		return nil, err
	}
	if updated_at.Valid {
		outside.Updated_at = updated_at.Time.String()
	}
	return &outside, nil
}

func (r *outsideRepasitory) GetOutside(id string) (*pb.Outside, error) {
	outside := pb.Outside{}
	var updated_at sql.NullTime
	query := `SELECT id,name,created_at,updated_at from outside where id=$1 and deleted_at is null`
	err := r.db.QueryRow(query, id).Scan(&outside.Id, &outside.Name, &outside.Created_at, &updated_at)
	if err != nil {
		return nil, err
	}
	if updated_at.Valid {
		outside.Updated_at = updated_at.Time.String()
	}
	fmt.Println(outside)
	return &outside, nil
}

func (r *outsideRepasitory) GetAllOutside() ([]*pb.Outside, error) {
	var outsides []*pb.Outside
	query := `SELECT id,name,created_at,updated_at where deleted_at is null`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		outside := pb.Outside{}
		var updated_at sql.NullTime
		err := rows.Scan(
			&outside.Id,
			&outside.Name,
			&outside.Created_at,
			&updated_at,
		)
		if err != nil {
			return nil, err
		}
		if updated_at.Valid {
			outside.Updated_at = updated_at.Time.String()
		}
		outsides = append(outsides, &outside)
	}
	return outsides, nil
}

func (r *outsideRepasitory) DeletedOutside(id string) (*pb.Outside, error) {
	outside, err := r.GetOutside(id)
	if err != nil {
		return nil, err
	}
	query := `UPDATE outside SET deleted_at = $2 where deleted_at is null and id=$1`
	_, err = r.db.Exec(query, id, time.Now().UTC())
	if err != nil {
		return nil, err
	}
	return outside, nil
}
