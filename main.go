package main

import (
	"fmt"
	"github.com/Harry-027/go-audio/cmd"
)

func main() {
	err := cmd.RootCmd.Execute()
	if err != nil {
		fmt.Println("An error occurred :: ", err)
		panic(err)
	}
}
