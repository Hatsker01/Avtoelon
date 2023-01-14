package v1

import (
	"net/http"

	"github.com/Avtoelon/pkg/logger"
	"github.com/Avtoelon/pkg/structs"
	"github.com/Avtoelon/storage/postgres"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateCategory ...
// @Summary CreateCategory
// @Description This API for creating category
// @Tags category
// @Accept json
// @Produce json
// @Param category body  structs.CategoryCreateReq true "CreateCategoryReq"
// @Success 200 {object} structs.Category
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/category [post]
func (h *handlerV1) CreateCategory(c *gin.Context) {
	var body structs.CategoryCreateReq
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while creating category", logger.Error(err))
		return
	}
	response, err := postgres.NewCategoryRepasitory(h.db).CreateCategory(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while creating category", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// UpdateCategory ...
// @Summary UpdateCategory
// @Description This API for updating category
// @Tags category
// @Accept json
// @Produce json
// @Param category body  structs.UpdateCategory true "UpdateCategory"
// @Success 200 {object} structs.Category
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/category [put]
func (h *handlerV1) UpdateCategory(c *gin.Context) {
	var body structs.Category

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while blinding json", logger.Error(err))
		return
	}

	response, err := postgres.NewCategoryRepasitory(h.db).UpdateCategory(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while updating category", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// GetCategory ...
// @Summary GetCategory
// @Description This API for getting category by id
// @Tags category
// @Accept json
// @Produce json
// @Param id path string true "Category_id"
// @Success 200 {object} structs.Category
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/category/{id} [get]
func (h *handlerV1) GetCategory(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id:=c.Param("id")

	response,err:=postgres.NewCategoryRepasitory(h.db).GetCategory(id)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err.Error(),
		})
		h.log.Error("failed while getting category by id",logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted,response)
}

// GetAllCategory ...
// @Summary GetAllCategory
// @Description This API for getting all categories
// @Tags category
// @Accept json
// @Produce json
// @Success 200 {object} structs.Categories
// @Failuer 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/categories [get]
func (h *handlerV1) GetAllCategory(c *gin.Context){
	response,err:=postgres.NewCategoryRepasitory(h.db).GetAllCategory()
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err.Error(),
		})
		h.log.Error("failed while hetting all categories",logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted,response)
}

// DeleteCategory ...
// @Summary DeleteCategory
// @Description This API for deleting category by id
// @Tags category
// @Accept json
// @Produce json
// @Param id path string true "Category_Id"
// @Success 200 {object} structs.Category
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/category/{id} [delete]
func (h *handlerV1) DeleteCategory(c *gin.Context){
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id:=c.Param("id")

	response,err:=postgres.NewCategoryRepasitory(h.db).DeleteCategory(id)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err.Error(),
		})
		h.log.Error("failed while deleting category",logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted,response)
}


