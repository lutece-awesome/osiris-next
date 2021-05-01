package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	port      = flag.Int("port", 9877, "The http server port")
	staticDir = flag.String("static", "./test/http/static", "Static file path")
)

func main() {
	flag.Parse()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(*staticDir))))
	log.Printf("test http server listen to :%d\n", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		panic(fmt.Sprintf("can not run http server with %v", err))
	}
}
