package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/time", timeHandler)

	http.ListenAndServe(":8080", mux)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	current_time := time.Now().Format(time.RFC3339)
	fmt.Fprint(w, current_time)
}
