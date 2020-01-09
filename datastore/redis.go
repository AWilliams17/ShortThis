package datastore

import (
	"errors"
	"github.com/gomodule/redigo/redis"
	"github.com/lithammer/shortuuid"
	"log"
	"time"
)

var Pool *redis.Pool
var errNoUUID = errors.New("UUID was not found")

type UrlEntry struct {
	UUID      string `redis:"UUID"`
	Url       string `redis:"URL"`
	CreatedAt string `redis:"CreatedAt"`
}

func AddUrlEntry(url string) string {
	conn := Pool.Get()
	defer conn.Close()

	urlUUID := shortuuid.New()[:8]

	_, err := conn.Do("HMSET", "uuid:"+urlUUID, "url", url, "createdAt", time.Now())
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Do("EXPIRE", "uuid:"+urlUUID, 1800)
	if err != nil {
		log.Fatal(err)
	}

	return urlUUID
}

func FindUrlEntry(uuid string) (*UrlEntry, error) {
	conn := Pool.Get()
	defer conn.Close()

	values, err := redis.Values(conn.Do("HGETALL", "uuid:"+uuid))
	if err != nil {
		return nil, err
	} else if len(values) == 0 {
		return nil, errNoUUID
	}

	var urlEntry UrlEntry
	err = redis.ScanStruct(values, &urlEntry)
	if err != nil {
		return nil, err
	}

	return &urlEntry, nil
}
