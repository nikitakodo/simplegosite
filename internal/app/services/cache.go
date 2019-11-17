package services

import (
	"encoding/json"
	"github.com/go-redis/redis/v7"
)

type Cache struct {
	Client *redis.Client
	Prefix string
}

func (cache *Cache) Set(key string, data interface{}) error {
	status := cache.Client.Set(
		cache.Prefix+key,
		data,
		0,
	)
	_, err := status.Result()

	return err
}

func (cache *Cache) SetStruct(key string, data interface{}) error {

	encodedData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return cache.Set(key, encodedData)
}

func (cache *Cache) Get(key string) (*string, error) {
	status := cache.Client.Get(cache.Prefix + key)
	v, err := status.Result()

	if err != nil && err != redis.Nil {
		return nil, err
	}

	return &v, nil
}
