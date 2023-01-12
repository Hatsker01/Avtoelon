package structs

type Transmission struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type CreateTrans struct {
	Name string `json:"name" binding:"required"`
}

type UpdateTrans struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type Transmissions struct {
	Transmissions []Transmission `json"transmissions"`
}
