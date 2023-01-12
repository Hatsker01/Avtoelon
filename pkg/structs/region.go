package structs

type Region struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type CreateRegion struct {
	Name string `json:"name" binding:"required"`
}

type UpdateRegionReq struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Regions struct {
	Regions []Region `json:"regions"`
}
