// Copyright (C) 2018 Robert A. Wallis, All Rights Reserved

// goserve is a very simple localhost:8000 static file server, for developing static sites.
//
// Usage: run `goserve` in a folder.  It's files will be available at http://localhost:8000/
package main

import (
	"flag"
	"fmt"
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
	fmt.Println("Listening at", *port)
	log.Fatal(http.ListenAndServe(*port, nil))
}
