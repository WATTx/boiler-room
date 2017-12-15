package main

import (
	"flag"
	"log"
)

var (
	port = flag.String("port", "8080", "HTTP server port")
)

func main() {
	flag.Parse()
	log.Printf("Starting boiler room api...")

	s := NewServer(*port)
	panic(s.ListenAndServe())
}
