package xdocker_test

import (
	"log"
	"testing"

	"github.com/avila-r/xgo/pkg/xdocker"
	"github.com/docker/docker/client"
)

func Test_ListAll(t *testing.T) {
	c := xdocker.NewClient(client.FromEnv)

	defer c.Close()

	containers, err := c.ListAll()

	if err != nil {
		t.Error(err)
	}

	for _, container := range containers {
		log.Print(container.ID)
	}
}

func Test_GetById(t *testing.T) {
	c := xdocker.NewClient(client.FromEnv)

	defer c.Close()

	ctr, err := c.GetById("1ee5b75cef83b5cad1a078db42229ad0acebf83288fedbccc4e082bcc8018f9a")

	if err != nil {
		t.Error(err)
	}

	log.Printf("%s %s\n", ctr.ID, ctr.Image)
}
