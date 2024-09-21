package cache_test

import (
	"testing"

	"github.com/avila-r/xgo/pkg/cache"
	"github.com/redis/go-redis/v9"
)

type Data struct {
	Name string `json:"name"`
}

func Test_CacheClient(t *testing.T) {
	options, _ := redis.ParseURL("redis://localhost:6379")

	client := cache.NewClient[Data](options)

	client.Cache("my-key", &Data{Name: "my name"})

	data := client.Uncache("my-key")

	if data == nil || data.Name != "my name" {
		t.Errorf("Expected 'my name', got: %v", data)
	}
}

func Test_CacheServer(t *testing.T) {
	options, _ := redis.ParseURL("redis://localhost:6379")

	server := cache.NewServer(options)

	server.Cache("my-key", "name")

	data := server.Uncache("my-key")

	if data != "name" {
		t.Errorf("Expected 'name', got: %v", data)
	}
}
