package v1

import (
	"net/http"

	"github.com/Avtoelon/pkg/logger"
	"github.com/Avtoelon/pkg/structs"
	"github.com/Avtoelon/storage/postgres"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// Create Marc for cars ...
// @Summary Create Marc for cars
// @Descrption This API for creating marc for cars
// @Tags marc
// @Accept json
// @Produce json
// @Param marc body structs.CreateMarc true "CreateMarc"
// @Success 200 {object} structs.Marc  
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/marc [post]
func (h *handlerV1) CreateMarc(c *gin.Context) {
	var body structs.CreateMarc

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while binding json", logger.Error(err))
		return
	}

	response, err := postgres.NewMarcsRepo(h.db).Create(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while creting marc for cars", logger.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}

// Update Marc ...
// @Summary Update Marc
// @Description This Api for updating marc
// @Tags marc
// @Accept json
// @Produce json
// @Param marc body structs.UpdateMarcReq true "UpdateMarc"
// @Success 200 {object} structs.Marc
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/marc [put]
func (h *handlerV1) UpdateMarc(c *gin.Context) {
	var body structs.UpdateMarcReq

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while updating marc ", logger.Error(err))
		return
	}

	response, err := postgres.NewMarcsRepo(h.db).Update(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while updating marc ", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// Get Marc By Id ...
// @Summary Get Marc By Id
// @Description This API for getting marc by ID
// @Tags marc
// @Accept json
// @Produce json
// @Param id path string true "Marc_Id"
// @Success 200 {object} structs.Marc
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/marc/{id} [get]
func (h *handlerV1) GetMarc(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	response, err := postgres.NewMarcsRepo(h.db).Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting marc by id", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// Get All Marc ...
// @Summary Get All Marc
// @Description This API for getting all marcs
// @Tags marc
// @Accept json
// @Produce json
// @Success 200 {object} structs.Marcs
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/marcs [get]
func (h *handlerV1) GetAllMarc(c *gin.Context) {
	response, err := postgres.NewMarcsRepo(h.db).GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting all marcs", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// Delete Marc By ID ...
// @Summary Delete Marc By ID
// @Description This API for deleting marc by ID
// @Tags marc
// @Accept json
// @Produce json
// @Param id path string true "Marc_Id"
// @Success 200 {object} structs.Marc
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/marc/{id} [delete]
func (h *handlerV1) DeleteMarc(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	response, err := postgres.NewMarcsRepo(h.db).Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while deleting marc by id", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// Get Marc Models ...
// @Summary Get Marc Model
// @Description This API for getting marc models
// @Tags marc 
// @Accept json
// @Produce json
// @Param id path string true "Marc_Id"
// @Success 200 {object} structs.GetMarcModels
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/marcModel/{id} [get]
func (h *handlerV1) GetMarcModel(c *gin.Context){
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	response,err:=postgres.NewMarcsRepo(h.db).GetMarcModels(id)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err.Error(),
		})
		h.log.Error("failed while getting marc models",logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted,response)
}