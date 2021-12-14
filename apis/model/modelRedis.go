package model

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func InitRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}

func RemoveKeyRedis(player string) {
	client := InitRedis()

	iter := client.Scan(ctx, 0, player, 0).Iterator()

	for iter.Next(ctx) {
		client.Del(ctx, iter.Val())
	}

}

func GetHandRedis(num string, gain string, seat string, player string) []Game {
	client := InitRedis()

	existGain, _ := client.Exists(ctx, player+"Gain"+gain).Result()
	existSeat, _ := client.Exists(ctx, player+"Seat"+seat).Result()

	if existGain == 0 {
		InsertGainRedis(gain, player)
	}
	if existSeat == 0 {
		InsertSeatRedis(seat, player)
	}
	client.ZInterStore(ctx, "inter", &redis.ZStore{Keys: []string{player + "Gain" + gain, player + "Seat" + seat}}).Result()
	results, _ := client.ZRange(ctx, "inter", 0, -1).Result()
	client.Del(ctx, "inter")

	games := []Game{}
	n, _ := strconv.Atoi(num)
	for i := 0; i < n && i < len(results); i++ {
		result := results[i]
		g := Game{}
		json.Unmarshal([]byte(result), &g)
		games = append(games, g)
	}
	return games
}

func InsertGainRedis(gain string, player string) {
	games := GetGainDB(gain, player)

	client := InitRedis()

	for i, game := range games {
		gameStr, _ := json.Marshal(game)
		client.ZAdd(ctx, player+"Gain"+gain, &redis.Z{Score: float64(i), Member: gameStr})
	}
}

func InsertSeatRedis(seat string, player string) {
	games := GetSeatDB(seat, player)

	client := InitRedis()

	for i, game := range games {
		gameStr, _ := json.Marshal(game)
		client.ZAdd(ctx, player+"Seat"+seat, &redis.Z{Score: float64(i), Member: gameStr})
	}
}

func GetProfitRedis(player string) []string {
	client := InitRedis()

	exist, _ := client.Exists(ctx, player+"Profit").Result()

	if exist == 0 {
		InsertProfitRedis(player)
	}

	result, _ := client.LRange(ctx, player+"Profit", 0, -1).Result()

	return result
}

func InsertProfitRedis(player string) {
	client := InitRedis()

	profits := GetProfitDB(player)

	for _, profit := range profits {
		client.RPush(ctx, player+"Profit", fmt.Sprint(profit))
	}
}

func GetCountAllRedis(player string) string {
	client := InitRedis()

	exist, _ := client.Exists(ctx, player+"CountAll").Result()

	if exist == 0 {
		InsertCountAllRedis(player)
	}

	result, _ := client.Get(ctx, player+"CountAll").Result()

	return result
}

func InsertCountAllRedis(player string) {
	client := InitRedis()

	result := GetCountAllDB(player)

	client.Set(ctx, player+"CountAll", fmt.Sprint(result), 0)
}

func GetCountThreeBetRedis(player string) string {
	client := InitRedis()

	exist, _ := client.Exists(ctx, player+"CountThreeBet").Result()

	if exist == 0 {
		InsertCountThreeBetRedis(player)
	}

	result, _ := client.Get(ctx, player+"CountThreeBet").Result()

	return result
}

func InsertCountThreeBetRedis(player string) {
	client := InitRedis()

	result := GetCountThreeBetDB(player)

	client.Set(ctx, player+"CountThreeBet", fmt.Sprint(result), 0)
}
