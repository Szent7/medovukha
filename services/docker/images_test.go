package docker

import (
	"context"
	"errors"
	"io"
	ts "medovukha/api/rest/v1/types"
	"strings"
	"testing"

	"github.com/docker/docker/api/types/image"
	"github.com/stretchr/testify/assert"
)

func TestPullImage(t *testing.T) {
	mockClient := new(MockDockerClient)

	//test: image found and pulled
	mockClient.On("ImagePull", context.Background(), "docker/welcome-to-docker", image.PullOptions{}).Return(io.NopCloser(strings.NewReader("creating web-test")), nil).Once()

	err := PullImage(mockClient, context.Background(), "docker/welcome-to-docker")
	assert.Nil(t, err)
	mockClient.AssertExpectations(t)

	//test: image not pulled
	mockClient.On("ImagePull", context.Background(), "docker/welcome-to-docker", image.PullOptions{}).Return(io.NopCloser(strings.NewReader("creating web-test")), errors.New("ImagePull error")).Once()

	err = PullImage(mockClient, context.Background(), "docker/welcome-to-docker")
	assert.EqualError(t, err, "ImagePull error")
	mockClient.AssertExpectations(t)
}

func TestGetImageList(t *testing.T) {
	mockClient := new(MockDockerClient)

	// test: image found
	mockClient.On("ImageList", context.Background(), image.ListOptions{All: true, ContainerCount: true}).Return([]image.Summary{
		{ID: "1234567890ab"},
	}, nil).Once()

	result, err := GetImageList(mockClient)
	assert.Equal(t, []ts.ImageBaseInfo{{Id: "1234567890ab"}}, result)
	assert.Nil(t, err)
	mockClient.AssertExpectations(t)

	// test: image list empty
	mockClient.On("ImageList", context.Background(), image.ListOptions{All: true, ContainerCount: true}).Return([]image.Summary{}, nil).Once()

	result, err = GetImageList(mockClient)
	assert.Equal(t, []ts.ImageBaseInfo{}, result)
	assert.Nil(t, err)
	mockClient.AssertExpectations(t)

	// test: ImageList throw error
	mockClient.On("ImageList", context.Background(), image.ListOptions{All: true, ContainerCount: true}).Return([]image.Summary{}, errors.New("ImageList error")).Once()

	result, err = GetImageList(mockClient)
	assert.Nil(t, result)
	assert.EqualError(t, err, "ImageList error")
	mockClient.AssertExpectations(t)
}
