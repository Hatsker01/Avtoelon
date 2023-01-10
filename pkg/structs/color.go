package structs

type Color struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type CreateColorReq struct {
	Name string `json:"name" binding:"required"`
}

type UpdateColor struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type Colors struct {
	Colors []Model `json:"colors"`
}
