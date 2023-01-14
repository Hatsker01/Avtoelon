package v1

import (
	"net/http"

	"github.com/Avtoelon/pkg/logger"
	"github.com/Avtoelon/pkg/structs"
	"github.com/Avtoelon/storage/postgres"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateAdditionally ...
// @Summary CreateAdditionally
// @Description This API for creating additionally for cars
// @Tags additionally
// @Accept json
// @Produce json
// @Param add body structs.CreateAdditional true "CreateAdditionally"
// @Success 200 {object} structs.Additional
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/additionally [post]
func (h *handlerV1) CreateAdd(c *gin.Context) {
	var body structs.CreateAdditional
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(), 	 
		})
		h.log.Error("failed while binding json", logger.Error(err))
		return
	}

	response, err := postgres.NewAdditionalsRepo(h.db).Create(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while create additionally", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// UpdateAdditionally ...
// @Summary UpdateAdditionally
// @Description This API for updating additionally
// @Tags additionally
// @Accept json
// @Produce json
// @Param add body structs.UpdateAdditionalReq true "UpdateAdditionally"
// @Success 200 {object} structs.Additional
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/additionally [put]
func (h *handlerV1) UpdateAdditionally(c *gin.Context) {
	var body structs.UpdateAdditionalReq
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while binding json", logger.Error(err))
		return
	}

	response, err := postgres.NewAdditionalsRepo(h.db).Update(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while updating additional", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// GetAdditional ...
// @Summary GetAdditionally
// @Description This API for getting additional by id
// @Tags additionally
// @Accept json
// @Produce json
// @Param id path string true "Additionally_Id"
// @Success 200 {object} structs.Additional
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/additionally/{id} [get]
func (h *handlerV1) GetAdditional(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	response, err := postgres.NewAdditionalsRepo(h.db).Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting Additional by id", logger.Error(err))
		return
	}

	c.JSON(http.StatusAccepted, response)
}

// GetAllAdditional ...
// @Summary GetAllAdditional
// @Description This API for getting all addittionals
// @Tags additionally
// @Accept json
// @Produce json
// @Success 200 {object} structs.Additionals
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/additionals [get]
func (h *handlerV1) GetAllAdditional(c *gin.Context) {
	response, err := postgres.NewAdditionalsRepo(h.db).GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting all additionals", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// DeleteAdditional ...
// @Summary DeleteAdditional
// @Description This API for deleting additional by id
// @Tags additionally
// @Accept json
// @Produce json
// @Param id path string true "Additional_Id"
// @Success 200 {object} structs.Additional
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/additionally/{id} [delete]
func (h *handlerV1) DeleteAdditional(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	response, err := postgres.NewAdditionalsRepo(h.db).Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while deleting additional by id", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}
