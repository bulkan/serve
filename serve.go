package main

import (
    "net/http"
    "log"
    "flag"
)

func main() {

    dirPtr := flag.String("dir", ".", "directory to serve")

    flag.Parse()

    err := http.ListenAndServe(":4242",
        http.FileServer(http.Dir(*dirPtr)))

    if err != nil {
        log.Printf("Error running web server for static assets: %v", err)
    }
}
