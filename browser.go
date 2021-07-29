package main

import (
	"net/url"
	"os"
	"os/exec"

	"github.com/taybart/log"
)

const (
	Default = "firefox"
)

// https://taybart.com
// https://zoom.us/test

func openBrowser() {

	log.Info("Opening something", os.Args)
	arg := os.Args[1]
	var handler string
	if _, err := os.Stat(arg); !os.IsNotExist(err) { // is file
		log.Info("argument was a file")
		handler = Default
	} else {
		log.Info("argument probably a url")
		u, err := url.Parse(arg)
		if err != nil {
			log.Fatal("could not parse url")
		}
		log.Info("parsed url", u.Host)
		switch u.Host {
		case "zoom.us":
			handler = "zoom"
		default:
			handler = Default
		}
	}

	cmd := exec.Command(handler, os.Args[1:]...)
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Process.Release(); err != nil {
		log.Fatal("cmd.Process.Release failed: ", err)
	}
}
