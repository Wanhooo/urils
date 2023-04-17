package database

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"go.uber.org/zap"
	"time"
	"urils/application/config"
)

var RedisPool0 *redis.Pool
var RedisPool1 *redis.Pool

func InitRedisPool(cfg *config.RedisConfig) *redis.Pool {
	return &redis.Pool{
		MaxActive:   cfg.MaxFreeConn,            // 允许分配最大连接数
		MaxIdle:     cfg.MaxConn,                // 最大空闲连接数
		IdleTimeout: time.Duration(cfg.TimeOut), // 连接时间限制
		Dial: func() (redis.Conn,  error) {      // 创建连接的函数
			c,err := redis.Dial(
				"tcp",
				fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
				redis.DialPassword(cfg.Password),   // 如果需要密码，就写
				redis.DialDatabase(cfg.Database),	// 如果是其他库，改成int类型的数字（1, 2 ...，15）
			)
			if err != nil {
				zap.S().Errorf("redis connection fail: %v", err)
				return nil, err
			}
			return c,nil
		},
	}
}

