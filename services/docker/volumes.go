package docker

import (
	"context"
	"medovukha/api/rest/v1/types"

	"github.com/docker/docker/api/types/volume"
)

func GetVolumeList(cli IDockerClient) ([]types.VolumeBaseInfo, error) {
	ctx := context.Background()

	volumes, err := cli.VolumeList(ctx, volume.ListOptions{})
	if err != nil {
		return nil, err
	}

	volList := make([]types.VolumeBaseInfo, len(volumes.Volumes))
	for i, volume := range volumes.Volumes {
		volList[i] = types.VolumeBaseInfo{
			Name:       volume.Name,
			Driver:     volume.Driver,
			Mountpoint: volume.Mountpoint,
			Created:    volume.CreatedAt,
		}
	}

	return volList, nil
}
