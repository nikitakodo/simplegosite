package services

import (
	"encoding/json"
	"github.com/go-redis/redis/v7"
)

type Cache struct {
	Client *redis.Client
	Prefix string
}

func (cache *Cache) SetStruct(key string, data interface{}) error {

	encodedData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	status := cache.Client.Set(
		cache.Prefix+key,
		encodedData,
		0,
	)
	_, err = status.Result()

	return err
}

func (cache *Cache) GetMarshalStruct(key string) (*string, error) {
	status := cache.Client.Get(cache.Prefix + key)
	v, err := status.Result()
	return &v, err
}
