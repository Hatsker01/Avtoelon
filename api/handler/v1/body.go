package v1

import (
	"net/http"

	"github.com/Avtoelon/pkg/logger"
	"github.com/Avtoelon/pkg/structs"
	"github.com/Avtoelon/storage/postgres"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateBody ...
// @Summary CreateBody
// @Description This API for creating body
// @Tags body
// @Accept json
// @Produce json
// @Param body body structs.CreateBody true "CreateBody"
// @Success 200 {object} structs.Body
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/body [post]
func (h *handlerV1) CreateBody(c *gin.Context) {
	var body structs.CreateBody
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while  blinding body", logger.Error(err))
		return
	}
	response, err := postgres.NewBodyRepo(h.db).CreateBody(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while creating body", logger.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}

// UpdateBody ...
// @Summary UpdateBody
// @Description This API for updating body
// @Tags body
// @Accept json
// @Produce json
// @Param body body structs.UpdateBody true "Update_Body"
// @Success 200 {object} structs.Body
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/body [put]
func (h *handlerV1) UpdateBody(c *gin.Context) {
	var body structs.UpdateBody

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while blinding json", logger.Error(err))
		return
	}

	response, err := postgres.NewBodyRepo(h.db).UpdateBody(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while updating body", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// GetBody ...
// @Summary GetBody
// @Description This API for getting body by id
// @Tags body
// @Accept json
// @Produce json
// @Param id path string true "Body_Id"
// @Success 200 {object} structs.Body
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/body/{id} [get]
func (h *handlerV1) GetBody(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	response, err := postgres.NewBodyRepo(h.db).GetBody(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("error while getting body by id", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// GetAllBody ...
// @Summary GetAllBody
// @Description This API for getting body by id
// @Tags body
// @Accept json
// @Produce json
// @Success 200 {object} structs.Bodies
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/bodies [get]
func (h *handlerV1) GetAllBody(c *gin.Context) {
	response, err := postgres.NewBodyRepo(h.db).GetAllBody()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting all bodies", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// DeleteBody ...
// @Summary DeleteBody
// @Description This API for deleting body
// @Tags body
// @Accept json
// @Produce json
// @Param id path string true "Body_Id"
// @Success 200 {object} structs.Body
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/body/{id} [delete]
func (h *handlerV1) DeleteBody(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	response, err := postgres.NewBodyRepo(h.db).DeleteBody(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while deleting body", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// Get Car By Body Id ...
// @Summary Get Car By Id
// @Desctiption This API for getting car By Body Id
// @Tags body
// @Accept json
// @Produce json
// @Param id path string true "Body_Id"
// @Success 200 {object} structs.Car
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/body/car/{id} [get]
func (h *handlerV1) GetCarByBody(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	response, err := postgres.NewBodyRepo(h.db).GetCarByBody(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting car by body id", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}
