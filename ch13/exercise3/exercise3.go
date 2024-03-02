package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

type TimeResponse struct {
	DayOfWeek  string `json: "day_of_week"`
	DayOfMonth int    `json: "day_of_month"`
	Month      string `json: "month"`
	Year       int    `json: "year"`
	Hour       int    `json: "hour"`
	Minute     int    `json: "minute"`
	Second     int    `json: "second"`
}

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

	now := time.Now()

	timeResponse := TimeResponse{
		DayOfWeek:  now.Weekday().String(),
		DayOfMonth: now.Day(),
		Month:      now.Month().String(),
		Year:       now.Year(),
		Hour:       now.Hour(),
		Minute:     now.Minute(),
		Second:     now.Second(),
	}

	jsonResponse, err := json.Marshal(timeResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, string(jsonResponse))
}

func getIp(r *http.Request) string {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "unknown"
	}

	return ip
}
