package cache

import (
	"errors"
	"fmt"
	"gopkg.in/redis.v5"
	"sync"
	"time"
)

var (
	redisClients = make(map[string]*redis.Client)
	ones         sync.Once
)

/**
 * redis 初始化,需要在调用前初始化，只需初始化一次
 */
func InitRedisClients(redisConfigs *RedisCacheConf) {
	ones.Do(func() {
		if redisConfigs == nil || len(redisConfigs.RedisConf) == 0 {
			panic("InitRedisClients error, config is nil")
		}
		for redisName, config := range redisConfigs.RedisConf {
			opts := &redis.Options{
				Addr:     config.Addr,
				Password: config.Password,
				DB:       config.DB,
			}
			if config.ReadTimeout > 0 {
				opts.ReadTimeout = time.Duration(config.ReadTimeout) * time.Second
			}
			if config.WriteTimeout > 0 {
				opts.WriteTimeout = time.Duration(config.WriteTimeout) * time.Second
			}
			if config.PoolSize > 0 {
				opts.PoolSize = config.PoolSize
			}
			if config.ConnTimeout > 0 {
				opts.DialTimeout = time.Duration(config.ConnTimeout) * time.Second
			}
			if config.IdleTimeout > 0 {
				opts.IdleTimeout = time.Duration(config.IdleTimeout) * time.Second
			}
			redisClient := redis.NewClient(opts)
			redisClients[redisName] = redisClient
		}
	})
}

/**
 * 通过名称获取redisclient,需要先调用初始化
 */
func GetRedisClient(name string) (*redis.Client, error) {
	client, ok := redisClients[name]
	if !ok {
		fmt.Printf("redis %s not exists", name)
		return nil, errors.New("redis not exists")
	}
	return client, nil
}
