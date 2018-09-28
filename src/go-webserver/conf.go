package main

import (
	"os"
	"path"
	"path/filepath"
)

type Configuration struct {
	LogFile string
	Port    string
	WebRoot string

	binPath string
	cfgFile string
	cfgPath string
}

func canonicalize(fn string) string {
	if fn == "" || path.IsAbs(fn) {
		return fn
	}
	p, _ := filepath.Abs(path.Join(cf.binPath, fn))
	return p
}

var cf Configuration

func loadConfig() {
	cf.binPath = path.Dir(os.Args[0])
	cf.Port = "8080"
	cf.WebRoot = "../webroot"
	cf.LogFile = "../log/log"
	cf.WebRoot = canonicalize(cf.WebRoot)
	cf.LogFile = canonicalize(cf.LogFile)
}
