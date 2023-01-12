package v1

import (
	"net/http"

	"github.com/Avtoelon/pkg/logger"
	"github.com/Avtoelon/pkg/structs"
	"github.com/Avtoelon/storage/postgres"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateCity ...
// @Summary CreateCity
// @Description This API for creating city
// @Tags city
// @Accept json
// @Produce json
// @Param city body structs.CreateCity true "CreateCity"
// @Success 200 {object} structs.City
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/city [post]
func (h *handlerV1) CreateCity(c *gin.Context) {
	var body structs.CreateCity

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while binding json", logger.Error(err))
		return
	}

	response, err := postgres.NewCitiesRepo(h.db).Create(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while creating new city", logger.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}

// UpdateCity ...
// @Summary UpdateCity
// @Description This API for updating city
// @Tags city
// @Accept json
// @Produce json
// @Param city body structs.UpdateCityReq true "UpdateCity"
// @Success 200 {object} structs.City
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/city [put]
func (h *handlerV1) UpdateCity(c *gin.Context) {
	var body structs.UpdateCityReq

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while binding json", logger.Error(err))
		return
	}

	response, err := postgres.NewCitiesRepo(h.db).Update(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while updating city", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// GetCityById ...
// @Summary GetCityById
// @Description This API for getting city by Id
// @Tags city
// @Accept json
// @Produce json
// @Param id path string true "City_Id"
// @Success 200 {object} structs.City
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/city/{id} [get]
func (h *handlerV1) GetCity(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	response, err := postgres.NewCitiesRepo(h.db).Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting city by Id", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// GetAllCities ...
// @Summary GetAllCities
// @Description This API for getting all cities
// @Tags city
// @Accept json
// @Produce json
// @Success 200 {object} structs.Cities
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/cities [get]
func (h *handlerV1) GetAllCities(c *gin.Context) {
	response, err := postgres.NewCitiesRepo(h.db).GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting all cities ", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// DeleteCity ...
// @Summary DeleteCity
// @Description This API for deleting city by id
// @Tags city
// @Accept json
// @Produce json
// @Param id path string true "City_Id"
// @Success 200 {object} structs.City
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/city/{id} [delete]
func (h *handlerV1) DeleteCity(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	response, err := postgres.NewCitiesRepo(h.db).Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while deleting city by id", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}
