package helpers

import (
	"context"
	"encoding/json"

	"github.com/Pratham-Mishra04/fampay/fampay-backend/config"
	"github.com/Pratham-Mishra04/fampay/fampay-backend/initializers"
	"github.com/Pratham-Mishra04/fampay/fampay-backend/models"
	"github.com/redis/go-redis/v9"
)

var ctx = context.TODO()

func GetFromCache(key string) []models.Video {
	data, err := initializers.RedisClient.Get(ctx, key).Result()
	if err != nil {
		if err != redis.Nil {
			config.Logger.Warnw("Error Getting from cache", "Error:", err)
		}
		return nil
	}

	videos := []models.Video{}
	if err = json.Unmarshal([]byte(data), &videos); err != nil {
		config.Logger.Warnw("Error while unmarshaling videos", "Error:", err)
		return nil
	}
	return videos
}

func SetToCache(key string, videos []models.Video) {
	data, err := json.Marshal(videos)
	if err != nil {
		config.Logger.Warnw("Error while marshaling videos", "Error:", err)
	}

	if err := initializers.RedisClient.Set(ctx, key, data, initializers.CacheExpirationTime).Err(); err != nil {
		config.Logger.Warnw("Error Setting to cache", "Error:", err)
	}
}

func RemoveFromCache(key string) {
	err := initializers.RedisClient.Del(ctx, key).Err()
	if err != nil && err != redis.Nil {
		config.Logger.Warnw("Error Removing from cache", "Error:", err)
	}
}

func FlushCache() {
	err := initializers.RedisClient.FlushAll(ctx).Err()
	if err != nil {
		config.Logger.Warnw("Error flushing cache", "Error", err)
	}
}
