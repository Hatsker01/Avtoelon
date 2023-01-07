package structs

type Body struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type CreateBody struct {
	Name string `json:"name" binding:"required"`
}

type UpdateBody struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type Bodies struct{
	Bodies []Body `json:"bodies"`
}
