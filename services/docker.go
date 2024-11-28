package services

import (
	"context"
	"fmt"
	"io"
	v1 "medovukha/api/rest/v1"
	"net/http"
	"os"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/gin-gonic/gin"
)

func GetContainerList(c *gin.Context) {
	ctx := context.Background()
	cli := CreateDockerClient()
	defer cli.Close()

	containers, err := cli.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		panic(err)
	}
	conList := make([]v1.ContainerBaseInfo, len(containers))
	for i, container := range containers {
		conList[i] = v1.ContainerBaseInfo{
			Id:        container.ID,
			Names:     container.Names,
			ImageName: container.Image,
			Ports:     container.Ports,
			Created:   container.Created,
			State:     container.State,
		}
	}
	c.IndentedJSON(http.StatusOK, conList)
}

func CreateTestContainer(c *gin.Context) {
	ctx := context.Background()
	cli := CreateDockerClient()
	defer cli.Close()

	imageName := "docker/welcome-to-docker"

	out, err := cli.ImagePull(ctx, imageName, image.PullOptions{})
	if err != nil {
		panic(err)
	}
	defer out.Close()
	io.Copy(os.Stdout, out)

	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			// 80 - изначальный порт welcome-to-docker
			"80/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "9990", // новый порт
				},
			},
		},
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
		Tty:   false,
	}, hostConfig, nil, nil, "web-test")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		panic(err)
	}

	fmt.Println(resp.ID)
}

func CreateDockerClient() *client.Client {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	return cli
}

func PauseContainerByID(c *gin.Context) {
	ctx := context.Background()
	cli := CreateDockerClient()
	defer cli.Close()

	containers, err := cli.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		panic(err)
	}
	conList := make([]v1.ContainerBaseInfo, len(containers))
	for i, container := range containers {
		conList[i] = v1.ContainerBaseInfo{
			Id:        container.ID,
			Names:     container.Names,
			ImageName: container.Image,
			Ports:     container.Ports,
			Created:   container.Created,
			State:     container.State,
		}
	}

	var containerId struct {
		Id string `json:"id"`
	}
	if err := c.BindJSON(&containerId); err != nil {
		return
	}

	for _, container := range conList {
		if container.Id == containerId.Id {
			if err := cli.ContainerPause(ctx, container.Id); err != nil {
				panic(err)
			}
			fmt.Println("Paused: ", container.Id)
			return
		}
	}
	fmt.Println("Not found: ", containerId.Id)
}

func UnpauseContainerByID(c *gin.Context) {
	ctx := context.Background()
	cli := CreateDockerClient()
	defer cli.Close()

	containers, err := cli.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		panic(err)
	}
	conList := make([]v1.ContainerBaseInfo, len(containers))
	for i, container := range containers {
		conList[i] = v1.ContainerBaseInfo{
			Id:        container.ID,
			Names:     container.Names,
			ImageName: container.Image,
			Ports:     container.Ports,
			Created:   container.Created,
			State:     container.State,
		}
	}

	var containerId struct {
		Id string `json:"id"`
	}
	if err := c.BindJSON(&containerId); err != nil {
		return
	}

	for _, container := range conList {
		if container.Id == containerId.Id {
			if err := cli.ContainerUnpause(ctx, container.Id); err != nil {
				panic(err)
			}
			fmt.Println("Unpaused: ", container.Id)
			return
		}
	}
	fmt.Println("Not found: ", containerId.Id)
}

func KillContainerByID(c *gin.Context) {
	ctx := context.Background()
	cli := CreateDockerClient()
	defer cli.Close()

	containers, err := cli.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		panic(err)
	}
	conList := make([]v1.ContainerBaseInfo, len(containers))
	for i, container := range containers {
		conList[i] = v1.ContainerBaseInfo{
			Id:        container.ID,
			Names:     container.Names,
			ImageName: container.Image,
			Ports:     container.Ports,
			Created:   container.Created,
			State:     container.State,
		}
	}

	var containerId struct {
		Id string `json:"id"`
	}
	if err := c.BindJSON(&containerId); err != nil {
		return
	}

	for _, container := range conList {
		if container.Id == containerId.Id {
			if err := cli.ContainerKill(ctx, container.Id, ""); err != nil {
				panic(err)
			}
			fmt.Println("Killed: ", container.Id)
			return
		}
	}
	fmt.Println("Not found: ", containerId.Id)
}

func StartContainerByID(c *gin.Context) {
	ctx := context.Background()
	cli := CreateDockerClient()
	defer cli.Close()

	containers, err := cli.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		panic(err)
	}
	conList := make([]v1.ContainerBaseInfo, len(containers))
	for i, container := range containers {
		conList[i] = v1.ContainerBaseInfo{
			Id:        container.ID,
			Names:     container.Names,
			ImageName: container.Image,
			Ports:     container.Ports,
			Created:   container.Created,
			State:     container.State,
		}
	}

	var containerId struct {
		Id string `json:"id"`
	}
	if err := c.BindJSON(&containerId); err != nil {
		return
	}

	for _, con := range conList {
		if con.Id == containerId.Id {
			if err := cli.ContainerStart(ctx, con.Id, container.StartOptions{}); err != nil {
				panic(err)
			}
			fmt.Println("Started: ", con.Id)
			return
		}
	}
	fmt.Println("Not found: ", containerId.Id)
}

func RemoveContainerByID(c *gin.Context) {
	ctx := context.Background()
	cli := CreateDockerClient()
	defer cli.Close()

	containers, err := cli.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		panic(err)
	}
	conList := make([]v1.ContainerBaseInfo, len(containers))
	for i, container := range containers {
		conList[i] = v1.ContainerBaseInfo{
			Id:        container.ID,
			Names:     container.Names,
			ImageName: container.Image,
			Ports:     container.Ports,
			Created:   container.Created,
			State:     container.State,
		}
	}

	var containerId struct {
		Id string `json:"id"`
	}
	if err := c.BindJSON(&containerId); err != nil {
		return
	}

	for _, con := range conList {
		if con.Id == containerId.Id {
			if err := cli.ContainerRemove(ctx, con.Id, container.RemoveOptions{
				RemoveVolumes: true,
				RemoveLinks:   false,
				Force:         false,
			}); err != nil {
				panic(err)
			}
			fmt.Println("Started: ", con.Id)
			return
		}
	}
	fmt.Println("Not found: ", containerId.Id)
}

/*
func CopyPorts(ports *[]types.Port) []v1.Ports {
	newPort := make([]v1.Ports, len(*ports))
	for _, port := range *ports {
		newPort = append(newPort, v1.Ports{
			Ip:          port.IP,
			PrivatePort: port.PrivatePort,
			PublicPort:  port.PublicPort,
			Type:        port.Type,
		})
	}
	return newPort
}
*/
