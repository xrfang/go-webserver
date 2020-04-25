package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	audit "github.com/xrfang/go-audit"
	res "github.com/xrfang/go-res"
)

func main() {
	conf := flag.String("conf", "", "configuration file")
	ver := flag.Bool("version", false, "show version info")
	init := flag.Bool("init", false, "initialize configuration")
	flag.Usage = func() {
		fmt.Printf("WebServer - Go WebServer Template %s\n", verinfo())
		fmt.Printf("\nUSAGE: %s OPTIONS\n", filepath.Base(os.Args[0]))
		fmt.Println("\nOPTIONS")
		flag.PrintDefaults()
	}
	flag.Parse()
	if *ver {
		fmt.Println(verinfo())
		return
	}
	if *init {
		fmt.Println("TODO: initialize configuration")
		return
	}
	loadConfig(*conf)
	if !cf.DbgMode {
		audit.Assert(res.Extract(cf.WebRoot, res.OverwriteIfNewer))
	}
	audit.ExpVars(map[string]interface{}{
		"config":  cf,
		"version": _G_REVS + "." + _G_HASH,
	})
	audit.SetLogFile(cf.LogFile)
	audit.ExpLogs()
	audit.SetDebugging(cf.DbgMode)
	setupRoutes()
	svr := http.Server{
		Addr:         ":" + cf.Port,
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
	}
	audit.Assert(svr.ListenAndServe())
}
