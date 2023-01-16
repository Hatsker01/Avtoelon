package postgres

import (
	"database/sql"
	"time"

	pb "github.com/Avtoelon/pkg/structs"
	"github.com/Avtoelon/storage/repo"
	"github.com/jmoiron/sqlx"
)

type categoryRepasitory struct {
	db *sqlx.DB
}

func NewCategoryRepasitory(db *sqlx.DB) repo.CategoryRepoInterface {
	return &categoryRepasitory{
		db: db,
	}
}

func (r *categoryRepasitory) CreateCategory(category *pb.CategoryCreateReq) (*pb.Category, error){
	newCategory:=pb.Category{}
	query:=`INSERT INTO category(name,created_at) VALUES ($1,$2) RETURNING id,name,created_at`
	err:=r.db.QueryRow(query,category.Name,time.Now().UTC()).Scan(
		&newCategory.Id,
		&newCategory.Name,
		&newCategory.Created_at,
	)
	if err!=nil{
		return nil,err
	}
	return &newCategory,nil
}

func (r *categoryRepasitory) UpdateCategory(category *pb.Category)(*pb.Category,error){
	upCategory:=pb.Category{}
	query:=`UPDATE category SET name=$2,update_at=$3 where deleted _at is null and id = $1 RETURNING id,name,created_at,updated_at`
	err:=r.db.QueryRow(query,category.Id,category.Name,time.Now().UTC()).Scan(
		&upCategory.Id,
		&upCategory.Name,
		&upCategory.Created_at,
		&upCategory.Updated_at,
	)
	if err!=nil{
		return nil,err
	}
	return &upCategory,nil
}

func (r *categoryRepasitory) GetCategory(id string)(*pb.Category,error){
	category:=pb.Category{}
	var updated_at sql.NullTime
	query:=`SELECT id,name,created_at,updated_at from category where deleted_at is null and id=$1`
	err:=r.db.QueryRow(query,id).Scan(
		&category.Id,
		&category.Name,
		&category.Created_at,
		&updated_at,
	)
	if err!=nil{
		return nil,err
	}
	if updated_at.Valid{
		category.Updated_at=updated_at.Time.String()
	}
	return &category,nil
}

func (r *categoryRepasitory) GetAllCategory()([]*pb.Category,error){
	var categories []*pb.Category
	query:=`SELECT id,name,category_id,updated_at from category where deleted_at is null`
	rows,err:=r.db.Query(query)
	if err!=nil{
		return nil,err
	}
	for rows.Next(){
		category:= pb.Category{}
		var updated_at sql.NullTime
		err:=rows.Scan(
			&category.Id,
			&category.Name,
			&category.Created_at,
			&updated_at,
		)
		if err!=nil{
			return nil,err
		}
		if updated_at.Valid{
			category.Updated_at=updated_at.Time.String()
		}
		categories = append(categories, &category)
	}
	return categories,nil
}

func (r *categoryRepasitory) DeleteCategory(id string)(*pb.Category,error){
	category,err:=r.GetCategory(id)
	if err!=nil{
		return nil,err
	}
	query:=`UPDATE category SET deleted_at =$2 where deleted_at is null and id=$1`
	_,err=r.db.Exec(query,id,time.Now().UTC())
	if err!=nil{
		return nil,err
	}
	return category,nil
}

