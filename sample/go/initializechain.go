package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

var chiledInitPath = "/child_initialize"

type initChiledData struct {
	Foo string
}

func handleInitChild(m *http.ServeMux) {
	m.HandleFunc(chiledInitPath, postChildInitialize)
}

// initChilds("isu2:5000", "isu3:5000")

func initChilds(hosts ...string) error {
	data, err := json.Marshal(&initChiledData{
		Foo: "test",
	})
	if err != nil {
		return err
	}
	for _, host := range hosts {
		u := "http://" + host + chiledInitPath
		res, err := http.Post(u, "application/json", bytes.NewReader(data))
		if err != nil {
			return err
		}
		io.Copy(ioutil.Discard, res.Body)
		res.Body.Close()
	}
	return nil
}

func postChildInitialize(w http.ResponseWriter, r *http.Request) {
	data := &initChiledData{}
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// TODO: something

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "ok")
}
