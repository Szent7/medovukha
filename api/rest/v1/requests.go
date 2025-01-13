package v1

import (
	"fmt"
	"net/http"

	"medovukha/api/rest/v1/types"
	"medovukha/services/docker"

	"github.com/gin-gonic/gin"
)

func CreateTestContainer(c *gin.Context) {
	cli, err := docker.CreateDockerClient()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Docker client error"})
		fmt.Printf("Docker client error: %s\n", err.Error())
		return
	}
	defer cli.Close()

	if err := docker.CreateTestContainer(cli); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "CreateTestContainer error"})
		fmt.Printf("CreateTestContainer error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Created TestContainer"})
}

func GetContainerList(c *gin.Context) {
	cli, err := docker.CreateDockerClient()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Docker client error"})
		fmt.Printf("Docker client error: %s\n", err.Error())
		return
	}
	defer cli.Close()

	conList, err := docker.GetContainerBaseInfoList(cli)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "GetContainerList error"})
		fmt.Printf("GetContainerList error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, conList)
}

func GetImageList(c *gin.Context) {
	cli, err := docker.CreateDockerClient()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Docker client error"})
		fmt.Printf("Docker client error: %s\n", err.Error())
		return
	}
	defer cli.Close()

	imgList, err := docker.GetImageList(cli)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "GetImageList error"})
		fmt.Printf("GetImageList error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, imgList)
}

func GetVolumeList(c *gin.Context) {
	cli, err := docker.CreateDockerClient()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Docker client error"})
		fmt.Printf("Docker client error: %s\n", err.Error())
		return
	}
	defer cli.Close()

	imgList, err := docker.GetVolumeList(cli)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "GetVolumeList error"})
		fmt.Printf("GetVolumeList error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, imgList)
}

func GetNetworkList(c *gin.Context) {
	cli, err := docker.CreateDockerClient()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Docker client error"})
		fmt.Printf("Docker client error: %s\n", err.Error())
		return
	}
	defer cli.Close()

	imgList, err := docker.GetNetworkList(cli)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "GetNetworkList error"})
		fmt.Printf("GetNetworkList error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, imgList)
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
		fmt.Printf("Docker client error: %s\n", err.Error())
		return
	}
	defer cli.Close()

	if err := docker.PauseContainerByID(cli, containerId.Id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "PauseContainerByID error"})
		fmt.Printf("PauseContainerByID error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Paused: " + containerId.Id})
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
		fmt.Printf("Docker client error: %s\n", err.Error())
		return
	}
	defer cli.Close()

	if err := docker.UnpauseContainerByID(cli, containerId.Id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "UnpauseContainerByID error"})
		fmt.Printf("UnpauseContainerByID error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Unpaused: " + containerId.Id})
}

func StopContainerByID(c *gin.Context) {
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
		fmt.Printf("Docker client error: %s\n", err.Error())
		return
	}
	defer cli.Close()

	if err := docker.StopContainerByID(cli, containerId.Id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "StopContainerByID error"})
		fmt.Printf("StopContainerByID error: %s\n", err.Error())
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
		fmt.Printf("Docker client error: %s\n", err.Error())
		return
	}
	defer cli.Close()

	if err := docker.KillContainerByID(cli, containerId.Id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "KillContainerByID error"})
		fmt.Printf("KillContainerByID error: %s\n", err.Error())
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
		fmt.Printf("Docker client error: %s\n", err.Error())
		return
	}
	defer cli.Close()

	if err := docker.StartContainerByID(cli, containerId.Id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "StartContainerByID error"})
		fmt.Printf("StartContainerByID error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Started: " + containerId.Id})
}

func RestartContainerByID(c *gin.Context) {
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
		fmt.Printf("Docker client error: %s\n", err.Error())
		return
	}
	defer cli.Close()

	if err := docker.RestartContainerByID(cli, containerId.Id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "RestartContainerByID error"})
		fmt.Printf("RestartContainerByID error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Restarted: " + containerId.Id})
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
		fmt.Printf("Docker client error: %s\n", err.Error())
		return
	}
	defer cli.Close()

	if err := docker.RemoveContainerByID(cli, containerId.Id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "RemoveContainerByID error"})
		fmt.Printf("RemoveContainerByID error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Removed: " + containerId.Id})
}
