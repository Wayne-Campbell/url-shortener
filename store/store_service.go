package store

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

//define the struct wrapper around raw redis client
type StorageService struct{
	redisClient *redis.Client
}

//top level declarations for the storeService and Redis context
var(
	storeService = &StorageService{}
	ctx = context.Background()
)

const CacheDuration = 6 * time.Hour

//initializing store service and returning a store pointer
func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB:		0,
	})
pong, err := redisClient.Ping(ctx).Result()
if err != nil {
	panic(fmt.Sprintf("Error init Redis: %v", err))
}

fmt.Printf("\nRedis started succesfully: pong message = {%s}", pong)
storeService.redisClient = redisClient
return storeService

}

//mapping between old url and new url

func SaveUrlMapping(shortUrl, originalUrl, userId string){
err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
if err!= nil {
	panic(fmt.Sprintf("Failed saving key url | Error: %v shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}
}
//retrieving intial URL only
func RetrieveInitialUrl(shortUrl string) string {
result, err := storeService.redisClient.Get(ctx, shortUrl).Result()
if err != nil {
	panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
}
return result
}
