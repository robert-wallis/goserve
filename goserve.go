// Copyright (C) 2018 Robert A. Wallis, All Rights Reserved
package main

import (
	"flag"
	"log"
	"net/http"
)

var port = flag.String("port", "localhost:8000", "server binding")

func main() {
	flag.Parse()
	var fileServer = http.FileServer(http.Dir("."))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RequestURI, r.RemoteAddr)
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		fileServer.ServeHTTP(w, r)
	})
	log.Fatal(http.ListenAndServe(*port, nil))
}
