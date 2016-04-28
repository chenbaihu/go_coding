package main

import (
	"fmt"
	"strconv"
	"sync"
)

//带mutex的struct必须是指针receivers
//如果你定义的struct中带有mutex,那么你的receivers必须是指针

var wg sync.WaitGroup

func getKVFromRedis(keys []string, kvMap map[string]string) bool {
	for i := 0; i < len(keys); i++ {
		//fmt.Printf("%s\n", keys[i])
		kvMap[keys[i]] = "value" + strconv.FormatInt((int64)(i), 10)
	}
	return true
}

func main() {
	keys := make([]string, 1000)
	keysLen := len(keys)
	for i := 0; i < keysLen; i++ {
		//fmt.Printf("%d\n", i)
		keys[i] = strconv.FormatInt((int64)(i), 10)
	}

	sliceNum := keysLen / 50
	kvMap := make([]map[string]string, sliceNum)

	wg.Add(sliceNum)
	for i := 0; i < sliceNum; i++ {
		go func(sliceID, startIndex, endIndex int) {
			fmt.Printf("%d<-->%d\t%d\t%v\n",
				startIndex,
				endIndex,
				len(keys[startIndex:endIndex]),
				keys[startIndex:endIndex])
			subKVMap := make(map[string]string)
			getKVFromRedis(keys[startIndex:endIndex], subKVMap) //map默认传引用
			kvMap[sliceID] = subKVMap
			wg.Done()
		}(i, i*50, (i+1)*50)
	}
	wg.Wait()

	for id, subKVMap := range kvMap {
		fmt.Printf("id=%d\tsubKVMap=%s\n", id, subKVMap)
	}
}
