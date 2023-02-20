package main

import (
	"fmt"
	"lb/utils"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeRoute)

	// ListenAndServe returns only an err if failed. so, it must be handled
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		utils.HandleError(err)
	}
}

func homeRoute(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Request Method Not Supported", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/plain")

	// writing RESPONSE using fmt package
	fmt.Fprint(w, "Server is Spinning...")
}
