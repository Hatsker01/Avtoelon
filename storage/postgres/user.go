package postgres

import (
	"time"

	pb "github.com/Avtoelon/pkg/structs"
	"github.com/Avtoelon/storage/repo"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

type usersRepasitory struct {
	db *sqlx.DB
}

func NewUsersRepasitory(db *sqlx.DB) repo.UsersRepoInterface {
	return &usersRepasitory{
		db: db,
	}
}

func (r *usersRepasitory) Create(user *pb.CreateUser) (*pb.User, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	newUser := pb.User{}
	query := `INSERT INTO users(id,phone,password,created_at) VALUES ($1,$2,$3,$4) RETURNING id,phone,created_at`
	err = r.db.QueryRow(query, id, user.Phone, user.Password, time.Now().UTC()).Scan(
		&newUser.Id,
		&newUser.Phone,
		&newUser.Password,
		&newUser.Created_at,
	)
	if err != nil {
		return nil, err
	}
	return &newUser, nil
}

func (r *usersRepasitory) Get(id string) (*pb.User, error) {
	user := pb.User{}
	query := `SELECT id,phone,password,created_at from users where id=$1 and deleted_at is null`
	err := r.db.QueryRow(query, id).Scan(
		&user.Id,
		&user.Phone,
		&user.Password,
		&user.Created_at,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *usersRepasitory) GetAll() ([]*pb.User, error) {
	var users []*pb.User
	query := `SELECT id,name,password,created_at from users where deleted_at is null`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		user := pb.User{}
		err := rows.Scan(
			&user.Id,
			&user.Phone,
			&user.Password,
			&user.Created_at,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (r *usersRepasitory) Delete(id string) (*pb.User, error) {
	user, err := r.Get(id)
	if err != nil {
		return nil, err
	}
	query := `UPDATE users SET deleted_at=$2 where id=$1 and deleted_at is null`
	_, err = r.db.Exec(query, id, time.Now().UTC())
	if err != nil {
		return nil, err
	}
	return user, nil
}
