package structs

type Marc struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type CreateMarc struct {
	Name string `json:"name" binding:"required"`
}

type UpdateMarcReq struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Marcs struct {
	Marcs []Marc `json:"marcs"`
}
