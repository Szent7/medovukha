package main

import (
	"medovukha/api/rest/middlewares"
	restapi "medovukha/api/rest/v1"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Static("/_app/immutable/", "./build/_app/immutable/")
	router.NoRoute(func(c *gin.Context) {
		c.File("./build/index.html")
	})

	//! dev headers
	router.Use(middlewares.CORSMiddleware())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	rest := router.Group("/rest")
	{
		v1 := rest.Group("/v1")
		{
			v1.POST("/createtest", restapi.CreateTestContainer)
			v1.POST("/pausecontainerbyid", restapi.PauseContainerByID)
			v1.POST("/unpausecontainerbyid", restapi.UnpauseContainerByID)

			v1.POST("/killcontainerbyid", restapi.KillContainerByID)
			v1.POST("/startcontainerbyid", restapi.StartContainerByID)

			v1.POST("/removecontainerbyid", restapi.RemoveContainerByID)

			v1.POST("/stopcontainerbyid", restapi.StopContainerByID)

			v1.POST("/restartcontainerbyid", restapi.RestartContainerByID)

			v1.GET("/getvolumelist", restapi.GetVolumeList)
			v1.GET("/getnetworklist", restapi.GetNetworkList)
			v1.GET("/getcontainerlist", restapi.GetContainerList)
			v1.GET("/getimagelist", restapi.GetImageList)
		}
	}

	router.Run("0.0.0.0:10015")
}
