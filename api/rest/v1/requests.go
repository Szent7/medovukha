package v1

import (
	"fmt"
	"net/http"
	"os"

	"medovukha/api/rest/v1/types"
	"medovukha/services/docker"

	containers "medovukha/services/docker/containers"
	images "medovukha/services/docker/images"
	networks "medovukha/services/docker/networks"
	volumes "medovukha/services/docker/volumes"
	git "medovukha/services/git"

	"github.com/gin-gonic/gin"
)

// Containers
func CreateTestContainer(c *gin.Context) {
	cli, err := docker.CreateDockerClient()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Docker client error"})
		fmt.Printf("Docker client error: %s\n", err.Error())
		return
	}
	defer cli.Close()

	if err := containers.CreateTestContainer(cli); err != nil {
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

	conList, err := containers.GetContainerBaseInfoList(cli)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "GetContainerList error"})
		fmt.Printf("GetContainerList error: %s\n", err.Error())
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

	if check, err := containers.CheckIsMedovukhaId(containerId.Id); err != nil {
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

	if err := containers.PauseContainerByID(cli, containerId.Id); err != nil {
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

	if check, err := containers.CheckIsMedovukhaId(containerId.Id); err != nil {
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

	if err := containers.UnpauseContainerByID(cli, containerId.Id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "UnpauseContainerByID error"})
		fmt.Printf("UnpauseContainerByID error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Unpaused: " + containerId.Id})
}

func KillContainerByID(c *gin.Context) {
	var containerId struct {
		Id string `json:"id"`
	}
	if err := c.BindJSON(&containerId); err != nil {
		return
	}

	if check, err := containers.CheckIsMedovukhaId(containerId.Id); err != nil {
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

	if err := containers.KillContainerByID(cli, containerId.Id); err != nil {
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

	if check, err := containers.CheckIsMedovukhaId(containerId.Id); err != nil {
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

	if err := containers.StartContainerByID(cli, containerId.Id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "StartContainerByID error"})
		fmt.Printf("StartContainerByID error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Started: " + containerId.Id})
}

func StopContainerByID(c *gin.Context) {
	var containerId struct {
		Id string `json:"id"`
	}
	if err := c.BindJSON(&containerId); err != nil {
		return
	}

	if check, err := containers.CheckIsMedovukhaId(containerId.Id); err != nil {
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

	if err := containers.StopContainerByID(cli, containerId.Id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "StopContainerByID error"})
		fmt.Printf("StopContainerByID error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Stopped: " + containerId.Id})
}

func RestartContainerByID(c *gin.Context) {
	var containerId struct {
		Id string `json:"id"`
	}
	if err := c.BindJSON(&containerId); err != nil {
		return
	}

	if check, err := containers.CheckIsMedovukhaId(containerId.Id); err != nil {
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

	if err := containers.RestartContainerByID(cli, containerId.Id); err != nil {
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

	if check, err := containers.CheckIsMedovukhaId(containerId.Id); err != nil {
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

	if err := containers.RemoveContainerByID(cli, containerId.Id); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "RemoveContainerByID error"})
		fmt.Printf("RemoveContainerByID error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Removed: " + containerId.Id})
}

//Images

func GetImageList(c *gin.Context) {
	cli, err := docker.CreateDockerClient()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Docker client error"})
		fmt.Printf("Docker client error: %s\n", err.Error())
		return
	}
	defer cli.Close()

	imgList, err := images.GetImageList(cli)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "GetImageList error"})
		fmt.Printf("GetImageList error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, imgList)
}

func BuildImageByRepo(c *gin.Context) {
	var repoUri struct {
		URI  string   `json:"Uri"`
		Tags []string `json:"Tags"`
	}
	if err := c.BindJSON(&repoUri); err != nil {
		return
	}

	cli, err := docker.CreateDockerClient()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Docker client error"})
		fmt.Printf("Docker client error: %s\n", err.Error())
		return
	}
	defer cli.Close()

	tempDir, err := os.MkdirTemp("", "docker-repo")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Created tempDir: " + tempDir)
	defer os.RemoveAll(tempDir)
	defer fmt.Println("Deleted tempDir: " + tempDir)

	if err := git.CloneRepo(&git.RepoCloner{}, repoUri.URI, tempDir); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "CloneRepo error"})
		fmt.Printf("CloneRepo error: %s\n", err.Error())
		return
	}

	imageListIdsOld, err := images.GetImageListIDs(cli)
	if err != nil {
		fmt.Println(err.Error())
	}
	newImageId, err := images.BuildImage(cli, &images.TarArchiver{}, tempDir, repoUri.Tags)
	if err != nil {
		fmt.Println(err.Error())
	}
	imageListIdsNew, err := images.GetImageListIDs(cli)
	if err != nil {
		fmt.Println(err.Error())
	}
	imageListIdsOld = append(imageListIdsOld, newImageId)
	for _, img := range imageListIdsOld {
		for i := len(imageListIdsNew) - 1; i >= 0; i-- {
			if imageListIdsNew[i] == img {
				imageListIdsNew = append(imageListIdsNew[:i], imageListIdsNew[i+1:]...)
			}
		}
	}
	undeletedImages := images.RemoveIntermediateImages(cli, imageListIdsNew)
	if len(undeletedImages) != 0 {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "some images were not removed"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "created: " + newImageId})
}

// Networks
func GetNetworkList(c *gin.Context) {
	cli, err := docker.CreateDockerClient()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Docker client error"})
		fmt.Printf("Docker client error: %s\n", err.Error())
		return
	}
	defer cli.Close()

	imgList, err := networks.GetNetworkList(cli)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "GetNetworkList error"})
		fmt.Printf("GetNetworkList error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, imgList)
}

// Volumes
func GetVolumeList(c *gin.Context) {
	cli, err := docker.CreateDockerClient()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Docker client error"})
		fmt.Printf("Docker client error: %s\n", err.Error())
		return
	}
	defer cli.Close()

	imgList, err := volumes.GetVolumeList(cli)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "GetVolumeList error"})
		fmt.Printf("GetVolumeList error: %s\n", err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, imgList)
}
