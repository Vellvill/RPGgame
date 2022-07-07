package cache

import (
	"Consumer/internal/config"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

const redisErrCon = "can not connect to redis"

type connErr struct {
	pgxConnErr string
	attempts   int
	time       int
}

func (c connErr) Error() string {
	return fmt.Sprintf("%d attempts, %d time, %s", c.attempts, c.time, redisErrCon)
}

type RClient interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	WithTimeout(timeout time.Duration) *redis.Client
	Ping(ctx context.Context) *redis.StatusCmd
}

func NewRedisClient(config *config.Config, t, att int) (RClient, error) {
	client := redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr: fmt.Sprintf("%s:%s",
			config.Redis.RHost,
			config.Redis.RPort),
		Password: config.Redis.RPass,
		DB:       0,
	})
	r := rClient{client}
	err := r.PingR(t, att)
	if err != nil {
		return nil, err
	}
	return r, nil
}

type rClient struct {
	RClient
}

func (r rClient) PingR(t, att int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(t)*time.Second)
	defer cancel()
	for i := att; i < 0; i-- {
		if r.Ping(ctx).Err() != nil {
			continue
		}
		return nil
	}
	return connErr{
		pgxConnErr: redisErrCon,
		attempts:   att,
		time:       t,
	}
}
