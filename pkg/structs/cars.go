package structs

type Car struct {
	Id              string   `json:"id"`
	User_Id         string   `json:"user_id"`
	Category_Id     int      `json:"category_id"`
	Marc_Id         int      `json:"marc_id"`
	Model_Id        int      `json:"model_id"`
	Position_Id     int      `json:"position_id"`
	Body_Id         int      `json:"body_id"`
	Date            string   `json:"date"`
	Price           int      `json:"price"`
	Auction         bool     `json:"auction"`
	Enginee         string   `json:"enginee"`
	Oil_Id          int      `json:"oil_id"`
	Transmission_id int      `json:"transmission_id"`
	Milage          int      `json:"milage"`
	Color_id        int      `json:"color_id"`
	Drive_unit_id   int      `json:"drive_unit_id"`
	Outside_Id      []string `json:"outside_id"`
	Optic_Id        []string `json:"optic_id"`
	Salon_Id        []string `json:"salon_id"`
	Media_Id        []string `json:"media_id"`
	Options_Id      []string `json:"options_id"`
	Additionally_Id []string `json:"additional_id"`
	Add_Info        string   `json:"add_info"`
	Region_Id       int      `json:"region_id"`
	City_Id         int      `json:"city_id"`
	Phone           string   `json:"phone"`
	Created_at      string   `json:"created_at"`
	Updated_at      string   `json:"updated_at"`
	Deleted_at      string   `json:"deleted_at"`
}
type GetCar struct {
	Id           string   `json:"id"`
	Category     string   `json:"category"`
	Marc         string   `json:"marc"`
	Model        string   `json:"model"`
	Position     string   `json:"position"`
	Body         string   `json:"body"`
	Date         string   `json:"date"`
	Price        int      `json:"price"`
	Auction      bool     `json:"auction"`
	Enginee      string   `json:"enginee"`
	Oil          string   `json:"oil"`
	Transmission string   `json:"transmission"`
	Milage       int      `json:"milage"`
	Color        string   `json:"color"`
	Drive_Unit   string   `json:"drive_unit"`
	Outside      []string `json:"outside"`
	Optic        []string `json:"optic"`
	Salon        []string `json:"salon"`
	Media        []string `json:"media"`
	Option       []string `json:"option"`
	Additional   []string `json:"additional"`
	Add_Info     string   `json:"add_info"`
	Region       string   `json:"region"`
	City         string   `json:"city"`
	Phone        string   `json:"phone"`
	Created_at   string   `json:"created_at"`
	Updated_at   string   `json:"updated_at"`
}

type CreateCarReq struct {
	User_Id         string   `json:"user_id" binding:"required"`
	Category_Id     int      `json:"category_id" binding:"required"`
	Marc_Id         int      `json:"marc_id" binding:"required"`
	Model_Id        int      `json:"model_id" binding:"required"`
	Position_Id     int      `json:"position_id" binding:"required"`
	Body_Id         int      `json:"body_id" binding:"required"`
	Date            string   `json:"date" binding:"required"`
	Price           int      `json:"price" binding:"required"`
	Auction         bool     `json:"auction" binding:"required"`
	Enginee         string   `json:"enginee" binding:"required"`
	Oil_Id          int      `json:"oil_id" binding:"required"`
	Transmission_id int      `json:"transmission_id" binding:"required"`
	Milage          int      `json:"milage" binding:"required"`
	Color_id        int      `json:"color_id" binding:"required"`
	Drive_unit_id   int      `json:"drive_unit_id" binding:"required"`
	Outside_Id      []string `json:"outside_id" binding:"required"`
	Optic_Id        []string `json:"optic_id" binding:"required"`
	Salon_Id        []string `json:"salon_id" binding:"required"`
	Media_Id        []string `json:"media_id" binding:"required"`
	Options_Id      []string `json:"options_id" binding:"required"`
	Additionally_Id []string `json:"additional_id" binding:"required"`
	Add_Info        string   `json:"add_info" binding:"required"`
	Region_Id       int      `json:"region_id" binding:"required"`
	City_Id         int      `json:"city_id" binding:"required"`
	Phone           string   `json:"phone" binding:"required"`
}

type UpdateCar struct {
	Id              string `json:"id" binding:"required"`
	Category_Id     int    `json:"category_id" binding:"required"`
	Marc_Id         int    `json:"marc_id" binding:"required"`
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
