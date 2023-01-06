package v1

import (
	"net/http"

	"github.com/Avtoelon/pkg/logger"
	"github.com/Avtoelon/pkg/structs"
	"github.com/Avtoelon/storage/postgres"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateOutside ....
// @Summary CreateOutside
// @Description This API for creating outside type for cars
// @Tags outside
// @Accept json
// @Produce json
// @Param CreateOutside body structs.CreateOutside true "CreateOutside"
// @Success 200 {object} structs.CreateOutside
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/outside [post]
func (h *handlerV1) CreateOutside(c *gin.Context) {
	var body structs.CreateOutside

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while blinding json", logger.Error(err))
		return
	}
	response, err := postgres.NewOutsideRepasitory(h.db).CreateOutside(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while creating outside for cars", logger.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}

// UpdateOutside ...
// @Summary UpdateOutside
// @Description This API for updating
// @Tags outside
// @Accept json
// @Produce json
// @Param UpdateOutside body structs.UpdateOutsideReq true "UpdateOutside"
// @Success 200 {object} structs.Outside
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/outside [put]
func (h *handlerV1) UpdateOutside(c *gin.Context) {
	var body structs.Outside

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while blinding json", logger.Error(err))
		return
	}

	response, err := postgres.NewOutsideRepasitory(h.db).UpdateOutside(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while updating outside", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// GetOutside
// @Summary GetOutside
// @Description This API for getting outside by ID
// @Tags outside
// @Accept json
// @Produce json
// @Param id path string true "Outside_ID"
// @Success 200 {object} structs.Outside
// @Success 400 {object} structs.StandardErrorModel
// @Success 500 {object} structs.StandardErrorModel
// @Router /v1/outside/{id} [get]
func (h *handlerV1) GetOutside(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	response, err := postgres.NewOutsideRepasitory(h.db).GetOutside(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("error while getting outside by ID", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// GetAllOutside ...
// @Summary GetAllOutside
// @Description This API for getting all outsides
// @Tags outside
// @Accept json
// @Produce json
// @Success 200 {object} structs.Outsides
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/outside/getAll [get]
func (h *handlerV1) GetAllOutside(c *gin.Context) {
	response, err := postgres.NewOutsideRepasitory(h.db).GetAllOutside()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting all outsides", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// DeletedOutside
// @Summary DeletedOutside
// @Description This API for deleting outside by id
// @Tags outside
// @Accept json
// @Produce json
// @Param id path string true "Outside_ID"
// @Success 200 {object} structs.Outside
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/outside/:id [delete]
func (h *handlerV1) DeletedOutside(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	response, err := postgres.NewOutsideRepasitory(h.db).DeletedOutside(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while deleting outside by id", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}
