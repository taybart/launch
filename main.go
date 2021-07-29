package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/taybart/args"
	"github.com/taybart/env"
	"github.com/taybart/log"
)

func main() {
	env.Set([]string{
		"VERSION?",
	})

	app := args.App{
		Name:    "launch",
		Version: env.Get("VERSION"),

		Author: "taylor <taybart@gmail.com>",
		About:  "launcher for things",
		Args: map[string]*args.Arg{
			"browser": {
				Short:   "b",
				Long:    "browser",
				Help:    "Send to browser handler",
				Default: false,
			},
		},
	}

	if len(os.Args) <= 1 {
		fmt.Println("Usage: launch <command> <optional parameters>")
		os.Exit(1)
	}
	cmd := exec.Command(os.Args[1], os.Args[2:]...)

	// this is the important part to avoid killing the child process instantly
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}
