package main

import (
	"fmt"
	"os"
)

func init() {
	builtin["env"] = getEnv
}

func getEnv(args []string) {
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
	fmt.Println()
}
