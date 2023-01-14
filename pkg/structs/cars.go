package structs

type Car struct {
	Id              string  `json:"id"`
	Category_Id     int     `json:"category_id"`
	Model_Id        int     `json:"model_id"`
	Body_Id         int     `json:"body_id"`
	Date            string  `json:"date"`
	Price           int     `json:"price"`
	Auction         bool    `json:"auction"`
	Enginee         string  `json:"enginee"`
	Oil_Id          int     `json:"oil_id"`
	Transmission_id int     `json:"transmission_id"`
	Milage          int     `json:"milage"`
	Color_id        int     `json:"color_id"`
	Drive_unit_id   int     `json:"drive_unit_id"`
	Outside_Id      []uint8 `json:"outside_id"`
	Optic_Id        []uint8 `json:"optic_id"`
	Salon_Id        []uint8 `json:"salon_id"`
	Media_Id        []uint8 `json:"media_id"`
	Options_Id      []uint8 `json:"options_id"`
	Additionally_Id []uint8 `json:"additional_id"`
	Add_Info        string  `json:"add_info"`
	Region_Id       int     `json:"region_id"`
	City_Id         int     `json:"city_id"`
	Phone           string  `json:"phone"`
	Created_at      string  `json:"created_at"`
	Updated_at      string  `json:"updated_at"`
	Deleted_at      string  `json:"deleted_at"`
}

type CreateCarReq struct {
	Category_Id     int    `json:"category_id" binding:"required"`
	Model_Id        int    `json:"model_id" binding:"required"`
	Body_Id         int    `json:"body_id" binding:"required"`
	Date            string `json:"date" binding:"required"`
	Price           int    `json:"price" binding:"required"`
	Auction         bool   `json:"auction" binding:"required"`
	Enginee         string `json:"enginee" binding:"required"`
	Oil_Id          int    `json:"oil_id" binding:"required"`
	Transmission_id int    `json:"transmission_id" binding:"required"`
	Milage          int    `json:"milage" binding:"required"`
	Color_id        int    `json:"color_id" binding:"required"`
	Drive_unit_id   int    `json:"drive_unit_id" binding:"required"`
	Outside_Id      []int  `json:"outside_id" binding:"required"`
	Optic_Id        []int  `json:"optic_id" binding:"required"`
	Salon_Id        []int  `json:"salon_id" binding:"required"`
	Media_Id        []int  `json:"media_id" binding:"required"`
	Options_Id      []int  `json:"options_id" binding:"required"`
	Additionally_Id []int  `json:"additional_id" binding:"required"`
	Add_Info        string `json:"add_info" binding:"required"`
	Region_Id       int    `json:"region_id" binding:"required"`
	City_Id         int    `json:"city_id" binding:"required"`
	Phone           string `json:"phone" binding:"required"`
}

type UpdateCar struct {
	Id              string `json:"id" binding:"required"`
	Category_Id     int    `json:"category_id" binding:"required"`
	Model_Id        int    `json:"model_id" binding:"required"`
	Body_Id         int    `json:"body_id" binding:"required"`
	Date            string `json:"date" binding:"required"`
	Price           int    `json:"price" binding:"required"`
	Auction         bool   `json:"auction" binding:"required"`
	Enginee         string `json:"enginee" binding:"required"`
	Oil_Id          int    `json:"oil_id" binding:"required"`
	Transmission_id int    `json:"transmission_id" binding:"required"`
	Milage          int    `json:"milage" binding:"required"`
	Color_id        int    `json:"color_id" binding:"required"`
	Drive_unit_id   int    `json:"drive_unit_id" binding:"required"`
	Outside_Id      []int  `json:"outside_id" binding:"required"`
	Optic_Id        []int  `json:"optic_id" binding:"required"`
	Salon_Id        []int  `json:"salon_id" binding:"required"`
	Media_Id        []int  `json:"media_id" binding:"required"`
	Options_Id      []int  `json:"options_id" binding:"required"`
	Additionally_Id []int  `json:"additional_id" binding:"required"`
	Add_Info        string `json:"add_info" binding:"required"`
	Region_Id       int    `json:"region_id" binding:"required"`
	City_Id         int    `json:"city_id" binding:"required"`
	Phone           string `json:"phone" binding:"required"`
}

type Cars struct {
	Cars []Car `json:"cars"`
}
