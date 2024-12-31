package docker

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types/image"
)

func PullImage(cli IDockerClient, ctx context.Context, imageName string) error {
	out, err := cli.ImagePull(ctx, imageName, image.PullOptions{})
	if err != nil {
		return err
	}
	defer out.Close()
	io.Copy(os.Stdout, out)
	return nil
}
