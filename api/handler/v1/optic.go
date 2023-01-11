package v1

import (
	"net/http"

	"github.com/Avtoelon/pkg/logger"
	"github.com/Avtoelon/pkg/structs"
	"github.com/Avtoelon/storage/postgres"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateOptic ...
// @Summary CreateOptic
// @Description This API for creating Optic for cars
// @Tags optic
// @Accept json
// @Produce json
// @Param optic body structs.CreateOptic true "CreateOptic"
// @Success 200 {object} structs.Optic
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/optic [post]
func (h *handlerV1) CreateOptic(c *gin.Context) {
	var body structs.CreateOptic
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while blinding json optic", logger.Error(err))
		return
	}

	response, err := postgres.NewOpticRepasitory(h.db).Create(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while creating optic", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

//  UpdateOptic ...
// @Summary UpdateOptic
// @Description This API for creating optic for cars
// @Tags optic
// @Accept json
// @Produce json
// @Param optic body structs.UpdateOpticReq true "Update_Optic"
// @Success 200 {object} structs.Optic
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/optic [put]
func (h *handlerV1) UpdateOptic(c *gin.Context) {
	var body structs.UpdateOpticReq
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while blinding json optic", logger.Error(err))
		return
	}
	response, err := postgres.NewOpticRepasitory(h.db).Update(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while updating optic", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// GetOptic ...
// @Summary GetOpticById
// @Description This API for getting optic by ID
// @Tags optic
// @Accept json
// @Produce json
// @Param id path string true "Optic_Id"
// @Success 200 {object} structs.Optic
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/optic/{id} [get]
func (h *handlerV1) GetOptic(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	response, err := postgres.NewOpticRepasitory(h.db).Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting optic by id", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// GetAllOptics ...
// @Summary GetAllOptics
// @Description This API for getting all optics
// @Tags optic
// @Accept json
// @Produce json
// @Success 200 {object} structs.Optics
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/optics [get]
func (h *handlerV1) GetAllOptics(c *gin.Context) {
	response, err := postgres.NewOpticRepasitory(h.db).GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting all models", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// Delete Optic ....
// @Summary DeleteOptic
// @Description This API for deleting optic by id
// @Tags optic
// @Accept json
// @Produce json
// @Param id path string true "Optic_id"
// @Success 200 {object} structs.Optic
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/optic/{id} [delete]
func (h *handlerV1) DeleteOptic(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	response, err := postgres.NewOpticRepasitory(h.db).Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while deleting optic by id", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}
