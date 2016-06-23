package main

import (
	"fmt"
	"sort"
)

type MapSorter struct {
	Keys []string
	Vals []int64
}

func NewMapSorter(m map[string]int64) *MapSorter {
	ms := &MapSorter{
		Keys: make([]string, 0, len(m)),
		Vals: make([]int64, 0, len(m)),
	}
	for k, v := range m {
		ms.Keys = append(ms.Keys, k)
		ms.Vals = append(ms.Vals, v)
	}
	return ms
}

func (ms *MapSorter) Sort() {
	sort.Sort(ms)
}

func (ms *MapSorter) Len() int {
	return len(ms.Keys)
}

func (ms *MapSorter) Less(i, j int) bool {
	return ms.Vals[i] > ms.Vals[j]
}

func (ms *MapSorter) Swap(i, j int) {
	ms.Vals[i], ms.Vals[j] = ms.Vals[j], ms.Vals[i]
	ms.Keys[i], ms.Keys[j] = ms.Keys[j], ms.Keys[i]

}

func main() {
	m := map[string]int64{
		"10": 10,
		"1":  1,
		"2":  2,
		"7":  7,
		"3":  3,
		"5":  5,
		"4":  4,
	}
	ms := NewMapSorter(m)
	ms.Sort()
	for i := 0; i < ms.Len(); i++ {
		fmt.Printf("key:%s;value:%d\n", ms.Keys[i], ms.Vals[i])
	}
}
