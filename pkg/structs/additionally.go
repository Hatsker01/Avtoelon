package structs

type Additional struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type CreateAdditional struct {
	Name string `json:"name" binding:"required"`
}

type UpdateAdditionalReq struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Additionals struct {
	Additionals []Additional `json:"additionals"`
}
