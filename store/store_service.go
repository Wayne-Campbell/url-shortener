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
