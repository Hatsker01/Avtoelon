package v1

import (
	"net/http"
	"strconv"

	"github.com/Avtoelon/pkg/logger"
	"github.com/Avtoelon/pkg/structs"
	"github.com/Avtoelon/storage/postgres"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateCar ...
// @Summary CreateCar
// @Description This API for creating new car for selling
// @Tags car
// @Accept json
// @Produce json
// @Param car body structs.CreateCarReq true "CreateCar"
// @Success 200 {object} structs.Car
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/cars [post]
func (h *handlerV1) CreateCar(c *gin.Context) {
	var body structs.Car

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while to blind json", logger.Error(err))
		return
	}
	response, err := postgres.NewCarsRepo(h.db).CreateCar(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while creating new car", logger.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}

// UpdateCar ...
// @Summary UpdateCar
// @Description This API for updating car
// @Tags car
// @Accept json
// @Produce json
// @Param updateCar body structs.UpdateCar true "UpdateCar"
// @Success 200 {object} structs.Car
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/cars/updateCar [put]
func (h *handlerV1) UpdateCar(c *gin.Context) {
	// er:=CheckClaims(h,c)
	// if er==nil{
	// 	c.JSON(http.StatusUnauthorized,gin.H{
	// 		"error":"error while checking token",
	// 	})
	// 	h.log.Error("failed while checking token")
	// 	return
	// }
	var body structs.Car
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while blinding json", logger.Error(err))
		return
	}
	response, err := postgres.NewCarsRepo(h.db).UpdateCar(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while updating car", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// GetCar ...
// @Summary GetCarById
// @Description This API for getting car by ID
// @Tags car
// @Accept json
// @Produce json
// @Param id path string true "Car_id"
// @Success 200 {object} structs.Car
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/cars/getById/{id} [get]
func (h *handlerV1) GetCar(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")

	response, err := postgres.NewCarsRepo(h.db).GetCar(guid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting car by id", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// GetAllCars ...
// @Summary GetAllCars
// @Description This API for getting all cars
// @Tags car
// @Accept json
// @Produce json
// @Success 200 {object} structs.Cars
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/cars/getAll [get]
func (h *handlerV1) GetAllCars(c *gin.Context) {
	response, err := postgres.NewCarsRepo(h.db).GetAllCars()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting all cars info", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// DeleteCar ....
// @Summary DeleteCar
// @Description This API for deleting car by id
// @Tags car
// @Accept json
// @Produce json
// @Param id path string true "Car_Id"
// @Success 200 {object} structs.Car
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/cars/{id} [delete]
func (h *handlerV1) DeleteCar(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	response, err := postgres.NewCarsRepo(h.db).DeleteCar(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while deleting car by id", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// GetCar ...
// @Summary GetCarById
// @Description This API for getting car by ID
// @Tags car
// @Accept json
// @Produce json
// @Param id path string true "Car_id"
// @Success 200 {object} structs.GetCar
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/cars/getInfo/{id} [get]
func (h *handlerV1) GetCarInfo(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")

	response, err := postgres.NewCarsRepo(h.db).GetCar(guid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting car by id", logger.Error(err))
		return
	}

	car := structs.GetCar{}

	category, err := postgres.NewCategoryRepasitory(h.db).GetCategory(strconv.Itoa(response.Category_Id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting category by id", logger.Error(err))
		return
	}
	car.Category = category.Name
	model, err := postgres.NewModelRepasitory(h.db).GetModel(strconv.Itoa(response.Model_Id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting model by id", logger.Error(err))
		return
	}
	car.Model = model.Name
	body, err := postgres.NewBodyRepo(h.db).GetBody(strconv.Itoa(response.Body_Id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting body by id", logger.Error(err))
		return
	}
	car.Body = body.Name
	oil, err := postgres.NewOilRepasitory(h.db).Get(strconv.Itoa(response.Oil_Id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting oil by id", logger.Error(err))
		return
	}
	car.Oil = oil.Name
	trans, err := postgres.NewTransmissionRepasitory(h.db).Get(strconv.Itoa(response.Transmission_id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting trans by id", logger.Error(err))
		return
	}
	car.Transmission = trans.Name
	color, err := postgres.NewColorRepasitory(h.db).Get(strconv.Itoa(response.Color_id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting color by id", logger.Error(err))
		return
	}
	car.Color = color.Name
	drive_unit, err := postgres.NewDriveUnitRepasitory(h.db).Get(strconv.Itoa(response.Drive_unit_id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting drive_unit by id", logger.Error(err))
		return
	}
	car.Drive_Unit = drive_unit.Name
	// for i, m := range response.Optic_Id {
	// 	fmt.Println(i, "-", int(m))
	// }
	var outsides []string

	for i := 0; i < len(response.Outside_Id); i++ {
		id, err := strconv.Atoi(response.Outside_Id[i])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			h.log.Error("failed while converting id", logger.Error(err))
			return
		}
		outside, err := postgres.NewOutsideRepasitory(h.db).GetOutside(strconv.Itoa(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			h.log.Error("failed while getting outside by id", logger.Error(err))
			return
		}
		outsides = append(outsides, outside.Name)
	}
	car.Outside = outsides

	var optics []string
	for i := 0; i < len(response.Optic_Id); i++ {
		id, err := strconv.Atoi(response.Optic_Id[i])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			h.log.Error("failed while converting id", logger.Error(err))
			return
		}
		optic, err := postgres.NewOpticRepasitory(h.db).Get(strconv.Itoa(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			h.log.Error("failed while getting outside by id", logger.Error(err))
			return
		}
		optics = append(optics, optic.Name)
	}
	car.Optic = optics

	var salons []string
	for i := 0; i < len(response.Salon_Id); i++ {
		id, err := strconv.Atoi(response.Salon_Id[i])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			h.log.Error("failed while converting id", logger.Error(err))
			return
		}
		salon, err := postgres.NewSalonsRepo(h.db).Get(strconv.Itoa(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			h.log.Error("failed while getting outside by id", logger.Error(err))
			return
		}
		salons = append(salons, salon.Name)
	}
	car.Salon = salons

	var medias []string
	for i := 0; i < len(response.Media_Id); i++ {
		id, err := strconv.Atoi(response.Media_Id[i])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			h.log.Error("failed while converting id", logger.Error(err))
			return
		}
		media, err := postgres.NewMediasRepo(h.db).Get(strconv.Itoa(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			h.log.Error("failed while getting outside by id", logger.Error(err))
			return
		}
		medias = append(medias, media.Name)
	}
	car.Media = medias

	var options []string
	for i := 0; i < len(response.Options_Id); i++ {
		id, err := strconv.Atoi(response.Options_Id[i])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			h.log.Error("failed while converting id", logger.Error(err))
			return
		}
		option, err := postgres.NewOptionsRepo(h.db).Get(strconv.Itoa(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			h.log.Error("failed while getting outside by id", logger.Error(err))
			return
		}
		options = append(options, option.Name)
	}
	car.Option = options

	var adds []string
	for i := 0; i < len(response.Additionally_Id); i++ {
		id, err := strconv.Atoi(response.Additionally_Id[i])
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			h.log.Error("failed while converting id", logger.Error(err))
			return
		}
		add, err := postgres.NewAdditionalsRepo(h.db).Get(strconv.Itoa(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			h.log.Error("failed while getting outside by id", logger.Error(err))
			return
		}
		adds = append(adds, add.Name)
	}
	car.Additional = adds

	region, err := postgres.NewRegionsRepo(h.db).Get(strconv.Itoa(response.Region_Id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting region by id", logger.Error(err))
		return
	}
	car.Region = region.Name

	city, err := postgres.NewCitiesRepo(h.db).Get(strconv.Itoa(response.City_Id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting city by id", logger.Error(err))
		return
	}
	car.City = city.Name

	marc, err := postgres.NewMarcsRepo(h.db).Get(strconv.Itoa(response.Marc_Id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting marc by id", logger.Error(err))
		return
	}
	car.Marc = marc.Name

	car.Id = response.Id
	car.Date = response.Date
	car.Price = response.Price
	car.Auction = response.Auction
	car.Enginee = response.Enginee
	car.Milage = response.Milage
	car.Add_Info = response.Add_Info
	car.Phone = response.Phone
	car.Created_at = response.Created_at
	car.Updated_at = response.Updated_at

	c.JSON(http.StatusAccepted, car)
}

// Get User Cars ...
// @Summary Get User Cars
// @Description This API for getting users' cars
// @Tags car
// @Accept json
// @Produce json
// @Param id path string true "User_Id"
// @Success 200 {object} structs.Car
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/cars/user/{id} [get]
func (h *handlerV1) GetUserCar(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	response, err := postgres.NewCarsRepo(h.db).UserCars(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting users' car", logger.Error(err))
		return
	}

	c.JSON(http.StatusAccepted, response)
}

// Get Cars By Price ...
// @Summary Get Car By Price
// @Description This API for getting cars by price
// @Tags car
// @Accept json
// @Produce json
// @Param high path bool true "High"
// @Success 200 {object} structs.Cars
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/cars/{high} [get]
func (h *handlerV1) GetCarByPrice(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	param := c.Param("high")
	high, err := strconv.ParseBool(param)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while parsing string to bool", logger.Error(err))
		return
	}
	response, err := postgres.NewCarsRepo(h.db).GetCarByPrice(high)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting car by price", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}
