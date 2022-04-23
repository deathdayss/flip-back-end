package repository

import (
	"fmt"

	"github.com/deathdayss/flip-back-end/models"
	"github.com/go-redis/redis"
)

var zsetKey string = "Search:Rank"

func AddWord(word string) error {
	reply, err := models.RDB.ZIncr(zsetKey, redis.Z{Member: word}).Result()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println(reply)
	return nil
}

func SearchRank() ([]string, error) {
	ret, err := models.RDB.ZRevRangeWithScores(zsetKey, 0, 10).Result()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	ans := make([]string, len(ret))
	for i := range ret {
		ans[i] = ret[i].Member.(string)
	}
	return ans, nil
}