package structs

type Oil struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type CreateOil struct {
	Name string `json:"name" binding:"required"`
}

type UpdateOil struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type Oils struct {
	Oils []Oil `json"oils"`
}
