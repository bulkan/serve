package main

import (
    "net/http"
    "log"
    "fmt"
    "flag"
)

func main() {

    dirPtr := flag.String("dir", ".", "directory to serve")
    portPtr := flag.String("port", "4242", "port number")

    flag.Parse()

    fmt.Printf("Serving %s on port %s\n", *dirPtr, *portPtr)

    if (*portPtr)[0] != ':' {
        *portPtr = ":" + *portPtr
    }

    dir := http.Dir(*dirPtr)

    err := http.ListenAndServe(":4242",
        http.FileServer(dir))

    if err != nil {
        log.Printf("Error running web server for static assets: %v", err)
    }
}
