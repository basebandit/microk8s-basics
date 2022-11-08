// Copyright  Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func main() {
	port := os.Args[1]
	l, close := httpListener(fmt.Sprintf(":%v", port))
	defer close()

	m := new(MyHandler)
	http.Handle("/", m)
	http.HandleFunc("/ready", readyHandler)
	fmt.Printf("Kubia server listening on %v\n", l.Addr().(*net.TCPAddr).Port)
	http.Serve(l, nil)
}

func httpListener(port string) (l net.Listener, close func()) {
	l, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	fmt.Println("Kubia server starting...")
	return l, func() {
		_ = l.Close()
	}
}

type MyHandler struct{}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request from %v\n", r.RemoteAddr)
	w.WriteHeader(http.StatusOK)
	host, err := os.Hostname()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Ooops, an unexpected error occurred %v\n", err.Error())
	}
	fmt.Fprintf(w, "You've hit %v\n", host)
}

func readyHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK\n")
}
