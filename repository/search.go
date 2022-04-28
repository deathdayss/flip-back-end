package repository

import (
	"fmt"
	"time"

	"github.com/deathdayss/flip-back-end/models"
	"github.com/go-redis/redis"
)

var gameKey string = "Search:Rank:Game"
var personKey string = "Search:Rank:Person"
var historyKey string = "Search:History:"
func AddWord(word string, mode string) error {
	var zsetKey string
	switch mode {
	case "person" : zsetKey = personKey
	case "game": zsetKey = gameKey
	}
	reply, err := models.RDB.ZIncrBy(zsetKey, 1.0, word).Result()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println(reply)
	return nil
}

func AddSearchHistory(email string, key string) error {
	var histroyIDKey string = historyKey + email
	script := redis.NewScript(`local len = redis.call('zcard', KEYS[1])
		if len == 10 then
			local rm = redis.call('zremrangebyrank', KEYS[1], 0, 0)
			if rm ~= 1 then
				return -1
			end
		end
		local ad = redis.call('zadd', KEYS[1], ARGV[2], ARGV[1])
		if ad ~= 1 and ad ~= 0 then
			return -1
		end
		if len == 0 then
			redis.call('expire', KEYS[1], 3600)
		end
		return 1`)
	result, err := script.Run(models.RDB, []string{histroyIDKey}, key, float64(time.Now().Unix())).Result()
	if err != nil {
		return err
	}
	if result == -1 {
		return fmt.Errorf("the script run fail")
	}
	return nil
}

func GetSearchHistory(email string) ([]string, error) {
	var histroyIDKey string = historyKey + email
	ans, err := models.RDB.ZRevRange(histroyIDKey, 0, -1).Result()
	if err != nil {
		return nil, err
	}
	return ans, nil
}

func SearchRank(mode string) ([]string, error) {
	var zsetKey string
	switch mode {
	case "person" : zsetKey = personKey
	case "game": zsetKey = gameKey
	}
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