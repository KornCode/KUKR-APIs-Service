package cache

import (
	"context"
	"fmt"

	config "github.com/KornCode/KUKR-APIs-Service/pkg/configs"
	"github.com/KornCode/KUKR-APIs-Service/pkg/logs"
	"github.com/go-redis/redis/v9"
)

var rd redisInstance

type redisInstance struct {
	Client *redis.Client
}

func NewConnectionRedisDB(conf config.RedisDB) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.Host, conf.Port),
		Password: conf.Password,
		DB:       0,
	})
	// client := redis.NewClient(&redis.Options{
	// 	Addr:     fmt.Sprintf("%s:%s", "127.0.0.1", "6379"),
	// 	Password: conf.Password,
	// 	DB:       0,
	// })

	if err := client.Ping(context.Background()).Err(); err != nil {
		logs.Error("RedisDB is Not connected")

		return nil, err
	}

	fmt.Println("RedisDB is Connected")

	rd = redisInstance{
		Client: client,
	}

	return client, nil
}

func CloseConnectionRedisDB() error {
	if rd.Client == nil {
		return nil
	}

	if err := rd.Client.Close(); err != nil {
		return err
	}

	return nil
}
