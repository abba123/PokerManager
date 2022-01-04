package model

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var ctx = context.Background()

func InitRedis() *redis.Client {
	viper.AutomaticEnv()
	redisUrl := viper.GetString("REDIS") + ":6379"
	client := redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}

func RemoveKeyRedis(username string) {

	client := InitRedis()

	iter := client.Scan(ctx, 0, username+"*", 0).Iterator()

	for iter.Next(ctx) {
		client.Del(ctx, iter.Val())
	}

}

func GetHandRedis(num string, gain string, seat string, username string) []Game {
	client := InitRedis()

	existGain, _ := client.Exists(ctx, username+"Gain"+gain).Result()
	existSeat, _ := client.Exists(ctx, username+"Seat"+seat).Result()

	if existGain == 0 {
		InsertGainRedis(gain, username)
	}
	if existSeat == 0 {
		InsertSeatRedis(seat, username)
	}
	client.ZInterStore(ctx, "inter", &redis.ZStore{Keys: []string{username + "Gain" + gain, username + "Seat" + seat}}).Result()
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

func InsertGainRedis(gain string, username string) {
	games := GetGainDB("pokerdb",gain, username)

	client := InitRedis()

	for i, game := range games {
		gameStr, _ := json.Marshal(game)
		client.ZAdd(ctx, username+"Gain"+gain, &redis.Z{Score: float64(i), Member: gameStr})
	}
}

func InsertSeatRedis(seat string, username string) {
	games := GetSeatDB("pokerdb",seat, username)

	client := InitRedis()
	for i, game := range games {
		gameStr, _ := json.Marshal(game)
		client.ZAdd(ctx, username+"Seat"+seat, &redis.Z{Score: float64(i), Member: gameStr})
	}
}

func GetProfitRedis(username string, player string) []string {
	client := InitRedis()

	exist, _ := client.Exists(ctx, username+player+"Profit").Result()

	if exist == 0 {
		InsertProfitRedis(username, player)
	}

	result, _ := client.LRange(ctx, username+player+"Profit", 0, -1).Result()

	return result
}

func InsertProfitRedis(username string, player string) {
	client := InitRedis()

	profits := GetProfitDB("pokerdb",username, player)

	for _, profit := range profits {
		client.RPush(ctx, username+player+"Profit", fmt.Sprint(profit))
	}
}

func GetActionRedis(stage string, action string, username string, player string) string {
	client := InitRedis()

	exist, _ := client.Exists(ctx, username+player+stage+action).Result()

	if exist == 0 {
		InsertActionRedis(stage, action, username, player)
	}

	result, _ := client.Get(ctx, username+player+stage+action).Result()

	return result
}

func InsertActionRedis(stage string, action string, username string, player string) {
	client := InitRedis()

	result := GetActionDB("pokerdb",stage, action, username, player)

	client.Set(ctx, username+player+stage+action, fmt.Sprint(result), 0)
}

func GetPlayerRedis(username string) []string {
	client := InitRedis()

	exist, _ := client.Exists(ctx, username+"playerlist").Result()

	if exist == 0 {
		InsertPlayerRedis(username)
	}

	result, _ := client.LRange(ctx, username+"playerlist", 0, -1).Result()

	return result
}

func InsertPlayerRedis(username string) {
	client := InitRedis()

	results := GetPlayerDB("pokerdb",username)

	for _, result := range results {
		client.RPush(ctx, username+"playerlist", fmt.Sprint(result))
	}

}
