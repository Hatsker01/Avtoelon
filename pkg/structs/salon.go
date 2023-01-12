package  structs

type Salon struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type CreateSalon struct {
	Name string `json:"name" binding:"required"`
}

type UpdateSalonReq struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Salons struct{
	Salons []Salon `json:"salons"`
}
