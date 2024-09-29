package cache_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/avila-r/xgo/pkg/cache"
	"github.com/go-redis/redismock/v9"
)

var (
	v = struct {
		Key   string
		Value string
	}{
		Key:   "test-key",
		Value: "test-value",
	}
)

func Test_Insert(t *testing.T) {
	c, mock := redismock.NewClientMock()

	expected, _ := json.Marshal(v.Value)

	mock.ExpectSet(v.Key, expected, 0).SetVal("OK")

	// Custom client
	client := &cache.Client{
		Advanced: c,
		Ctx:      context.Background(),
	}

	insert := cache.Register{
		Key:  v.Key,
		Data: v.Value,
	}

	if err := client.Insert(&insert); err != nil {
		t.Errorf("failed at caching data - %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock's expectations weren't met in order: %v", err)
	}
}

func Test_Query(t *testing.T) {
	c, mock := redismock.NewClientMock()

	// Custom client
	client := &cache.Client{
		Advanced: c,
		Ctx:      context.Background(),
	}

	expected, _ := json.Marshal(v.Value)

	mock.ExpectSet(v.Key, expected, 0).SetVal("OK")

	insert := cache.Register{
		Key:  v.Key,
		Data: v.Value,
	}

	if err := client.Insert(&insert); err != nil {
		t.Errorf("failed at caching data - %v", err)
	}

	mock.ExpectGet(v.Key).SetVal(string(expected))

	query := cache.Query{
		Key: v.Key,
	}

	if err := client.Get(&query); err != nil {
		t.Errorf("failed at wrapping data - %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("mock's expectations weren't met in order: %v", err)
	}

	if query.Result != v.Value {
		t.Errorf("invalid uncache result. Expected %v, but got %v", v.Value, query.Result)
	}
}

func Test_Cache(t *testing.T) {
	c, mock := redismock.NewClientMock()

	mock.ExpectSet(v.Key, v.Value, 0).SetVal("OK")

	client := &cache.Client{
		Advanced: c,
		Ctx:      context.Background(),
	}

	if err := client.Cache(v.Key, v.Value); err != nil {
		t.Errorf("failed at caching data - %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock's expectations weren't met in order: %v", err)
	}
}

func Test_Uncache(t *testing.T) {
	c, mock := redismock.NewClientMock()

	mock.ExpectSet(v.Key, v.Value, 0).SetVal("OK")

	client := &cache.Client{
		Advanced: c,
		Ctx:      context.Background(),
	}

	if err := client.Cache(v.Key, v.Value); err != nil {
		t.Errorf("failed at caching data - %v", err)
	}

	mock.ExpectGet(v.Key).SetVal(v.Value)

	r, err := client.Uncache(v.Key)

	if err != nil {
		t.Errorf("failed at uncaching data - %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("mock's expectations weren't met in order: %v", err)
	}

	if r != v.Value {
		t.Errorf("invalid uncache result. Expected %v, but got %v", v.Value, r)
	}
}
