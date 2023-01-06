package structs

type Outside struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type CreateOutside struct{
	Name string `json:"name" binding:"required"`
}



