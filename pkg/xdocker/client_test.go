package xdocker_test

import (
	"testing"

	"github.com/avila-r/xgo/pkg/xdocker"
	"github.com/docker/docker/client"
)

func Test_ListAll(t *testing.T) {
	c := xdocker.NewClient(client.FromEnv)

	defer c.Close()

	if err := c.ListAll(); err != nil {
		t.Error(err)
	}
}
