package structs

type City struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type CreateCity struct {
	Name string `json:"name" binding:"required"`
}

type UpdateCityReq struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Cities struct {
	Cities []City `json:"cities"`
}
