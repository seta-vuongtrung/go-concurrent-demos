package main

import (
	"log"
	"net/http"
)

// go run main.go
// curl --location 'localhost:3000'
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	addr := ":3000"

	log.Println("listen on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
