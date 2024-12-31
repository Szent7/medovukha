package docker

import (
	"context"
	"errors"
	"io"
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
