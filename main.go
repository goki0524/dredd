package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/message", index)

	p := port()
	fmt.Println("Listening on Port", p)
	http.ListenAndServe(p, nil)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	return ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello World!!")
}
