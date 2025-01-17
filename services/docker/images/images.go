package images

import (
	"context"
	"fmt"
	"io"
	"log"
	ts "medovukha/api/rest/v1/types"
	dc "medovukha/services/docker"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/pkg/archive"
)

func PullImage(cli dc.IDockerClient, ctx context.Context, imageName string) error {
	out, err := cli.ImagePull(ctx, imageName, image.PullOptions{})
	if err != nil {
		return err
	}
	defer out.Close()
	io.Copy(os.Stdout, out)
	return nil
}

func GetImageList(cli dc.IDockerClient) ([]ts.ImageBaseInfo, error) {
	images, err := GetRawImageList(cli)
	if err != nil {
		return nil, err
	}

	imgList := make([]ts.ImageBaseInfo, len(images))
	for i, image := range images {
		imgList[i] = ts.ImageBaseInfo{
			Id:      image.ID,
			Tags:    image.RepoTags,
			Size:    image.Size,
			Created: image.Created,
		}
	}

	return imgList, nil
}

func GetRawImageList(cli dc.IDockerClient) ([]image.Summary, error) {
	ctx := context.Background()

	images, err := cli.ImageList(ctx, image.ListOptions{All: true, ContainerCount: true})
	if err != nil {
		return nil, err
	}

	return images, nil
}

func GetImageIDFromTag(cli dc.IDockerClient, tag string) (string, error) {
	images, err := GetRawImageList(cli)
	imageID := ""
	if err != nil {
		return imageID, err
	}
	for _, image := range images {
		if len(image.RepoTags) != 0 {
			for _, repoTag := range image.RepoTags {
				if tag == repoTag {
					imageID = image.ID
					return imageID, nil
				}
			}
		}
	}
	return imageID, nil
}

func GetImageListIDs(cli dc.IDockerClient) ([]string, error) {
	images, err := GetRawImageList(cli)
	if err != nil {
		return nil, err
	}
	imageListIds := make([]string, len(images))
	for i, image := range images {
		imageListIds[i] = image.ID
	}
	return imageListIds, nil
}

// ? Perhaps archiving should be made into a separate function (ITarArchiver)
func BuildImage(cli dc.IDockerClient, tarArchiver ITarArchiver, path string, tags []string) (string, error) {
	ctx := context.Background()

	if len(tags) < 1 {
		return "", ts.ErrEmptyTags
	}

	tar, err := tarArchiver.TarWithOptions(path, &archive.TarOptions{})
	if err != nil {
		return "", err
	}
	defer tar.Close()

	out, err := cli.ImageBuild(ctx, tar, types.ImageBuildOptions{
		Dockerfile: "Dockerfile",
		Tags:       tags,
		Remove:     true,
	})
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	defer out.Body.Close()
	io.Copy(os.Stdout, out.Body)

	return GetImageIDFromTag(cli, tags[0])
}

// RemoveIntermediateImages get a list of id images to be deleted
// returns a list of undeleted images
// if all images have been deleted, the list will be empty
func RemoveIntermediateImages(client dc.IDockerClient, IDs []string) []string {
	opt := image.RemoveOptions{Force: true}
	var err error = nil
	undeletedIDs := make([]string, 0, len(IDs))
	for _, image := range IDs {
		_, err = client.ImageRemove(context.Background(), image, opt)
		if err != nil {
			undeletedIDs = append(undeletedIDs, image)
			log.Printf("failed to remove image %s: %v", image, err)
		} else {
			log.Printf("removed intermediate image %s", image)
		}
	}
	return undeletedIDs
}
