package structs

type Car struct{
	Id string `json:"id"`
	Category_Id int `json:"category_id"`
	Model_Id int `json:"model_id"`
	Body_Id int	`json:"body_id"`
	Date string	`json:"date"`
	Price int	`json:"price"`
	Auction bool	`json:"auction"`
	Enginee string	`json:"enginee"`
	Oil_Id int	`json:"oil_id"`
	Transmission_id int	`json:"transmission_id"`
	Milage int	`json:"milage"`
	Color_id int	`json:"color_id"`
	Drive_unit_id int	`json:"drive_unit_id"`
	Outside_Id []int	`json:"outside_id"`
	Optic_Id []int	`json:"optic_id"`
	Salon_Id []int	`json:"salon_id"`
	Media_Id []int	`json:"media_id"`
	Options_Id []int	`json:"options_id"`
	Additionally_Id []int	`json:"additional_id"`
	Add_Info string	`json:"add_info"`
	Region_Id int	`json:"region_id"`
	City_Id int	`json:"city_id"`
	Phone string	`json:"phone"`
    Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Deleted_at string `json;"deleted_at"`
}

