package main

import (
	"net/http"
)

func setupRoutes() {
	http.HandleFunc("/", home)
}
