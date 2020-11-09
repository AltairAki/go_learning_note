package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

// GetFibonacciserie function
func GetFibonacciserie(n int) []int64 {
	fidList := make([]int64, 2, n)
	fidList[0] = 1
	fidList[1] = 1
	for i := 2; i < n; i++ {
		fidList = append(fidList, fidList[i-2]+fidList[i-1])
	}
	return fidList
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome!"))
}

func createFBS(w http.ResponseWriter, r *http.Request) {
	var fbs []int64
	for i := 0; i < 10000; i++ {
		fbs = GetFibonacciserie(100)
	}
	w.Write([]byte(fmt.Sprintf("%v", fbs)))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/fb", createFBS)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
