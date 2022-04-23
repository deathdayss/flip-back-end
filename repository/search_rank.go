package repository

import (
	"fmt"

	"github.com/deathdayss/flip-back-end/models"
)

var gameKey string = "Search:Rank:Game"
var personKey string = "Search:Rank:Person"
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