package v1

import (
	"net/http"

	"github.com/Avtoelon/pkg/logger"
	"github.com/Avtoelon/pkg/structs"
	"github.com/Avtoelon/storage/postgres"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateDriveUnit
// @Summary CreateDiriveUnit
// @Description This API for creating drive unit
// @Tags drive_unit
// @Accept json
// @Produce json
// @Param drive body structs.DriveUnitCreateReq true "Drive_unit"
// @Success 200 {object} structs.Drive_Unit
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/drive [post]
func (h *handlerV1) CreateDrive(c *gin.Context) {
	var body structs.DriveUnitCreateReq
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while creating drive_unit", logger.Error(err))
		return
	}

	response, err := postgres.NewDriveUnitRepasitory(h.db).Create(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while creating new drive_unit", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// UpdateDriveUnit ...
// @Summary UpdateDriveUnit
// @Description This API for updating drive_unit
// @Tags drive_unit
// @Accept json
// @Produce json
// @Param drive body structs.UpdateDriveUnit true "UpdateDriveUnit"
// @Success 200 {object} structs.Drive_Unit
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/drive [put]
func (h *handlerV1) UpdateDriveUnit(c *gin.Context) {
	var body structs.UpdateDriveUnit
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while updating drive unit", logger.Error(err))
		return
	}
	response, err := postgres.NewDriveUnitRepasitory(h.db).Update(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while updating drive_unit", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// GetDriveUnit ...
// @Summary GetDriveUnit
// @Description This API for getting drive_unit by id
// @Tags drive_unit
// @Accept json
// @Produce json
// @Param id path string true "Drive_Unit_Id"
// @Success 200 {object} structs.Drive_Unit
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/drive/{id} [get]
func (h *handlerV1) GetDriveUnit(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	response,err:=postgres.NewDriveUnitRepasitory(h.db).Get(id)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err.Error(),
		})
		h.log.Error("failed while getting drive_unit by id",logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted,response)
}

// GetAllDriveUnits ...
// @Summary GetAllDriveUnit
// @Description This API for getting all drive_units
// @Tags drive_unit
// @Accept json
// @Produce json
// @Success 200 {object} structs.DriveUnits
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/drives [get]
func (h *handlerV1) GetAllDriveUnits(c *gin.Context){
	response,err:=postgres.NewDriveUnitRepasitory(h.db).GetAll()
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err.Error(),
		})
		h.log.Error("failed while getting all drive_units",logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted,response)
}

// DeleteDriveUnit ...
// @Summary DeleteDriveUnit
// @Description This API for deleting drive_unit by id
// @Tags drive_unit
// @Accept json
// @Produce json
// @Param id path string true "Drive_Unit_Id"
// @Success 200 {object} structs.Drive_Unit
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/drive/{id} [delete]
func (h *handlerV1) DeleteDriveUnit(c *gin.Context){
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id :=c.Param("id")

	response,err:=postgres.NewDriveUnitRepasitory(h.db).Delete(id)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err.Error(),
		})
		h.log.Error("failed while deleting drive_unit",logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted,response)
}
