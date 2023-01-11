package structs

type Drive_Unit struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type DriveUnitCreateReq struct{
	Name string `json:"name" binding:"required"`
}

type UpdateDriveUnit struct{
	Id int `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type DriveUnits struct{
	DriveUnits []Drive_Unit `json:"drive_units"`
}