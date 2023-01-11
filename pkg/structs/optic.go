package  structs

type Optic struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type CreateOptic struct {
	Name string `json:"name" binding:"required"`
}

type UpdateOpticReq struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Optics struct{
	Outsides []Outside `json:"Optics"`
}
