package v1

import (
	"net/http"

	"github.com/Avtoelon/pkg/logger"
	"github.com/Avtoelon/pkg/structs"
	"github.com/Avtoelon/storage/postgres"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// Create User ...
// @Summary Create user
// @Description This API for creating user
// @Tags user
// @Accept json
// @Produce json
// @Param user body structs.CreateUser true "Create_User"
// @Success 200 {object} structs.User
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/user [post]
func (h *handlerV1) CreateUser(c *gin.Context) {
	var body structs.CreateUser

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while binding json", logger.Error(err))
		return
	}

	response, err := postgres.NewUsersRepasitory(h.db).Create(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while creating new user", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// Get User By ID ...
// @Summary Get User By Id
// @Description This API for getting user by ID
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User_Id"
// @Success 200 {object} structs.User
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/user/{id} [get]
func (h *handlerV1) GetUser(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	response, err := postgres.NewUsersRepasitory(h.db).Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting user by id", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// Get All User ...
// @Summary Get All User
// @Description This API for getting all users
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} structs.Users
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/users [get]
func (h *handlerV1) GetAllUser(c *gin.Context) {
	response, err := postgres.NewUsersRepasitory(h.db).GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while getting all users", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}

// Delete User By Id ...
// @Summary Delete User By Id
// @Description This API for deleting user by ID
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User_Id"
// @Success 200 {object} structs.User
// @Failure 400 {object} structs.StandardErrorModel
// @Failure 500 {object} structs.StandardErrorModel
// @Router /v1/user/{id} [delete]
func (h *handlerV1) DeleteUser(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	response, err := postgres.NewUsersRepasitory(h.db).Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while deleting user by id", logger.Error(err))
		return
	}
	c.JSON(http.StatusAccepted, response)
}
