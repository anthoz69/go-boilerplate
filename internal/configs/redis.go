package configs

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func NewRedisClient(v *viper.Viper) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", v.GetString("redis.host"), v.GetString("redis.port")),
		Password: v.GetString("redis.password"),
		DB:       v.GetInt("redis.db"),
	})
}
