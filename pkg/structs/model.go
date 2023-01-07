package structs

type Model struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type CreateModelReq struct{
	Name string `json:"name" binding:"required"`
}

type UpdateModel struct{
	Id int `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type Models struct{
	Models []Model `json:"models"`
}