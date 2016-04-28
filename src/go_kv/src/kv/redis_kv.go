package kv

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"math/rand"
	"time"
)

type RedisClient struct {
	Servers        []string
	ConnTimeoutMs  int
	WriteTimeoutMs int
	ReadTimeoutMs  int

	MaxIdle      int
	MaxActive    int
	IdleTimeoutS int
	Password     string

	current_index int
	pool          *redis.Pool
}

func (client *RedisClient) Close() {
	client.pool.Close()
}

func (client *RedisClient) Init() error {
	if len(client.Servers) == 0 {
		return fmt.Errorf("invalid Redis config servers:%s", client.Servers)
	}

	client.pool = &redis.Pool{
		MaxIdle:     client.MaxIdle,
		IdleTimeout: time.Duration(client.IdleTimeoutS) * time.Second,
		MaxActive:   client.MaxActive,
		Dial: func() (redis.Conn, error) {
			var c redis.Conn
			var err error
			for i := 0; i < len(client.Servers); i++ {
				//随机挑选一个IP
				index := RandIntn(len(client.Servers))
				client.current_index = index
				c, err = redis.DialTimeout("tcp", client.Servers[index],
					time.Duration(client.ConnTimeoutMs)*time.Millisecond,
					time.Duration(client.ReadTimeoutMs)*time.Millisecond,
					time.Duration(client.WriteTimeoutMs)*time.Millisecond)
				if err != nil {
					fmt.Printf("warning=[redis_connect_failed] num=[%d] server=[%s] err=[%s]",
						i, client.Servers[index], err.Error())
				}
				//支持密码认证
				if len(client.Password) > 0 {
					if _, err_pass := c.Do("AUTH", client.Password); err_pass != nil {
						c.Close()
					}
				}
				if err == nil {
					fmt.Printf("info=[redis_connect_ok] num=[%d] server=[%s] err=[%s]",
						i, client.Servers[index])
					break
				}
			}
			return c, err
		},
	}

	return nil
}

func (client *RedisClient) Set(key string, value []byte) error {
	conn := client.pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	if err != nil {
		fmt.Printf("error=[redis_set_failed] server=[%s] key=[%s] err=[%s]",
			client.Servers[client.current_index], key, err.Error())

		conn_second := client.pool.Get()
		defer conn_second.Close()
		_, err = conn_second.Do("SET", key, value)
		if err != nil {
			fmt.Printf("second error=[redis_set_failed] server=[%s] key=[%s] err=[%s]",
				client.Servers[client.current_index], key, err.Error())
			return err
		}
	}

	return nil
}

func (client *RedisClient) Get(key string) ([]byte, error) {
	conn := client.pool.Get()
	defer conn.Close()

	value, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		if err.Error() == "redigo: nil returned" {
			fmt.Printf("error=[redis_get_failed] server=[%s] key=[%s] err=[%s]",
				client.Servers[client.current_index], key, err.Error())
			return nil, err
		} else {
			fmt.Printf("error=[redis_get_failed] server=[%s] key=[%s] err=[%s]",
				client.Servers[client.current_index], key, err.Error())
		}

		conn_second := client.pool.Get()
		defer conn_second.Close()

		value, err = redis.Bytes(conn_second.Do("GET", key))
		if err != nil {
			if err.Error() == "redigo: nil returned" {
				fmt.Printf("second error=[redis_get_failed] server=[%s] key=[%s] err=[%s]",
					client.Servers[client.current_index], key, err.Error())
			} else {
				fmt.Printf("second error=[redis_get_failed] server=[%s] key=[%s] err=[%s]",
					client.Servers[client.current_index], key, err.Error())
			}
			return nil, err
		}
	}

	return value, nil
}

func (client *RedisClient) Rpush(key string, value string) (int64, error) {
	conn := client.pool.Get()
	defer conn.Close()

	list_len, err := redis.Int64(conn.Do("RPUSH", key, value))
	if err != nil {
		fmt.Printf("error=[redis_rpush_failed] server=[%s] key=[%s] err=[%s]",
			client.Servers[client.current_index], key, err.Error())

		conn_second := client.pool.Get()
		defer conn_second.Close()

		list_len, err = redis.Int64(conn_second.Do("RPUSH", key, value))
		if err != nil {
			fmt.Printf("second error=[redis_rpush_failed] server=[%s] key=[%s] err=[%s]",
				client.Servers[client.current_index], key, err.Error())
			return -1, err
		}
	}

	return list_len, nil
}

func (client *RedisClient) Lpop(key string) (string, error) {
	conn := client.pool.Get()
	defer conn.Close()

	value, err := redis.String(conn.Do("LPOP", key))
	if err != nil {
		if err.Error() == "redigo: nil returned" {
			fmt.Printf("error=[redis_lpop_failed] server=[%s] key=[%s] err=[%s]",
				client.Servers[client.current_index], key, err.Error())
			return "", err
		} else {
			fmt.Printf("error=[redis_lpop_failed] server=[%s] key=[%s] err=[%s]",
				client.Servers[client.current_index], key, err.Error())
		}

		conn_second := client.pool.Get()
		defer conn_second.Close()

		value, err = redis.String(conn_second.Do("LPOP", key))
		if err != nil {
			if err.Error() == "redigo: nil returned" {
				fmt.Printf("second error=[redis_lpop_failed] server=[%s] key=[%s] err=[%s]",
					client.Servers[client.current_index], key, err.Error())
			} else {
				fmt.Printf("second error=[redis_lpop_failed] server=[%s] key=[%s] err=[%s]",
					client.Servers[client.current_index], key, err.Error())
			}
			return "", err
		}
	}

	return value, nil
}

func (client *RedisClient) Llen(key string) (int64, error) {
	conn := client.pool.Get()
	defer conn.Close()

	value, err := redis.Int64(conn.Do("LLEN", key))
	if err != nil {
		fmt.Printf("error=[redis_llen_failed] server=[%s] key=[%s] err=[%s]",
			client.Servers[client.current_index], key, err.Error())

		conn_second := client.pool.Get()
		defer conn_second.Close()

		value, err = redis.Int64(conn_second.Do("LLEN", key))
		if err != nil {
			fmt.Printf("second error=[redis_llen_failed] server=[%s] key=[%s] err=[%s]",
				client.Servers[client.current_index], key, err.Error())
			return -1, err
		}
	}
	return value, nil
}

func (client *RedisClient) Del(keys []interface{}) (int64, error) {
	conn := client.pool.Get()
	defer conn.Close()

	value, err := redis.Int64(conn.Do("DEL", keys...))
	if err != nil {
		fmt.Printf("error=[redis_del_failed] server=[%s] keys=[%v] err=[%s]",
			client.Servers[client.current_index], keys, err.Error())

		conn_second := client.pool.Get()
		defer conn_second.Close()

		value, err = redis.Int64(conn_second.Do("DEL", keys...))
		if err != nil {
			fmt.Printf("second error=[redis_del_failed] server=[%s] keys=[%v] err=[%s]",
				client.Servers[client.current_index], keys, err.Error())
			return -1, err
		}
	}

	return value, nil
}

//var rand_gen = rand.New(rand.NewSource(time.Now().UnixNano()))
func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandInt() int {
	//return rand_gen.Int()
	return rand.Int()
}

func RandIntn(max int) int {
	//return rand_gen.Intn(max)
	return rand.Intn(max)
}

func NowInS() int64 {
	return time.Now().Unix()
}

func NowInNs() int64 {
	return time.Now().UnixNano()
}
