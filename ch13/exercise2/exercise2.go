package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	log.SetOutput(os.Stderr)
	log.SetFlags(log.LstdFlags)

	router := http.NewServeMux()
	router.HandleFunc("/", currentTimeHandler)

	server := &http.Server{
		Addr:        ":8080",
		Handler:     router,
		ReadTimeout: 5 * time.Second,
	}

	fmt.Println("Server listening on http://localhost:8080")
	server.ListenAndServe()
}

func currentTimeHandler(w http.ResponseWriter, r *http.Request) {
	ip := getIp(r)
	log.Printf("Request from Ip: %s, Method: %s, Path: %s\n", ip, r.Method, r.URL.Path)

	now := time.Now().Format(time.RFC3339)
	fmt.Fprintln(w, now)
}

func getIp(r *http.Request) string {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "unknown"
	}

	return ip
}
