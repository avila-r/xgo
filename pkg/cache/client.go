package cache

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	ErrNoProvidedKey = errors.New("no key was provided at Insert{} request")
	ErrInvalidData   = errors.New("invalid data was provided at Insert{} request")
	ErrNilQuery      = errors.New("provided query must be non-nil")
)

type Client struct {
	Advanced *redis.Client
	Ctx      context.Context
}

func NewClient(options *redis.Options) *Client {
	return &Client{
		Advanced: redis.NewClient(options),
		Ctx:      context.Background(),
	}
}

func (c *Client) Insert(i *Register) error {
	if i.Key == "" {
		return ErrNoProvidedKey
	}

	if i.Data == nil {
		return ErrInvalidData
	}

	json, err := json.Marshal(i.Data)

	if err != nil {
		return err
	}

	if err := c.Advanced.Set(c.Ctx, i.Key, json, i.ExpirationTime).Err(); err != nil {
		return err
	}

	return nil
}

func (c *Client) Get(q *Query) error {
	if q == nil {
		return ErrNilQuery
	}

	v, err := c.Advanced.Get(c.Ctx, q.Key).Result()

	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(v), &q.Result); err != nil {
		return err
	}

	return nil
}

func (c *Client) Cache(key string, v string, expiration ...time.Duration) error {
	var exp time.Duration

	if len(expiration) == 0 {
		exp = 0
	} else {
		exp = expiration[0]
	}

	if err := c.Advanced.Set(c.Ctx, key, v, exp).Err(); err != nil {
		return err
	}

	return nil
}

func (c *Client) Uncache(key string) (string, error) {
	return c.Advanced.Get(c.Ctx, key).Result()
}
