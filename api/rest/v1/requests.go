package v1

import (
	"net/http"

	"medovukha/api/rest/v1/types"
	"medovukha/services/docker"

	"github.com/gin-gonic/gin"
)

func CreateTestContainer(c *gin.Context) {
	cli, err := docker.CreateDockerClient()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Docker client error"})
		return
	}
	defer cli.Close()

	if err := docker.CreateTestContainer(cli); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "CreateTestContainer error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Created TestContainer"})
}

func GetContainerList(c *gin.Context) {
	cli, err := docker.CreateDockerClient()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Docker client error"})
		return
	}
	defer cli.Close()

	conList, err := docker.GetContainerBaseInfoList(cli)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "GetContainerList error"})
		return
	}

	c.IndentedJSON(http.StatusOK, conList)
}

func PauseContainerByID(c *gin.Context) {
	var containerId struct {
		Id string `json:"id"`
	}
	if err := c.BindJSON(&containerId); err != nil {
		return
	}

	if check, err := docker.CheckIsMedovukhaId(containerId.Id); err != nil {
		return
	} else if check {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": types.ErrContainerIsMedovukha.Error()})
		return
	}

	cli, err := docker.CreateDockerClient()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Docker client error"})
		return
	}
	defer cli.Close()

	if err := docker.PauseContainerByID(cli, containerId.Id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "PauseContainerByID error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Stopped: " + containerId.Id})
}

func UnpauseContainerByID(c *gin.Context) {
	var containerId struct {
		Id string `json:"id"`
	}
	if err := c.BindJSON(&containerId); err != nil {
		return
	}

	if check, err := docker.CheckIsMedovukhaId(containerId.Id); err != nil {
		return
	} else if check {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": types.ErrContainerIsMedovukha.Error()})
		return
	}

	cli, err := docker.CreateDockerClient()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Docker client error"})
		return
	}
	defer cli.Close()

	if err := docker.UnpauseContainerByID(cli, containerId.Id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "PauseContainerByID error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Stopped: " + containerId.Id})
}

func KillContainerByID(c *gin.Context) {
	var containerId struct {
		Id string `json:"id"`
	}
	if err := c.BindJSON(&containerId); err != nil {
		return
	}

	if check, err := docker.CheckIsMedovukhaId(containerId.Id); err != nil {
		return
	} else if check {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": types.ErrContainerIsMedovukha.Error()})
		return
	}

	cli, err := docker.CreateDockerClient()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Docker client error"})
		return
	}
	defer cli.Close()

	if err := docker.KillContainerByID(cli, containerId.Id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "KillContainerByID error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Killed: " + containerId.Id})
}

func StartContainerByID(c *gin.Context) {
	var containerId struct {
		Id string `json:"id"`
	}
	if err := c.BindJSON(&containerId); err != nil {
		return
	}

	if check, err := docker.CheckIsMedovukhaId(containerId.Id); err != nil {
		return
	} else if check {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": types.ErrContainerIsMedovukha.Error()})
		return
	}

	cli, err := docker.CreateDockerClient()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Docker client error"})
		return
	}
	defer cli.Close()

	if err := docker.StartContainerByID(cli, containerId.Id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "StartContainerByID error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "started: " + containerId.Id})
}

func RemoveContainerByID(c *gin.Context) {
	var containerId struct {
		Id string `json:"id"`
	}
	if err := c.BindJSON(&containerId); err != nil {
		return
	}

	if check, err := docker.CheckIsMedovukhaId(containerId.Id); err != nil {
		return
	} else if check {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": types.ErrContainerIsMedovukha.Error()})
		return
	}

	cli, err := docker.CreateDockerClient()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Docker client error"})
		return
	}
	defer cli.Close()

	if err := docker.RemoveContainerByID(cli, containerId.Id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "RemoveContainerByID error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "removed: " + containerId.Id})
}
