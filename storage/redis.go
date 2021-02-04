package storage

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

var (
	ErrNotFound = errors.New("not found")
)

// Redis wraps redis client
type Redis struct {
	client  *redis.Client
	timeout time.Duration
}

func New(address string, timeout time.Duration, poolSize int) (*Redis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:        address,
		PoolSize:    poolSize,
		ReadTimeout: timeout,
	})

	if err := client.Ping().Err(); err != nil {
		return nil, errors.Wrap(err, "failed to ping redis")
	}

	return &Redis{
		client:  client,
		timeout: timeout,
	}, nil
}

// Get retrieves a value from redis
func (s *Redis) Get(key string) ([]byte, error) {
	res, err := s.client.Get(key).Bytes()
	if err == redis.Nil {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to get from redis")
	}
	return res, nil
}

// Set sets the key/value with ttl
func (s *Redis) Set(key string, value []byte, ttl time.Duration) error {
	if err := s.client.Set(key, value, ttl).Err(); err != nil {
		return errors.Wrap(err, "failed to set redis value")
	}
	return nil
}
