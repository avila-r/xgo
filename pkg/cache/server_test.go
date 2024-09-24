package cache_test

import (
	"context"
	"testing"

	"github.com/avila-r/xgo/pkg/cache"
	"github.com/go-redis/redismock/v9"
)

func Test_Cache(t *testing.T) {
	data := struct { Key string; Value string }{
		Key:   "test-key",
		Value: "test-value",
	}

	c, mock := redismock.NewClientMock()

	mock.ExpectSet(data.Key, data.Value, 0).SetVal("OK")

	server := &cache.Server{
		Client: c,
		Ctx:    context.Background(),
	}

	if ok := server.Cache(data.Key, data.Value); !ok {
		t.Errorf("Cache() failed. Expected true, received false")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock's expectations weren't met in order: %v", err)
	}
}

func Test_Uncache(t *testing.T) {
	data := struct { Key string; Value string }{
		Key:   "test-key",
		Value: "test-value",
	}

	c, mock := redismock.NewClientMock()

	mock.ExpectGet(data.Key).SetVal(data.Value)

	server := &cache.Server{
		Client: c,
		Ctx:    context.Background(),
	}

	v, err := server.Uncache(data.Key)

	if err != nil {
		t.Errorf("Uncache() failed. Some unexpected error occurred: %v", err)
	}

	if v != data.Value {
		t.Errorf("Uncache() failed. Expected %v, received %v.", data.Value, v)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock's expectations weren't met in order: %v", err)
	}
}
