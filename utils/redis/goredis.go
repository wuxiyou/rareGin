package myredis

import (
	"github.com/garyburd/redigo/redis"
	//"inspection/tools"
)


func(this *RedisCoon) Set(key,val string)(interface{},error){
	rc := this.Pool.Get()
	defer rc.Close()
	n,err := rc.Do("SET",redis_prefix + key,val)
	if err != nil {
		//tools.LogError("MyRedis Set Error:",err.Error())
		return nil,err
	}
	return n,err
}

func (this *RedisCoon)Get(key string) (string,error) {
	rc := this.Pool.Get()
	defer rc.Close()
	val,err := redis.String(rc.Do("GET",redis_prefix + key))
	if err != nil {
		//tools.LogError("MyRedis Get Error:",err.Error())
		return "",err
	}
	//tools.LogInfo(val)
	return val,err
}


func(this *RedisCoon)Expire(key string,secondTime int64)(interface{},error){
	rc := this.Pool.Get()
	defer rc.Close()

	n,err := rc.Do("EXPIRE",redis_prefix + key,secondTime)
	if err != nil {
		//tools.LogError("MyRedis Expire Error:",err.Error())
		return nil,err
	}
	return n,err
}




