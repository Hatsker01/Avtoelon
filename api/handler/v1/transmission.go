package v1

import (
	"net/http"

	"github.com/Avtoelon/pkg/logger"
	"github.com/Avtoelon/pkg/structs"
	"github.com/Avtoelon/storage/postgres"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateTransmission ...
// @Summary CreateTransmission
// @Description This API forcreating transmission
// @Tags transmission
// @Accept json
// @Produce json
// @Param transmission body structs.CreateTrans true "Transmission"
// @Success 200 {object} structs.Transmission
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/transmission [post]
func (h *handlerV1) CreateTrans(c *gin.Context) {
	var body structs.CreateTrans

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while blinding json", logger.Error(err))
		return
	}

	response, err := postgres.NewTransmissionRepasitory(h.db).Create(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while creating transmission", logger.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}

// UpdateTransmission
// @Summary UpdateTransmission
// @Description This API for updating transmission
// @Tags transmission
// @Accept json
// @Produce json
// @Param trans body structs.UpdateTrans true "Update_Transmission"
// @Success 200 {object} structs.Transmission
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/transmission [put]
func (h *handlerV1) UpdateTrans(c *gin.Context) {
	var body structs.UpdateTrans

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while updating transmission", logger.Error(err))
		return
	}

	response, err := postgres.NewTransmissionRepasitory(h.db).Update(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while updating transmission", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// Get Transmission By Id ...
// @Summary GetTransmissionById
// @Description This API for getting transmission by id
// @Tags transmission
// @Accept json
// @Produce json
// @Param id path string true "Transmission_id"
// @Success 200 {object} structs.Transmission
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/transmission/{id} [get]
func (h *handlerV1) GetTransmission(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	response, err := postgres.NewTransmissionRepasitory(h.db).Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting transmision by id", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// GetAllTransmission
// @Summary GetAllTransmission
// @Description This API for getting all transmission
// @Tags transmission
// @Accept json
// @Produce json
// @Success 200 {object} structs.Transmissions
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/transmissions [get]
func (h *handlerV1) GetAllTrans(c *gin.Context) {
	response, err := postgres.NewTransmissionRepasitory(h.db).GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting all transmissions", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// DeleteTransmission
// @Summary DeleteTransmission
// @Description This API for deleting transmission
// @Tags transmission
// @Accept json
// @Produce json
// @Param id path string true "Transmission_id"
// @Success 200 {object} structs.Transmission
// @Success 400 {object} structs.StandardErrorModel
// @Success 500 {object} structs.StandardErrorModel
// @Router /v1/transmission/{id} [delete]
func (h *handlerV1) DeleteTransmission(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	response, err := postgres.NewTransmissionRepasitory(h.db).Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while deleting transmission by id", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}
