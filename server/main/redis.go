package main

import (
	"chatroom/server/model/github.com/gomodule/redigov1.8.1/redis"
	"time"
)

var pool *redis.Pool

//初始化连接池
func initPool(address string, maxIdle, maxActive int, idleTimeout time.Duration) {
	pool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
	}
}
