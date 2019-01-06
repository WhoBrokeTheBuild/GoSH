package main

import (
	"os"
	"path/filepath"
)

func init() {
	builtin["cd"] = cmdCd
}

func cmdCd(args []string) {
	if len(args) < 2 {
		cwd, _ = os.Getwd()
	} else if filepath.IsAbs(args[1]) {
		cwd = args[1]
	} else {
		cwd = filepath.Join(cwd, args[1])
	}
}
