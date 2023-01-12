package structs

type Option struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type CreateOption struct {
	Name string `json:"name" binding:"required"`
}

type UpdateOptionReq struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Options struct {
	Options []Option `json:"Options"`
}
