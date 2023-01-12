package v1

import (
	"net/http"

	"github.com/Avtoelon/pkg/logger"
	"github.com/Avtoelon/pkg/structs"
	"github.com/Avtoelon/storage/postgres"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateMedia ...
// @Summary CreateMedia
// @Description This API for creating media for car
// @Tags media
// @Accept json
// @Produce json
// @Param media body structs.CreateMedia true "CreateMedia"
// @Success 200 {object} structs.Media
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/media [post]
func (h *handlerV1) CreateMedia(c *gin.Context) {
	var body structs.CreateMedia

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while blinding json", logger.Error(err))
		return
	}

	response, err := postgres.NewMediasRepo(h.db).Create(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while creating media", logger.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}

// UpdatedMedia ...
// @Summary UpdateMedia
// @Description This API for updating media
// @Tags media
// @Accept json
// @Produce json
// @Param media body structs.UpdateMediaReq true "Update_Media"
// @Success 200 {object} structs.Media
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/media [put]
func (h *handlerV1) UpdateMedia(c *gin.Context) {
	var body structs.UpdateMediaReq

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while blinding json media", logger.Error(err))
		return
	}

	response, err := postgres.NewMediasRepo(h.db).Update(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while updating media", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// GetMedia...
// @Summary GetMedia
// @Description This API for getting media by id
// @Tags media
// @Accept json
// @Produce json
// @Param id path string true "Media_Id"
// @Success 200 {object} structs.Media
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/media/{id} [get]
func (h *handlerV1) GetMedia(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	response, err := postgres.NewMediasRepo(h.db).Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting media by id", logger.Error(err))
		return
	}

	c.JSON(http.StatusAccepted, response)
}

// GetAllMedia ...
// @Summary GetAllMedia
// @Description This API for getting all medias
// @Tags media
// @Accept json
// @Produce json
// @Success 200 {object} structs.Medias
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/medias [get]
func (h *handlerV1) GetAllMedias(c *gin.Context) {
	response, err := postgres.NewMediasRepo(h.db).GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting all medias", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// DeleteMedia ...
// @Summary DeleteMedia
// @Description This API for deleting media by id
// @Tags media
// @Accept json
// @Produce json
// @Param id path string true "Media_Id"
// @Success 200 {object} structs.Media
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/media/{id} [delete]
func (h *handlerV1) DeleteMedia(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	response, err := postgres.NewMediasRepo(h.db).Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while deleting media by id", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}
