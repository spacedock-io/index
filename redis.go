package main

import (
  "log"
  "github.com/boj/redistore"
  "github.com/garyburd/redigo/redis"
  "github.com/stretchr/objx"
  "time"
)

var (
  RedisConn redis.Conn
  RediStore *redistore.RediStore
)

func NewRedis(config objx.Map) {
  var err error

  // Bug(Colton): Again, the objx bug raises its head. Int is being picked up
  // as Float64. Still got to look into why this is happening.
  ct, wt, rt :=
    time.Duration(int(config.Get("redis.connectTimeout").Float64(5))) *
      time.Second,
    time.Duration(int(config.Get("redis.writeTimeout").Float64(30))) *
      time.Second,
    time.Duration(int(config.Get("redis.readTimeout").Float64(30))) *
      time.Second

  RedisConn, err = redis.DialTimeout("tcp",
    config.Get("redis.host").Str(":6379"), ct, wt, rt)
  if err != nil {
    log.Fatalln(err)
  }

  RediStore = redistore.NewRediStore(10, "tcp",
    config.Get("redis.host").Str(":6379"), "",
    []byte(config.Get("redis.secret").Str("SECRET")))
}
