package main

import (
	"fmt"
	"sort"
)

func init() {
	builtin["help"] = cmdHelp
}

func cmdHelp([]string) {
	commands := []string{}
	for cmd := range builtin {
		commands = append(commands, cmd)
	}

	sort.Strings(commands)

	for _, cmd := range commands {
		fmt.Println(cmd)
	}
}
