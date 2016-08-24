package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	host, _ := os.Hostname()
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		fmt.Fprintf(w, fmt.Sprintf("Hello from server %s!", host))
		elapse := time.Since(now) / time.Millisecond
		log.Printf("%s %s %dms\n", r.Method, r.RequestURI, elapse)
	})

	log.Println(fmt.Sprintf("Server running on %s:%s", host, port))

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		panic(err)
	}
}
