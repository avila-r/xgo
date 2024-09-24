package cache

import "time"

type Insert[T any] struct {
	Key            string
	Data           T
	ExpirationTime time.Duration
}

type Query struct {
	Key string
}
