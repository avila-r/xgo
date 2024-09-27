package cache

import (
	"encoding/json"
	"log"
	"time"

	"github.com/icza/gog"
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
	json, err := json.Marshal(i.Data)

	if err != nil {
		return false
	}

	expiration := func() time.Duration {
		if i.ExpirationTime < 0 {
			return i.ExpirationTime
		}

		return time.Minute * 10
	}()

	if err := c.server.Client.Set(c.server.Ctx, i.Key, gog.First(json), expiration).Err(); err != nil {
		log.Fatalf("error: %v", err.Error())
		return false
	}

	return true
}

// Returns nil if ok, error otherwise
func (c *Cacher[T]) CacheOrError(i Insert[T]) error {
	json, err := json.Marshal(i.Data)

	if err != nil {
		return err
	}

	expiration := func() time.Duration {
		if i.ExpirationTime < 0 {
			return i.ExpirationTime
		}

		return time.Minute * 10
	}()

	if err := c.server.Client.Set(c.server.Ctx, i.Key, gog.First(json), expiration).Err(); err != nil {
		return err
	}

	return nil
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
