package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/Sami1309/go-grpc-server/middleware"

)

func Router() *gin.Engine {
	//Initialize gin router
	router := gin.Default()
	router.Use(cors.Default())

	//Retrieve list of all avaliable data types
	router.GET("/list", middleware.GetTypes)

	//Retrieve list of all of a specific data type (feature, label, training set, etc)
	router.GET("/:type", middleware.GetType)

	//Retrieve all versions of metadata for specific object
	router.GET("/:type/:name", middleware.GetObject)

	//Get metrics for service
	router.GET("/metrics", middleware.PrometheusHandler())

	return router

}
