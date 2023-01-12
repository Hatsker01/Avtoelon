package structs

type Media struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type CreateMedia struct {
	Name string `json:"name" binding:"required"`
}

type UpdateMediaReq struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Medias struct{
	Medias []Media `json:"medias"`
}
