package main

import (
	"fmt"
	"kv"
	"os"
)

func main() {
	mcHosts := make([]string, 3)
	mcHosts[0] = "127.0.0.1:11311"
	mcHosts[1] = "127.0.0.1:11311"
	mcHosts[2] = "127.0.0.1:11311"
	mcKV := &kv.MemcacheClient{
		Servers:   mcHosts,
		TimeoutMs: 100, //100ms
	}
	err := mcKV.Init()
	if err != nil {
		fmt.Printf("init one memcache failed, hosts:%v, err:%s\n", mcHosts, err.Error())
		os.Exit(1)
	}
	defer mcKV.Close()

	//TODO get and set value

	redisHosts := make([]string, 5)
	redisHosts[0] = "127.0.0.1:6379"
	redisHosts[1] = "127.0.0.1:6379"
	redisHosts[2] = "127.0.0.1:6379"
	redisHosts[3] = "127.0.0.1:6379"
	redisHosts[4] = "127.0.0.1:6379"
	redisKV := &kv.RedisClient{
		Servers:        redisHosts,
		ConnTimeoutMs:  100, //100ms
		WriteTimeoutMs: 200, //200ms
		ReadTimeoutMs:  200, //200ms
		MaxIdle:        50,
		MaxActive:      1000,
		IdleTimeoutS:   900,
		Password:       "",
	}
	err = redisKV.Init()
	if err != nil {
		fmt.Printf("init redis failed, err:%s", err.Error())
		return
	}
	defer redisKV.Close()

}
