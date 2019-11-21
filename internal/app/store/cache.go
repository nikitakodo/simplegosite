package store

import (
	"encoding/json"
	"github.com/go-redis/redis/v7"
	"time"
)

type Cache struct {
	Client *redis.Client
	Prefix string
}

func (cache *Cache) Set(key string, data interface{}, expiration time.Duration) error {
	_, err := cache.Client.Set(
		cache.Prefix+key,
		data,
		expiration,
	).Result()

	return err
}

func (cache *Cache) SetStruct(key string, data interface{}) error {

	encodedData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return cache.Set(key, encodedData, 0)
}

func (cache *Cache) Get(key string) (*string, error) {
	v, err := cache.Client.Get(cache.Prefix + key).Result()

	if err != nil && err != redis.Nil {
		return nil, err
	}

	return &v, nil
}

func (cache *Cache) Del(key string) error {
	_, err := cache.Client.Del(cache.Prefix + key).Result()
	if err != nil && err != redis.Nil {
		return err
	}

	return nil
}
