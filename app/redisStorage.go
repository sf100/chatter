package app

import (
	log "code.google.com/p/log4go"
	"errors"
	"github.com/garyburd/redigo/redis"
	"github.com/sf100/chatter/ketama"
	"strconv"
	"strings"
)

var RedisNoConnErr = errors.New("can't get a redis conn")

type RedisStorage struct {
	pool map[string]*redis.Pool
	ring *ketama.HashRing
}

var rs *RedisStorage

// initRedisStorage initialize the redis pool and consistency hash ring.
func InitRedisStorage() {
	log.Info("Connecting Redis....")
	var (
		err error
		w   int
		nw  []string
	)

	redisPool := map[string]*redis.Pool{}
	ring := ketama.NewRing(Conf.RedisKetamaBase)

	for n, addr := range Conf.RedisSource {
		nw = strings.Split(n, ":")
		if len(nw) != 2 {
			err = errors.New("node config error, it's nodeN:W")
			log.Error("strings.Split(\"%s\", :) failed (%v)", n, err)
			panic(err)
		}

		w, err = strconv.Atoi(nw[1])
		if err != nil {
			log.Error("strconv.Atoi(\"%s\") failed (%v)", nw[1], err)
			panic(err)
		}

		tmp := addr

		redisPool[nw[0]] = &redis.Pool{
			MaxIdle:     Conf.RedisMaxIdle,
			MaxActive:   Conf.RedisMaxActive,
			IdleTimeout: Conf.RedisIdleTimeout,
			Dial: func() (redis.Conn, error) {
				conn, err := redis.Dial("tcp", tmp)
				if err != nil {
					log.Error("redis.Dial(\"tcp\", \"%s\") error(%v)", tmp, err)
					return nil, err
				}

				return conn, err
			},
		}
		ring.AddNode(nw[0], w)
	}

	ring.Bake()
	rs = &RedisStorage{pool: redisPool, ring: ring}

	log.Info("Redis connected")
}

// 获取 Redis 连接.
func (s *RedisStorage) getConn(key string) redis.Conn {
	if len(s.pool) == 0 {
		return nil
	}

	node := s.ring.Hash(key)
	p, ok := s.pool[node]
	if !ok {
		log.Warn("key: \"%s\" hit redis node: \"%s\" not in pool", key, node)
		return nil
	}
	log.Debug("key: \"%s\" hit redis node: \"%s\"", key, node)

	return p.Get()
}
