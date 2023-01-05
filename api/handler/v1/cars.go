package v1

import (
	"net/http"

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

	response,err:=postgres.NewCarsRepo(h.db).GetCar(guid)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err.Error(),
		})
		h.log.Error("failed while getting car by id",logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted,response)
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
func (h *handlerV1) GetAllCars(c *gin.Context){
	response,err:=postgres.NewCarsRepo(h.db).GetAllCars()
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err.Error(),
		})
		h.log.Error("failed while getting all cars info",logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted,response)
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
// @Router /v1/cars/:id [delete]
func (h *handlerV1) DeleteCar(c *gin.Context){
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id:=c.Param("id")
	
	response,err:=postgres.NewCarsRepo(h.db).DeleteCar(id)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err.Error(),
		})
		h.log.Error("failed while deleting car by id",logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted,response)
}
