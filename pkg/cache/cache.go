package cache

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type Server struct {
	client *redis.Client
	ctx    context.Context
}

func NewServer(options *redis.Options) *Server {
	return &Server{
		client: redis.NewClient(options),
		ctx:    context.Background(),
	}
}

func (s *Server) Cache(key string, value string) {
	if err := s.client.Set(s.ctx, key, value, 0).Err(); err != nil {
		panic(err.Error())
	}
}

func (s *Server) Uncache(key string) string {
	result, err := s.client.Get(s.ctx, key).Result()

	if err != nil {
		panic(err.Error())
	}

	return result
}

type CacheClient[T any] struct {
	server *Server
}

func NewClient[T any](options *redis.Options) *CacheClient[T] {
	return &CacheClient[T]{NewServer(options)}
}

func Client[T any](server *Server) *CacheClient[T] {
	return &CacheClient[T]{server}
}

func (c *CacheClient[T]) Cache(key string, data *T) {
	json, _ := json.Marshal(data)

	if err := c.server.client.Set(c.server.ctx, key, json, 0).Err(); err != nil {
		panic(err.Error())
	}
}

func (c *CacheClient[T]) Uncache(key string) *T {
	v, err := c.server.client.Get(c.server.ctx, key).Result()

	if err != nil {
		panic(err.Error())
	}

	var result T

	if err := json.Unmarshal([]byte(v), &result); err != nil {
		panic(err.Error())
	}

	return &result
}
