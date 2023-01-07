package structs

type Category struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type CategoryCreateReq struct{
	Name string `json:"name" binding:"required"`
}

type UpdateCategory struct{
	Id int `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type Categories struct{
	Categories []Category `json:"categories"`
}