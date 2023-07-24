package store

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
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