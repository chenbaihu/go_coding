package main

import (
	"fmt"
)

func ShowFlags() {
	fmt.Printf("ShowFlags: Port=%d\tPercent=%f\tMainConf=%s\tAsync=%v\n", Port, Percent, MainConf, Async)
}
