package main

import (
	restapi "medovukha/api/rest/v1"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Static("/assets", "./dist/assets")
	router.NoRoute(func(c *gin.Context) {
		c.File("./dist/index.html")
	})

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

			v1.GET("/getcontainerlist", restapi.GetContainerList)
		}
	}

	router.Run("0.0.0.0:10015")
}
