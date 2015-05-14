package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("", "/home/chenbaihu/", "test_env")

	//cmd.Path = "/bin/ls"
	//cmd.Args = append(cmd.Args, "/home/chenbaihu/")
	//cmd.Env = append(cmd.Env, "test_env")
	fmt.Printf("cmd=%v\n", cmd)

	er := cmd.Run()
	if er == nil {
		fmt.Printf("cmd=%v Run failed\n", cmd)
		return
	}

	//result := make([]byte, 4096)
	result, err := cmd.Output()
	if err == nil {
		fmt.Printf("cmd=%v Output faild\n", cmd)
		return
	}
	fmt.Printf("result=%s, result=%v\n", result, result)

	cmd = exec.Command("tr", "a-z", "A-Z")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
}
