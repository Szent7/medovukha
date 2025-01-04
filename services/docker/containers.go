package docker

import (
	"context"
	"fmt"
	"medovukha/api/rest/v1/types"
	"os"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
)

func GetContainerBaseInfoList(cli IDockerClient) ([]types.ContainerBaseInfo, error) {
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
		if check, err := CheckIsMedovukhaId(conList[i].Id); err != nil {
			return nil, err
		} else {
			conList[i].IsMedovukha = check
		}
	}

	return conList, nil
}

func CreateTestContainer(cli IDockerClient) error {
	ctx := context.Background()

	imageName := "docker/welcome-to-docker"

	if err := PullImage(cli, ctx, imageName); err != nil {
		return err
	}

	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			// original port
			"80/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "9990", // new port
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

func PauseContainerByID(cli IDockerClient, id string) error {
	ctx := context.Background()

	conList, err := GetContainerBaseInfoList(cli)
	if err != nil {
		return err
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
	return types.ErrContainerNotFound
}

func UnpauseContainerByID(cli IDockerClient, id string) error {
	ctx := context.Background()

	conList, err := GetContainerBaseInfoList(cli)
	if err != nil {
		return err
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
	return types.ErrContainerNotFound
}

func KillContainerByID(cli IDockerClient, id string) error {
	ctx := context.Background()

	conList, err := GetContainerBaseInfoList(cli)
	if err != nil {
		return err
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
	return types.ErrContainerNotFound
}

func StartContainerByID(cli IDockerClient, id string) error {
	ctx := context.Background()

	conList, err := GetContainerBaseInfoList(cli)
	if err != nil {
		return err
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
	return types.ErrContainerNotFound
}

func RemoveContainerByID(cli IDockerClient, id string) error {
	ctx := context.Background()

	conList, err := GetContainerBaseInfoList(cli)
	if err != nil {
		return err
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
	return types.ErrContainerNotFound
}

func CheckIsMedovukhaId(id string) (bool, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return false, err
	}
	//the hostname is part of the full container ID
	contains := strings.Contains(id, hostname)
	i := strings.Index(id, hostname)
	if contains && i == 0 {
		return true, nil
	}
	return false, nil
}
