package db

import (
	"github.com/go-redis/redis"
)

type Client struct {
	*redis.Client
}

func NewClient() *Client {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

    return &Client{client}
}