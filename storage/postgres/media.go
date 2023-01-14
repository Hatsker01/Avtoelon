package postgres

import (
	"database/sql"
	"time"

	"github.com/Avtoelon/storage/repo"
	"github.com/jmoiron/sqlx"

	pb "github.com/Avtoelon/pkg/structs"
)

type mediasRepasitory struct {
	db *sqlx.DB
}

func NewMediasRepo(db *sqlx.DB) repo.MediasRepoInterface {
	return &mediasRepasitory{
		db: db,
	}
}

func (r *mediasRepasitory) Create(media *pb.CreateMedia) (*pb.Media, error) {
	newMedia := pb.Media{}
	query := `INSERT INTO media(name,created_at) VALUES ($1,$2) RETURNING id,name,created_at`
	err := r.db.QueryRow(query, media.Name, time.Now().UTC()).Scan(
		&newMedia.Id,
		&newMedia.Name,
		&newMedia.Created_at,
	)
	if err != nil {
		return nil, err
	}
	return &newMedia, nil
}

func (r *mediasRepasitory) Update(upMedia *pb.UpdateMediaReq) (*pb.Media, error) {
	media := pb.Media{}
	query := `UPDATE media SET name=$2,updated_at=$3 where id=$1 and deleted_at is null RETURNING id,name,created_at,updated_at`
	err := r.db.QueryRow(query, upMedia.Id, upMedia.Name, time.Now().UTC()).Scan(
		&media.Id,
		&media.Name,
		&media.Created_at,
		&media.Updated_at,
	)
	if err != nil {
		return nil, err
	}
	return &media, nil
}

func (r *mediasRepasitory) Get(id string) (*pb.Media, error) {
	media := pb.Media{}
	var updated_at sql.NullTime
	query := `SELECT id,name,created_at,updated_at from media where deleted_at is null and id=$1`
	err := r.db.QueryRow(query, id).Scan(
		&media.Id,
		&media.Name,
		&media.Created_at,
		&updated_at,
	)
	if err != nil {
		return nil, err
	}
	if updated_at.Valid {
		media.Updated_at = updated_at.Time.String()
	}
	return &media, nil
}

func (r *mediasRepasitory) GetAll() ([]*pb.Media, error) {
	var medias []*pb.Media
	query := `SELECT id,name,created_at,updated_at from media where deleted_at is null`
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		media := pb.Media{}
		var updated_at sql.NullTime
		err := rows.Scan(
			&media.Id,
			&media.Name,
			&media.Created_at,
			&updated_at,
		)
		if err != nil {
			return nil, err
		}
		if updated_at.Valid {
			media.Updated_at = updated_at.Time.String()
		}
		medias = append(medias, &media)
	}
	return medias, nil
}

func (r *mediasRepasitory) Delete(id string) (*pb.Media, error) {
	media, err := r.Get(id)
	if err != nil {
		return nil, err
	}
	query := `UPDATE media SET deleted_at=$2 where deleted_at is null and id=$1 `
	_, err = r.db.Exec(query, id, time.Now().UTC())
	if err != nil {
		return nil, err
	}
	return media, nil
}
