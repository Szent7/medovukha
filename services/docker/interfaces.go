package docker

import (
	"context"
	"io"

	ts "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/network"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
)

type IDockerClient interface {
	ContainerList(ctx context.Context, options container.ListOptions) ([]ts.Container, error)
	ContainerPause(ctx context.Context, containerID string) error

	ImagePull(ctx context.Context, refStr string, options image.PullOptions) (io.ReadCloser, error)

	ContainerCreate(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig,
		platform *ocispec.Platform, containerName string) (container.CreateResponse, error)

	ContainerStart(ctx context.Context, containerID string, options container.StartOptions) error

	ContainerUnpause(ctx context.Context, containerID string) error

	ContainerKill(ctx context.Context, containerID, signal string) error

	ContainerRemove(ctx context.Context, containerID string, options container.RemoveOptions) error
}
