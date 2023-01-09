package api

import (
	_ "github.com/Avtoelon/api/docs"
	v1 "github.com/Avtoelon/api/handler/v1"
	"github.com/Avtoelon/config"
	"github.com/Avtoelon/middleware"
	"github.com/Avtoelon/pkg/logger"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Option struct {
	Conf   config.Config
	Logger logger.Logger
}

func New(option Option) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(),
		gin.Recovery(),
		middleware.New(GinCorsMiddleware()),
	)
	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger: option.Logger,
		Cfg:    option.Conf,
	})
	api := router.Group("/v1")

	api.POST("/cars", handlerV1.CreateCar)
	api.PUT("/cars/updateCar", handlerV1.UpdateCar)
	api.GET("/cars/getById/:id", handlerV1.GetCar)
	api.GET("/cars/getAll", handlerV1.GetAllCars)
	api.DELETE("/cars/:id", handlerV1.DeleteCar)

	//Outside
	api.POST("/outside", handlerV1.CreateOutside)
	api.PUT("/outside", handlerV1.UpdateOutside)
	api.GET("/outside/:id", handlerV1.GetOutside)
	api.GET("/outside/getAll", handlerV1.GetAllOutside)
	api.DELETE("/outside/:id", handlerV1.DeletedOutside)

	//Model
	api.POST("/model", handlerV1.CreateModel)
	api.PUT("/model", handlerV1.UpdateModel)
	api.GET("/model/:id", handlerV1.GetModel)
	api.GET("/model/getAll", handlerV1.GetAllModels)
	api.DELETE("/model/:id", handlerV1.DeleteModel)

	//Body
	api.POST("/body", handlerV1.CreateBody)
	api.PUT("/body", handlerV1.UpdateBody)
	api.GET("/body/:id", handlerV1.GetBody)
	api.GET("/bodies", handlerV1.GetAllBody)
	api.DELETE("/body/:id", handlerV1.DeleteBody)

	//Oil
	api.POST("/oil", handlerV1.CreateOil)
	api.PUT("/oil", handlerV1.UpdateOil)
	api.GET("/oil/:id", handlerV1.GetOil)
	api.GET("/oils", handlerV1.GetAllOils)
	api.DELETE("/oil/:id", handlerV1.DeleteOil)

	//Transmission
	api.POST("/transmission", handlerV1.CreateTrans)
	api.PUT("/transmission", handlerV1.UpdateTrans)
	api.GET("/transmission/:id", handlerV1.GetTransmission)
	api.GET("/transmissions", handlerV1.GetAllTrans)
	api.DELETE("/transmission/:id", handlerV1.DeleteTransmission)

	url := ginSwagger.URL("swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router

}
