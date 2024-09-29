package cache

import "time"

type Register struct {
	Key  string
	Data any

	// If 0, data won't have expiration time
	ExpirationTime time.Duration
}

type Query struct {
	Key    string
	Result any
}
