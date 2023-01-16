package structs

type Position struct {
	Id         int    `json:"id"`
	Model_Id   int    `json;"model_id"`
	Name       string `json:"name"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type CreatePosition struct {
	Model_Id int    `json:"model_id" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

type UpdatePostionReq struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Positions struct {
	Positions []Position `json:"positions"`
}
