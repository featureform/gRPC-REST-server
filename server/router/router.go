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

	//Retrieve list of spaces in embeddingstore (rpc ListEntries)
	//router.GET("/spaces", middleware.GetListEntriesAll)

	//Retrieve all metadata from all spaces
	//router.GET("/spaces/*metadata", middleware.GetSpacesMetadata)

	//Retrieve information about individual space (rpc GetSpaceEntry)
	//router.GET("/spaces/:name/*metadata", middleware.GetSpaceMetadata)

	//router.GET("/spaces/:name", middleware.GetEmbeddings)

	//Get specific embedding (rpc Get)
	router.GET("/spaces/:name/:key", middleware.GetEmbedding)

	//Get nearest neighbors for embedding (rpc NearestNeighbor)
	router.GET("/spaces/:name/:key/*nn", middleware.GetNearestNeighbors)

	//Get metrics for service
	router.GET("/metrics", middleware.PrometheusHandler())

	return router

}
