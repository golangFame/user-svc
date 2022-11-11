package cache

import (
	"github.com/BzingaApp/user-svc/internal/genesis"
	"github.com/gomodule/redigo/redis"
	"time"
)

type cache struct {
	*genesis.Service
}

type redisStore struct {
	pool              *redis.Pool
	defaultExpiration time.Duration
	*cache
}
