package main

import (
	"fmt"
)

type Recode struct {
	args     []interface{}
	expected string
}

func send2(cmd string, args []interface{}) error {
	fmt.Printf("cmd=%s\targs=%v\n", cmd, args)
	return nil
}

func send(cmd string, args ...interface{}) error {
	send2(cmd, args)
	return nil
}

func recodeWrite(recodes []Recode) error {
	for _, tt := range recodes {
		send(tt.args[0].(string), tt.args[1:]...)
		fmt.Printf("recode expected=%v\n", tt.expected)
	}
	return nil
}

func main() {
	var recodes [10]Recode
	for i := 0; i < len(recodes); i++ {
		//recode[i].args = make([]interface{}, 0)
		recodes[i].args = []interface{}{"SET", "key", "value"}
		recodes[i].expected = "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n"
	}
	recodeWrite(recodes[:])
}
