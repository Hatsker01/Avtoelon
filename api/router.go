package api

import (
	_ "github.com/Avtoelon/api/docs"
	v1 "github.com/Avtoelon/api/handler/v1"
	"github.com/Avtoelon/config"
	"github.com/Avtoelon/middleware"
	"github.com/Avtoelon/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Option struct {
	Db     *sqlx.DB
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
		Db:     option.Db,
		Logger: option.Logger,
		Cfg:    option.Conf,
	})
	api := router.Group("/v1")

	api.POST("/cars", handlerV1.CreateCar)
	api.PUT("/cars/updateCar", handlerV1.UpdateCar)
	api.GET("/cars/getById/:id", handlerV1.GetCar)
	api.GET("/cars/getAll", handlerV1.GetAllCars)
	api.DELETE("/cars/:id", handlerV1.DeleteCar)
	api.GET("/cars/getInfo/:id", handlerV1.GetCarInfo)

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

	//Category
	api.POST("/category", handlerV1.CreateCategory)
	api.PUT("/category", handlerV1.UpdateCategory)
	api.GET("/category/:id", handlerV1.GetCategory)
	api.GET("/categories", handlerV1.GetAllCategory)
	api.DELETE("/category/:id", handlerV1.DeleteCategory)

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

	//Color
	api.POST("/color", handlerV1.CreateColor)
	api.PUT("/color", handlerV1.UpdateColor)
	api.GET("/color/:id", handlerV1.GetColor)
	api.GET("/colors", handlerV1.GetAllColors)
	api.DELETE("/color/:id", handlerV1.DeleteColor)

	//Drive_Unit
	api.POST("/drive", handlerV1.CreateDrive)
	api.PUT("/drive", handlerV1.UpdateDriveUnit)
	api.GET("/drive/:id", handlerV1.GetDriveUnit)
	api.GET("/drives", handlerV1.GetAllDriveUnits)
	api.DELETE("/drive/:id", handlerV1.DeleteDriveUnit)

	//Optic
	api.POST("/optic", handlerV1.CreateOptic)
	api.PUT("/optic", handlerV1.UpdateOptic)
	api.GET("/optic/:id", handlerV1.GetOptic)
	api.GET("/optics", handlerV1.GetAllOptics)
	api.DELETE("/optic/:id", handlerV1.DeleteOptic)

	//Salon
	api.POST("/salon", handlerV1.CreateSalon)
	api.PUT("/salon", handlerV1.UpdateSalon)
	api.GET("/salon/:id", handlerV1.GetSalonById)
	api.GET("/salons", handlerV1.GetAllSalons)
	api.DELETE("/salon/:id", handlerV1.DeleteSalon)

	//Media
	api.POST("/media", handlerV1.CreateMedia)
	api.PUT("/media", handlerV1.UpdateMedia)
	api.GET("/media/:id", handlerV1.GetMedia)
	api.GET("/medias", handlerV1.GetAllMedias)
	api.DELETE("/media/:id", handlerV1.DeleteMedia)

	//Options
	api.POST("/option", handlerV1.CreateOption)
	api.PUT("/option", handlerV1.UpdateOption)
	api.GET("/option/:id", handlerV1.GetOption)
	api.GET("/options", handlerV1.GetAllOptions)
	api.DELETE("/option/:id", handlerV1.DeleteOption)

	//Additionally
	api.POST("/additionally", handlerV1.CreateAdd)
	api.PUT("/additionally", handlerV1.UpdateAdditionally)
	api.GET("/additionally/:id", handlerV1.GetAdditional)
	api.GET("/additionals", handlerV1.GetAllAdditional)
	api.DELETE("/additionally/:id", handlerV1.DeleteAdditional)

	//Region
	api.POST("/region", handlerV1.CreateRegion)
	api.PUT("/region", handlerV1.UpdateRegion)
	api.GET("/region/:id", handlerV1.GetRegion)
	api.GET("/regions", handlerV1.GetAllRegions)
	api.DELETE("/region/:id", handlerV1.DeleteRegion)

	//City
	api.POST("/city", handlerV1.CreateCity)
	api.PUT("/city", handlerV1.UpdateCity)
	api.GET("/city/:id", handlerV1.GetCity)
	api.GET("/cities", handlerV1.GetAllCities)
	api.DELETE("/city/:id", handlerV1.DeleteCity)

	//Marc
	api.POST("/marc", handlerV1.CreateMarc)
	api.PUT("/marc", handlerV1.UpdateMarc)
	api.GET("/marc/:id", handlerV1.GetMarc)
	api.GET("/marcs", handlerV1.GetAllMarc)
	api.DELETE("/marc/:id", handlerV1.DeleteMarc)

	url := ginSwagger.URL("swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router

}
