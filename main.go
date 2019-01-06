package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	builtin = map[string]func([]string){}
	cwd     = "./"
)

func init() {
	builtin["exit"] = cmdExit
	builtin["echo"] = cmdEcho
	builtin["pwd"] = cmdPwd
}

func main() {
	if len(os.Args) > 1 {
		cwd = os.Args[1]
	} else {
		cwd, _ = os.Getwd()
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for scanner.Scan() {
		line := process(scanner.Text())
		args := strings.Fields(line)
		if len(args) == 0 {
			continue
		}

		if cmd, found := builtin[args[0]]; found {
			cmd(args)
		} else {
			fmt.Fprintf(os.Stderr, "Command not found '%v'\n", line)
		}

		fmt.Print("> ")
	}
}

func process(line string) string {
	re := regexp.MustCompile("\\$([a-zA-Z_]*)")
	line = re.ReplaceAllStringFunc(line, func(name string) string {
		return os.Getenv(name[1:])
	})
	return line
}

func cmdExit([]string) {
	os.Exit(0)
}

func cmdEcho(args []string) {
	fmt.Println(strings.Join(args[1:], " "))
}

func cmdPwd([]string) {
	abs, _ := filepath.Abs(cwd)
	fmt.Println(abs)
}
