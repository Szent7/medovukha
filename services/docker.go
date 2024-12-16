package services

import (
	"context"
	"errors"
	"fmt"
	"io"
	"medovukha/api/rest/v1/types"
	"os"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func CreateDockerClient() (*client.Client, error) {
	return client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
}

func GetContainerList(cli *client.Client) ([]types.ContainerBaseInfo, error) {
	ctx := context.Background()

	containers, err := cli.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		return nil, err
	}

	conList := make([]types.ContainerBaseInfo, len(containers))
	for i, container := range containers {
		conList[i] = types.ContainerBaseInfo{
			Id:        container.ID,
			Names:     container.Names,
			ImageName: container.Image,
			Ports:     container.Ports,
			Created:   container.Created,
			State:     container.State,
		}
	}

	return conList, nil
}

func CreateTestContainer(cli *client.Client) error {
	ctx := context.Background()

	imageName := "docker/welcome-to-docker"

	out, err := cli.ImagePull(ctx, imageName, image.PullOptions{})
	if err != nil {
		return err
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
		return err
	}

	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		return err
	}

	fmt.Println(resp.ID)
	return nil
}

func PauseContainerByID(cli *client.Client, id string) error {
	ctx := context.Background()

	containers, err := cli.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		return err
	}
	conList := make([]types.ContainerBaseInfo, len(containers))
	for i, container := range containers {
		conList[i] = types.ContainerBaseInfo{
			Id:        container.ID,
			Names:     container.Names,
			ImageName: container.Image,
			Ports:     container.Ports,
			Created:   container.Created,
			State:     container.State,
		}
	}

	for _, container := range conList {
		if container.Id == id {
			if err := cli.ContainerPause(ctx, container.Id); err != nil {
				return err
			}
			fmt.Println("Paused: ", container.Id)
			return nil
		}
	}
	fmt.Println("Not found: ", id)
	return errors.New("container not found")
}

func UnpauseContainerByID(cli *client.Client, id string) error {
	ctx := context.Background()

	containers, err := cli.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		return err
	}

	conList := make([]types.ContainerBaseInfo, len(containers))
	for i, container := range containers {
		conList[i] = types.ContainerBaseInfo{
			Id:        container.ID,
			Names:     container.Names,
			ImageName: container.Image,
			Ports:     container.Ports,
			Created:   container.Created,
			State:     container.State,
		}
	}

	for _, container := range conList {
		if container.Id == id {
			if err := cli.ContainerUnpause(ctx, container.Id); err != nil {
				return err
			}
			fmt.Println("Unpaused: ", container.Id)
			return nil
		}
	}
	fmt.Println("Not found: ", id)
	return errors.New("container not found")
}

func KillContainerByID(cli *client.Client, id string) error {
	ctx := context.Background()

	containers, err := cli.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		return err
	}

	conList := make([]types.ContainerBaseInfo, len(containers))
	for i, container := range containers {
		conList[i] = types.ContainerBaseInfo{
			Id:        container.ID,
			Names:     container.Names,
			ImageName: container.Image,
			Ports:     container.Ports,
			Created:   container.Created,
			State:     container.State,
		}
	}

	for _, container := range conList {
		if container.Id == id {
			if err := cli.ContainerKill(ctx, container.Id, ""); err != nil {
				return err
			}
			fmt.Println("Killed: ", container.Id)
			return nil
		}
	}
	fmt.Println("Not found: ", id)
	return errors.New("container not found")
}

func StartContainerByID(cli *client.Client, id string) error {
	ctx := context.Background()

	containers, err := cli.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		return err
	}
	conList := make([]types.ContainerBaseInfo, len(containers))
	for i, container := range containers {
		conList[i] = types.ContainerBaseInfo{
			Id:        container.ID,
			Names:     container.Names,
			ImageName: container.Image,
			Ports:     container.Ports,
			Created:   container.Created,
			State:     container.State,
		}
	}

	for _, con := range conList {
		if con.Id == id {
			if err := cli.ContainerStart(ctx, con.Id, container.StartOptions{}); err != nil {
				return err
			}
			fmt.Println("Started: ", con.Id)
			return nil
		}
	}
	fmt.Println("Not found: ", id)
	return errors.New("container not found")
}

func RemoveContainerByID(cli *client.Client, id string) error {
	ctx := context.Background()

	containers, err := cli.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		return err
	}
	conList := make([]types.ContainerBaseInfo, len(containers))
	for i, container := range containers {
		conList[i] = types.ContainerBaseInfo{
			Id:        container.ID,
			Names:     container.Names,
			ImageName: container.Image,
			Ports:     container.Ports,
			Created:   container.Created,
			State:     container.State,
		}
	}

	for _, con := range conList {
		if con.Id == id {
			if err := cli.ContainerRemove(ctx, con.Id, container.RemoveOptions{
				RemoveVolumes: true,
				RemoveLinks:   false,
				Force:         false,
			}); err != nil {
				return err
			}
			fmt.Println("Started: ", con.Id)
			return nil
		}
	}
	fmt.Println("Not found: ", id)
	return errors.New("container not found")
}

/*
func CopyPorts(ports *[]types.Port) [] types.Ports {
	newPort := make([] types.Ports, len(*ports))
	for _, port := range *ports {
		newPort = append(newPort,  types.Ports{
			Ip:          port.IP,
			PrivatePort: port.PrivatePort,
			PublicPort:  port.PublicPort,
			Type:        port.Type,
		})
	}
	return newPort
}
*/
