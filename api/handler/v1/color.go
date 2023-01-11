package v1

import (
	"net/http"

	"github.com/Avtoelon/pkg/logger"
	"github.com/Avtoelon/pkg/structs"
	"github.com/Avtoelon/storage/postgres"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateColor ...
// @Summary CreateColor
// @Description This API for creating color
// @Tags color
// @Accept json
// @Produce json
// @Param color body structs.CreateColorReq true "CreateColor"
// @Success 200 {object} structs.Color
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/color [post]
func (h *handlerV1) CreateColor(c *gin.Context) {
	var body structs.CreateColorReq
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while blinding json", logger.Error(err))
		return
	}

	response, err := postgres.NewColorRepasitory(h.db).Create(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while creating color", logger.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}

// UpdateColor ...
// @Summary UpdateColor
// @Description This API for updating color
// @Tags color
// @Accept json
// @Produce json
// @Param color body structs.UpdateColor true "UpdateColor"
// @Success 200 {object} structs.Color
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/color [put]
func (h *handlerV1) UpdateColor(c *gin.Context) {
	var body structs.UpdateColor

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while blinding json", logger.Error(err))
		return
	}

	response, err := postgres.NewColorRepasitory(h.db).Update(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while updating color", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// GetColorById ...
// @Summary GetColorById
// @Description This API for getting color by id
// @Tags color
// @Accept json
// @Produce json
// @Param id path string true "Color_id"
// @Success 200 {object} structs.Color
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/color/{id} [get]
func (h *handlerV1) GetColor(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	response, err := postgres.NewColorRepasitory(h.db).Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting color by id", logger.Error(err))
		return
	}

	c.JSON(http.StatusAccepted, response)
}

// GetAllColors ...
// @Summary GetAllColors
// @Description This APi for gettig all Colors
// @Tags color
// @Accept json
// @Produce json
// @Success 200 {object} structs.Colors
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/colors [get]
func (h *handlerV1) GetAllColors(c *gin.Context) {
	response, err := postgres.NewColorRepasitory(h.db).GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting all colors", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// DeleteColor ...
// @Summary DeleteColor
// @Description This API for deleting color by id
// @Tags color
// @Accept json
// @Produce json
// @Param id path string true "Color_id"
// @Success 200 {object} structs.Color
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/color/{id} [delete]
func (h *handlerV1) DeleteColor(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	response, err := postgres.NewColorRepasitory(h.db).Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while deleting color", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}
