package v1

import (
	"net/http"

	"github.com/Avtoelon/pkg/logger"
	"github.com/Avtoelon/pkg/structs"
	"github.com/Avtoelon/storage/postgres"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateOil ..
// @Summary CreateOil
// @Description This API for creating oil
// @Tags oil
// @Accept json
// @Produce json
// @Param oil body structs.CreateOil true "Create_oil"
// @Success 200 {object} structs.Oil
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/oil [post]
func (h *handlerV1) CreateOil(c *gin.Context) {
	var body structs.CreateOil

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while blinding oil", logger.Error(err))
		return
	}

	response, err := postgres.NewOilRepasitory(h.db).Create(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while creating oil", logger.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}

// UpdateOil ...
// @Summary UpdateOil
// @Description This API for updating oil
// @Tags oil
// @Accept json
// @Produce json
// @Param oil body structs.UpdateOil true "Update_oil"
// @Success 200 {object} structs.Oil
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/oil [put]
func (h *handlerV1) UpdateOil(c *gin.Context) {
	var body structs.UpdateOil

	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while binding updating oil", logger.Error(err))
		return
	}

	response, err := postgres.NewOilRepasitory(h.db).Update(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while updating oil", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// GetOilById ...
// @Summary GetOilById
// @Description This API for getting oil by ID
// @Tags oil
// @Accept json
// @Produce json
// @Param id path string true "Oil_id"
// @Success 200 {object} structs.Oil
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/oil/{id} [get]
func (h *handlerV1) GetOil(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	response, err := postgres.NewOilRepasitory(h.db).Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting oil by id", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// GetAllOils ...
// @Summary GetAllOils
// @Description This API for getting all oils
// @Tags oil
// @Accept json
// @Produce json
// @Success 200 {object} structs.Oils
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/oils [get]
func (h *handlerV1) GetAllOils(c *gin.Context) {
	response, err := postgres.NewOilRepasitory(h.db).GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting all oils", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// DeleteOil
// @Summary DeleteOil
// @Description This API for deleting oil by ID
// @Tags oil
// @Accept json
// @Produce json
// @Param id path string true "Oil_id"
// @Success 200 {object} structs.Oil
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/oil/{id} [delete]
func (h *handlerV1) DeleteOil(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	response, err := postgres.NewOilRepasitory(h.db).Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while deleting oil", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}
