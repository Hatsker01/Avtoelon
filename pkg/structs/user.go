package structs

type CreateUser struct{
	Phone string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Users struct{
	Users []User `json:"users"`
}

type User struct{
	Id string `json:"id"`
	Phone string `json:"phone"`
	Password string `json:"password"`
	Created_at string `json:"created_at"`
}
