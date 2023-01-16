package v1

import (
	"net/http"

	"github.com/Avtoelon/pkg/logger"
	"github.com/Avtoelon/pkg/structs"
	"github.com/Avtoelon/storage/postgres"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateSalon ...
// @Summary CreateSalon
// @Description This API for creating salon
// @Tags salon
// @Accept json
// @Produce json
// @Param salon body structs.CreateSalon true "CreateSalon"
// @Success 200 {object} structs.Salon
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/salon [post]
func (h *handlerV1) CreateSalon(c *gin.Context) {
	var body structs.CreateSalon
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while blinding json", logger.Error(err))
		return
	}

	response, err := postgres.NewSalonsRepo(h.db).Create(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while creating salon", logger.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}

// UpdateSalon ...
// @Sumamry UpdateSalon
// @Description This API for updating salon for cars
// @Tags salon
// @Accept json
// @Produce json
// @Param salon body structs.UpdateSalonReq true "UpdateSalon"
// @Success 200 {object} structs.Salon
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/salon [put]
func (h *handlerV1) UpdateSalon(c *gin.Context) {
	var body structs.UpdateSalonReq

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while biling json", logger.Error(err))
		return
	}

	response, err := postgres.NewSalonsRepo(h.db).Update(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while Updating salon", logger.Error(err))
		return
	}

	c.JSON(http.StatusAccepted, response)
}

// GetSalonById
// @Summary GetSalonById
// @Description This API for getting salon by id
// @Tags salon
// @Produce json
// @Accept json
// @Param id path string true "Salon_Id"
// @Success 200 {object} structs.Salon
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/salon/{id} [get]
func (h *handlerV1) GetSalonById(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	response, err := postgres.NewSalonsRepo(h.db).Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting salon by id", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// GetAllSalons ...
// @Summary GetAllSalons
// @Description This API for getting all salons
// @Tags salon
// @Accept json
// @Produce json
// @Success 200 {object} structs.Salons
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/salons [get]
func (h *handlerV1) GetAllSalons(c *gin.Context) {
	response, err := postgres.NewSalonsRepo(h.db).GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting all salons", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// DeleteSalon ...
// @Summary DeleteSalon
// @Description This API for deleting salon by id
// @Tags salon
// @Accept json
// @Produce json
// @Param id path string true "Salon_id"
// @Success 200 {object} structs.Model
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/salon/{id} [delete]
func (h *handlerV1) DeleteSalon(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	response, err := postgres.NewSalonsRepo(h.db).Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while deleting salon by id", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}
