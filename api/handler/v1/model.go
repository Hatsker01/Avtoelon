package v1

import (
	"net/http"

	"github.com/Avtoelon/pkg/logger"
	"github.com/Avtoelon/pkg/structs"
	"github.com/Avtoelon/storage/postgres"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateModel ...
// @Summary CreateModel
// @Description This API for creating model
// @Tags model
// @Accept json
// @Produce json
// @Param model body structs.CreateModelReq true "CreateModelReq"
// @Success 200 {object} structs.Model
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/model [post]
func (h *handlerV1) CreateModel(c *gin.Context) {
	var body structs.CreateModelReq

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while blinding json", logger.Error(err))
		return
	}
	response, err := postgres.NewModelRepasitory(h.db).CreateModel(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while creating new model", logger.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}

// UpdateModel ...
// @Summary UpdateModel
// @Description This API for updating model
// @Tags model
// @Accept json
// @Produce json
// @Param model body structs.UpdateModel true "UpdateModel"
// @Success 200 {object} structs.Model
// @Failure 400 {object} structs.StandardErrorModel
// @Failuer 500 {object} structs.StandardErrorModel
// @Router /v1/model [put]
func (h *handlerV1) UpdateModel(c *gin.Context) {
	var body structs.UpdateModel

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while blinding json", logger.Error(err))
		return
	}
	response, err := postgres.NewModelRepasitory(h.db).UpdateModel(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while updating model", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// GetModel ...
// @Summary GetModel
// @Description This API for getting model by id
// @Tags model
// @Accept json
// @Produce json
// @Param id path string true "Model_id"
// @Success 200 {object} structs.Model
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/model/{id} [get]
func (h *handlerV1) GetModel(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	response, err := postgres.NewModelRepasitory(h.db).GetModel(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting model by id", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// GetAllModels ...
// @Summary GetAllModels
// @Description This API for getting all models
// @Tags model
// @Accept json
// @Produce json
// @Success 200 {object} structs.Models
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/model/getAll [get]
func (h *handlerV1) GetAllModels(c *gin.Context) {
	response, err := postgres.NewModelRepasitory(h.db).GetAllModels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting all models", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// DeleteModel ...
// @Summary DeleteModel
// @Description This API for deleting model by id
// @Tags model
// @Accept json
// @Produce json
// @Param id path string true "Model_Id"
// @Success 200 {object} structs.Model
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/model/{id} [delete]
func (h *handlerV1) DeleteModel(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")
	response, err := postgres.NewModelRepasitory(h.db).DeleteModel(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while deleting model by id", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// Get Car By Model Id
// @Summary Get Car By Model Id
// @Description This API for getting car by Model Id
// @Tags model
// @Accept json
// @Produce json
// @Param id path string true "Model_Id"
// @Success 200 {object} structs.Car
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/model/car/{id} [get]
func (h *handlerV1) GetCarByModel(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	response, err := postgres.NewModelRepasitory(h.db).GetCarByModel(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting car by model id", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}
