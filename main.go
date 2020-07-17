package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	for {
		log.Print("Just keep on logging")
		time.Sleep(time.Second)
	}
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
