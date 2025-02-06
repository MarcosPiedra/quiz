package concurrentmap

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
)

type ConcurrentMap[T any] struct {
	data map[int]T
	mu   sync.RWMutex
	id   int32
}

func NewConcurrentMap[T any]() *ConcurrentMap[T] {
	return &ConcurrentMap[T]{
		data: make(map[int]T),
		id:   0,
	}
}

func (c *ConcurrentMap[T]) Count(ctx context.Context) (int, error) {
	select {
	case <-ctx.Done():
		return 0, errors.New("operation cancelled")
	default:
		c.mu.Lock()
		defer c.mu.Unlock()
		return len(c.data), nil
	}
}

func (c *ConcurrentMap[T]) Set(key int, value T, ctx context.Context) error {
	select {
	case <-ctx.Done():
		return errors.New("operation cancelled")
	default:
		c.mu.Lock()
		defer c.mu.Unlock()
		c.data[key] = value
		return nil
	}
}

func (c *ConcurrentMap[T]) Get(key int, ctx context.Context) (T, bool, error) {
	select {
	case <-ctx.Done():
		var zero T
		return zero, false, errors.New("operation cancelled")
	default:
		c.mu.RLock()
		defer c.mu.RUnlock()
		val, exists := c.data[key]
		return val, exists, nil
	}
}

func (c *ConcurrentMap[T]) GetAll(ctx context.Context) ([]T, error) {
	select {
	case <-ctx.Done():
		return nil, errors.New("operation cancelled")
	default:
		c.mu.RLock()
		defer c.mu.RUnlock()
		var values []T
		for _, v := range c.data {
			values = append(values, v)
		}
		return values, nil
	}
}

func (c *ConcurrentMap[T]) NextId() int {
	return int(atomic.AddInt32(&c.id, 1))
}
