package v1

import (
	"net/http"

	"medovukha/services"

	"github.com/gin-gonic/gin"
)

func CreateTestContainer(c *gin.Context) {
	cli, err := services.CreateDockerClient()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Docker client error"})
		return
	}
	defer cli.Close()

	if err := services.CreateTestContainer(cli); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "CreateTestContainer error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Created TestContainer"})
}

func GetContainerList(c *gin.Context) {
	cli, err := services.CreateDockerClient()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Docker client error"})
		return
	}
	defer cli.Close()

	conList, err := services.GetContainerList(cli)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "GetContainerList error"})
		return
	}

	c.IndentedJSON(http.StatusOK, conList)
}

func PauseContainerByID(c *gin.Context) {
	cli, err := services.CreateDockerClient()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Docker client error"})
		return
	}
	defer cli.Close()

	var containerId struct {
		Id string `json:"id"`
	}
	if err := c.BindJSON(&containerId); err != nil {
		return
	}

	if err := services.PauseContainerByID(cli, containerId.Id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "PauseContainerByID error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Stopped: " + containerId.Id})
}

func UnpauseContainerByID(c *gin.Context) {
	cli, err := services.CreateDockerClient()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Docker client error"})
		return
	}
	defer cli.Close()

	var containerId struct {
		Id string `json:"id"`
	}
	if err := c.BindJSON(&containerId); err != nil {
		return
	}

	if err := services.UnpauseContainerByID(cli, containerId.Id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "PauseContainerByID error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Stopped: " + containerId.Id})
}

func KillContainerByID(c *gin.Context) {
	cli, err := services.CreateDockerClient()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Docker client error"})
		return
	}
	defer cli.Close()

	var containerId struct {
		Id string `json:"id"`
	}
	if err := c.BindJSON(&containerId); err != nil {
		return
	}

	if err := services.KillContainerByID(cli, containerId.Id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "KillContainerByID error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Killed: " + containerId.Id})
}

func StartContainerByID(c *gin.Context) {
	cli, err := services.CreateDockerClient()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Docker client error"})
		return
	}
	defer cli.Close()

	var containerId struct {
		Id string `json:"id"`
	}
	if err := c.BindJSON(&containerId); err != nil {
		return
	}

	if err := services.StartContainerByID(cli, containerId.Id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "StartContainerByID error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "started: " + containerId.Id})
}

func RemoveContainerByID(c *gin.Context) {
	cli, err := services.CreateDockerClient()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Docker client error"})
		return
	}
	defer cli.Close()

	var containerId struct {
		Id string `json:"id"`
	}
	if err := c.BindJSON(&containerId); err != nil {
		return
	}

	if err := services.RemoveContainerByID(cli, containerId.Id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "RemoveContainerByID error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "removed: " + containerId.Id})
}
