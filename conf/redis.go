package conf

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"time"
)

// redis 客户端
var rdCilent *redis.Client
var nDuration = 30 * 24 * 60 * 60 * time.Second

// 要对 redis 客户端进行包装，使得客户端操作方法被封装起来，不会暴露给外部
type RedisClient struct {
}

// InitRedis
//
//	@Description: 初始化Redis客户端，利用viper对象获取配置信息返回一个redis操作的客户端对象
//	@return *RedisClient
//	@return error
func InitRedis() (*RedisClient, error) {
	rdCilent = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.url"),
		Password: "",
		DB:       0,
	})

	_, err := rdCilent.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return &RedisClient{}, nil
}

// 利用初始化中生成的rdClient对象，对redis进行操作
func (r *RedisClient) Set(key string, value interface{}) error {
	return rdCilent.Set(context.Background(), key, value, nDuration).Err()
}

func (r *RedisClient) Get(key string) (string, error) {
	return rdCilent.Get(context.Background(), key).Result()
}

func (r *RedisClient) Del(key string) error {
	return rdCilent.Del(context.Background(), key).Err()
}
