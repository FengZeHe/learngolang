package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/basicprojectv2/internal/domain"
	"github.com/redis/go-redis/v9"
	"time"
)

var ErrKeyNotExist = redis.Nil

type UserCache interface {
	Get(ctx context.Context, id string) (domain.User, error)
	Set(ctx context.Context, du domain.User) error
}

type RedisUserCache struct {
	cmd        redis.Cmdable
	expiration time.Duration
}

func NewUserCache(cmd redis.Cmdable) UserCache {
	return &RedisUserCache{cmd: cmd, expiration: time.Minute * 15}
}

func (c *RedisUserCache) key(uid string) string {
	return fmt.Sprintf("user:info:%s", uid)
}

func (c *RedisUserCache) Get(ctx context.Context, uid string) (domain.User, error) {
	key := c.key(uid)
	data, err := c.cmd.Get(ctx, key).Result()
	if err != nil {
		return domain.User{}, err
	}
	var u domain.User
	err = json.Unmarshal([]byte(data), &u)

	return u, err
}

func (c *RedisUserCache) Set(ctx context.Context, du domain.User) error {
	//intId, err := strconv.Atoi(du.ID)
	key := c.key(du.ID)
	data, err := json.Marshal(du)
	if err != nil {
		return err
	}
	return c.cmd.Set(ctx, key, data, c.expiration).Err()
}
