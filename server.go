package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "80", "The port on which to server")

	var home string
	flag.StringVar(&home, "home", ".", "The path of the directory to serve as root")

	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(home)))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
