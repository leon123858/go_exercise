package main

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	controller "github.com/leon123858/committee-meeting-assistan/data-api/controller"
	docs "github.com/leon123858/committee-meeting-assistan/data-api/docs"
	utils "github.com/leon123858/committee-meeting-assistan/data-api/utils"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			CMA data api
//	@version		0.0.1
//	@description	This is the API server about data access for CMA system

//	@contact.name	system designer
//	@contact.email	a0970785699@gmail.com
//	@contact.url	https://github.com/leon123858/committee-meeting-assistan

//	@Schemes	http https
//	@Accept		json
//	@Produce	json

//	@securityDefinitions.apikey	Bearer
//	@in							header
//	@name						Authorization
//	@description				Type your api key
// println(c.Request.Header.Get("Authorization"))

func main() {
	// init
	if os.Getenv("GO_ENV") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	config := *(utils.GetConfig(gin.Mode()))
	utils.InitDB()
	utils.InitFirebase()
	// router
	router := gin.Default()
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          1 * time.Minute,
		Credentials:     false,
		ValidateHeaders: false,
	}))

	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = config.HOST
	// path: http://localhost:8080/swagger/index.html
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Use(utils.AuthUser)

	albums := router.Group("/api/albums")
	{
		albums.GET("/", controller.GetAlbums)
		albums.POST("/", controller.PostAlbums)
		albums.GET("/:id", controller.GetAlbumByID)
	}
	// note: should not add domain or will crash in docker container
	router.Run(":8080")
}
