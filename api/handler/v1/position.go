package v1

import (
	"net/http"

	"github.com/Avtoelon/pkg/logger"
	"github.com/Avtoelon/pkg/structs"
	"github.com/Avtoelon/storage/postgres"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// Create Position ...
// @Summary Create Position
// @Description This API for creating position
// @Tags position
// @Accept json
// @Produce json
// @Param position body structs.CreatePosition true "CreatePosition"
// @Success 201 {object} structs.Position
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/position [post]
func (h *handlerV1) CreatePosition(c *gin.Context) {
	var body structs.CreatePosition

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while creating position", logger.Error(err))
		return
	}

	response, err := postgres.NewPositionRepasitory(h.db).Create(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while creating new position for model", logger.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}

// Update Position ...
// @Summary UpdatePosition
// @Description This API for updating position
// @Tags position
// @Accept json
// @Produce json
// @Param position body structs.UpdatePostionReq true "UpdatePosition"
// @Success 200 {object} structs.Position
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/position [put]
func (h *handlerV1) UpdatePosition(c *gin.Context) {
	var body structs.UpdatePostionReq

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while binding json", logger.Error(err))
		return
	}

	response, err := postgres.NewPositionRepasitory(h.db).Update(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while updating position of model", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// Get Position By Id ...
// @Summary Get Position By Id
// @Description This API for getting position by Id
// @Tags position
// @Accept json
// @Produce json
// @Param id path string true "Position_Id"
// @Success 200 {object} structs.Position
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/position/{id} [get]
func (h *handlerV1) GetPosition(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id:=c.Param("id")

	response,err:=postgres.NewPositionRepasitory(h.db).Get(id)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err.Error(),
		})
		h.log.Error("failed while getting position by id",logger.Error(err))
		return 
	}
	c.JSON(http.StatusAccepted,response)
}


// Get All Positions ...
// @Summary Get All Positions
// @Description This API for getting all positions
// @Tags position
// @Accept json
// @Produce json
// @Success 200 {object} structs.Positions
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/positions [get]
func (h *handlerV1) GetAllPositions(c *gin.Context){
	response,err:=postgres.NewPositionRepasitory(h.db).GetAll()
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err.Error(),
		})
		h.log.Error("failed while geting all positions",logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted,response)
}

// Delete Position By Id
// @Summary Delete Position
// @Description This API for deleting position by ID
// @Tags position 
// @Accept json
// @Produce json
// @Param id path string true "Position_Id"
// @Success 200 {object} structs.Position
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/position/{id} [delete]
func(h *handlerV1) DeletePosition(c *gin.Context){
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id:=c.Param("id")

	response,err:=postgres.NewPositionRepasitory(h.db).Delete(id)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err.Error(),
		})
		h.log.Error("failed while deleting position by Id")
	}
}