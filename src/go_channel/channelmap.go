package main

import (
	"fmt"
	"sync"
)

type myMap struct {
	m map[string]interface{}
	sync.Mutex
}

func (m *myMap) push(key string, e interface{}) interface{} {
	m.Lock()
	defer m.Unlock()
	if v, exist := m.m[key]; exist {
		return v
	}
	m.m[key] = e
	return nil
}

func (m *myMap) pop(key string) interface{} {
	m.Lock()
	defer m.Unlock()
	if v, exist := m.m[key]; exist {
		m.m[key] = nil
		return v
	}
	return nil
}

func newMap() *myMap {
	return &myMap{m: make(map[string]interface{})}
}

func hello() {
	fmt.Println("hello fun test")
}

func main() {
	m := newMap()
	fmt.Println(m.push("hello", "world"))
	fmt.Println(m.push("hello", "world"))
	fmt.Println(m.pop("hello"))
	fmt.Println(m.pop("hello"))

	fmt.Println(m.push("test", 1))
	fmt.Println(m.push("test", 1))
	fmt.Println(m.pop("test"))
	fmt.Println(m.pop("test"))

	fmt.Println(m.push("func", hello))
	fmt.Println(m.push("func", hello))
	fmt.Println(m.push("func", hello))
	m.pop("func")
}
