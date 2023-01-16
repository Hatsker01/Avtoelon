package postgres

import (
	"database/sql"
	"time"

	pb "github.com/Avtoelon/pkg/structs"
	"github.com/Avtoelon/storage/repo"
	"github.com/jmoiron/sqlx"
)

type positionRepasitory struct {
	db *sqlx.DB
}

func NewPositionRepasitory(db *sqlx.DB) repo.PositionRepoInterface {
	return &positionRepasitory{
		db: db,
	}
}

func (r *positionRepasitory) Create(pos *pb.CreatePosition) (*pb.Position, error) {
	newPos := pb.Position{}
	query := `INSERT INTO position(model_id,name,created_at) VALUES ($1,$2,$3) RETURNING id,model_id,name,created_at`
	err := r.db.QueryRow(query, pos.Model_Id, pos.Name, time.Now().UTC()).Scan(
		&newPos.Id,
		&newPos.Model_Id,
		&newPos.Name,
		&newPos.Created_at,
	)
	if err != nil {
		return nil, err
	}
	return &newPos, nil
}

func (r *positionRepasitory) Update(upPos *pb.UpdatePostionReq) (*pb.Position, error) {
	pos := pb.Position{}
	query := `UPDATE position SET name=$2,updated_at=$3 where id=$1 and deleted_at is null RETURNING id,model_id,name,created_at,updated_at`
	err := r.db.QueryRow(query, upPos.Id, upPos.Name, time.Now().UTC()).Scan(
		&pos.Id,
		&pos.Model_Id,
		&pos.Name,
		&pos.Created_at,
		&pos.Updated_at,
	)
	if err != nil {
		return nil, err
	}
	return &pos, nil
}

func (r *positionRepasitory) Get(id string) (*pb.Position, error) {
	pos := pb.Position{}
	var updated_at sql.NullTime
	query := `SELECT id,model_id,name,created_at,updated_at from position where id=$1 and deleted_at is null`
	err := r.db.QueryRow(query, id).Scan(
		&pos.Id,
		&pos.Model_Id,
		&pos.Name,
		&pos.Created_at,
		&updated_at,
	)
	if err != nil {
		return nil, err
	}
	if updated_at.Valid {
		pos.Updated_at = updated_at.Time.String()
	}
	return &pos, nil
}

func (r *positionRepasitory) GetAll() ([]*pb.Position, error) {
	var positions []*pb.Position
	query := `SELECT id,model_id,name,created_at,updated_at from position where deleted_at is null`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		pos := pb.Position{}
		var updated_at sql.NullTime
		err := rows.Scan(
			&pos.Id,
			&pos.Model_Id,
			&pos.Name,
			&pos.Created_at,
			&updated_at,
		)
		if err != nil {
			return nil, err
		}
		if updated_at.Valid {
			pos.Updated_at = updated_at.Time.String()
		}
		positions = append(positions, &pos)
	}
	return positions, nil
}

func (r *positionRepasitory) Delete(id string) (*pb.Position, error) {
	pos, err := r.Get(id)
	if err != nil {
		return nil, err
	}
	query := `UPDATE posotion SET deleted_at = $1 where deleted_at is null and id=$1`
	_, err = r.db.Exec(query, id)
	if err != nil {
		return nil, err
	}
	return pos, nil
}
