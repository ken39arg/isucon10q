package main

import (
	"net"
	"net/http"
)

func runUnix(svr http.Handler) error {
	listener, err := net.Listen("unix", "/var/run/app.sock")
	if err != nil {
		return err
	}
	defer listener.Close()

	return http.Serve(listener, svr)
}

func runTCP(svr http.handler) error {
	s := &http.Server{Addr: ":5000", Handler: svr}
	return s.ListenAndServe()
}
