package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(([]byte("Hello World!")))
	})
	http.HandleFunc("/time/", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		timeStr := fmt.Sprintf("{\"time\":\"%s\"}", t)
		w.Write(([]byte(timeStr)))
		// w.WriteHeader(200)
	})
	// http.ListenAndServe(addr, handler)
	// http.ListenAndServeTLS(addr, certFile, keyFile, handler)
	http.ListenAndServe(":8080", nil)
}
