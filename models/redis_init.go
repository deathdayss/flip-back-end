package models

import (
	"fmt"

	"github.com/go-redis/redis"
)

var RDB *redis.Client

func InitRedisClient() (err error) {
    RDB = redis.NewClient(&redis.Options{
        Addr:     "175.178.159.131:6379",
        Password: "Cptbtptp1790340626.", // no password set
        DB:       0,  // use default DB
    })

    _, err = RDB.Ping().Result()
    if err != nil {
        fmt.Println(err)
        return err
    }
    return nil
}