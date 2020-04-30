package db

import (
    "github.com/go-redis/redis/v7"
)

type RedisDatabase struct {
    client *redis.Client
}

func (r RedisDatabase) Connect() error {
    r.client = redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })

    err := r.client.Ping().Err()
    return err
}
