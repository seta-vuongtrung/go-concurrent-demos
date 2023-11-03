package main

import (
	"log"
	"net/http"
)

// go run main.go
// curl --location 'localhost:3000?name=seta-vuongtrung'
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		names := r.URL.Query()["name"]
		var name string
		if len(names) == 1 {
			name = names[0]
		}
		w.Write([]byte("Hello, " + name))
	})

	addr := ":3000"

	log.Println("listen on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
