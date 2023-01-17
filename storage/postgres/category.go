package postgres

import (
	"database/sql"
	"time"

	pb "github.com/Avtoelon/pkg/structs"
	"github.com/Avtoelon/storage/repo"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type categoryRepasitory struct {
	db *sqlx.DB
}

func NewCategoryRepasitory(db *sqlx.DB) repo.CategoryRepoInterface {
	return &categoryRepasitory{
		db: db,
	}
}

func (r *categoryRepasitory) CreateCategory(category *pb.CategoryCreateReq) (*pb.Category, error) {
	newCategory := pb.Category{}
	query := `INSERT INTO category(name,created_at) VALUES ($1,$2) RETURNING id,name,created_at`
	err := r.db.QueryRow(query, category.Name, time.Now().UTC()).Scan(
		&newCategory.Id,
		&newCategory.Name,
		&newCategory.Created_at,
	)
	if err != nil {
		return nil, err
	}
	return &newCategory, nil
}

func (r *categoryRepasitory) UpdateCategory(category *pb.Category) (*pb.Category, error) {
	upCategory := pb.Category{}
	query := `UPDATE category SET name=$2,update_at=$3 where deleted _at is null and id = $1 RETURNING id,name,created_at,updated_at`
	err := r.db.QueryRow(query, category.Id, category.Name, time.Now().UTC()).Scan(
		&upCategory.Id,
		&upCategory.Name,
		&upCategory.Created_at,
		&upCategory.Updated_at,
	)
	if err != nil {
		return nil, err
	}
	return &upCategory, nil
}

func (r *categoryRepasitory) GetCategory(id string) (*pb.Category, error) {
	category := pb.Category{}
	var updated_at sql.NullTime
	query := `SELECT id,name,created_at,updated_at from category where deleted_at is null and id=$1`
	err := r.db.QueryRow(query, id).Scan(
		&category.Id,
		&category.Name,
		&category.Created_at,
		&updated_at,
	)
	if err != nil {
		return nil, err
	}
	if updated_at.Valid {
		category.Updated_at = updated_at.Time.String()
	}
	return &category, nil
}

func (r *categoryRepasitory) GetAllCategory() ([]*pb.Category, error) {
	var categories []*pb.Category
	query := `SELECT id,name,category_id,updated_at from category where deleted_at is null`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		category := pb.Category{}
		var updated_at sql.NullTime
		err := rows.Scan(
			&category.Id,
			&category.Name,
			&category.Created_at,
			&updated_at,
		)
		if err != nil {
			return nil, err
		}
		if updated_at.Valid {
			category.Updated_at = updated_at.Time.String()
		}
		categories = append(categories, &category)
	}
	return categories, nil
}

func (r *categoryRepasitory) DeleteCategory(id string) (*pb.Category, error) {
	category, err := r.GetCategory(id)
	if err != nil {
		return nil, err
	}
	query := `UPDATE category SET deleted_at =$2 where deleted_at is null and id=$1`
	_, err = r.db.Exec(query, id, time.Now().UTC())
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (r *categoryRepasitory) GetCarByCategory(id string) (*pb.Car, error) {
	newCar := pb.Car{}
	var updated_at sql.NullTime
	query := `SELECT id,user_id,category_id,marc_id,model_id,position_id,body_id,date,price,auction,enginee,oil_id,transmission_id,milage,
	color_id,drive_unit_id,outside_id,optic_id,salon_id,media_id,options_id,additionally_id,add_info,region_id,city_id,
	phone,created_at,updated_at from cars JOIN category ON cars.category_id=category.id and cars.deleted_at is null and category.id=$1`
	err := r.db.QueryRow(query, id).Scan(
		&newCar.Id,
		&newCar.User_Id,
		&newCar.Category_Id,
		&newCar.Marc_Id,
		&newCar.Model_Id,
		&newCar.Position_Id,
		&newCar.Body_Id,
		&newCar.Date,
		&newCar.Price,
		&newCar.Auction,
		&newCar.Enginee,
		&newCar.Oil_Id,
		&newCar.Transmission_id,
		&newCar.Milage,
		&newCar.Color_id,
		&newCar.Drive_unit_id,
		pq.Array(&newCar.Outside_Id),
		pq.Array(&newCar.Optic_Id),
		pq.Array(&newCar.Salon_Id),
		pq.Array(&newCar.Media_Id),
		pq.Array(&newCar.Options_Id),
		pq.Array(&newCar.Additionally_Id),
		&newCar.Add_Info,
		&newCar.Region_Id,
		&newCar.City_Id,
		&newCar.Phone,
		&newCar.Created_at,
		&updated_at,
	)
	if err!=nil{
		return nil,err
	}
	if updated_at.Valid{
		newCar.Updated_at=updated_at.Time.String()
	}
	return &newCar,nil
}
