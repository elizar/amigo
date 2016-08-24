package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		host, _ := os.Hostname()
		fmt.Fprintf(w, fmt.Sprintf("Hello from server %s!", host))
		elapse := time.Since(now) / time.Millisecond
		log.Printf("%s %s %dms", r.Method, r.RequestURI, elapse)
	})

	fmt.Println("Server running...")
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		panic(err)
	}
}
