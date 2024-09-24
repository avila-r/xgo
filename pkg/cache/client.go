package cache

import (
	"encoding/json"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cacher[T any] struct {
	server *Server
}

func NewClient[T any](options *redis.Options) *Cacher[T] {
	return &Cacher[T]{NewServer(options)}
}

// Returns true if ok, false otherwise
func (c *Cacher[T]) Cache(i Insert[T]) bool {
	json, _ := json.Marshal(i.Data)

	expiration := func() time.Duration {
		if i.ExpirationTime < 0 {
			return i.ExpirationTime
		}

		return time.Minute * 10
	}()

	if err := c.server.Client.Set(c.server.Ctx, i.Key, json, expiration).Err(); err != nil {
		log.Fatalf("error: %v", err.Error())
		return false
	}

	return true
}

func (c *Cacher[T]) Uncache(q Query) (*T, error) {
	v, err := c.server.Client.Get(c.server.Ctx, q.Key).Result()

	if err != nil {
		return nil, err
	}

	var response T

	if err := json.Unmarshal([]byte(v), &response); err != nil {
		return nil, err
	}

	return &response, nil
}
