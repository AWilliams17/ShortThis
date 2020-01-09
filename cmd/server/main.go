package main

import (
	"github.com/awilliams17/ShortThis/api"
	"github.com/awilliams17/ShortThis/datastore"
	"github.com/gomodule/redigo/redis"
	"github.com/labstack/echo/middleware"
	"time"
)

func main() {
	datastore.Pool = &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}

	e := api.CreateRouter()
	e.Use(middleware.Logger())

	e.Use(middleware.Recover())

	e.Start(":8000")
}
