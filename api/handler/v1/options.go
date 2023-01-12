package v1

import (
	"net/http"

	"github.com/Avtoelon/pkg/logger"
	"github.com/Avtoelon/pkg/structs"
	"github.com/Avtoelon/storage/postgres"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateOption ...
// @Summary CreateOptions
// @Description This API for creating option fot cars
// @Tags option
// @Accept json
// @Produce json
// @Param option body structs.CreateOption true "CreateOption"
// @Success 200 {object} structs.Option
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/option [post]
func (h *handlerV1) CreateOption(c *gin.Context) {
	var body structs.CreateOption
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusAccepted, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while blinding json", logger.Error(err))
		return
	}
	response, err := postgres.NewOptionsRepo(h.db).Create(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while creating option for cars", logger.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}

// UpdateOption ...
// @Summary UpdateOption
// @Description This API for creating option
// @Tags option
// @Accept json
// @Produce json
// @Param option body structs.UpdateOptionReq true "Update_Option"
// @Success 200 {object} structs.Option
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/option [put]
func (h *handlerV1) UpdateOption(c *gin.Context) {
	var body structs.UpdateOptionReq

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while binding json", logger.Error(err))
		return
	}

	response, err := postgres.NewOptionsRepo(h.db).Update(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while updating option", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// GetOptionById ...
// @Summary GetOptionById
// @Description This API for getting option by Id
// @Tags option
// @Accept json
// @Produce json
// @Param id path string true "Option_Id"
// @Success 200 {object} structs.Option
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/option/{id} [get]
func (h *handlerV1) GetOption(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")
	response, err := postgres.NewOptionsRepo(h.db).Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting option by id", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// GetAllOptions ...
// @Summary GetAllOptions
// @Description This API for getting all options for cars
// @Tags option
// @Accept json
// @Produce json
// @Success 200 {object} structs.Options
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/options [get]
func (h *handlerV1) GetAllOptions(c *gin.Context) {
	response, err := postgres.NewOptionsRepo(h.db).GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting all options", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// DeleteOption ...
// @Summary DeleteOption
// @Description This API for deleting options by id
// @Tags option
// @Accept json
// @Produce json
// @Param id path string true "Option_id"
// @Success 200 {object} structs.Option
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/option/{id} [delete]
func (h *handlerV1) DeleteOption(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")
	response, err := postgres.NewOptionsRepo(h.db).Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while deleting option by id", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}
