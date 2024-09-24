package cache

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

type Server struct {
	Client *redis.Client
	Ctx    context.Context
}

func NewServer(options *redis.Options) *Server {
	return &Server{
		Client: redis.NewClient(options),
		Ctx:    context.Background(),
	}
}

func (s *Server) Cache(key string, value string) bool {
	if err := s.Client.Set(s.Ctx, key, value, 0).Err(); err != nil {
		log.Fatalf("error: %v", err.Error())
		return false
	}

	return true
}

func (s *Server) Uncache(key string) (string, error) {
	result, err := s.Client.Get(s.Ctx, key).Result()

	return result, err
}
