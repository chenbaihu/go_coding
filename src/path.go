package main

import (
	"fmt"
	"path"
)

func main() {
	p := "/home/chenbaihu/code_test/go_coding/path.go"
	fmt.Println("file_name_in_path:", path.Base(p))

	fmt.Println("dir_name_in_path:", path.Dir(p))

	fmt.Println("ext_name_in_path:", path.Ext(p))

	if path.IsAbs(p) {
		fmt.Println("is abs")
	}

	if !path.IsAbs("./path.go") {
		fmt.Println("is not abs")
	}

	p2 := path.Join("home", "chenbaihu", "code_test", "go_coding", "path.go")
	println(p2)

	match, err := path.Match("*", p)
	if !match {
		fmt.Println("path Match false")
	}
	if err == nil {
		fmt.Println("path Match err:", err)
	}

	dirname, filename := path.Split(p)
	fmt.Printf("dirname=%s\tfilename=%s\n", dirname, filename)
}
