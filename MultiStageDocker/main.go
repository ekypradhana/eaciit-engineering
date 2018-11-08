package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello there..."))
	})
	fmt.Println("Application is listening on port : 8080")
	http.ListenAndServe(":8080", mux)
}
