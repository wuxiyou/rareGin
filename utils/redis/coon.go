package myredis

import (
	"github.com/goini"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
	"strconv"
	"math/rand"
)

type RedisCoon struct {
	Pool   *redis.Pool
}
var ConfigCentor *goini.Config
var redis_prefix string
func Init() (*RedisCoon){
	// 从配置文件获取redis配置并连接
	ConfigCentor = goini.SetConfig("./config/config.ini")
	redis_host := ConfigCentor.GetValue("redis", "redis_host")
	redis_port := ConfigCentor.GetValue("redis", "redis_port")

	redis_database := ConfigCentor.GetValue("redis", "redis_database")
	redis_maxidle := ConfigCentor.GetValue("redis", "redis_maxidle")
	redis_maxactive := ConfigCentor.GetValue("redis", "maxactive")
	redis_idle_timeout := ConfigCentor.GetValue("redis", "redis_idle_timeout")
	redis_maxidle_int,_ := strconv.Atoi(redis_maxidle)
	redis_maxactive_int,_ := strconv.Atoi(redis_maxactive)
	redis_idle_timeout_int,_ := strconv.Atoi(redis_idle_timeout)
	data_str := fmt.Sprintf("%s:%s", redis_host, redis_port)
	redis_prefix = ConfigCentor.GetValue("redis","redis_prefix")

	// 建立连接池
	this := new(RedisCoon)
	this.Pool = &redis.Pool{
		MaxIdle:     redis_maxidle_int,
		MaxActive:   redis_maxactive_int,
		IdleTimeout: time.Duration(rand.Intn(redis_idle_timeout_int)) * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", data_str)
			if err != nil {
				//tools.LogError("Redis Init Error:", err.Error())
				return nil, err
			}
			// 选择db
			c.Do("SELECT", redis_database)
			return c, nil
		},
	}
	return this
}


func (this *RedisCoon) PoolClose() {
	this.Pool.Close()
}


