package redis

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/joho/godotenv"
)

var (
	_          = godotenv.Load()
	timeout    = time.Duration(10) * time.Second // 10sec timeout for redis
	expiration = getExpTime()
)

type Service struct {
	client *redis.Client
}

func NewRedisSessionService() *Service {
	s := new(Service)
	s.client = getClient()

	return s
}

func (r *Service) SetSession(username string) (string, error) {
	key := utils.UUIDv4()
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err := r.client.SetEX(ctx, key, username, expiration).Err()

	return key, err
}

func (r *Service) GetSession(token string) *string {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res, err := r.client.Get(ctx, token).Result()

	if err != nil {
		return nil
	}

	if res == "" {
		return nil
	}

	return &res
}

func getExpTime() time.Duration {
	exp, err := strconv.Atoi(os.Getenv("SESSION_EXP"))

	if err != nil {
		panic("Failed to get session expiration time.")
	}

	return time.Duration(exp*24) * time.Hour
}

func getClient() *redis.Client {
	rdb := redis.NewClient(
		&redis.Options{
			Addr: os.Getenv("REDIS_URL"),
		},
	)

	return rdb
}
