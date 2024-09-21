package xdocker

import (
	"context"
	"fmt"

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

func (c *DockerClient) ListAll() error {
	containers, err := c.client.ContainerList(c.ctx, container.ListOptions{All: true})

	if err != nil {
		return err
	}

	for _, ctr := range containers {
		fmt.Printf("%s %s (status: %s)\n", ctr.ID, ctr.Image, ctr.Status)
	}

	return nil
}
