package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/pflag"
)

func init() {
	builtin["ls"] = cmdLs
}

var (
	printLink       = color.New(color.FgCyan).PrintfFunc()
	printDirectory  = color.New(color.FgBlue).PrintfFunc()
	printExecutable = color.New(color.FgGreen).PrintfFunc()
)

const modeExe = 0111

func cmdLs(args []string) {
	var list bool

	f := pflag.NewFlagSet("", pflag.ContinueOnError)
	f.BoolVarP(&list, "list", "l", false, "")
	f.Parse(args[1:])

	dir := f.Arg(0)
	if dir == "" {
		dir = cwd
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	for _, file := range files {
		format := "%v "
		if list {
			format = "%v\n"
		}

		if file.IsDir() {
			printDirectory(format, file.Name())
		} else if file.Mode()&os.ModeSymlink > 0 {
			printLink(format, file.Name())
		} else if file.Mode()&modeExe > 0 {
			printExecutable(format, file.Name())
		} else {
			fmt.Printf(format, file.Name())
		}
	}

	if !list {
		fmt.Println()
	}
}
