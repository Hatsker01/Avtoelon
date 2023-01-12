package postgres

import (
	"database/sql"
	"time"

	pb "github.com/Avtoelon/pkg/structs"
	"github.com/Avtoelon/storage/repo"
	"github.com/jmoiron/sqlx"
)

type regionsRepasitory struct {
	db *sqlx.DB
}

func NewRegionsRepo(db *sqlx.DB) repo.RegionsRepoInterface {
	return &regionsRepasitory{
		db: db,
	}
}

func (r *regionsRepasitory) Create(region *pb.CreateRegion) (*pb.Region, error) {
	query := `INSERT INTO region (name,created_at) VALUES ($1,$2) RETURNING id,name,created_at`
	newRegion := pb.Region{}
	err := r.db.QueryRow(query, region.Name, time.Now().UTC()).Scan(
		&newRegion.Id,
		&newRegion.Name,
		&newRegion.Created_at,
	)
	if err != nil {
		return nil, err
	}
	return &newRegion, nil
}

func (r *regionsRepasitory) Update(upRegion *pb.UpdateRegionReq) (*pb.Region, error) {
	query := `UPDATE region SET name=$2,update_at=$3 where id=$1 and deleted_at is null RETURNING id,name,created_at,updated_at`
	region := pb.Region{}
	err := r.db.QueryRow(query, upRegion.Id, upRegion.Name, time.Now().UTC()).Scan(
		&region.Id,
		&region.Name,
		&region.Created_at,
		&region.Updated_at,
	)
	if err != nil {
		return nil, err
	}
	return &region, nil
}

func (r *regionsRepasitory) Get(id string) (*pb.Region, error) {
	region := pb.Region{}
	var updated_at sql.NullTime
	query := `SELECT id,name,created_at,updated_at from region where deleted_at is null and id=$1`
	err := r.db.QueryRow(query, id).Scan(
		&region.Id,
		&region.Name,
		&region.Created_at,
		&updated_at,
	)
	if err != nil {
		return nil, err
	}
	if updated_at.Valid {
		region.Updated_at = updated_at.Time.String()
	}
	return &region, nil
}

func (r *regionsRepasitory) GetAll() ([]*pb.Region, error) {
	var regions []*pb.Region
	query := `SELECT id,name,created_at,updated_at from region where deleted_at is null`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		region := pb.Region{}
		var updated_at sql.NullTime
		err := rows.Scan(
			&region.Id,
			&region.Name,
			&region.Created_at,
			&updated_at,
		)
		if err != nil {
			return nil, err
		}
		if updated_at.Valid {
			region.Updated_at = updated_at.Time.String()
		}
		regions = append(regions, &region)
	}
	return regions, nil
}

func (r *regionsRepasitory) Delete(id string) (*pb.Region, error) {
	region, err := r.Get(id)
	if err != nil {
		return nil, err
	}
	query := `UPDATE region SET deleted_at = $2 where deleted_at is null and id=$1`
	_, err = r.db.Exec(query, id, time.Now().UTC())
	if err != nil {
		return nil, err
	}
	return region, nil
}
