package docker

import (
	"context"
	"io"
	"medovukha/api/rest/v1/types"
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

func GetImageList(cli IDockerClient) ([]types.ImageBaseInfo, error) {
	ctx := context.Background()

	images, err := cli.ImageList(ctx, image.ListOptions{All: true, ContainerCount: true})
	if err != nil {
		return nil, err
	}

	imgList := make([]types.ImageBaseInfo, len(images))
	for i, image := range images {
		imgList[i] = types.ImageBaseInfo{
			Id:      image.ID,
			Tags:    image.RepoTags,
			Size:    image.Size,
			Created: image.Created,
		}
	}

	return imgList, nil
}
