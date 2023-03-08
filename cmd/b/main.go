package main

import (
	"fmt"
	"github.com/lishimeng/app-starter/buildscript"
)

func main() {
	err := buildscript.Generate("oauth2",
		"lishimeng",
		"cmd/oauth/main.go", false)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}
}
