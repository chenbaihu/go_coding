package kv

import (
	"errors"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"time"
)

type MemcacheClient struct {
	Servers   []string
	TimeoutMs int

	client *memcache.Client
}

func (client *MemcacheClient) Close() {
	if client.client != nil {
		//todo, fix close
	}
}

func (client *MemcacheClient) Init() error {
	if len(client.Servers) == 0 || 0 > client.TimeoutMs {
		err := errors.New(fmt.Sprintf("invalid memcache config: servers %s timeout %d",
			client.Servers, client.TimeoutMs))
		return err
	}

	//client.client = memcache.New(memcache.RANDOM, client.Servers...)
	client.client = memcache.New(client.Servers...)
	client.client.Timeout = time.Duration(client.TimeoutMs) * time.Millisecond

	return nil
}

func (client *MemcacheClient) GetMulti(keys []string) (map[string][]byte, error) {
	if client.client == nil {
		err := errors.New("internal error - MemcacheClient not init")
		return nil, err
	}

	dataMap, err := client.client.GetMulti(keys)
	if err != nil {
		if err.Error() == "memcache: cache miss" {
			fmt.Printf("error=[mc_getmulti_failed] err=[%s]\n", err.Error())
			return nil, err
		} else {
			fmt.Printf("error=[mc_getmulti_failed] err=[%s]\n", err.Error())
			return nil, err
		}
	}

	result := make(map[string][]byte)
	for k, v := range dataMap {
		result[k] = v.Value
	}
	return result, nil
}

func (client *MemcacheClient) Set(key string, data []byte, expiretime int32) error {
	if client.client == nil {
		err := errors.New("internal error - MemcacheClient not init")
		return err
	}

	item := &memcache.Item{Key: key, Value: data, Expiration: expiretime}
	err := client.client.Set(item)
	if err != nil {
		fmt.Printf("error=[mc_set_failed] key=[%s] err=[%s]\n", key, err.Error())
	}
	return err
}

func (client *MemcacheClient) Get(key string) ([]byte, error) {
	if client.client == nil {
		err := errors.New("internal error - MemcacheClient not init")
		return nil, err
	}

	data, err := client.client.Get(key)
	if err != nil {
		if err.Error() == "memcache: cache miss" {
			fmt.Printf("error=[mc_get_failed] key=[%s] err=[%s]\n", key, err.Error())
			return nil, err
		} else {
			fmt.Printf("error=[mc_get_failed] key=[%s] err=[%s]\n", key, err.Error())
			return nil, err
		}
	}
	return data.Value, nil
}

func (client *MemcacheClient) Delete(key string) error {
	if client.client == nil {
		err := errors.New("internal error - MemcacheClient not init")
		return err
	}

	err := client.client.Delete(key)
	if err != nil {
		if err.Error() == "memcache: cache miss" {
			fmt.Printf("error=[mc_delete_failed] key=[%s] err=[%s]\n", key, err.Error())
			return err
		} else {
			fmt.Printf("error=[mc_delete_failed] key=[%s] err=[%s]\n", key, err.Error())
		}
	}
	return err
}
