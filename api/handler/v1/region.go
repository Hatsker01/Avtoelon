package v1

import (
	"net/http"

	"github.com/Avtoelon/pkg/logger"
	"github.com/Avtoelon/pkg/structs"
	"github.com/Avtoelon/storage/postgres"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateRegion ...
// @Summary CreateRegion
// @Description This API for creating region
// @Tags region
// @Accept json
// @Produce json
// @Param region body structs.CreateRegion true "CreateRegion"
// @Success 200 {object} structs.Region
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/region [post]
func (h *handlerV1) CreateRegion(c *gin.Context) {
	var body structs.CreateRegion
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while blinding json", logger.Error(err))
		return
	}

	response, err := postgres.NewRegionsRepo(h.db).Create(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while creating region", logger.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}

// UpdateRegion ...
// @Summary UpdateRegion
// @Description This API for updating region
// @Tags region
// @Accept json
// @Produce json
// @Param region body structs.UpdateRegionReq true "UpdateRegion"
// @Success 200 {object} structs.Region
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/region [put]
func (h *handlerV1) UpdateRegion(c *gin.Context) {
	var body structs.UpdateRegionReq

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while blinding json", logger.Error(err))
		return
	}

	response, err := postgres.NewRegionsRepo(h.db).Update(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while updating region ", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// GetRegionById ...
// @Summary GetRegionById
// @Description This API for getting region by Id
// @Tags region
// @Accept json
// @Produce json
// @Param id path string true "Region_Id"
// @Success 200 {object} structs.Region
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/region/{id} [get]
func (h *handlerV1) GetRegion(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	response, err := postgres.NewRegionsRepo(h.db).Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while hetting region by id", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// GetAllRegions ...
// @Summary GetAllRegions
// @Description This API for getting all regions
// @Tags region
// @Accept json
// @Produce json
// @Success 200 {object} structs.Regions
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/regions [get]
func (h *handlerV1) GetAllRegions(c *gin.Context) {
	response, err := postgres.NewRegionsRepo(h.db).GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting all regions", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// DeletingRegion ...
// @Summary DeletingRegion
// @Description This API for deleting region by id
// @Tags region
// @Accept json
// @Produce json
// @Param id path string true "Region_Id"
// @Success 200 {object} structs.Region
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/region/{id} [delete]
func (h *handlerV1) DeleteRegion(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	response, err := postgres.NewRegionsRepo(h.db).Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while deleting region by id", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// Get Car By Region ...
// @Summary Get Car By Region Id
// @Description This API for getting car by region Id
// @Tags region
// @Accept json
// @Produce json
// @Param id path string true "Region_Id"
// @Success 200 {object} structs.Car
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/region/car/{id} [get]
func (h *handlerV1) GetCarByRegion(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	response, err := postgres.NewRegionsRepo(h.db).GetCarByRegion(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting car by region", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}
