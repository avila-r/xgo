package xdocker

import (
	"context"
	"errors"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type DockerClient struct {
	client *client.Client
	ctx    context.Context
}

func NewClient(options ...client.Opt) *DockerClient {
	c, _ := client.NewClientWithOpts(options...)

	return &DockerClient{
		client: c,
		ctx:    context.Background(),
	}
}

func (c *DockerClient) Close() {
	c.client.Close()
}

func (c *DockerClient) ListAll() ([]types.Container, error) {
	return c.client.ContainerList(c.ctx, container.ListOptions{All: true})
}

func (c *DockerClient) GetById(id string) (*types.Container, error) {
	containers, err := c.client.ContainerList(c.ctx, container.ListOptions{All: true})

	if err != nil {
		return nil, err
	}

	var (
		container *types.Container
	)

	for _, ctr := range containers {
		if ctr.ID == id {
			container = &ctr
		}
	}

	if container == nil {
		return nil, errors.New("container not found")
	}

	return container, nil
}
