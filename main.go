package main

import (
	"medovukha/services"

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
			v1.POST("/createtest", services.CreateTestContainer)
			v1.POST("/pausecontainerbyid", services.PauseContainerByID)
			v1.POST("/unpausecontainerbyid", services.UnpauseContainerByID)

			v1.POST("/killcontainerbyid", services.KillContainerByID)
			v1.POST("/startcontainerbyid", services.StartContainerByID)

			v1.POST("/removecontainerbyid", services.RemoveContainerByID)

			v1.GET("/getcontainerlist", services.GetContainerList)
		}
	}

	router.Run("0.0.0.0:10015")
}
