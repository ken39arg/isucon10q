package main

import (
	"net/http"
	_ "net/http/pprof"
)

func runProfile() {
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()
}
